---
type: Blog
title: How to use the AWS Load Balancer Controller to connect multiple EKS clusters with existing Application Load Balancers
subheading: Exposing services in AWS EKS clusters via Load Balancers can be done in many ways. In this post we explore using the AWS Load Balancer Controller to dynamically bind nodes to existing Application Load Balancers.
authors:
- jlarfors
tags:
- AWS
- HashiCorp
- Terraform
date: 2022-10-19
image: "/static/blog/aws-load-balancer-controller-with-existing-alb/aws-load-balancer-controller-with-existing-alb.png"
featured: true

---

## Background

In a project I am working on we manage the AWS infrastructure with Terraform (Application Load Balancers, Elastic Kubernetes Service clusters, Security Groups, etc.). We also have one requirement; the Application Load Balancers (ALBs) need to be treated like pets, primarily because another team manages DNS records. This is why we cannot use Kubernetes controllers to dynamically manage the ALBs and update DNS records. Thus, the problem statement can be summarised as: how to manage AWS EKS clusters and ALBs with Terraform, and attach EKS nodes to ALB TargetGroups.

### Connect nodes to ALBs using Terraform

Our initial implementation used Terraform to attach the AWS AutoScalingGroups to TargetGroups.

When using **self-managed node groups** you can pass a list of `target_group_arns` to have any nodes part of the AutoScalingGroup to auto-register themselves as targets to the given TargetGroup ARNs. Nice and easy. [Check the docs](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/autoscaling_group#target_group_arns). We use the community [AWS EKS Terraform module](https://registry.terraform.io/modules/terraform-aws-modules/eks/aws/latest) which supports [this argument](https://registry.terraform.io/modules/terraform-aws-modules/eks/aws/latest/submodules/self-managed-node-group#input_target_group_arns).

When using **EKS managed node groups**, the option to pass `target_group_arns` is not available, and EKS will dynamically generate AutoScalingGroups based on your EKS node group definitions. The side effect here is that the AutoScalingGroup IDs are **not known** until they have been created. When working with Terraform this becomes a problem. It requires you to run Terraform apply with the `-target` option to first provision the EKS node group before you can reference the AutoScalingGroup and attach it to a TargetGroup.

Here’s a lovely GitHub issue with more details: [https://github.com/terraform-aws-modules/terraform-aws-eks/issues/1539](https://github.com/terraform-aws-modules/terraform-aws-eks/issues/1539)

This really is the core source of the problem we wanted to address; how to use EKS managed node groups and not have a hacky solution. And the solution we chose was the [AWS Load Balancer Controller](https://kubernetes-sigs.github.io/aws-load-balancer-controller/v2.4/).

## AWS Load Balancer Controller

The [AWS Load Balancer Controller](https://kubernetes-sigs.github.io/aws-load-balancer-controller/v2.4/) is a Kubernetes controller that can manage the lifecycle of AWS Load Balancers, TargetGroups, Listeners (and Rules), and connect them with nodes (and pods) in your Kubernetes cluster. Check out the [how it works](https://kubernetes-sigs.github.io/aws-load-balancer-controller/v2.4/how-it-works/) page for details on the design.

![AWS-Load-Balancer-diagram-1](/static/blog/aws-load-balancer-controller-with-existing-alb/aws-load-balancer-diagram-1.png)

Looking back at our use case, we want to use an existing Application Load Balancer that is managed by Terraform. If you search for this online, you will most certainly find another lovely GitHub issue: [https://github.com/kubernetes-sigs/aws-load-balancer-controller/issues/228](https://github.com/kubernetes-sigs/aws-load-balancer-controller/issues/228)

It’s a long issue, with lots of suggestions. Personally, I was not concerned with how much of the infrastructure we manage with Terraform vs Kubernetes; the primary goal was a simple solution that did not require too much customisation, and I found the idea of the [TargetGroupBinding](https://kubernetes-sigs.github.io/aws-load-balancer-controller/v2.4/guide/targetgroupbinding/targetgroupbinding/) Custom Resource Definition (CRD) quite appealing.

### TargetGroupBinding Custom Resource Definition

[TargetGroupBinding](https://kubernetes-sigs.github.io/aws-load-balancer-controller/v2.4/guide/targetgroupbinding/targetgroupbinding/) is a CRD that the AWS Load Balancer Controller installs. If you follow the most common use case and manage your ALBs with the AWS Load Balancer controller, it will create TargetGroupBindings under the hood even if you do not interact with them directly. Good news; it is a core feature of the AWS LB Controller, not an extension for people with an edge case. That gave me some confidence.

It requires an existing TargetGroup and IAM policies to lookup and attach/detach targets to the TargetGroup. There’s a mention of the required IAM policies required [here](https://kubernetes-sigs.github.io/aws-load-balancer-controller/v2.4/deploy/installation/#option-b-attach-iam-policies-to-nodes).

Here’s a sample `TargetGroupBinding` manifest taken from the docs:

```yaml
apiVersion: elbv2.k8s.aws/v1beta1
kind: TargetGroupBinding
metadata:
  name: ingress-nginx-binding
spec:
  # The service we want to connect to.
  # We use the Ingress Nginx Controller so let's point to that service.
  # NOTE: it was necessary for this TargetGroupBinding to be in the same namespace as the service.
  serviceRef:
    name: ingress-nginx
    port: 80
  # NOTE: need the ARN of the TargetGroup... Bit of a PITA.
  # Would be nice to use tags to look up the TargetGroup, for example.
  targetGroupARN: arn:aws:elasticloadbalancing:eu-west-1:123456789012:targetgroup/eks-abcdef/73e2d6bc24d8a067
```

For our case, this means managing the ALBs, Listeners, Rules and TargetGroups with Terraform. The AWS Load Balancer Controller would only be responsible for attaching nodes to the specified TargetGroups. This sounds like a clean separation of concerns.

### Multiple EKS clusters, same ALB

Expanding on our particular use case, we manage multiple EKS clusters that share Application Load Balancers. We already use ArgoCD [ApplicationSets](https://argocd-applicationset.readthedocs.io/en/stable/) to manage applications across clusters. We have a “root” cluster that runs core services, like ArgoCD, which connects to multiple other clusters. The below diagram is a high-level simplified illustration of the setup we want to achieve. It will be ArgoCD’s job to deploy the AWS Load Balancer Controller and `TargetGroupBinding` manifests to the different clusters.

![AWS-Load-Balancer-diagram-2](/static/blog/aws-load-balancer-controller-with-existing-alb/aws-load-balancer-diagram-2.png)

Let’s look at how we implemented this with the AWS Load Balancer Controller next.

## Implementation

### Terraform

We use Terraform to manage (amongst other things) the EKS clusters, ALBs and TargetGroups. For implementing the AWS Load Balancer Controller all we needed to do was create the necessary IAM role that can be assumed by a Kubernetes ServiceAccount. The following snippet show this:

```hcl
#
# Create IAM role policy granting the kubernetes service account AssumeRoleWithWebIdentity
#
data "aws_iam_policy_document" "aws_lb_controller" {
  statement {
    actions = ["sts:AssumeRoleWithWebIdentity"]

    principals {
      type = "Federated"
      identifiers = [
        "arn:aws:iam::${data.aws_caller_identity.current.account_id}:oidc-provider/${local.cluster_oidc_issuer}"
      ]
    }

    condition {
      test     = "StringEquals"
      variable = "${local.cluster_oidc_issuer}:sub"
      values   = ["system:serviceaccount:aws-lb-controller:aws-lb-controller"]
    }
  }
}

#
# Create an AWS IAM role that will be assumed by our kubernetes service account
#
resource "aws_iam_role" "aws_lb_controller" {
  name               = "${local.cluster_name}-aws-lb-controller"
  assume_role_policy = data.aws_iam_policy_document.aws_lb_controller.json
  inline_policy {
    name = "${local.cluster_name}-aws-lb-controller"
    policy = jsonencode(
      {
        "Version" : "2012-10-17",
        "Statement" : [
          {
            "Action" : [
              "ec2:DescribeVpcs",
              "ec2:DescribeSecurityGroups",
              "ec2:DescribeInstances",
              "elasticloadbalancing:DescribeTargetGroups",
              "elasticloadbalancing:DescribeTargetHealth",
              "elasticloadbalancing:ModifyTargetGroup",
              "elasticloadbalancing:ModifyTargetGroupAttributes",
              "elasticloadbalancing:RegisterTargets",
              "elasticloadbalancing:DeregisterTargets"
            ],
            "Effect" : "Allow",
            "Resource" : "*"
          }
        ]
      }
    )
  }
}
```

This has some dependencies so cannot be run “as is”. But if you need help configuring IAM Roles for Service Accounts (IRSA) then I already wrote a post on the topic which you can find [here](/blog/how-to-assume-an-aws-iam-role-from-a-service-account-in-eks-with-terraform/).

The `TargetGroupBinding` Kubernetes Custom Resource we need to create requires the TargetGroup ARN which is non deterministic. In our setup, we use Terraform to create the [Kubernetes secret](https://argo-cd.readthedocs.io/en/stable/operator-manual/declarative-setup/#clusters) that informs ArgoCD about connected clusters. Within that secret we can attach additional labels that can be accessed by the ArgoCD [ApplicationSets](https://argocd-applicationset.readthedocs.io/en/stable/), which means we have a very primitive way of passing data from Terraform to ArgoCD without ay extra tools. Note that Kubernetes labels have some restrictions on the [syntax and character set](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#syntax-and-character-set) that can be used, so we can’t just pass in arbitrary data.

Here is how we create the Kubernetes secret that essentially “register” a cluster with ArgoCD, and how we pass the TargetGroup name and ID via labels.

```hcl
locals {
 #
 # Extract the target group name and ID to use in ArgoCD secret
 #
 aws_lb_controller = merge(coalesce(local.apps.aws_lb_controller, { enabled = false }), {
    targetgroup_name = split("/", data.aws_lb_target_group.this.arn_suffix)[1]
    targetgroup_id   = split("/", data.aws_lb_target_group.this.arn_suffix)[2]
  })
}

#
# Get cluster TargetGroup ARNs
#
data "aws_lb_target_group" "this" {
  name = "madeupname-${local.cluster_name}"
}

#
# Create Kubernetes secret in root cluster where ArgoCD is running.
#
# The secret tells ArgoCD about a cluster and how to connect (e.g. credentials).
#
resource "kubernetes_secret" "argocd_cluster" {
  provider = kubernetes.root

  metadata {
    name      = "cluster-${local.cluster_name}"
    namespace = "argocd"
    labels = {
      # Tell ArgoCD that this secret defines a new cluster
      "argocd.argoproj.io/secret-type" = "cluster"
      "environment"                    = var.environment
      "aws-lb-controller/enabled"      = local.aws_lb_controller.enabled
      # Kubernetes labels do not allow ARN values, so pass the name and ID separately
      "aws-lb-controller/targetgroup-name" = local.aws_lb_controller.targetgroup_name
      "aws-lb-controller/targetgroup-id"   = local.aws_lb_controller.targetgroup_id
   ...
   ...

    }
  }

  data = {
    name   = local.cluster_name
    server = data.aws_eks_cluster.this.endpoint
    config = jsonencode({
      awsAuthConfig = {
        clusterName = data.aws_eks_cluster.this.id
        # Provide the rolearn that was created for this cluster, and which the
        # root ArgoCD role should be able to assume
        roleARN = aws_iam_role.argocd_access.arn
      }
      tlsClientConfig = {
        insecure = false
        caData   = data.aws_eks_cluster.this.certificate_authority.0.data
      }
    })
  }

  type = "Opaque"
}
```

That’s it for the Terraform config. We use these snippets in Terraform modules that get called for each cluster we create, keeping things DRY.

### ArgoCD

Let’s first look at the directory structure and we can work through the relevant files (note that in our code this exists alongside many other applications inside a Git repository).

```console
.
├── appset-aws-lb-controller.yaml
├── appset-targetgroupbindings.yaml
├── chart
│   ├── Chart.yaml
│   ├── README.md
│   ├── templates
│   │   └── targetgroupbindings.yaml
│   └── values.yaml
└── kustomization.yaml

2 directories, 7 files
```

We use [Kustomize](https://kustomize.io/) to connect our ArgoCD applications together (minimising the number of “app of apps” connections needed) and that’s what the `kustomization.yaml` file is for, and here it contains the two top-level `appset-*.yaml` files.

The `appset-aws-lb-controller.yaml` file contains the AWS Load Balancer Controller ApplicationSet which uses the [Helm chart](https://github.com/kubernetes-sigs/aws-load-balancer-controller/tree/main/helm/aws-load-balancer-controller) to install the controller on each of our clusters.

```yaml
# File: appset-aws-lb-controller.yaml
apiVersion: argoproj.io/v1alpha1
kind: ApplicationSet
metadata:
  name: aws-lb-controller
  namespace: argocd
spec:
  generators:
  # This is a little trick we use for nearly all our apps to control which apps
  # should be installed on which clusters. Terraform creates the secret that
  # contains these labels, so that's where the logic is controlled for setting
  # these labels to true/false.
    - clusters:
        selector:
          matchLabels:
            aws-lb-controller/enabled: "true"
  syncPolicy:
    preserveResourcesOnDeletion: false
  template:
    metadata:
      name: "aws-lb-controller-{{ name }}"
      namespace: argocd
    spec:
      project: madeupname
      destination:
        name: "{{ name }}"
        namespace: aws-lb-controller
      syncPolicy:
        automated:
          prune: true
          selfHeal: true
        syncOptions:
          - CreateNamespace=true
          - PruneLast=true

      source:
        repoURL: "https://aws.github.io/eks-charts"
        chart: aws-load-balancer-controller
        targetRevision: 1.4.4
        helm:
          releaseName: aws-lb-controller
          values: |
            clusterName: {{ name }}
            serviceAccount:
              create: true
              annotations:
                "eks.amazonaws.com/role-arn": "arn:aws:iam::123456789012:role/{{ name }}-aws-lb-controller"
              name: aws-lb-controller
            # We won't be using ingresses with this controller.
            createIngressClassResource: false
            disableIngressClassAnnotation: true

            resources:
              limits:
                cpu: 100m
                memory: 128Mi
              requests:
                cpu: 100m
                memory: 128Mi
```

Next up we have the `appset-targetgroupbindings.yaml` ApplicationSet which creates the TargetGroupBinding on each of our clusters. For this, we needed to template some values based on the cluster and the most straightforward way I have found with ArgoCD is to create a minimalistic Helm chart for our purpose. This is what the `chart/` directory is, which contains a single template file `targetgroupbindings.yaml`.

Let’s first look at the ApplicationSet which installs the Helm Chart:

```yaml
# File: appset-targetgroupbindings.yaml
apiVersion: argoproj.io/v1alpha1
kind: ApplicationSet
metadata:
  name: targetgroupbindings
  namespace: argocd
spec:
  generators:
    - clusters:
        selector:
          matchLabels:
            aws-lb-controller/enabled: "true"
  syncPolicy:
    preserveResourcesOnDeletion: false
  template:
    metadata:
      name: "targetgroupbindings-{{ name }}"
      namespace: argocd
    spec:
      project: madeupname
      destination:
        name: "{{ name }}"
        # TargetGroupBindings need to be in the same namespace as the service
        # they bind to
        namespace: ingress-nginx
      syncPolicy:
        automated:
          prune: true
          selfHeal: true
        syncOptions:
          - CreateNamespace=true
          - PruneLast=true

      source:
        repoURL: "<url-of-this-repo>"
        path: "apps/aws-lb-controller/chart"
        targetRevision: "master"
        helm:
          releaseName: targetgroupbinding
          values: |
            # AWS ARNs are not valid Kubernetes label values, so we had to split the ARN up and glue it back together here.
            targetGroupArn: "arn:aws:elasticloadbalancing:eu-west-1:123456789012:targetgroup/{{ metadata.labels.aws-lb-controller/targetgroup-name }}/{{ metadata.labels.aws-lb-controller/targetgroup-id }}"
```

Next we can look at the single template in our minimal Helm chart:

```yaml
# File: targetgroupbindings.yaml
apiVersion: elbv2.k8s.aws/v1beta1
kind: TargetGroupBinding
metadata:
  name: ingress-nginx
spec:
  targetType: instance
  serviceRef:
    name: ingress-nginx-controller
    port: 80
  targetGroupARN: {{ required "targetGroupArn required" .Values.targetGroupArn }}
  # By default, add all nodes to the cluster unless they have the label
  # exclude-from-lb-targetgroups set
  nodeSelector:
    matchExpressions:
      - key: exclude-from-lb-targetgroups
        operator: DoesNotExist
```

It’s not an ideal situation to use Helm charts for this purpose, but the inspiration came from the ArgoCD app of apps [example repository](https://github.com/argoproj/argocd-example-apps/tree/master/helm-guestbook). Anyway, I am fairly happy with this implementation and it works dynamically for any new clusters that we add to ArgoCD.

## Conclusion

In this post we looked at a fairly specific problem; binding EKS cluster nodes to existing Application Load Balancers using the `TargetGroupBinding` CRD from the AWS Load Balancer Controller. The motivation to make this write-up came from the number of people asking about this on GitHub, and I think this is quite a simple and elegant approach.

A point worth noting is that using the AWS Load Balancer Controller decouples your node management with your cluster management. Let’s say we wanted to use [Karpenter](https://karpenter.sh/) for autoscaling instead of the defacto cluster-autoscaler. Karpenter will not use AWS AutoScalingGroups but will instead create standalone EC2 instances based on the [Provisioners](https://karpenter.sh/v0.16.2/provisioner/) you define. This means our previous approach of attaching AutoScalingGroups with TargetGroups will not work as the EC2 instances Karpenter manages will not belong to the AutoScalingGroup and therefore not be automatically attached to the TargetGroup. The AWS Load Balancer Controller doesn’t care how the nodes are created; only that they belong to the cluster and match the label selectors defined. Probably we will look into Karpenter again in the near future for our project now that it [supports pod anti-affinity](https://github.com/aws/karpenter/issues/942), as this was previously a blocker for us.

If you have any suggestions or questions about this post, please leave a comment or [get in touch with us](/contact/)!
