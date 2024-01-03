---
type: Blog
title: "Nomad: The workload orchestrator you may have missed"
subheading: "Nomad allows you to deploy and manage applications across on-premise and cloud-based
platforms. Its power comes with its flexibility yet ease of operation and use."
authors:
- jlarfors
tags:
  - HashiCorp Nomad
  - HashiCorp
  - Containers
  - Continuous Delivery
date: 2020-09-05
image: /static/blog/2020-09-05/main.svg
featured: true
---

**The most well-known orchestrator today is Kubernetes and its adoption is rising, with its main rivals such as Docker Swarm and Mesos being replaced. Nomad is a lesser-known alternative that has been around since 2015 and is a very compelling alternative to Kubernetes.**

With Kubernetes having become the de facto orchestrator in the market you might be wondering
why there is a need for alternatives. Whilst Kubernetes might be able to meet all of your
requirements, it is not the most straightforward tool to adopt, and often provides way more
features than what teams might typically need.

Administrators ("ops" team) will need to deploy, manage and troubleshoot Kubernetes clusters
and developers will need to learn how to deploy and manage applications on Kubernetes
clusters.

Of course, using a managed Kubernetes service (such as those provided by Google, Amazon
and Azure) greatly simplifies the operations side of a Kubernetes cluster, but it still requires
some fundamental in-house knowledge that needs to be developed. Furthermore, there are a
significant number of teams using on-premise or other hosted solutions that do not have access
to such managed Kubernetes services and need to deploy and maintain their own Kubernetes
clusters, or those that choose to for cost reasons (on-premise hardware is cheap!).

## What is Nomad?

Nomad by HashiCorp (first released in 2015) is a workload orchestrator that provides a wide
variety of capabilities with one very noteworthy trait: it is straightforward to deploy.
This is by design, as Nomad only delivers the function that you might be looking for in an
orchestrator - namely deploying and managing applications. Interestingly it supports both
containerized and non-containerized applications.

What do we mean by this? Well, if you need extra networking, storage or secret management
capabilities you can bring this to the cluster yourself. Conveniently, HashiCorp provides native
support for tools like Consul (for networking) and Vault (for secrets management) to make
adding these capabilities quite manageable.

A quick walkthrough of the Nomad architecture will help us elaborate.

![Nomad diagram](/static/blog/2020-09-05/nomad-diagram.svg)

## Nomad Architecture

Nomad (like the other HashiCorp tools) is a single binary. When the binary is run on a node in a
cluster it runs in agent mode, of which it is either a server or a client. The servers are the
"brains" of the cluster, and the clients are where workloads run.

So what's the setup process? Well, we would highly recommend creating a Consul cluster first.
Then Nomad connects itself within the Consul cluster, which is a really pleasant experience to
go through - especially having worked with setting up Docker Swarm and Kubernetes clusters
previously.

Consul does not just make setting up a Nomad cluster easier, but once you are deploying
applications on your Nomad cluster it provides service discovery and a really slick DNS server
to help route your requests to the relevant services in the cluster. You can of course run your
Nomad cluster without Consul and use something like a reverse proxy to handle your requests
and networking, but for the low effort it takes to setup a Consul cluster it should be heavily
considered.

As a side note, compare this to the setup of a Kubernetes cluster where, if you want a cluster,
you need all of the cluster - you cannot simply opt out of some of the Kubernetes components.
This makes the operational overhead very high.

## Getting started with Nomad

If you wanted to get started playing around with Nomad, we would recommend familiarising
yourself with Consul first.

This is optional, but we recommend Consul with Nomad - you can follow the great online
[tutorials by HashiCorp](https://learn.hashicorp.com/consul). And once you are familiar with
Consul, you can start looking at the Nomad tutorials, also by [HashiCorp](https://learn.hashicorp.com/nomad).

These tutorials will take you through some fundamental concepts and give you a pretty solid
base for experimenting further with the setup.

Let's take a look at an example Nomad job to give you a flavour of how "jobs" in Nomad look. A
nice place to start would be to add secrets management with Vault to your setup, so let's
inspect a very basic example of this along with some nice comments for your reference.

```hcl
job "vault" {
  // specify the nomad data centers which are eligible for task placement
  datacenters = ["dc1"]
  // specify the nomad scheduler to use
  type = "service"

  // tasks within the same group are co-located on the same nomad client
  group "vault" {
  // have only one instance of the group tasks running
    count = 1

    network {
      // create a network namespace for this group
      mode = "bridge"
      // define a port called "http" with static port 8200
      port "http" {
        static = 8200
        to = 8200
      }

      // can specify some requirements for network speed...
      mbits = 20
    }

    // create a consul service
    service {
      name = "vault"
      port = "http"
      // specify a health check for the consul service

      check {
        name = "alive"
        type = "http"
        path = "/ui"
        interval = "10s"
        timeout = "10s"
      }
    }

    // task for running the vault docker container
    task "vault" {
      driver = "docker"
      config {
        image = "vault:1.4.2"
      }

      resources {
        cpu = 500
        memory = 512
      }
    }
  }
}
```

What's elegant about Nomad jobs is that you can capture so much information about the
deployment in one file. This is a very basic example which does not make use of other job
features, which are [documented online](https://www.nomadproject.io/docs/job-specification).
How does Nomad work?

Nomad uses the HashiCorp Configuration Language (HCL) which is used by most of the other
HashiCorp tools. It provides a really nice balance between being human readable (and usable)
and machine-friendly.

From a developer perspective, Nomad can be run in "dev" mode which means you can run your
own Nomad cluster locally in no time to test your jobs. If you have a dev or test environment
configured you simply set up the environment variable NOMAD_ADDR and configure an ACL
token for authentication and submit your jobs remotely.

The Nomad CLI also provides handy features like getting logs, exec'ing into a container,
checking the status of a job, and more. If you are interested in tools or frameworks from the
community for Nomad there is [a nice curated list here](https://github.com/jippi/awesome-nomad).

We have adopted Levant for some of our projects and are eager to get projects like Toast
running in our setups, sending Slack notifications to keep our teams up-to-date with changes to
the deployment environment.

With Levant we had a need for mounting config files into our containers. Nomad provides a
template stanza that enables you to create and mount files into containers. Levant can be used
to template whole files with this approach - together this provides a nice workflow as an
alternative to ConfigMaps in Kubernetes and Docker Configs in Docker Swarm.

## Who should use Nomad?

At Verifa, our view on Nomad is that it is one of the most overslept tools we have come across.
Given the state of the current market - with Kubernetes rapidly gaining more adopters even
though it comes with so much complexity which many do not foresee - we think that more
people should consider Nomad as it fills a gap in the market with its main advantage being
operational ease.

Just to be clear, we are by no means suggesting that people should drop Kubernetes in place of
Nomad. However, those teams which have not yet begun their Kubernetes adventures, or those
who only need 10% of Kubernetes and want something with a more thriving community than
what Docker Swarm provides, then Nomad is a great option.
