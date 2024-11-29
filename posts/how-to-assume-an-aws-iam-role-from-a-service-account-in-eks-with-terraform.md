---
type: Blog
title: How to assume an AWS IAM role from a Service Account in EKS with Terraform
subheading: In this blog we'll look at how to assume an AWS IAM role from a Service
  Account that exists within an AWS Elastic Kubernetes Service (EKS) cluster. All
  provisioned with Terraform.
authors:
- jlarfors
tags:
- HashiCorp
- AWS
- Kubernetes
date: 2021-12-07
image: "/static/blog/2021-12-08/blog-eks_aws_iam.png"
featured: true

---

**When working with AWS Elastic Kubernetes Service (EKS) clusters, your pods will likely want to interact with other AWS services and possibly other EKS clusters. In a recent project we were setting up** [**ArgoCD**](https://argo-cd.readthedocs.io/en/stable/) **with multiple EKS clusters and our goal was to use Kubernetes Service Accounts to assume an AWS IAM role to authenticate with other EKS clusters. This led to some learning and discovering that we'd like to share with you.**

***

When running workloads in EKS, the running pods will operate under a service account which allows us to enforce [RBAC within a Kubernetes cluster](https://kubernetes.io/docs/reference/access-authn-authz/rbac/). Well, we are not going to talk more about that in this post, we want to talk about how we can do things _outside_ of our cluster and interact with other AWS services. The [AWS documentation](https://docs.aws.amazon.com/eks/latest/userguide/iam-roles-for-service-accounts.html) for this is fairly good if you want a reference point. There is also a [workshop](https://www.eksworkshop.com/beginner/110_irsa/preparation/) on this topic which might be useful to run through.

In our case, we were trying to communicate across EKS clusters to allow ArgoCD to manage multiple clusters and there is a pretty mammoth [GitHub issue](https://github.com/argoproj/argo-cd/issues/2347) with people struggling (and succeeding!) with this. That GitHub issue partly inspired this blog post - if it was an easy topic people would not struggle and a blog would not be necessary ;)

## The Plan

The simple breakdown of what we need:

1. An EKS cluster with an [IAM OIDC provider](https://docs.aws.amazon.com/eks/latest/userguide/enable-iam-roles-for-service-accounts.html)
2. A Kubernetes Service Account in the EKS cluster
3. An AWS IAM role which we are going to _assume_ (meaning we can do whatever that role is able to do)
4. An AWS IAM role policy that allows our Service Account (2.) to assume our AWS IAM role (3.)

We won't bore you with creating an EKS cluster and an IAM OIDC provider... Pick your poison for how you want to do this... We personally use Terraform and the [awesome EKS module](https://registry.terraform.io/modules/terraform-aws-modules/eks/aws/latest) that has a convenient input `enable_irsa` which creates the OIDC provider for us.

## Basic Deployment with Terraform

Before we create the Service Account and the IAM role we need to define the names of these as there's a bit of a cyclic dependency - the Service Account needs to know the role ARN, and the role policy needs to know the Service Account name and namespace (if we want to limit scope, which we do!).

### Locals

So let's define some locals to keep things simple and DRY.

```hcl
# locals.tf

locals {
  k8s_service_account_name      = "iam-role-test"
  k8s_service_account_namespace = "default"

  # Get the EKS OIDC Issuer without https:// prefix
  eks_oidc_issuer = trimprefix(data.aws_eks_cluster.eks.identity[0].oidc[0].issuer, "https://")
}
```

### IAM

And let's define the Terraform code that creates the IAM role with a policy allowing the service account to assume that role.

```hcl
# iam.tf

#
# Get the caller identity so that we can get the AWS Account ID
#
data "aws_caller_identity" "current" {}

#
# Get the EKS cluster we want to target
#
data "aws_eks_cluster" "eks" {
  name = "<cluster-name>"
}

#
# Create the IAM role that will be assumed by the service account
#
resource "aws_iam_role" "iam_role_test" {
  name               = "iam-role-test"
  assume_role_policy = data.aws_iam_policy_document.iam_role_test.json
}

#
# Create IAM policy allowing the k8s service account to assume the IAM role
#
data "aws_iam_policy_document" "iam_role_test" {
  statement {
    actions = ["sts:AssumeRoleWithWebIdentity"]

    principals {
      type = "Federated"
      identifiers = [
        "arn:aws:iam::${data.aws_caller_identity.current.account_id}:oidc-provider/${local.eks_oidc_issuer}"
      ]
    }

    # Limit the scope so that only our desired service account can assume this role
    condition {
      test     = "StringEquals"
      variable = "${local.eks_oidc_issuer}:sub"
      values = [
        "system:serviceaccount:${local.k8s_service_account_namespace}:${local.k8s_service_account_name}"
      ]
    }
  }
}
```

### Service Account and Pod

Then we need to create some Kubernetes resources. When working with Terraform it can make a lot of sense to use the [Terraform Kubernetes Provider](https://registry.terraform.io/providers/hashicorp/kubernetes/latest) to apply our Kubernetes resources, especially as Terraform knows the ARN of the role and we can reuse our locals. However, if you don't want yet another provider dependency in Terraform you can easily do this with vanilla Kubernetes.

#### Terraform Kubernetes

NOTE: you will need to configure the Kubernetes provider if you want to do this via Terraform

```hcl
# kubernetes.tf

#
# Create the Kubernetes service account which will assume the AWS IAM role
#
resource "kubernetes_service_account" "iam_role_test" {
  metadata {
    name      = local.k8s_service_account_name
    namespace = local.k8s_service_account_namespace
    annotations = {
      # This annotation is needed to tell the service account which IAM role it
      # should assume
      "eks.amazonaws.com/role-arn" = aws_iam_role.iam_role_test.arn
    }
  }
}

#
# Deploy Kubernetes Pod with the Service Account that can assume an AWS IAM role
#
resource "kubernetes_pod" "iam_role_test" {
  metadata {
    name      = "iam-role-test"
    namespace = local.k8s_service_account_namespace
  }

  spec {
    service_account_name = local.k8s_service_account_name
    container {
      name  = "iam-role-test"
      image = "amazon/aws-cli:latest"
      # Sleep so that the container stays alive
      # #continuous-sleeping
      command = ["/bin/bash", "-c", "--"]
      args    = ["while true; do sleep 5; done;"]
    }
  }
}
```

#### Vanilla Kubernetes

And now the same as above with vanilla Kubernetes YAML.

```yaml
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: iam-role-test
  namespace: default
  annotations:
    # TODO: replace ACCOUNT_ID with your account id
    eks.amazonaws.com/role-arn: arn:aws:iam::<ACCOUNT_ID>:role/iam-role-test
---
apiVersion: v1
kind: Pod
metadata:
  name: iam-role-test
  namespace: default
spec:
  serviceAccountName: iam-role-test
  containers:
    - name: iam-role-test
      image: amazon/aws-cli:latest
      # Sleep so that the container stays alive
      # #continuous-sleeping
      command: ["/bin/bash", "-c", "--"]
      args: ["while true; do sleep 5; done;"]
```

## Verify the setup

We can describe the pod (i.e. `kubectl describe pod iam-role-test`) and check the volumes, mounts and environment variables attached to the pod, but seeing as we way launched a pod with the AWS CLI, let's just get in there and check! Exec into the running container and execute the `aws` CLI:

```bash
# Exec into the running pod
kubectl exec -ti iam-role-test -- /bin/bash

# Check the AWS Security Token Service identity
bash-4.2# aws sts get-caller-identity
{
    "UserId": "AROA46FON4H773JH4MPJD:botocore-session-1637837863",
    "Account": "123456789101",
    "Arn": "arn:aws:sts::123456789101:assumed-role/iam-role-test/botocore-session-1637837863"
}

# Check the AWS environment variables
bash-4.2# env | grep "AWS_"
AWS_ROLE_ARN=arn:aws:iam::<ACCOUNT_ID>:role/iam-role-test
AWS_WEB_IDENTITY_TOKEN_FILE=/var/run/secrets/eks.amazonaws.com/serviceaccount/token
AWS_DEFAULT_REGION=eu-west-1
AWS_REGION=eu-west-1
```

As you can see, the AWS Service Token Service (STS) confirms that we have successfully assumed the role we wanted to! And if we check our environment variables we can see that these have been in injected when we started the pod, and the `AWS_WEB_IDENTITY_TOKEN_FILE` file is the part that is sensitive and mounted when we run the container.

## Alternative Approaches

If we remove the service account from the pod and use the default service account (which exists per namespace), we can see who AWS STS thinks we are:

```bash
# Exec into the running pod
kubectl exec -ti iam-role-test -- /bin/bash

# Check the AWS Security Token Service identity
bash-4.2# aws sts get-caller-identity
{
    "UserId": "AROA46FON4H72Q3SPL6SC:i-0d0aff479cf2e2405",
    "Account": "123456789101",
    "Arn": "arn:aws:sts::<ACCOUNT_ID>:assumed-role/<cluster-name>XXXXXXXXX/i-<node-instance-id>"
}


# Check the AWS environment variables
bash-4.2# env | grep "AWS_"
# ... it's empty!
```

Having created the cluster with the EKS Terraform module, it has created a role for our autoscaling node group... and what we could do is allow this role to assume another role which grants us the access that we might need...

However, I find managing Service Accounts in Kubernetes much easier than assigning roles to node groups and it seems this is the recommended approach based on searches online. However I thought it worth mentioning here anyway.

## Next Steps

Now that you can run a Pod with a service account that can assume an IAM role, just give that IAM role the permissions it needs to do what you want.

In our case, we needed the role to access another EKS cluster and so the role does not need _any_ _more policies_ in AWS, but it needs to be added to the `aws-auth` [ConfigMap that controls the RBAC](https://docs.aws.amazon.com/eks/latest/userguide/add-user-role.html) of the target cluster. But let's not delve into that here, there's already a ton of posts around that :)

To add a cluster in ArgoCD you can either use the `argocd` CLI, or do it with Kubernetes Secrets. Of course we did it with Kubernetes Secrets, and of course we did that with Terraform after we create the cluster!

```hcl
#
# Get the target cluster details to use in our secret
#
data "aws_eks_cluster" "target" {
  name = "<cluster_name>"
}

#
# Create a secret that represents a new cluster in ArgoCD.
#
# ArgoCD will use the provided config to connect and configure the target cluster
#
resource "kubernetes_secret" "cluster" {
  metadata {
    name      = "argocd-cluster-name"
    namespace = "argocd"
    labels = {
      # Tell ArgoCD that this secret defines a new cluster
      "argocd.argoproj.io/secret-type" = "cluster"
    }
  }

  data = {
  # Just a display name
    name   = data.aws_eks_cluster.target.id
    server = data.aws_eks_cluster.target.endpoint
    config = jsonencode({
      awsAuthConfig = {
        clusterName = data.aws_eks_cluster.target.id
        # NOTE: roleARN not needed as ArgoCD will already assume the role that
        # has access to the target cluster (added to aws-auth ConfigMap)
      }
      tlsClientConfig = {
        insecure = false
        caData   = data.aws_eks_cluster.target.certificate_authority.0.data
      }
    })
  }

  type = "Opaque"
}
```

And boom! We can now create EKS clusters with Terraform and register them with ArgoCD using our Service Account that can assume an AWS IAM Role that is added to the target cluster RBAC... Sometimes it's confusing just to write this stuff, but happy Terraform, Kubernetes and AWS'ing (and GitOps'ing with ArgoCD perhaps)!

## Useful Links

1. AWS EKS IAM Roles for Service Accounts (IRSA): [https://docs.aws.amazon.com/eks/latest/userguide/iam-roles-for-service-accounts.html](https://docs.aws.amazon.com/eks/latest/userguide/iam-roles-for-service-accounts.html "https://docs.aws.amazon.com/eks/latest/userguide/iam-roles-for-service-accounts.html")
2. Workshop on IAM Roles for Service Accounts (IRSA): [https://www.eksworkshop.com/beginner/110_irsa/preparation/](https://www.eksworkshop.com/beginner/110_irsa/preparation/ "https://www.eksworkshop.com/beginner/110_irsa/preparation/")
3. Create AWS IAM OIDC Provider for EKS: [https://docs.aws.amazon.com/eks/latest/userguide/enable-iam-roles-for-service-accounts.html](https://docs.aws.amazon.com/eks/latest/userguide/enable-iam-roles-for-service-accounts.html "https://docs.aws.amazon.com/eks/latest/userguide/enable-iam-roles-for-service-accounts.html")
4. Managing users or IAM roles in EKS: [https://docs.aws.amazon.com/eks/latest/userguide/add-user-role.html](https://docs.aws.amazon.com/eks/latest/userguide/add-user-role.html "https://docs.aws.amazon.com/eks/latest/userguide/add-user-role.html")
5. AWS EKS Terraform Module: [https://registry.terraform.io/modules/terraform-aws-modules/eks/aws/latest](https://registry.terraform.io/modules/terraform-aws-modules/eks/aws/latest "https://registry.terraform.io/modules/terraform-aws-modules/eks/aws/latest")
6. ArgoCD: [https://argo-cd.readthedocs.io/en/stable/](https://argo-cd.readthedocs.io/en/stable/ "https://argo-cd.readthedocs.io/en/stable/")
7. Related GitHub issue on ArgoCD: [https://github.com/argoproj/argo-cd/issues/2347](https://github.com/argoproj/argo-cd/issues/2347 "https://github.com/argoproj/argo-cd/issues/2347")
