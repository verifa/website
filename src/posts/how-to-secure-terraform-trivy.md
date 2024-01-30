---
type: Blog
title: How to secure Terraform code with Trivy
subheading: Learn how Trivy can be used to secure your Terraform code and integrated into your development workflow.
authors:
- mvainio
tags:
- DevOps
- Continuous Integration
- Terraform
- Open Source
- DevSecOps
date: 2024-01-24
image: "/static/blog/how-to-secure-terraform-code-trivy.png"
featured: true
jobActive: true

---

**In this blog post we will look at securing an AWS Terraform configuration using Trivy to check for known security issues. We will explore different ways of using Trivy, integrating it into your CI pipelines, practical issues you might face and solutions to those issues to get you started with improving the security of your IaC codebases.**

Terraform is a powerful tool with a thriving community that makes it easy to find ready-made modules and providers for practically any cloud platform or service that exposes an API. Also internally many companies have a great deal of modules available. One of the strengths of Terraform is that modules provide an abstraction. You don’t have to worry about what is underneath the module’s variables (interface); you provide the necessary values and off you go, but this might lead into some trouble security-wise. Especially in public cloud platforms, it’s easy to expose a VM, load balancer or an object storage bucket publicly to the internet, and when using an abstraction this can happen without you truly acknowledging it. If you are familiar with Terraform, you might say, “Well I will just review the plan before applying”. But if the module is presenting you a plan of creating/modifying 200 resources, are you really confident you can eyeball that information and catch a misconfiguration that would expose your infrastructure to an attacker?

At the time of writing there are around 16,000 modules available from [HashiCorp’s public registry](https://registry.terraform.io/browse/modules). In this post we will pick a couple of AWS modules and check for insecure configurations. For this “check”, we will use an open source tool called [Trivy](https://trivy.dev/).

## Security Scanners for Terraform

One of the big upsides of maintaining your infrastructure using an IaC approach is the fact that your infrastructure can be analysed by static analysis tools since your infrastructure is in plain text files. We can analyse the infrastructure before creating any resources to get quick feedback on the security posture and fix any issues before deployment. The only problem is that there are so many tools! After trying few alternatives, however, I have settled on a favourite that is both easy to use and effective at finding issues with built-in checks. In the past this favourite tool was `tfsec` , but quite recently the development efforts of the Tfsec project have been migrated into the Trivy project. Thus, it’s time to move over to Trivy although it’s not specialised to Terraform like Tfsec was.

> [!NOTE]
> I noticed a few differences between Tfsec and Trivy when comparing their results and I will make note of these later in the hands-on section. Based on the GitHub issues and PRs, both open and closed, I am confident Trivy will eventually match and surpass Tfsec in features and accuracy as the development team looks very keen on closing gaps between the two tools.

Worth noting that there are some great open-source alternatives to Trivy, but overall we have found Trivy to be both easy to use locally and to integrate into build pipelines.

There’s of course nobody stopping you from using multiple tools, and when you automate the checks, that might not be a big deal to implement in the end. However, the goal of this blog post is not to focus on comparing different tools. What we want to focus on is that **you should use a tool like this to perform security checks on your Terraform code,** and it’s quite trivial to accomplish in the end.

## Introduction to Trivy

Trivy is a Swiss army knife type of tool for security scanning of various types of artifacts and code. It can scan different targets such as your local filesystem or a container image from a container registry. It can also check for many kinds of security issues such as known vulnerabilities, exposed secrets and most relevant to this blog post; misconfigurations.

At the time of writing Trivy supports scanning of various IaC configurations such as Terraform, [CloudFormation](https://aws.amazon.com/cloudformation/) and [Azure Resource Manager](https://learn.microsoft.com/en-us/azure/azure-resource-manager/management/overview). So even if your organisation uses different tools across teams, Trivy might just be the right tool. Trivy comes with built-in checks for various cloud platforms and in this blog post we will only use the built-in checks, but you can also define your own [custom checks/policies](https://aquasecurity.github.io/trivy/latest/docs/scanner/misconfiguration/custom/).

Trivy can also scan for secrets which you should also use in the IaC context, but this is not really specific to the Terraform use-case. I suggest looking into the [Trivy documentation](https://aquasecurity.github.io/trivy/latest/docs) to discover all of it’s power beyond what I already covered here as this will naturally evolve over time.

Now it’s time to get our hands dirty and look at an example of how Trivy can save you from doing things that might put your organisation in jeopardy.

## Installing Trivy

For installation I suggest checking out the [installation guide in the documentation](https://aquasecurity.github.io/trivy/latest/getting-started/installation/) that covers all supported platforms. But for a quick start, here are a couple of commands that work for most folks:

```text
brew install trivy
```

For Debian/Ubuntu:

```text
apt install trivy
```

For Windows:

```text
choco install trivy
```

There are also pre-built packages available for various Linux distros, or grab the binary from GitHub releases: [https://github.com/aquasecurity/trivy/releases](https://github.com/aquasecurity/trivy/releases)

I highly suggest [verifying the signature](https://aquasecurity.github.io/trivy/latest/getting-started/signature-verification/) when installing, especially when you are using Trivy in your production build pipelines.

## Scanning an Example Terraform Module

Let’s create an example Terraform root module in order to get something to point Trivy at. Like I mentioned earlier, there are many open-source modules for Terraform that we can utilise in order to quickly build infrastructure. The [AWS modules](https://registry.terraform.io/namespaces/terraform-aws-modules) are especially popular, so I thought let’s write an example by utilising a couple of these modules with mostly their default configuration. Here’s what I came up with:

```hcl
#main.tf
terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

locals {
  common_tags = {
    Terraform = "true"
    Environment = "dev"
  }
}

module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "5.0.0"

  name = "my-vpc"
  cidr = "10.0.0.0/16"

  azs             = ["eu-west-1a", "eu-west-1b", "eu-west-1c"]
  private_subnets = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
  public_subnets  = ["10.0.101.0/24", "10.0.102.0/24", "10.0.103.0/24"]

  enable_nat_gateway = false

  tags = local.common_tags
}

data "aws_ami" "amazon_linux" {
  most_recent = true

  owners = ["amazon"]

  filter {
    name = "name"

    values = [
      "amzn2-ami-hvm-*-x86_64-gp2",
    ]
  }

  filter {
    name = "owner-alias"

    values = [
      "amazon",
    ]
  }
}

resource "aws_instance" "this" {
  ami           = data.aws_ami.amazon_linux.id
  instance_type = "t3.nano"
  subnet_id     = element(module.vpc.private_subnets, 0)
}

module "alb" {
  source  = "terraform-aws-modules/alb/aws"
  version = "8.7.0"

  name = "my-alb"

  load_balancer_type = "application"

  vpc_id             = module.vpc.vpc_id
  subnets            = module.vpc.public_subnets

  target_groups = [
    {
      name_prefix      = "pref-"
      backend_protocol = "HTTP"
      backend_port     = 80
      target_type      = "instance"
      targets = {
        my_ec2 = {
          target_id = aws_instance.this.id
          port      = 8080
        }
      }
    }
  ]

  http_tcp_listeners = [
    {
      port               = 80
      protocol           = "HTTP"
      target_group_index = 0
    }
  ]

  tags = local.common_tags
}
```

This configuration is ~100 LoC and it will create a VPC, an EC2 instance and an ALB. Naturally, the ALB also targets the EC2 instance. After creating the file, let’s initialise Terraform to download the external modules:

```bash
terraform init
```

> [!NOTE]
> If you are working with local modules then there is no need to run `terraform init` before the scan as all files are already present, but remote modules must be fetched in to the `.terraform` folder before a scan.

Simplest way to run a Trivy misconfiguration scan is to point it at your current folder:

```bash
trivy config .
```

Like mentioned earlier, we can also scan for secrets at the same time with Trivy:

```bash
trivy fs --scanners misconfig,secret .
```

Due to the focus on Terraform, I’ll use the `config` subcommand for the rest of the blog post, but in a CI pipeline I would run the secrets scanning definitely for the whole repository as well, not only in IaC folders.

Before showing the full results, I noticed there are some example configurations picked up by Trivy from the remote modules, such as this:

```bash
HIGH: IAM policy document uses sensitive action 'logs:CreateLogStream' on wildcarded resource '*'
═══════════════════════════════════════════════════════════════════════════════════════════════════════
You should use the principle of least privilege when defining your IAM policies.
This means you should specify each exact permission required without using wildcards,
as this could cause the granting of access to certain undesired actions, resources and principals.

See https://avd.aquasec.com/misconfig/avd-aws-0057
───────────────────────────────────────────────────────────────────────────────────────────────────────
 modules/vpc/vpc-flow-logs.tf:112
   via modules/vpc/vpc-flow-logs.tf:100-113 (data.aws_iam_policy_document.vpc_flow_log_cloudwatch[0])
    via modules/vpc/vpc-flow-logs.tf:97-114 (data.aws_iam_policy_document.vpc_flow_log_cloudwatch[0])
     via modules/vpc/examples/complete/main.tf:25-82 (module.vpc)
───────────────────────────────────────────────────────────────────────────────────────────────────────
  97   data "aws_iam_policy_document" "vpc_flow_log_cloudwatch" {
  ..
 112 [     resources = ["*"]
 ...
 114   }
```

If you look closely you notice that the source of the finding is a `main.tf` file in the examples folder of the VPC module: `via modules/vpc/examples/complete/main.tf:25-82 (module.vpc)`

This is not really our configuration and we should not include these files in the scan. This is also a difference between Tfsec and Trivy, when running Tfsec it does not pickup the examples folder.

However, we can easily resolve this by skipping all files under `examples` folders and then we should get proper report of findings:

```bash
trivy config . --skip-dirs '**/examples'
```

Here are the results:

```bash
.terraform/modules/alb/main.tf (terraform)
==========================================
Tests: 3 (SUCCESSES: 1, FAILURES: 2, EXCEPTIONS: 0)
Failures: 2 (UNKNOWN: 0, LOW: 0, MEDIUM: 0, HIGH: 2, CRITICAL: 0)

HIGH: Application load balancer is not set to drop invalid headers.
════════════════════════════════════════════════════════════════════════════════════════════════
Passing unknown or invalid headers through to the target poses a potential risk of compromise.

By setting drop_invalid_header_fields to true, anything that doe not conform to well known,
defined headers will be removed by the load balancer.

See https://avd.aquasec.com/misconfig/avd-aws-0052
────────────────────────────────────────────────────────────────────────────────────────────────
 .terraform/modules/alb/main.tf:23
   via .terraform/modules/alb/main.tf:5-63 (aws_lb.this[0])
────────────────────────────────────────────────────────────────────────────────────────────────
   5   resource "aws_lb" "this" {
   .
  23 [   drop_invalid_header_fields                  = var.drop_invalid_header_fields
  ..
  63   }
────────────────────────────────────────────────────────────────────────────────────────────────

HIGH: Load balancer is exposed publicly.
════════════════════════════════════════════════════════════════════════════════════════════════
There are many scenarios in which you would want to expose a load balancer to the wider internet,
but this check exists as a warning to prevent accidental exposure of internal assets.
You should ensure that this resource should be exposed publicly.

See https://avd.aquasec.com/misconfig/avd-aws-0053
────────────────────────────────────────────────────────────────────────────────────────────────
 .terraform/modules/alb/main.tf:12
   via .terraform/modules/alb/main.tf:5-63 (aws_lb.this[0])
────────────────────────────────────────────────────────────────────────────────────────────────
   5   resource "aws_lb" "this" {
   .
  12 [   internal           = var.internal
  ..
  63   }
────────────────────────────────────────────────────────────────────────────────────────────────

.terraform/modules/vpc/main.tf (terraform)
==========================================
Tests: 1 (SUCCESSES: 0, FAILURES: 1, EXCEPTIONS: 0)
Failures: 1 (UNKNOWN: 0, LOW: 0, MEDIUM: 1, HIGH: 0, CRITICAL: 0)

MEDIUM: VPC Flow Logs is not enabled for VPC
════════════════════════════════════════════════════════════════════════════════════════════════
VPC Flow Logs provide visibility into network traffic that traverses the VPC and can be used to
detect anomalous traffic or insight during security workflows.

See https://avd.aquasec.com/misconfig/avd-aws-0178
────────────────────────────────────────────────────────────────────────────────────────────────
 .terraform/modules/vpc/main.tf:29-52
────────────────────────────────────────────────────────────────────────────────────────────────
  29 ┌ resource "aws_vpc" "this" {
  30 │   count = local.create_vpc ? 1 : 0
  31 │
  32 │   cidr_block          = var.use_ipam_pool ? null : var.cidr
  33 │   ipv4_ipam_pool_id   = var.ipv4_ipam_pool_id
  34 │   ipv4_netmask_length = var.ipv4_netmask_length
  35 │
  36 │   assign_generated_ipv6_cidr_block     = var.enable_ipv6 && !var.use_ipam_pool ? true : null
  37 └   ipv6_cidr_block                      = var.ipv6_cidr
  ..
────────────────────────────────────────────────────────────────────────────────────────────────

main.tf (terraform)
===================
Tests: 3 (SUCCESSES: 1, FAILURES: 2, EXCEPTIONS: 0)
Failures: 2 (UNKNOWN: 0, LOW: 0, MEDIUM: 0, HIGH: 2, CRITICAL: 0)

HIGH: Instance does not require IMDS access to require a token
════════════════════════════════════════════════════════════════════════════════════════════════

IMDS v2 (Instance Metadata Service) introduced session authentication tokens which improve
security when talking to IMDS.
By default <code>aws_instance</code> resource sets IMDS session auth tokens to be optional.
To fully protect IMDS you need to enable session tokens by using <code>metadata_options</code>
block and its <code>http_tokens</code> variable set to <code>required</code>.

See https://avd.aquasec.com/misconfig/avd-aws-0028
────────────────────────────────────────────────────────────────────────────────────────────────
 main.tf:55-59
────────────────────────────────────────────────────────────────────────────────────────────────
  55 ┌ resource "aws_instance" "this" {
  56 │   ami           = data.aws_ami.amazon_linux.id
  57 │   instance_type = "t3.nano"
  58 │   subnet_id     = element(module.vpc.private_subnets, 0)
  59 └ }
────────────────────────────────────────────────────────────────────────────────────────────────

HIGH: Root block device is not encrypted.
════════════════════════════════════════════════════════════════════════════════════════════════
Block devices should be encrypted to ensure sensitive data is held securely at rest.

See https://avd.aquasec.com/misconfig/avd-aws-0131
────────────────────────────────────────────────────────────────────────────────────────────────
 main.tf:55-59
────────────────────────────────────────────────────────────────────────────────────────────────
  55 ┌ resource "aws_instance" "this" {
  56 │   ami           = data.aws_ami.amazon_linux.id
  57 │   instance_type = "t3.nano"
  58 │   subnet_id     = element(module.vpc.private_subnets, 0)
  59 └ }
────────────────────────────────────────────────────────────────────────────────────────────────
```

For brevity I shortened some of the long lines.

## Inspecting the Results

Unfortunately Trivy does not print a summary in the end like `tfsec` does which makes it nice to read the output from bottom to top. Trivy does offer different ways to modify the resulting report, but for the needs of this blog I quickly used `grep` to find a short summary of each finding:

```text
HIGH: Application load balancer is not set to drop invalid headers.
HIGH: Load balancer is exposed publicly.
MEDIUM: VPC Flow Logs is not enabled for VPC
HIGH: Instance does not require IMDS access to require a token
HIGH: Root block device is not encrypted.
```

Looking at the list, I think we want to change the configuration before deploying the resources. Next, let’s look at our options for resolving these findings.

## Resolving the Issues

We have two choices when it comes to resolving these issues so that we can have a nice clean report (until we change the configuration again). We can either resolve the issues by modifying our configuration or we can choose to accept the finding as something that is not relevant for our requirements and ignore the findings for future scans.

Since we use the public AWS modules, we cannot easily make changes besides the inputs without forking the source module, however we can change the EC2 instance which is defined directly in the `main.tf` . So let’s look into the issues related to the EC2 instance first.

There is a finding related to the AWS Instance Metadata Service (IMDS). The finding is related to making sure the instance uses the IMDSv2 instead of the legacy IMDSv1. Looking at the full report above, you can see that Trivy explicitly tells us what the problem is and how to resolve it:

```text
IMDS v2 (Instance Metadata Service) introduced session authentication tokens which improve
security when talking to IMDS.
By default <code>aws_instance</code> resource sets IMDS session auth tokens to be optional.
To fully protect IMDS you need to enable session tokens by using <code>metadata_options</code>
block and its <code>http_tokens</code> variable set to <code>required</code>.
```

You can read more of the security benefits and scenarios where this configuration matters in the [IMDSv2 announcement blog post by AWS](https://aws.amazon.com/blogs/security/defense-in-depth-open-firewalls-reverse-proxies-ssrf-vulnerabilities-ec2-instance-metadata-service/).

To resolve this, we are going to change the configuration in the following way, like the description suggested:

```diff
resource "aws_instance" "this" {
   ami           = data.aws_ami.amazon_linux.id
   instance_type = "t3.nano"
   subnet_id     = element(module.vpc.private_subnets, 0)
+
+  metadata_options {
+    http_tokens = "required"
+  }
 }
```

If you run the scan again you will notice the finding is gone:

```bash
trivy config . --skip-dirs '**/examples'
```

Let’s move onto the next issue that is related to the root disk being unencrypted for this EC2 instance. In reality I would configure encryption because it is a common compliancy requirement and AWS makes it very easy, but for the sake of the example let’s ignore this instead, saying that we are ok with running unencrypted root disk:

```diff
+#trivy:ignore:avd-aws-0131
 resource "aws_instance" "this" {
   ami           = data.aws_ami.amazon_linux.id
   instance_type = "t3.nano"
   subnet_id     = element(module.vpc.private_subnets, 0)

   metadata_options {
     http_tokens = "required"
   }
 }
```

Using the inline method for ignoring findings is the most intuitive way in my opinion, this might be familiar to you if you have worked with just about any code linter in the past.

Now the only remaining issues are related to the AWS modules which we did not author. Unfortunately, right now Trivy can’t figure out that the remote modules are downloaded under different path (`.terraform/modules`) than what is declared when specifying the source for the modules:

```hcl
module "alb" {
  source  = "terraform-aws-modules/alb/aws"
  version = "8.7.0"
...
```

Again, not an issue when using local modules since the paths nicely match between the findings report and the declaration in the Terraform configuration. This issue is actively being [discussed in the Trivy GitHub repository](https://github.com/aquasecurity/trivy/discussions/5872), so when you read this it might be fixed and I should update this blog.

Luckily Trivy has a cure for this even without us waiting for a fix. I quickly brewed a solution using an [advanced filtering mechanism](https://aquasecurity.github.io/trivy/v0.48/docs/configuration/filtering/#by-open-policy-agent) in Trivy that uses the [Rego](https://www.openpolicyagent.org/docs/latest/policy-language/) language:

```rego
package trivy

import data.lib.trivy

default ignore = false

ignore_avdid := {"AVD-AWS-0052", "AVD-AWS-0053"}

ignore_severities := {"LOW", "MEDIUM"}

ignore {
 input.AVDID == ignore_avdid[_]
}

ignore {
 input.Severity == ignore_severities[_]
}
```

This will resolve the issue with the VPC module because all `MEDIUM` findings are ignored, and the  findings in the ALB module are ignored explicitly by the finding IDs.

Now we can run a scan and include this policy from a local file:

```bash
trivy config . --skip-dirs '**/examples' --ignore-policy custom-policy.rego
```

Now all the findings should be resolved (of course if you try this yourself, there might be new built-in checks and you get a different list of findings).

While authoring the custom ignore policy, I couldn’t figure out how to connect the source of the finding (the ALB module) to the AVDID for a more fine-grained rule, but that does not seem like a big deal to me. You should not place all your IaC configuration into a single root module after all, in production I would separate the VPC creation from this Terraform root module and use an existing one that is part of a different root module.

## Scanning Terraform Plans

Another way to run the scan is to first create a plan and then convert the plan from the default binary format to JSON, and then point Trivy to scan this plan that contains a list of changes:

```bash
terraform plan --out tf.plan
terraform show -json tf.plan > tfplan.json
trivy config tfplan.json
```

I noticed that there’s one additional finding when running Trivy against the plan:

```text
CRITICAL: Listener for application load balancer does not use HTTPS.
═════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════
Plain HTTP is unencrypted and human-readable. This means that if a malicious actor was to eavesdrop on your connection,
they would be able to see all of your data flowing back and forth.

You should use HTTPS, which is HTTP over an encrypted (TLS) connection, meaning eavesdroppers cannot read your traffic.

See https://avd.aquasec.com/misconfig/avd-aws-0054
──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────
 main.tf:36
   via main.tf:34-48 (aws_lb_listener.frontend_http_tcp_ffdb4db32d85be4b5cd7539e4d3c6d16)
──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────
  ..
  36 [  protocol = "HTTP"
  ..
  48   }
────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────
```

I found it odd that this was not found in the previous scan and after some testing this seems to be caused by using remote modules (again!). I also noticed that running `tfsec` instead of Trivy will catch this issue. When using local modules Trivy can right away pick up this finding. Right now it seems necessary to scan both your plan and your configuration for the best accuracy, but I might revisit this in the future and see if the issue is resolved given the active discussion around remote modules support in Trivy.

Luckily, we can work around this also by scanning everything at once, so generate the plan and then scan the entire folder (then Trivy processes the configuration AND the plan):

```bash
terraform plan --out tf.plan
terraform show -json tf.plan > tfplan.json
trivy config . --skip-dirs '**/examples'
```

Trivy is smart enough to not duplicate the findings - awesome! With Tfsec it was not possible to scan Terraform plans, so this is great if scanning the plan fits your workflow better.

As we saw above, scanning the Terraform plan is more accurate than scanning just the files, but the downside is that you need to generate a plan for each Terraform root module and it takes much more time to generate a plan prior to running Trivy. You also need to be able to connect and authenticate to the providers for Terraform to generate the plan, although that’s not typically an issue.

One more thing to note about plans is that your inline `#trivy:ignore` comments will be ignored since that information will not make it into the plan, so if you are using plans primarily for your scanning, then you might need to get comfortable defining the Rego ignore policies instead.

## Running Trivy in CI

Including Trivy scans in your IaC repositories’ CI pipelines is a must. If you don’t have CI pipelines for your IaC… Well you should! Trivy offers integrations with many CI/CD tools, IDEs and other systems, see the [documentation for an up-to-date list](https://aquasecurity.github.io/trivy/latest/ecosystem/).

When using Trivy in CI it’s wise to use a [configuration file](https://aquasecurity.github.io/trivy/latest/docs/references/configuration/config-file/) instead of the command line flags, this makes it easy to reproduce the scan using same configuration locally if you need to investigate some new findings. If you are using GitHub Actions, there’s an [official Action](https://github.com/aquasecurity/trivy-action/tree/master) that you can use to integrate Trivy into your CI pipeline, here’s a simple example which uses a configuration file:

```yaml
name: build
on:
  push:
    branches:
    - main
  pull_request:
jobs:
  build:
    name: Build
    runs-on: ubuntu-20.04
    steps:
    - name: Checkout code
      uses: actions/checkout@v3

    - name: Run Trivy vulnerability scanner in fs mode
      uses: aquasecurity/trivy-action@master
      with:
        scan-type: 'fs'
        scan-ref: '.'
        trivy-config: trivy.yaml
```

As seen in the above example, in CI you likely want to run the `fs` scan which includes by default all the scanners, meaning Trivy will also scan for secrets and vulnerabilities, not only for misconfigurations.

However, keep in mind [this excellent blog post by my colleague Thierry](https://verifa.io/blog/keep-your-pipelines-simple). The Trivy action only really wraps the GitHub workflow YAML inputs to CLI flags. If you are using another CI/CD system, you can simply install and invoke the CLI as well, making transition between CI/CD tools extremely simple.

## Bonus for GitHub Users

If you have an open-source project in GitHub or you pay GitHub for the advanced security features, then you can also upload the Trivy scan results into the GitHub code scanning which you should be using if you are not already. Refer to the [Trivy Action’s README](https://github.com/aquasecurity/trivy-action?tab=readme-ov-file#using-trivy-with-github-code-scanning) to view a sample configuration of uploading the results. This will help you to track the findings in addition to gatekeeping with a pipeline that must pass always before merging code (or however your team works).

## Summary

In this blog post we rolled up our sleeves and looked into how to secure Terraform configuration by using static analysis. As an example and recommended tool we explored Trivy, but honestly the tool choice isn’t as important as the principle of integrating such checks into your workflow. I hope you can see the value of running a simple scan over your configuration. Thanks to extensive builtin checks in Trivy you can get actionable findings without spending time reviewing the configuration manually and magically knowing all the security intricacies of AWS infrastructure.

If you found something wrong with the content or something felt vague or awesome, leave us a comment! Additionally, if you’d like any help with Terraform and/or Trivy [please get in touch](https://verifa.io/contact/index.html)!
