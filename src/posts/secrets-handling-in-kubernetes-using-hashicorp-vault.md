---
type: Blog
title: Secrets handling in Kubernetes using HashiCorp Vault
subheading: In this blog we explore several approaches for getting secrets from HashiCorp Vault into Jenkins running in Kubernetes
authors:
- mvainio
tags:
- Kubernetes
- Vault
- Jenkins
date: 2022-10-12
image: "/blogs/secrets-handling-in-kubernetes-using-hashicorp-vault/secrets-handling-in-kubernetes-using-hashicorp-vault.png"
featured: true

---

When running Jenkins jobs using the Kubernetes plugin, there are many ways to fetch secrets from HashiCorp Vault. [In the previous post](/blog/secrets-handling-in-kubernetes-a-jenkins-story/) we stored the secrets in Kubernetes, but let’s look at options that don’t persist secrets in the cluster. 

In the earlier post the secrets were not stored as plain-text in the Jenkins controller, but still the secrets were stored together with the encryption key on the persistent volume. Using a secret manager allows us to increase the security posture further, and make our lives easier when there is one central location to store and modify the secrets, which also records an audit trail of all activities. When using Vault as secret manager the secrets are only fetched into the running worker pods in the beginning of a pipeline. When the pipeline finishes, the pod is terminated and the secrets are no longer anywhere in the cluster.

## Options for Integrating Jenkins with Vault

Here are some of the most popular options of integrating Jenkins with Vault:

### 1. Jenkins HashiCorp Vault plugin

Yet another [plugin](https://plugins.jenkins.io/hashicorp-vault-plugin/) to your Jenkins, yikes! When not using Kubernetes this is a good choice, but it requires providing static credentials to the Jenkins controller which we can avoid in Kubernetes by leveraging the built-in authentication: service accounts.

### 2. HashiCorp Vault Agent sidecar injector

The [sidecar injector](https://www.vaultproject.io/docs/platform/k8s/injector) is a mutating admission webhook controller which is installed to the Kubernetes cluster by using [the official Vault Helm chart](https://github.com/hashicorp/vault-helm). The authentication requires configuring the [Kubernetes auth method](https://www.vaultproject.io/docs/auth/kubernetes) and uses the pod’s service account as identity.

### 3. Use Vault API directly

Unlike the first two, this option has no external dependencies or software to install into the Kubernetes cluster. In order to use the API we must of course authenticate, for this we can use the Kubernetes auth method for the Kubernetes workers. If running Jenkins VM-based workers on EC2, Azure or GCE there are auth methods that can be used to authenticate against Vault. Since the option two also requires the Kubernetes auth method configuration, this seems like the minimalistic option that meets the requirements.

### Which option to go with?

There’s no one-size-fits-all solution or answer, however here’s a flow chart to help make the decision but it’s not exhaustive:

![Decision tree-integrate jenkins with vault](/blogs/secrets-handling-in-kubernetes-using-hashicorp-vault/decision-tree-integrate-jenkins-with-vault.png)

Based on this we decide to not use the Helm chart or Jenkins plugin and just setup the Kubernetes auth method for directly accessing the Vault API. Let’s see how that can be configured next.

## Demo: Using Kubernetes auth method in Jenkins pipelines

### Tools and versions

- Docker 20.10.17
- k3d v5.3.0
- kubectl v1.22.0
- Terraform v1.2.6
    - Vault provider 3.4.1
- jq-1.6

Where possible the version for images etc. is embedded into the code snippets.

### Setting up local Kubernetes cluster and Vault

A lot of tutorials install Vault inside the Kubernetes cluster, but we feel like this is a naive scenario that skips a lot of the setup necessary when integrating with an external Vault. To make the example more interesting let’s run an external Vault inside a Docker container, it’s important to note that the Vault and k3d cluster is running inside the same Docker network, thus the Vault can be reached using the name of the docker container. Neat!

Let’s start the Vault container and then create k3d cluster:

```bash
export DOCKER_NETWORK=k3d-vault-net
export VAULT_TOKEN=root

docker run -d \
  --cap-add=IPC_LOCK \
  -p 8200:8200 \
  --name=dev-vault \
  -e "VAULT_DEV_ROOT_TOKEN_ID=$VAULT_TOKEN" \
  --network ${DOCKER_NETWORK} \
  vault:1.11.0

export VAULT_ADDR=http://$(docker inspect dev-vault | jq -r ".[0].NetworkSettings.Networks.\"$DOCKER_NETWORK\".IPAddress"):8200

# write a secret into Vault KV engine for our demo
docker exec dev-vault /bin/sh -c "VAULT_TOKEN=$VAULT_TOKEN VAULT_ADDR=http://127.0.0.1:8200 vault kv put secret/jenkins/dev secretkey=supersecretvalue"

k3d cluster create jenkins \
  --network $DOCKER_NETWORK \
  --api-port $(ip route get 8.8.8.8 | awk '{print $7}'):16550
k3d kubeconfig get jenkins > jenkins-kubeconfig.yaml
export KUBECONFIG=$PWD/jenkins-kubeconfig.yaml
```

After the cluster is provisioned and nodes are in `Ready` state, create a service account which will be used by Vault to verify the service account tokens by calling the Kubernetes API:

```bash
kubectl create serviceaccount vault
kubectl create clusterrolebinding vault-reviewer-binding \
  --clusterrole=system:auth-delegator \
  --serviceaccount=default:vault

# set the variables needed to configure auth method
export VAULT_SA_NAME=$(kubectl get secrets --output=json | jq -r '.items[].metadata | select(.name|startswith("vault-token-")).name')
export TF_VAR_token_reviewer_jwt=$(kubectl get secret $VAULT_SA_NAME --output json | jq -r .data.token | base64 --decode)
export TF_VAR_kubernetes_ca_cert=$(kubectl config view --raw --minify --flatten --output='jsonpath={.clusters[].cluster.certificate-authority-data}')
export TF_VAR_kubernetes_host=$(kubectl config view --raw --minify --flatten --output='jsonpath={.clusters[].cluster.server}')
```

### Configuring Vault using Terraform

In order to configure Vault we will use Terraform instead of just running bunch of commands. We already set the necessary environment variables in above code to reach our Vault instance, next we can copy the following code to a `main.tf` file:

```hcl
terraform {
  required_providers {
    vault = {
      source = "hashicorp/vault"
      version = "3.4.1"
    }
  }
}

provider "vault" {
  # Configured with environment variables:
  # VAULT_ADDR
  # VAULT_TOKEN
}

variable "kubernetes_host" {
  type = string
  description = "URL for the Kubernetes API."
}

variable "kubernetes_ca_cert" {
  type = string
  description = "Base64 encoded CA certificate of the cluster."
}

variable "token_reviewer_jwt" {
  type = string
  description = "JWT token of the Vault Service Account."
}

resource "vault_auth_backend" "this" {
  type = "kubernetes"
}

resource "vault_kubernetes_auth_backend_config" "example" {
  backend                = vault_auth_backend.this.path
  kubernetes_host        = var.kubernetes_host
  kubernetes_ca_cert     = base64decode(var.kubernetes_ca_cert)
  token_reviewer_jwt     = var.token_reviewer_jwt
  issuer                 = "api"
  disable_iss_validation = "true" # k8s API checks it
}

resource "vault_policy" "jenkins-dev" {
  name = "jenkins-dev"

  policy = <<EOT
path "secret/data/jenkins/dev" {
  capabilities = ["read"]
}
EOT
}

resource "vault_kubernetes_auth_backend_role" "jenkins-dev" {
  backend                          = vault_auth_backend.this.path
  role_name                        = "jenkins-dev"
  bound_service_account_names      = ["jenkins-dev"]
  bound_service_account_namespaces = ["jenkins-dev"]
  token_ttl                        = 3600
  token_policies                   = ["jenkins-dev"]
}The heart of the configuration is in the `vault_kubernetes_auth_backend_role` resource: 
```

```bash
  bound_service_account_names      = ["jenkins-dev"]
  bound_service_account_namespaces = ["jenkins-dev"]
  token_ttl                        = 3600
  token_policies                   = ["jenkins-dev"]
```

This configuration is what glues the Kubernetes auth and Vault policy together for AuthN and AuthZ, we can also make the Vault token short-lived here as we likely only need it when starting a new Jenkins pipeline.

When done digesting the configuration, `terraform`  init and apply the configuration:

```bash
terraform init && terraform apply
```

### Installing Jenkins and example pipeline

Spin up Jenkins with the Helm chart:

```bash
helm install jenkins jenkins/jenkins --namespace jenkins-dev --create-namespace --version 4.1.14
```

Create service account in the namespace created by Helm as part of the Jenkins installation:

```bash
kubectl create serviceaccount jenkins-dev --namespace jenkins-dev
```

In order to access the Jenkins UI, follow the instructions printed from the Helm install:

```bash
1. Get your 'admin' user password by running:
  kubectl exec --namespace jenkins-dev -it svc/jenkins -c jenkins -- /bin/cat /run/secrets/additional/chart-admin-password && echo
2. Get the Jenkins URL to visit by running these commands in the same shell:
  echo http://127.0.0.1:8080
  kubectl --namespace jenkins-dev port-forward svc/jenkins 8080:8080
```

Here’s an example pipeline for reading the secrets in Vault using the Vault CLI:

```groovy
pipeline {
    agent {
        kubernetes {
            yaml '''
apiVersion: v1
kind: Pod
spec:
  serviceAccount: jenkins-dev
  containers:
  - name: build
    image: ubuntu
    command:
    - sleep
    args:
    - infinity
  - name: vault
	    image: hashicorp/vault:1.11.0
    env:
    - name: VAULT_ADDR
      value: "http://dev-vault:8200"
    command:
    - sleep
    args:
    - infinity
'''
            defaultContainer 'build'
        }
    }
    stages {
        stage('Main') {
            steps {
                container('vault') {
                    sh '''
                    SA_TOKEN=$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)
                    export VAULT_TOKEN=$(vault write -field=token auth/kubernetes/login role=jenkins-dev jwt=$SA_TOKEN)
                    vault kv get secret/jenkins/dev > secrets.txt
                    '''
                }
                container('build') {
                    sh 'cat secrets.txt'
                }
            }
        }
    }
}
```

Note the `serviceAccount` field and name of the role in the Vault CLI command, this is the configuration in Jenkins in order to access the Vault secrets. For ease of use and brevity we will use the Vault CLI instead of the API directly in the pipeline, but to drop the dependency on an additional container/binary the Vault API can be used directly with `curl` for example. 

## Summary/Conclusion

Looking at the pipeline, there’s actually nothing Jenkins specific to the solution presented here. We could run any pod in the `jenkins-dev` namespace with service account `jenkins-dev` to access the secrets that the policy in Vault allows. No hard coded and long-lived credentials laying around! No need for complex groovy syntax, Helm charts or Jenkins plugins. This keeps the pipeline flexible in case you (or the management) decide to ditch Jenkins and go for another tool to run your CI/CD pipelines.


This is a follow up blog to [Secrets handling in Kubernetes - A Jenkins story](/blog/secrets-handling-in-kubernetes-a-jenkins-story/), where we explore some ways of getting secrets into Jenkins which we deploy in Kubernetes, without the use of an external secrets manager.