---
type: Blog
title: What to consider when migrating to Kubernetes
subheading: Let's look at the key considerations when moving existing applications
  to Kubernetes.
authors:
- alarfors
tags:
- Kubernetes
- GitOps
- Open Source
- Containers
date: 2021-09-06
image: "/blogs/2021-09-07/kubernetes-blog-image-01.svg"
featured: true
jobActive: true
---

**You've probably heard of Kubernetes by now. It's the container-orchestration system that everyone is talking about. In this article we'll look at what you should consider when moving your existing application stack to (or developing a new application stack on) Kubernetes.**

## Key Concepts

By the end of this article, we hope to have given you a solid grounding in the following key concepts:

1. When moving to Kubernetes, you do not build the exact same system
2. The difficulty of moving to Kubernetes will depend largely on:
   * The complexity of your Application Stack
   * Where you choose to host your Kubernetes Cluster
3. Simply “moving to Kubernetes” is not the end. You must be prepared to:
   * Design a system that is reliable, maintainable and scalable
   * Operate and maintain the resulting system

### Here's what we'll cover

* What is Kubernetes?
* What should I consider when moving to Kubernetes?
* What improvements can you expect from moving to Kubernetes?
* How do you make the most of Kubernetes?
* When is Kubernetes not the right choice?

***

## What is Kubernetes?

Before we dive in, let's define what Kubernetes is exactly.

> A common definition of Kubernetes is **_a container orchestration system_.**

What are containers? And what does it mean to orchestrate them?

### Containers

Containers are a type of virtualization. A running container virtualizes its own operating system (OS) layer, so that the applications running in the container are (somewhat) isolated from the host operating system. This brings an array of benefits, such as:

* Portability (“Run anywhere”) - the container contains its own dependencies and ensures they are installed and configured correctly
* Consistent environment - the runtime environment of the container is always the same each time it is run. It is (mostly) isolated from changes to the host environment.
* Isolation - the resources available to the container can be isolated from the host environment and provide developers with a safe sandbox environment.

Containers are very commonly used in microservices architectures. A container will typically run a single application, so a microservices architecture based on containers will typically contain multiple containers, which form the complete application solution. When most people think of containers they think of Docker, but Docker is just one option.

### Orchestration

Kubernetes is our Container Orchestrator. We orchestrate our system by telling the orchestrator (Kubernetes) what it is we want to run, including:

* How many of which containers?
* Network interfaces between these containers
* File systems and how they are mounted
* Etc.

This forms our **desired state**. Kubernetes is responsible for ensuring that our desired state is fulfilled, and the desired state is defined through Kubernetes resource definitions. For example, if a container fails, it will be restarted if our desired state specifies a minimum number of running instances which is now not met because one instance just stopped.

***

## What should I consider when moving to Kubernetes?

Now that we have defined Kubernetes, what do you need to consider when moving your existing stack to Kubernetes? Or if you are developing a new application stack on Kubernetes?

### Learning the “Building Blocks” a.k.a. Resource Types

Kubernetes is a framework that allows you to build systems. Although abstract, the Resource Types in Kubernetes are well-defined, which promotes best-practices when designing and constructing systems. It is important to understand the Resource Types in order to build a system that is reliable, maintainable and scalable.

Detailing all the resource types of Kubernetes is outside the scope of this article, but in order to design your cluster and stack correctly, you will need to understand concepts such as:

* scheduling pods which run the containers that you either build yourself or come from a 3rd party (e.g. databases, reverse proxies)
* creating **Services** and **Network Policies** to correctly route network traffic
* using **Persistent Volumes** to manage disk storage availability to **Pods**
* creating **Secrets** to securely manage security-critical data like passwords
* using **Namespaces** to isolate parts of your stack from each other
* designing **Node Pools** and using **Taints/Tolerations** to schedule **Pods** onto appropriate **Nodes**

These are just some examples of knowing the Resource Types available to you and how to correctly use them to design the optimal solution.

### What are you moving to Kubernetes?

> “The difficulty of moving to Kubernetes will depend largely on your Application Stack” is one of the key takeaways from this article.

Many people migrating to Kubernetes are coming from a system running on a “traditional rack”, either physical or virtual servers. Typically there will be multiple applications installed on a single OS, such that they share the OS-level resources like network and storage.

In Kubernetes, applications are typically one-per-Pod. Each Pod can be thought of as an independent VM, and so it is necessary to explicitly design solutions for applications that require communication via network and/or disk storage.

#### Persistent Storage

A common challenge in Kubernetes is how to manage persistent storage. The correct solution will depend on the requirements of the persistent storage and what is available in the service hosting the Kubernetes cluster. Most Cloud hosting providers will offer a series of persistent disk solutions, such as AWS EBS, GCE PD and Azure Disk.

#### Cloud Native - what's that?

Cloud native means to design a system such that it leverages the services offered by cloud providers. A good example of this is external databases. Databases are a common requirement in application solutions, and due to their persistent storage requirements they are unsuitable to run as containers. Instead, for a Kubernetes-hosted application with database requirements, it's recommended to use an external, cloud-hosted database.

This is a great example of breaking out components (in this case, a database) of an application as part of a cloud-native approach. You should identify cases like this in your application early in your migration to Kubernetes.

#### COTS in Kubernetes

Many commercial off-the-shelf (COTS) applications can be run as containers, but some cannot. If you are deploying a COTS application you should investigate thoroughly whether it can be deployed on Kubernetes. This may include not only technical restrictions but also license restrictions. For example, some COTS applications are licensed by fixed-host such as by MAC address or fixed hostname, which can make them very difficult to deploy in Kubernetes, perhaps also violating license agreements.

#### Types of Application Stack

What type of Application Stack are you deploying? Here are a few examples of challenges you might face:

* **End-user Applications**

A typical challenge with end-user applications is that the workloads can be unpredictable and vary greatly depending on external factors. This is a perfect use-case of the auto-scaling and load balancing capabilities of Kubernetes.

End-user applications are often also exposed to the internet, which raises network security concerns. Rather than the traditional “big firewall” approach of a single firewall protecting the application, Kubernetes typically exposes more attack surfaces than a traditional single entry-point. Security is, of course, well thought out in Kubernetes, but it's important to understand what attack surfaces exist, and ensure they are secured when exposing your application to the internet.

These challenges exist whether Kubernetes is used to deploy the application or not. However, it is important to understand how Kubernetes works and how it can be configured for optimal performance and security.

* **Internal Tooling**

Internal tooling typically handles a lot of Intellectual Property, such as source code and software artifacts. It's important to secure these from unauthorized access, so once again security is a major factor in deploying this type of application solution in Kubernetes.

A commonly seen issue in CI/CD systems is unmanaged Artifact Storage; some teams who are inexperienced might rely on the build artifact storage of Jenkins to manage all their build artifacts, for example. When moving such a system to Kubernetes, we should aim for a stateless system with as few requirements on persistent storage as possible. Containers can be thought of as non-persistent, they run when required and might be destroyed when suitable, so we should not rely on their file systems existing post-build. This can, of course, be configured, but always consult best-practice documents. A common solution to this issue is to introduce an Artifact-management system such as JFrog Artifactory.

Another issue often seen in CI/CD tooling is the number of external/proprietary tools being used in a pipeline, such as test tools, coverage analyzers, static analyzers and metrics gathering. Some of these may not work well in a containerized environment. Pay special attention to proprietary tools whose license restrictions may prevent running those tools in a containerized environment, as described earlier in this article.

* **Pure Workloads**

“Pure workload” applications, such as Machine Learning or Artificial Intelligence, are in some ways, probably the easiest to design and implement in Kubernetes. In many cases, the workload can be estimated and as such, auto-scaling and load balancing can be designed to meet the main requirements of this application.

### How will you secure your Kubernetes Deployment?

Kubernetes has been designed with security in mind and it is one of the top priorities of the project. However, building security into your deployment is still something you need to perform yourself - it does not happen automatically!

#### Securing a Cluster

Containerization isolates applications from each other by design, which can improve the security of a solution. Consider, for example, a solution deployed on a single host with an application front-end and a database back-end. There is a security risk where vulnerabilities in the application front-end could be used to access sensitive data in the database back-end due to the front-end processes running on the system which also hosts the disks for the database back-end. However, if the applications were running in containers, the front-end application would only have access to the disk mounted to the container it is running in, thereby reducing the security risk.

Containerization isn't perfect, however. There is a security vulnerability in Docker known as “Container Breakout” whereby a program running inside a container can overcome the isolation mechanisms and gain access to the underlying host of the container. If our program is running as root, this allows the container breakout program to also operate as root on the host operating system! A key rule to follow when deploying in K8s is: **don't run your containers as root.**

Managing your own cluster is one matter, but one of the key features of virtualization is that it allows us to run many systems on a single hardware system. In the case of Kubernetes, this is often seen as many clusters running on some underlying hardware platform. Even computational Nodes, when provisioned by a Cloud Provider, are often Virtual Machines running on a server rack. This means that, for any cluster which shares the underlying hardware with another cluster, the security of each cluster is inter-dependent of the other clusters.

Fortunately, there are control mechanisms built into Kubernetes, and monitoring available as third-party tools which can protect your cluster. If using a Cloud Provider, these will typically be offered to you - study the documentation closely and build in security to your deployments from the beginning. Cloud Providers typically also have sets of Policies you can enable which will help keep your clusters secure.

#### Securing Environments

There are times when you must isolate parts of a cluster; in Kubernetes this is done with **Namespaces**. A Namespace can be thought of as a Virtual Cluster, they cannot be nested within each other and a Kubernetes resource can only belong to one Namespace. A typical use case of Namespaces is that of environments - for example, splitting up development, staging and production is often done with Namespaces. A development environment should always be completely isolated (or as much as it is possible) from a production environment, and the use of Namespaces achieves this.

Most clusters will have a kube-system namespace which contains the Kubernetes resources which operate the cluster, which includes DNS, monitoring and the agent on each Node which manages communication with the API server. So another rule to follow is: **unless you know what you are doing, don't make changes to the kube-system namespace.**

#### Securing Pods (from each other)

It is common to run multiple Pods on a single Node, and a common issue you may encounter is that Node resources run out. You may find that Pods cannot be scheduled as there are no available Nodes. Assuming your system is not under-scaled, the error may be an application with a memory leak, a vulnerable application being targeted by hackers, or simply a poorly designed application.

We can control the resources available to Pods with **Limits and Requests**. By limiting the resources available to a Pod, we can ensure it does not use excessive resources on our Node, thereby affecting the other Pods sharing the Node. Limits control the maximum resources we allow a Pod to use, whereas Requests provide Kubernetes with the minimum resources our Pod must have available. Requests can be a dynamic way to ensure that Pods are correctly scheduled across cluster Nodes.

Consider a two-Node cluster which must run three Pods, two small and one large. The Nodes are sized such that the system can only run with two small Pods on one Node, and the large Pod on the other Node. If the two small Pods are assigned to a Node each, the big Pod cannot be assigned as neither Node now has enough available resources. By specifying a Request for resources for the big Pod, it is scheduled to an empty Node before the smaller Nodes are scheduled. (Note: this example is greatly simplified, but illustrates the problem. Taints and tolerations are the more common way to manage what-goes-where).

We have covered security in Kubernetes clusters only very briefly here, and only discussed points which should apply almost universally, regardless of the solution you are building. Be sure to study your system closely and implement appropriate security measures.

### Where will you host your Kubernetes clusters?

If you are looking to deploy your first Kubernetes cluster, chances are you will either deploy it:

1. In the cloud, with a Cloud Provider like GCP/AWS/Azure
2. Locally (self-hosted) using Minikube

For production-scale deployments, there are a few considerations to each platform.

#### Cloud-Hosted

When deploying to Kubernetes in the cloud, the Cloud Provider (GCP/AWS/Azure) will take care of the basic cluster setup and operations for you. Their Kubernetes services (GKE/EKS/AKS respectively) create the backbones of the cluster (control plane, DNS etc.) that you would otherwise need to build yourself.

This is great, as it means we can focus on building our solution, but Cloud-hosted Kubernetes does come with some disadvantages:

* Your technical configuration is more restricted than in a self-hosted setup; Cloud Providers will restrict access and/or functionality within the cluster. For example, some Cloud Providers prevent SSH access to the Node running the Control Plane, or they restrict which type of Operating System can be run on the Cluster Nodes.
* A cloud provider will restrict you in choosing the exact locale of the physical machines operating your cluster. You may be able to select a region or country, but not which city or which building the machines are located in.
* A cloud provider will also restrict your options as to what else runs on the same underlying hardware. You may be able to rent dedicated hardware for your cluster, but if you use the standard cluster operations of a Cloud Provider, chances are there will be other workloads from other customers sharing the same underlying hardware.
* The type of solution you are building may restrict you from placing the workloads/data in the cloud - this is something you need to investigate in your project.

#### Self-hosted

Self-hosting a cluster is considerably more work, both initially and in maintenance, than using a Cloud Provider's Kubernetes service. If you're required to take this approach because Cloud Hosted Kubernetes is not an option for you, it's recommended you look at tools and frameworks that help you setup and run your Cluster, such as:

* [Rancher](https://rancher.com/  "Rancher website")
* [KinD](https://kind.sigs.k8s.io/  "kinD website") - most suitable for testing, prototyping & local development

A great read if you're self-hosting Kubernetes is [“Kubernetes the Hard Way”](https://github.com/kelseyhightower/kubernetes-the-hard-way  "Github link")

***

## What improvements can you expect?

So far we have mostly discussed thresholds and obstacles you can come to expect when building a solution in Kubernetes, but Kubernetes wouldn't be as popular as it is if it did not provide benefits! You may have a clear idea of the improvements you seek already, but let's briefly cover the general improvements anyone migrating to Kubernetes can expect.

### 1. Uptime & Reliability

Have you ever worked with an unstable application which would sometimes enter an unresponsive state and require restarting? Perhaps the root cause was difficult to identify and it was just simpler to restart the server and/or application every X months to keep the system online. Java applications are a common culprit of this, they can reach their JVM memory limits and stop responding or behaving as expected.

Kubernetes has features which deal with this type of issue. Health checks can be configured to automatically monitor the state of an application (for example, the return code when making an HTTP request) and if the health check fails, the Pod is killed and a new Pod created. The result is the system self-healing by destroying unresponsive parts of it and replacing it with new ones. Note the importance of Pods being ephemeral (the opposite of durable) to support this type of functionality. Any Pod should be able to be killed and replaced with a new one.

### 2. Traceability

Kubernetes lends itself very well to defining a solution as code (“configuration as code” and “infrastructure as code”) - and code which is version controlled becomes very easily traceable. Changes between versions can be easily identified, and should be associated with a description of the intended change and the author of the change. Add a workflow to this and changes to the system become well-managed, with development work performed in branches and not being committed to the production system until pull requests are approved into the master branch.

Kubernetes keeps a history of the state of the system, meaning you can access a list of events, view failed Pods, access system logs etc. With each part (Pod) of the system being well defined, faults in the system can be traced to the root cause of the failure.

### 3. Scalability

One of the major changes which occurs when moving to Kubernetes from a traditional system is decoupling; hardware (even virtualized hardware) is decoupled from applications, applications are decoupled from each other and so on. This decoupling makes scalability much simpler than in a tightly-coupled system. Nodes can be increased in size to offer more system resources, computational Pods can be increased in number (replicas) and load balancers can direct traffic to where the load is lowest.

### 4. Vendor Agnosticity

Kubernetes is a kind of framework in which you can define your system. Although developed by Google, Kubernetes is vendor-agnostic, and it is up to each platform vendor to support it. As a result, you may build a system in Kubernetes and have very little work required to move it from one cloud provider to another. This prevents vendor lock-in.

### 5. Maintainability

As with traceability, with everything defined as code, maintainability becomes much easier. Changes to the system are easier to develop, test and execute, and can be more easily shared with a team of administrators.

Once trained in Kubernetes, an administrator can quickly understand how a system is built. This makes knowledge transfer and long-term support of a system easier. Given a completely unknown system to us, we can with a few calls to kubectl get a complete description of all parts of the system, from computational Pods to Network to Persistent Storage.

### 6. Community Support

When building in Kubernetes, we are building on layers of open-source software. Indeed, Kubernetes itself is open-source and therefore maintained by the open-source community. When we create Pods, they run Containers which often have an open-source image as a base (usually some variant of Linux). Common applications like nginx, logstash and tomcat are all provided as ready-to-go container images. If we use Helm Charts, then we have a huge selection of applications and their configurations which we can simply pull and use.

The result of this is that most of the solution we construct is coming from the open-source community, meaning it has been tried and tested, by thousands, if not millions of users already. For many open-source projects there is a strong community of developers, releasing regular updates both for feature, performance and security improvements. With our solution built in Kubernetes, making use of this software and these updates is easier than ever before.

***

## How do you make the most of Kubernetes?

Making the optimal configuration in Kubernetes will very much depend on the type of application you are building, but there are some general rules that almost all application types can follow.

### Use Auto-scaling

With our applications decoupled from our hardware, and scaling easily available to us, we should use the auto-scaling features offered whenever possible and suitable. By scaling up a system at peak loads, we can offer optimal performance for users, whilst scaling down as load decreases ensures we do not pay for excessive resources.

### Use Deployments

One of the main types we should be using when building a solution in Kubernetes is the Deployment type. Rather than defining individual Pods, with a Deployment we define a desired state of the system, and Kubernetes manages the resulting Pods for us. We can roll back to a previous deployment, or when updating we can use the rolling update feature to ensure 100% uptime.

### Control Security

IT systems have traditionally been built with a firewall blocking network access between a sensitive and non-sensitive part of the network, such as the internal corporate network and the internet, for example. In most cases, the overhead of managing a firewall between each application within a network has been too great, and so within each network security zone all traffic has simply been permitted.

With Kubernetes, our system is structured into individual pieces (Pods) and we must define the network interfaces between these. Rather than being an overhead, this is an opportunity to strictly control the network traffic between our different applications. Using Ingress and Egress rules we can explicitly define who-should-talk-to-who-about-what in our Cluster. Use this to your advantage to build a system that is secure and does not depend on a single firewall, which once breached, leaves the complete system exposed.

### GitOps

GitOps is a development workflow coupled with a system defined as code. We have covered the topic of GitOps at length in our blog [“GitOps: All you need is Git (and Continuous Deployment)"](https://verifa.io/blog/gitops-all-you-need-is-git-and-continuous-deployment/) and [“What is GitOps anyway?"](https://verifa.io/blog/what-is-gitops-anyway/) podcast, but to summarize:

* The complete system is defined as code
* Changes to the code (system) are made on development branches
  * These changes are typically tested and staged in separate environments to the production environment
* Once approved, “staged changes” are merged into the master branch via a pull request
* Whenever a change occurs on the master branch (and they only occur when a pull request is merged), an operator deploys the latest version of the system to production

With this workflow, we ensure that changes to the system are tested and well-managed, and also that the master branch always reflects what is currently deployed in our production system. GitOps can be enormously helpful to development teams, it helps define the workflow and removes ambiguity in the system (“what is deployed where?") - but it is not for everyone.

***

## When is Kubernetes not the right choice?

There are, of course, some situations where Kubernetes is not the right choice for you.

### Low ROI

There is an investment required to develop something in Kubernetes. If the return on this investment is low, which is typically if you do not benefit from the improvements Kubernetes can bring to a system, then Kubernetes is probably not the right choice for you.

### Cannot containerize

If you have applications in your stack which cannot be containerized, then Kubernetes will not be a good framework in which to build and deploy your stack.

### Guaranteed latency, RTOS applications

Real-time Operating Systems (RTOSs) are systems designed to process data within very strict time measures. Although it may change in the future, Kubernetes is currently not suitable for applications such as these.

### Connected hardware

Hardware-in-the-Loop (HIL) tests are very common in the embedded software industry, and are an example of when you may need a piece of hardware connected to your system. Although it is possible to have connected hardware in a Kubernetes solution, you should consider carefully if Kubernetes is the right framework for your solution, as it may introduce complexities which make your system worse, not better.

### Low-level systems engineering - Kernel interaction

If your work requires you to interact with the Kernel of the operating system, Kubernetes and other containerized systems are not a suitable solution.

***

## Conclusion

Hopefully this article has helped you break down the key considerations when moving to Kubernetes, the main benefits, as well as when Kubernetes is not the right solution. Below are further resources to help you in your Kubernetes migration. Alternatively, you can [get in touch with us](/contact) should you or your team need any help with Kubernetes. Looking forward to hearing from you!

## Further resources

[Kubernetes website](https://kubernetes.io/)

[Kubernetes set up docs](https://kubernetes.io/docs/setup/)
