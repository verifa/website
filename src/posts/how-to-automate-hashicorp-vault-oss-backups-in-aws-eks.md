---
type: Blog
title: How to automate HashiCorp Vault OSS backups in AWS EKS
subheading: In this post we will walk through an implementation using a Kubernetes CronJob to take daily snapshots and store them in an AWS S3 bucket for safe keeping.
authors:
- jlarfors
tags:
- HashiCorp
- Vault
- AWS
- Kubernetes
date: 2022-04-26
image: "/blogs/automate-hashicorp-vault-oss-backups-aws-eks.png"
featured: true

---

**[HashiCorp Vault](https://www.vaultproject.io/) is an API-driven tool for storing and retrieving static and dynamic secrets. Vault can be deployed in a Kubernetes cluster using the [official Helm chart](‣). The recommended storage for Vault in Kubernetes is the [integrated raft storage]([https://www.vaultproject.io/docs/configuration/storage/raft](https://www.vaultproject.io/docs/configuration/storage/raft)) and frequent snapshots of Vault should be taken and stored, making it possible to restore Vault in case of data loss.**

In this post we will walk through an implementation using a [Kubernetes CronJob]([https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/](https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/)) to take daily snapshots and store them in an AWS S3 bucket for safe keeping. Note that Vault Enterprise makes backups a [native feature]([https://www.vaultproject.io/docs/enterprise/automated-integrated-storage-snapshots](https://www.vaultproject.io/docs/enterprise/automated-integrated-storage-snapshots)) that should be used if you have that version.

## Write a Kubernetes CronJob

Let’s start with the CronJob and go backwards from there, because in order for the CronJob to work we will need to authenticate with both HashiCorp Vault and an AWS S3 bucket.

```yaml
---
apiVersion: batch/v1
kind: CronJob
	metadata:
  name: vault-snapshot-cronjob
spec:
	# Set your desired cron schedule
  schedule: "0 2 * * 1-5"
  successfulJobsHistoryLimit: 10
  failedJobsHistoryLimit: 3
  jobTemplate:
    spec:
      template:
        spec:
					# Use a ServiceAccount that we will create next (keep reading!)
          serviceAccountName: vault-snapshot
          volumes:
						# Create an empty drive to share the snapshot across containers
            - name: share
              emptyDir: {}
          initContainers:
						# Run an init container that creates the the snapshot of Vault
            - name: vault-snapshot
							# Choose an appropriate Vault version (e.g. same as your Vault setup)
              image: vault:1.9.4
              command: ["/bin/sh", "-c"]
              args:
								# 1. Get the ServiceAccount token which we will use to authenticate against Vault
								# 2. Login to Vault using the SA token at the endpoint where the Kubernetes auth engine
								#    has been enabled
								# 3. Use the Vault CLI to store a snapshot in our empty volume
                - |
                  SA_TOKEN=$(cat /var/run/secrets/kubernetes.io/serviceaccount/token);
                  export VAULT_TOKEN=$(vault write -field=token auth/kubernetes/login jwt=$SA_TOKEN role=vault-snapshot);
                  vault operator raft snapshot save /share/vault.snap;
              env:
								# Set the Vault address using the Kubernetes service name
                - name: VAULT_ADDR
                  value: http://vault.vault.svc.cluster.local:8200
              volumeMounts:
                - mountPath: /share
                  name: share
          containers:
						# Run a container with the AWS CLI and copy the snapshot to our S3 bucket
            - name: aws-s3-backup
              image: amazon/aws-cli:2.2.14
              command:
                - /bin/sh
              args:
                - -ec
								# Copy the snapshot file to an S3 bucket called hashicorp-vault-snapshots
                - aws s3 cp /share/vault.snap s3://hashicorp-vault-snapshots/vault_$(date +"%Y%m%d_%H%M%S").snap;
              volumeMounts:
                - mountPath: /share
                  name: share
          restartPolicy: OnFailure
```

Writing the CronJob is probably the easiest part. Now we need to ensure that the two containers we are running (`vault-snapshot` and `aws-s3-backup`) can authenticate with Vault and AWS. For this, we will rely on a ServiceAccount.

## Authentication with Vault and AWS

### Define a Kubernetes ServiceAccount

Let’s define a Kubernetes ServiceAccount called `vault-snapshot` that we referenced in the above CronJob.

```yaml
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: vault-snapshot
  annotations:
		# Assume the AWS role hashicorp-vault-snapshotter
    eks.amazonaws.com/role-arn: arn:aws:iam::<ACCOUNT_ID>:role/hashicorp-vault-snapshotter
```

Notice how we add the annotation to assume the AWS role `hashicorp-vault-snapshotter`. For details on assuming AWS IAM roles from EKS, please read our [blog post on that topic](./how-to-assume-an-aws-iam-role-from-a-service-account-in-eks-with-terraform).

### Define an AWS IAM Role

Let’s define the AWS IAM role `hashicorp-vault-snapshotter` and make the `vault-snapshot`ServiceAccount a trusted entity that can assume that role.

```yaml
locals {
  vault_cluster = "<eks-cluster-name>"
  k8s_service_account_name      = "vault-snapshot"
  k8s_service_account_namespace = "vault"
  vault_cluster_oidc_issuer_url = trimprefix(data.aws_eks_cluster.vault_cluster.identity[0].oidc[0].issuer, "https://")
}

# 
# Might as well create the S3 bucket whilst we are at it...
# 
resource "aws_s3_bucket" "snapshots" {
  bucket = "hashicorp-vault-snapshots"
}

#
# Get the caller identity so that we can get the AWS Account ID
#
data "aws_caller_identity" "current" {}

# 
# Get the cluster that vault is running in
# 
data "aws_eks_cluster" "vault_cluster" {
  name = local.vault_cluster
}

#
# Create the IAM role that will be assumed by the service account
#
resource "aws_iam_role" "snapshot" {
  name               = "hashicorp-vault-snapshotter"
  assume_role_policy = data.aws_iam_policy_document.snapshot.json

  inline_policy {
    name = "hashicorp-vault-snapshot"
    policy = jsonencode({
      Version = "2012-10-17"
      Statement = [
        {
          Effect = "Allow",
          Action = [
            "s3:PutObject",
            "s3:GetObject",
          ],
					# Refer to the S3 bucket we created along the way
          Resource = ["${aws_s3_bucket.snapshots.arn}/*"]
        }
      ]
    }) 
  }
}

#
# Create IAM policy allowing the k8s service account to assume the IAM role
#
data "aws_iam_policy_document" "snapshot" {
  statement {
    actions = ["sts:AssumeRoleWithWebIdentity"]

    principals {
      type = "Federated"
      identifiers = [
        "arn:aws:iam::${data.aws_caller_identity.current.account_id}:oidc-provider/${local.vault_cluster_oidc_issuer_url}"
      ]
    }

    # Limit the scope so that only our desired service account can assume this role
    condition {
      test     = "StringEquals"
      variable = "${local.vault_cluster_oidc_issuer_url}:sub"
      values = [
        "system:serviceaccount:${local.k8s_service_account_namespace}:${local.k8s_service_account_name}"
      ]
    }
  }
}
```

### Configure Vault Kubernetes Auth Engine

So far we have a Kubernetes ServiceAccount which can assume an AWS IAM role which has access to S3. What’s missing is the authentication with Vault.

You could use the [Vault AWS Auth Engine]([https://www.vaultproject.io/docs/auth/aws](https://www.vaultproject.io/docs/auth/aws)) and use the same AWS role for that. However, in our case we use Vault to provide secrets to Kubernetes workloads and therefore already have multiple EKS clusters authenticated with Vault so it made sense to reuse that logic, and that’s what we will show below.

This is quite an involved process, and could make it’s own blog post, but a summary of what we will do is:

1. Create a Kubernetes ServiceAccount with the ClusterRole `system:auth-delegator`
    1. This gives Vault the ability to authenticate and authorize Kubernetes ServiceAccount tokens that are used to authenticate with Vault
        1. Remember our initContainer passes the Kubernetes ServiceAccount token to Vault in exchange for an ordinary Vault Token
    2. Read more here: [https://kubernetes.io/docs/reference/access-authn-authz/rbac/#other-component-roles](https://kubernetes.io/docs/reference/access-authn-authz/rbac/#other-component-roles)
2. Enable a Vault Auth Engine of type `kubernetes` at the mount path `kubernetes`
    1. `kubernetes` is the default mount path, so you probably want to use something like `kube-<eks-cluster-name>` so that you can authenticate with multiple clusters
3. Configure the Kubernetes auth engine using the Kubernetes ServiceAccount we created in step 1

```yaml
locals {
  namespace = "vault-client"
}

# 
# Create kubernetes service account that vault can use to authenticate requests
# from the cluster
# 
resource "kubernetes_service_account" "this" {
  metadata {
    name      = "vault-auth"
    namespace = local.namespace
  }
  automount_service_account_token = "true"
}

# 
# Give the service account permissions to authenticate other service accounts
# 
resource "kubernetes_cluster_role_binding" "this" {
  metadata {
    name = "vault-token-auth"
  }
  role_ref {
    api_group = "rbac.authorization.k8s.io"
    kind      = "ClusterRole"
    name      = "system:auth-delegator"
  }
  subject {
    kind      = "ServiceAccount"
    name      = kubernetes_service_account.this.metadata[0].name
    namespace = local.namespace
  }
}

# 
# Get the secret created for the service account
# 
data "kubernetes_secret" "this" {
  metadata {
    name      = kubernetes_service_account.this.default_secret_name
    namespace = local.namespace
  }
}

# 
# Create the vault auth backend
# 
resource "vault_auth_backend" "this" {
  type = "kubernetes"
	# Make this something else for multiple clusters
  path = "kubernetes"
}

# 
# Configure the backend to use the service account we created, so that vault
# can verify requests made to this backend
# 
resource "vault_kubernetes_auth_backend_config" "this" {
  backend                = vault_auth_backend.this.path
	# Get the EKS endpoint from somewhere, like a `aws_eks_cluster` data block
  kubernetes_host        = var.cluster.endpoint
  kubernetes_ca_cert     = data.kubernetes_secret.auth.data["ca.crt"]
  token_reviewer_jwt     = data.kubernetes_secret.auth.data["token"]
  issuer                 = "api"
  disable_iss_validation = "true"
}
```

### Create a Vault Kubernetes Role

Now that we have a Kubernetes auth engine mounted and configured, we need to create a role in Vault so that a ServiceAccount in our Kubernetes cluster can actually do something!

We need this Vault policy. Let’s store it in a file such as `policies/sys-snapshot-read.hcl`

```yaml
path "sys/storage/raft/snapshot" {
  capabilities = ["read"]
}
```

And now the Terraform code to create the Vault role.

```yaml
#
# Create a Vault policy based of a template
#
resource "vault_policy" "this" {
  name = "vault-snapshot"
  policy = file("policies/sys-snapshot-read.hcl")
}

#
# Create a Vault role with our snapshot policy, that is bound
# to the vault-snapshot Kubernetes ServiceAccount in the
# vault-snapshot namespace.
#
# NOTE: Make sure you use the correct namespace and serviceaccount!
#
resource "vault_kubernetes_auth_backend_role" "this" {
  depends_on = [vault_policy.this]

  backend                          = vault_auth_backend.this.path
  role_name                        = "vault-snapshot"
  bound_service_account_names      = ["vault-snapshot"]
  bound_service_account_namespaces = ["vault-snapshot"]
  token_policies                   = ["vault-snapshot"]
  token_ttl                        = 3600
  audience                         = null
}
```

## Testing our Vault backup process

The CronJob is currently set to run daily, and we probably want to test this without waiting for the CronJob each time... I would be amazed if you get this working first time - if so, **you owe me at least one beer!**

```yaml
# Let's do our work in a separate namespace
kubectl create namespace vault-snapshot

# Set active namespace
kubens vault-snapshot

# Apply the CronJob from earlier if you haven't already
kubectl apply -f vault-snapshot-cronjob.yaml

# Create a Job from the CronJob to test that it works
kubectl create job --from=cronjob/vault-snapshot-cronjob test-1
# Do your thing and describe/debug the job...
kubectl describe job.batch/test-1
# Check logs from vault initContainer
kubectl logs test-1-<hash> -c vault-snapshot
# Check logs from aws container
kubectl logs test-1-<hash>

# Probably something failed, so repeat the above with test-2 :)
# Remember to cleanup your jobs afterwards.
```

Once you get this working you should have a snapshot stored in your AWS S3 bucket. That’s great, so how do you check that this can be restored?

## Restoring Vault Snapshot

We found the quickest and easiest way to test a restore was to spin up a dev instance of Vault in EKS without persistent storage, initialise the fresh vault instance and restore the snapshot.

```yaml
# First download the snapshot from S3, e.g. via the AWS Console (UI)
ls vault_20220325_082239.snap

# Create another namespace for this. Make sure this Vault instance will
# also have access to your AWS KMS (or however you auto-unseal Vault).
# And if you don't currently auto-unseal Vault in AWS EKS... leave a
# comment and I will help make your life easier :)
kubectl create namespace vault-dev

# Set active namespace
kubens vault-dev

# Deploy a dev instance of Vault without persistent storage, e.g.
helm install vault hashicorp/vault -f dev-values.yaml

# Check the Vault pod (it should not have started because Vault needs
# to be initialised)
kubectl get pods

# Intialise Vault
kubectl exec -n vault-dev -ti vault-dev-0 -- vault operator init
# Check the log... What we care about is the Root token

# Next let's setup port-forwarding so that we can access our dev instance
# without any Ingresses and extra hassle
kubectl port-forward svc/vault 8200:8200

# Setup our Vault variables
export VAULT_ADDR=http://localhost:8200
export VAULT_TOKEN=<root-token> # Root token from init command above

# Restore the snapshot
vault operator raft snapshot restore vault_20220325_082239.snap

# Browse to http://localhost:8200 or use your Terminal to verify that
# Vault has restored to the point you'd expect.
```

## Conclusion

This post has gone through setting up automated backups of HashiCorp Vault OSS running on AWS EKS using a Kubernetes CronJob, and storing the snapshots in an S3 bucket. There’s a lot of pieces to the puzzle, and hopefully this post has given some insight into how it can be setup in a secure way following The Principle of Least Privilege.

If you have any questions, feedback or want help with your Vault setup please leave us a comment!