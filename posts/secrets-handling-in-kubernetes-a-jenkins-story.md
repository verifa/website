---
type: Blog
title: Secrets handling in Kubernetes - A Jenkins story
subheading: In this blog we'll explore some ways of getting secrets into Jenkins which
  we deploy in Kubernetes, without the use of an external secrets manager.
authors:
- avijayan
tags:
- Kubernetes
- Jenkins
date: 2022-01-24
image: "/static/blog/2022-01-26/blog_k8s_jenkins_secrets.png"
featured: true

---
Passwords and keys are some of the most important tools used for authenticating and/or authorising applications. They provide access to sensitive systems, services, and information. Due to this, secrets management must account for and mitigate the risks to these secrets both in transit and at rest \[[1](https://www.beyondtrust.com/resources/glossary/secrets-management)\].

In this short write up we explore some ways of getting secrets into Jenkins which we deploy in Kubernetes. Jenkins typically is a central part of infrastructure and usually requires access to critical components like source control, testing tools, artifact management, etc. It's not Jenkins directly, but tools or tasks triggered from Jenkins which needs it. Therefore a natural choice is to make use of Jenkins secrets. However, putting all these secrets into Jenkins means a lot of focus should be placed on Jenkins secret management and Jenkins was not designed to be a secrets manager. Therefore, it makes a lot of sense to store this wide variety of secrets into “real” secrets managers like HashiCorp Vault to minimize risk. For the purposes of this post we consider options for using secrets with Jenkins _without_ an external secrets manager [(like HashiCorp Vault)](/blog/secrets-handling-in-kubernetes-using-hashicorp-vault/).

***

## Secrets through Kubernetes

If Jenkins is running inside a Kubernetes cluster, it is a natural approach to use Kubernetes secrets \[[2](https://kubernetes.io/docs/concepts/configuration/secret/)\], a built-in API object to store sensitive data. When running a Kubernetes pod (e.g. containing our Jenkins controller) secrets can be used either by mounting them as plaintext files or as environment variables.

Here we will discuss the situation with Jenkins as a workload for a Kubernetes cluster and how secrets are passed to it. There are multiple ways of passing secrets to Jenkins from Kubernetes natively \[[3](https://github.com/jenkinsci/configuration-as-code-plugin/blob/master/docs/features/secrets.adoc)\]. Let's take a look at them here:

### Environment variables

This is the simplest and thereby the most common approach that exposes secrets as environment variables into the container. The Jenkins configurations inside the pod can use those variables without any further processing. As Jenkins is designed to use its environment variables as variable substitutions within Jenkins, it makes this setup easy to configure. However, Jenkins does not know they are secrets and hence handles them as normal environment variables while logging, printing etc. making this approach a less secure one.

### Kubernetes/Docker secrets

This approach is slightly more secure than the previous one. The secrets are mounted as files to `/run/secrets/`, and then the filename can be used as the KEY. This is not a Jenkins or a Kubernetes secrets per se, but a Docker feature.\[[4](https://docs.docker.com/engine/swarm/secrets/)\]

### Using Jenkins plugin - kubernetes-credentials-provider-plugin

The Kubernetes Credentials Provider plugin is a Jenkins plugin to enable the retrieval of credentials directly from Kubernetes. This allows credentials to be bound to environment variables for use from build steps directly. \[[5](https://plugins.jenkins.io/credentials-binding/)\]

### Passing credentials as Jenkins secrets

This approach provides a way to natively encrypt the secrets exposed by Kubernetes to Jenkins. All of the above approaches do not provide ways to encrypt the secret. So, this approach can be used in combination with the above while implementing secret handling in a more secure way. \[[6](https://www.jenkins.io/doc/developer/security/secrets/)\].

Now that we have learnt enough about this, it's time to dive-in for a hands-on demonstration.

In the following steps, we will take a deeper look to see how one could store a credential with values `test/test123` with the above approaches. Also, by executing some commands we will also demonstrate how easy/difficult would be for an attacker to get hold of that credential.

For this, first the secret should be created in the Kubernetes cluster and then passed to the workload (Jenkins) in various ways. We will be doing this with a Helm chart \[5\].

## Demo: Secrets handing without Jenkins secrets

The above approaches of Environment variables, Kubernetes/Docker secrets and Kubernetes credentials provider plugin are easier to implement but they store passwords in plaintext. Let's see how it actually looks in a Kubernetes cluster with a practical example.

1. Let's assume that you have a Kubernetes cluster available and Helm v3 installed. Here we have a single node cluster from GKE.

```bash
$ kubectl get nodes
NAME                              STATUS   ROLES    AGE   VERSION
gke-jenkins-linux-468bbc97-5lxw   Ready    <none>   17m   v1.20.9-gke.1001
$ helm ls
NAME NAMESPACE REVISION UPDATED STATUS CHART APP VERSION
$ helm version
version.BuildInfo{Version:"v3.5.2", GitCommit:"167aac70832d3a384f65f9745335e9fb40169dc2", GitTreeState:"dirty", GoVersion:"go1.15.7"}
$
```

2. Let's create a Kubernetes secret with name: `test-user-cred` with `test/test123` values.

```bash
$ kubectl create secret generic test-user-cred --from-literal='username=test' --from-literal='password=test123'
secret/test-user-cred created
$
```

3. Now let's mount them as environment variables to the Jenkins master. For this, we need a Helm chart values file with the following content (according to the Jenkins Helm chart).

```bash
$ cat jenkins-values-env.yaml
controller:
  containerEnv:
  - name: username
    valueFrom:
    secretKeyRef:
      name: test-user-cred
      key: username
  - name: password
    valueFrom:
    secretKeyRef:
      name: test-user-cred
      key: password
```

4. Let's install the chart and check the values with `kubectl exec`.

```bash
$ helm install jenkins jenkins/jenkins -f jenkins-values-env.yaml
NAME: jenkins
LAST DEPLOYED: Fri Oct  1 10:07:34 2021
NAMESPACE: default
STATUS: deployed
REVISION: 1
..<lots of text>..

$ kubectl exec -it po/jenkins-0 -- bash -c "env |grep 'user\\|pass'"
Defaulting container name to jenkins.
Use 'kubectl describe pod/jenkins-0 -n default' to see all of the containers in this pod.
username=test
password=test123
$
```

5. This is the same case with Docker/Kubernetes secrets. Let's create a new revision with new values file mounting them as secrets. Let's do this with a new values file.

```bash
$ cat jenkins-values-mount.yaml
controller:
  additionalExistingSecrets:
  - name: test-user-cred
  keyName: username
  - name: test-user-cred
  keyName: password

$ helm upgrade jenkins jenkins/jenkins -f jenkins-values-mount.yaml
NAME: jenkins
LAST DEPLOYED: Fri Oct  1 10:09:34 2021
NAMESPACE: default
STATUS: deployed
REVISION: 2
..<lots of text>..

$ kubectl exec -it pod/jenkins-0 -- bash -c "cat /run/secrets/test-user-cred-username; echo; cat /run/secrets/test-user-cred-password"
Defaulting container name to jenkins.
Use 'kubectl describe pod/jenkins-0 -n default' to see all of the containers in this pod.
test
test123
$
```

## Demo: Secrets handling with Jenkins secrets

The approach of passing credentials with Jenkins secrets is harder to implement but avoids storing passwords as plaintext, instead using the built-in Hudson encryption which is natively supported inside Jenkins.\[[7](https://www.jenkins.io/doc/developer/security/secrets/)\]

### How it works

The idea is to encrypt the secrets first with Jenkins domain credentials. This way, the secrets are not in plaintext. Jenkins uses AES to encrypt and protect secrets, credentials, and their respective encryption keys. This is unique to every jenkins instance as it is generated on first start. Due to this, any secrets encrypted with this work out-of-box inside the same Jenkins server and the encrypted secrets cannot be decrypted with a different Jenkins server.

![Jenkins and Kubernetes secret encryption process](/static/blog/2022-01-26/secrets_jenkins_kubernetes_diagram.png)

From the above picture, any secret is first encrypted with Jenkins Hudson secret and then imported as Kubernetes secret. This way, they are not in plaintext anywhere.

### Using Hudson secrets in Jenkins with Kubernetes

In order to use our secrets from `test-user-cred`, we need to first encrypt with hudson secret. Note that, we need to have a Jenkins instance first running to get this working. Also, if you do not know where to find jenkins cli jar in the host, please refer to this \[8\]

1. Create hudson secret for user test with our existing Jenkins instance

```bash
$ kubectl exec -it po/jenkins-0 -c jenkins -- bash -c \\
    'echo \\'println(hudson.util.Secret.fromString("test").getEncryptedValue())\\' \\
    | java -jar /var/jenkins_home/war/WEB-INF/lib/cli-2.303.1.jar \\
    s http://0.0.0.0:8080 \\
    auth Admin:$(cat /run/secrets/chart-admin-password) \\
    groovy = '
{AQAAABAAAAAQGiN0B2weIsYfpg0LqBbM7WSBn9+zSBcH4OXyYpaVVig=}
$
```

2. Create hudson secret for password test123 with our existing Jenkins instance

```bash
$ kubectl exec -it po/jenkins-0 -c jenkins -- bash -c \\
    'echo \\'println(hudson.util.Secret.fromString("test123").getEncryptedValue())\\' \\
    | java -jar /var/jenkins_home/war/WEB-INF/lib/cli-2.303.1.jar \\
    s http://0.0.0.0:8080 \\
    auth Admin:$(cat /run/secrets/chart-admin-password) \\
    groovy = '
{AQAAABAAAAAQ16q66jh91Uc0hn502+qgdFpuWTZPBP99JDQU7ZqZpxg=}
$
```

Below we repeat the same procedure with Kubernetes secrets using Helm. The values file does not have to be changed.

3. Let's overwrite our Kubernetes secret with new values.

```bash
$ kubectl delete secret/test-user-cred
secret "test-user-cred" deleted
$
$ kubectl create secret generic test-user-cred --from-literal='user={AQAAABAAAAAQGiN0B2weIsYfpg0LqBbM7WSBn9+zSBcH4OXyYpaVVig=}' --from-literal='pass={AQAAABAAAAAQ16q66jh91Uc0hn502+qgdFpuWTZPBP99JDQU7ZqZpxg=}'
secret/test-user-cred created
$
```

4. Now upgrade the helm chart with the old values files and check the secrets from the disk and environment.

```bash
$ helm upgrade jenkins jenkins/jenkins -f ./jenkins-values-mount.yaml
Release "jenkins" has been upgraded. Happy Helming!
NAME: jenkins
LAST DEPLOYED: Fri Oct  1 11:08:45 2021
NAMESPACE: default
STATUS: deployed
REVISION: 3
..lots of text here...
$
$ kubectl exec -it po/jenkins-0 -- bash -c "cat /run/secrets/test-user-cred-username; echo; cat /run/secrets/test-user-cred-password"
Defaulting container name to jenkins.
Use 'kubectl describe pod/jenkins-0 -n default' to see all of the containers in this pod.
{AQAAABAAAAAQGiN0B2weIsYfpg0LqBbM7WSBn9+zSBcH4OXyYpaVVig=}
{AQAAABAAAAAQ16q66jh91Uc0hn502+qgdFpuWTZPBP99JDQU7ZqZpxg=}
$
$ helm upgrade jenkins jenkins/jenkins -f ./jenkins-values-env.yaml
..lots of text here...
$
$ kubectl exec -it po/jenkins-0 -- bash -c "env|grep 'user\\|pass'"
Defaulting container name to jenkins.
Use 'kubectl describe pod/jenkins-0 -n default' to see all of the containers in this pod.
username={AQAAABAAAAAQGiN0B2weIsYfpg0LqBbM7WSBn9+zSBcH4OXyYpaVVig=}
password={AQAAABAAAAAQ16q66jh91Uc0hn502+qgdFpuWTZPBP99JDQU7ZqZpxg=}
$
```

## Conclusion

We have now reached the end of an interesting round of comparisons of various ways available to share secrets from a Kubernetes cluster to our Jenkins server. As we all understand, one should take good care of secrets and add ways to protect as much as possible to minimise risks even if some parts of a system are compromised - security is a multi-layer practice after all. Not allowing passwords in plaintext at rest is a good rule of thumb. Here we explored, just by taking an extra step, that it was possible to use native encryption of the passwords stored in the disk or environment.

The scope of this blog is to explore more secure ways to present passwords for Jenkins inside a Kubernetes cluster without the use of external tools. There are secret managers dedicated for these purposes such as HashiCorp Vault or cloud specific such as AWS, Google cloud secrets manager, Azure key vault etc. One of the benefits of using external secret managers is that of decoupling sensitive entities outside of the application.

As a follow up to this blog, we went on to explore several approaches for getting secrets from HashiCorp Vault into Jenkins running in Kubernetes - [Secrets handling in Kubernetes using HashiCorp Vault](/blog/secrets-handling-in-kubernetes-using-hashicorp-vault/).

## Versions of tools

GKE kubernetes: v1.20.9-gke.1001

kubectl: 1.20.2

helm: 3.5.2

Jenkins helm chart: jenkins-3.5.19/2.303.1

## References

\[1\] Secrets in general: [https://www.beyondtrust.com/resources/glossary/secrets-management](https://www.beyondtrust.com/resources/glossary/secrets-management "https://www.beyondtrust.com/resources/glossary/secrets-management")

\[2\] Kubernetes secrets: [https://kubernetes.io/docs/concepts/configuration/secret/](https://kubernetes.io/docs/concepts/configuration/secret/ "https://kubernetes.io/docs/concepts/configuration/secret/")

\[3\] Jenkins configuration as code secrets: [https://github.com/jenkinsci/configuration-as-code-plugin/blob/master/docs/features/secrets.adoc](https://github.com/jenkinsci/configuration-as-code-plugin/blob/master/docs/features/secrets.adoc "https://github.com/jenkinsci/configuration-as-code-plugin/blob/master/docs/features/secrets.adoc")

\[4\] Docker secrets: [https://docs.docker.com/engine/swarm/secrets/](https://docs.docker.com/engine/swarm/secrets/ "https://docs.docker.com/engine/swarm/secrets/")

\[5\] Jenkins plugin - Kubernetes credentials provider: [https://plugins.jenkins.io/credentials-binding/](https://plugins.jenkins.io/credentials-binding/ "https://plugins.jenkins.io/credentials-binding/")

\[6\] Jenkins Hudson secret: [https://www.jenkins.io/doc/developer/security/secrets/](https://www.jenkins.io/doc/developer/security/secrets/ "https://www.jenkins.io/doc/developer/security/secrets/")

\[7\] Encryption of Jenkins secrets:

[https://www.jenkins.io/doc/developer/security/secrets/](https://www.jenkins.io/doc/developer/security/secrets/ "https://www.jenkins.io/doc/developer/security/secrets/")
