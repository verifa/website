---
type: Blog
title: A sneak peek into UpCloud Kubernetes Service
subheading: In this episode of The Continuous & Cloud Podcast, Jacob and Lauri chat with Ville Törhönen, Product Owner at UpCloud, about UpCloud’s upcoming managed Kubernetes Service (UKS).
authors:
- jlarfors
- lsuomalainen
tags:
- Kubernetes
- Cloud
- UpCloud
date: 2022-08-16
image: "/blogs/podcast-ep04-a-sneak-peek-into-upcloud-kubernetes-service.png"
featured: true
jobActive: true

---
<iframe title="Embedded podcast player" src="https://anchor.fm/verifa/embed/episodes/A-sneak-peek-into-UpCloud-Kubernetes-Service-e1miudr" height="151px" width="100%" frameborder="0" scrolling="no"></iframe>

<div class="blog-flex">

[![Listen on Spotify](/blogs/2021-03-30/listen-on-spotify.png)](https://open.spotify.com/show/12yStrneLdEsXn1Bjp6Myz)

[![Listen on Apple Podcasts](/blogs/2021-03-30/listen-on-apple-podcasts.png)](https://podcasts.apple.com/gb/podcast/the-verifa-podcast/id1561051552)

[![Listen on Google Podcasts](/blogs/2021-03-30/listen-on-google-podcasts.png)](https://www.google.com/podcasts?feed=aHR0cHM6Ly9hbmNob3IuZm0vcy81Mzg0NzE1Yy9wb2RjYXN0L3Jzcw==)

</div>

**Welcome to The Continuous & Cloud Podcast by Verifa, where we chat about continuous delivery, cloud architecture and most things inbetween.**

In this episode of The Continuous & Cloud Podcast, Jacob and Lauri chat with Ville Törhönen, Product Owner at UpCloud, about UpCloud’s upcoming managed Kubernetes Service (UKS).

## During this episode we discuss

- What is a managed Kubernetes service? \[00:39\]
- What is UpCloud Kubernetes Service (UKS)? \[03:10\]
    - Motivation behind UKS’ development
    - What’s under the hood?
    - How to create a Kubernetes cluster in UKS
- What do you get from UpCloud that supports the Kubernetes service? \[06:02\]
    - What do you get ‘out of the box’?
    - What can you customise?
- How has it been developing UKS? \[13:10\]
    - Timescale and team
    - Kubernetes community and open source projects
- Observability and monitoring in UKS \[20:42\]
- What are the main benefits of using UpCloud and UKS? \[23:51\]
- When will UKS be available? \[25:40\]
- Example use case: using UKS with Terraform for IaC \[27:11\]

## Mentioned in the podcast

[UpCloud Kubernetes Engineer vacancy](https://jobs.upcloud.com/o/kubernetes-engineer)

[UpCloud CSI driver on Github](https://github.com/UpCloudLtd/upcloud-csi) 


## About Ville Törhönen

Ville (MSc, Tech) is working as a Product Owner at UpCloud. He is responsible for the upcoming UpCloud Kubernetes Service (UKS) product offering but also Developer Experience on the platform. Previously he’s been building various businesses and products on top of container technologies in companies like Unity, Noice and Redhill Games.

## About UpCloud

[UpCloud](https://upcloud.com/) is a European Cloud Provider offering fast and reliable PaaS from 12 data centers around the globe. They are a strong believer in open source and openness in general. That’s why all of their services are available from [public repositories](https://github.com/UpCloudLtd). UpCloud’s mission is to make Infrastructure as Code easier for small and medium size businesses.

## Connect with today’s podcast crew on Linkedin

[Ville Törhönen](https://www.linkedin.com/in/vtorhonen/)

[Jacob Lärfors](https://www.linkedin.com/in/jlarfors/)

[Lauri Suomalainen](https://www.linkedin.com/in/lauri-suomalainen/)

## Transcript

### Hello and welcome!

**Jacob** \[00:15\]: Welcome to the Continuous and Cloud Podcast by Verifa. I'm your host, Jacob, and my co-host today is Lauri.

**Lauri** \[00:21\]: Hello.

**Jacob** \[00:22\]: And our guest today is Ville Törhönen from UpCloud.

**Ville** \[00:26\]: Hello.

**Jacob** \[00:27\]: And he's product owner of the product we're going to discuss today, which is the UpCloud Kubernetes Service.

**Ville** \[00:33\]: Yes. And thank you for having me.

**Jacob** \[00:35\]: It's great that you could be here. So I think let's just start, shall we?

**Lauri** \[00:38\]: Yeah, let's do it.

### What is a managed Kubernetes service?

**Jacob** \[00:39\]: Cool. So UKS, the short term UpCloud Kubernetes Service is a new service coming soon from UpCloud. And let's maybe start by talking about what Kubernetes services are, or managed Kubernetes services. So maybe, Ville, you'd like to tell us what your idea of managed Kubernetes services?

**Ville** \[01:02\]: Oh, of course. Well, the first thing that people want is Kubernetes, then the second part is that they want it running somewhere and then running it by themselves is a tedious job of its own. Often you need people who have a lot of experience running Kubernetes. You have a lot of people who have used Kubernetes, but not that many who have actually managed to run it from scratch or in production or in a larger scale. So what most customers or most users want is a managed service just because it's a hard thing to run.

Same applies to any of the popular sort of managed services such as databases or serverless functions or whatever, and the same thing applies here as well. And how we ended up developing any product for Kubernetes was that there was a lot of, or there is a lot of, customer demand for it. There are other ways of orchestrating containers, but obviously the most traction is in Kubernetes and the ecosystem benefits that it brings are so enormous that we thought that this is the product to be part of.

So we've been developing it from early, well, late last year, we did some research how we will actually start implementing it and we've been now doing it for this year, and we are aiming to do a launch by the end of this year, so that's the short story of it.

**Jacob** \[02:23\]: Awesome. Thank you. And just for clarity as well, this is being recorded now in, where are we? Beginning of August 2022. So maybe if you're listening to this in the future, UKS is already out, but if you're listening to it as soon as we publish, then something to look forward to.

**Ville** \[02:38\]: Yes.

**Lauri** \[02:40\]: Yes. And if I correctly understood, it's in pre-alpha, in closed alpha, right now?

**Ville** \[02:44\]: Yeah. So we've been working with our existing customers, such as you guys, developing the product. And one of the key elements that we use is a user centered design, so everything we do is based on customer use cases and we implement them then on smaller increments and smaller releases. So we work on an agile mindset, if you will.

**Jacob** \[03:08\]: Yeah. That's really great.

**Ville** \[03:09\]: Yeah.

### What is UpCloud Kubernetes Service (UKS)?

**Jacob** \[03:10\]: Yeah, so managed Kubernetes is you basically get a Kubernetes cluster and then it's managed by UpCloud, and I guess the two main components within Kubernetes cluster is the control plane and the worker nodes. So everything's in UpCloud, right?

**Ville** \[03:26\]: Yes. Everything is in UpCloud and we take care of the hard part, which is the control plane and everything related to running it and maintaining it in an operable state. And what customers or users will have is the... or the responsibility is the worker nodes, so the nodes that will actually run your workloads. And how you want to run it is up to you, you define what kind of machines you want to use and how many and all that jazz.

**Jacob** \[03:53\]: Awesome. So let's maybe talk a bit more about what's under the hood then. So let's say we want to...

**Lauri** \[04:01\]: Obviously create a cluster and the communication layer between that and so on. So what do you get out of the box?

**Ville** \[04:12\]: Good question. So another team that I'm a product owner for is Developer Experience and it goes really well hand-in-hand with this stuff. So how do users actually create a cluster? How do they maintain it? And what's the sort of day-to-day workflows and operations related to it? And you have to be flexible. We can't have one single way of creating a cluster because as we've seen with just virtual machines, there are different flavours of how people will deploy them.

Some want to use a comment line tool, someone might want to use something like Ansible or config management systems to do that. Some will want to use a more refined, let's say, infrastructure as code approach, where they maintain a state, such as in something with Terraform. And there are a plethora of ways of doing that. Some might even want to use UI for it, if you're really up for that stuff. But the main idea that we have is that you can create it in a number of ways. We will support most of the workflows, most common workflows, so pretty much what I just mentioned.

We do support UI workflows for creating clusters, of course. Not everyone wants to or has to have infrastructure as code if they really don't want to. So as long as they have a cluster up and they can connect it with any standard tooling, then that's what we support. So we take care of the control plane, you run the cluster and do whatever you want with it. So it's pretty much what other, let's say, competing Kubernetes as service approaches do. So yeah, you have a lot of flexibility and power in your hands.

**Jacob** \[05:45\]: That's great. So you get the Kubernetes API exposed and you can start using kubectl or whatever things you want to do with your Kubernetes cluster.

**Ville** \[05:54\]: Yes.

**Jacob** \[05:55\]: Start getting excited just thinking what I would do with a fresh cluster.

**Ville** \[06:00\]: Giving you ideas.

### What do you get from UpCloud that supports the Kubernetes service?

**Jacob** \[06:02\]: Exactly. Yeah. But maybe this is a good time to talk about the other things UpCloud a little bit too, because it's never just Kubernetes. You might have the cluster and now you can run your pods there, create your services, maybe deploy an ingress controller and get access to those. But then there's things like storage, for example, and well, probably running a database within Kubernetes isn't such a great idea unless you know what you're doing, and then there's maybe some fast system storage that you want as well, like persistent volumes and these sorts of things. Yeah, could you tell us what's available in UpCloud?

**Ville** \[06:39\]: Sure. So as we are developing a new product, the first and foremost we need to match a certain level of standards. So even though we have our own, let's say, targets of adding new features or new ways of working on top of Kubernetes, the absolute minimum is the standard set of features like you mentioned, persistent volumes, exposing services from the cluster, connecting to the cluster with standard tooling and stuff like that.

So for volumes, for example, we went ahead and implemented a container storage interface driver. So we have our own UpCloud CSI driver, which you can... or which is actually deployed automatically to the cluster. And it will handle all the persistent volume-related operations for you. So you can run a database in Kubernetes, I think that's totally fine. I do it myself as well, so.

**Lauri** \[07:31\]: That's an interesting part. Obviously, when you are building a service on top of another service that would be the app cloud, you're going to be on top of that, but sometimes you have to change the underlying infrastructure as well and underlying systems. So was this one part of it, having the volume's interfacing with it correctly, or are there any other things you have had to change from the underlying system?

**Ville** \[07:58\]: So actually from CSI perspective, we just thought that it's the only possible solution because other container orchestration systems have the same or they support the same interface. So now that we have implemented it once and it works for Kubernetes, it also works for other container orchestration systems, such as Nomad, which is something we've also tried and done, so get basically support for any container orchestration system by adding that driver. And it's still in the works. It's an Open Source project. You can go to our company GitHub org page, and you can find the source code. And it's being worked on, as we add new features to our storage backends we will also add then to the CSI driver, so it's a ongoing work as we speak.

**Jacob** \[08:47\]: Nice. And there is a managed database service in UpCloud as well.

**Ville** \[08:52\]: Yes.

**Jacob** \[08:52\]: Which is what I was trying to get out of you about not wanting to run a database in your Kubernetes cluster. Well, you don't have to because...

**Ville** \[09:02\]: Yeah, of course. Well, I can be the sales guy here as well. Yeah, of course. As I mentioned earlier, it's the same thing as with Kubernetes. So as we have managed services for a number of other things, you probably want to run your database also as a management. Same goes for load balancing because that's also another hard problem area of it, so often people want to off-load that to the providers.

**Jacob** \[09:26\]: Cool. So do you have an ingress controller that supports UpCloud now? Or is that... I mean, this isn't like a good way or a bad way of doing it because if you use the ingress, like an NGINX Ingress Controller, and you deploy a service type of load balancer in Google Cloud or EKS AWS, it'll create the load balancer for you. And of course that's the way you do it on day one, but at least from my experience, usually we manage the load balances with Terraform at the end of the day, anyway, because it gives you much better flexibility than doing it with Kubernetes. So I guess there's no right or wrong, but do you have that kind of thing yet in UKS to make it really easy for somebody to start with a...

**Ville** \[10:10\]: So not going too much, too deep, into the details. But yes, we are supporting exposing the services in a number of different ways just to make it easy to make your stuff running and then expose it to outside of the cluster. What comes to sort of deploying load balancers, let's say, over Terraform and then hooking that up into the cluster, I think that's somewhat of an anti-pattern because you already have Kubernetes, you can already define objects in there, so why would you do it in two separate places? So I think what's common in some Kubernetes setups is what you just described, but I think we can do better and make it a more of a developer-friendly approach.

**Jacob** \[10:58\]: Yeah.

**Lauri** \[10:58\]: Speaking of availability and resilience, obviously Kubernetes can autoscale on a pod level, I guess that you have implemented the horizontal autoscaling when it comes to nodes. And so what zone do you have available and can you tell us a little bit more about that? Do you have a vertical autoscaler there in the works?

**Ville** \[11:23\]: So what I can say is that we will support autoscaling, it's part of our plans. As for regions and high availability, there will be a number of different kinds of, let's say, flavours for you to choose from, but we don't have any preferences on the regions. So all the regions that we have will be eventually supported. So maybe if someone is, for example, on the journey of... in our beta phase, for example, you might not have all the zones, but eventually we will have all of them.

**Jacob** \[11:58\]: That's great.

**Ville** \[11:58\]: Yeah.

### Customising resources within UKS

**Jacob** \[12:00\]: Cool. And that nicely leads on to customisability as well because that's... well, we already talked a little bit about the storage classes and, well, at least persistent volumes and a little bit about the networking with load balances, but what about other customisations? One obvious thing that comes to mind is the worker nodes. So what kind of nodes do you want to have in your cluster? Are there like node groups or node sets or are they individual nodes and how... Yeah, what can you do?

**Ville** \[12:28\]: Yeah, we are looking into... we are exposing the node groups as a separate API object, so you will be able to modify your workers accordingly or according to your workloads. So you might want to have separate kind of, or different kinds of nodes for running stateless workloads or some other specific requirements and then have smaller or, whatever, bigger nodes to run your other, let's say, stateful or stateless workloads. So it's configurable as you would expect from a managed Kubernetes service. So the way I see it, it's a basic requirement to make such thing work definitely.

### How has it been developing UKS?

**Lauri** \[13:10\]: So how long had you been building the product and how many of you are in the team now?

**Ville** \[13:20\]: So as mentioned, we started the research work last year, started building the product early this year, and we have a team of four developers now working full time on it. Of course, since we are a small company, it's not just this team building it, we have plenty of others because it's very much tied into our existing infrastructure and the ways we do things there. So we are using other teams of course, it's not just four people, it's a bigger group of people, but in the end, the team is a four man team.

**Lauri** \[13:53\]: Right. How's the work been up to this point? It sounds to me that you have built this whole thing up in a relatively short time, which is quite impressive.

**Ville** \[14:04\]: Yeah. One of the things we've enjoyed is that the tooling and the community around Kubernetes has grown quite significantly over the last few years, so there's good community support on a number of things. In addition, the projects that revolve exactly around provisioning clusters or configuring them is quite big. So things like cluster API, which I think went into GA a few months ago, or in KubeCon in May, is one really good example of what kind of community pushes towards a single problem in Kubernetes. So even though we are a small team, we've still been able to create a product in a relatively short time because of all the existing or helping tools that's out there.

**Lauri** \[14:52\]: Sure. This is also an Open Source product, right?

**Ville** \[14:57\]: Eventually, yes, it will be Open Sourced. So as with other things such as the CSI driver I mentioned, it will be, will be Open Source as well. Yeah. We like Open Source.

**Jacob** \[15:09\]: Yeah. Well, that's nice. I mean, it's good that there's so many projects around and so much community around it, but it's also overwhelming a lot of the time when you look for something, like maybe we could talk about the networking, the container network interface that you chose. Yeah, if you can talk about that. There's a lot available. I mean, AWS developed their own, the AWS VPC CNI. There's Calico, which is very popular. And now there's the newer blocks on the kid as well... the new kids on the block, like Cilium.

**Ville** \[15:43\]: Cilium. Yeah. So we went with Cilium, of course we might support other ones, but just for, let's say, for performance reasons, we went with Cilium. And we've loved it so far, there are really cool things coming from that project. And it's something we are really having a close look on, close eye on all the time. So it's a good example of another good Kubernetes project that has solved problems for many. It's not just us that's using... I think there are other providers who also went with Cilium or at least a flavor of Cilium, to do things.

**Jacob** \[16:23\]: So is eBPF a common coffee break chat?

**Ville** \[16:29\]: Yeah. We've had some chats about eBPF. Oh, for sure. Yeah. It's a good thing. We could also talk about service mesh, but maybe that's for another chat of itself.

**Jacob** \[16:38\]: Yeah. I feel like we could do a whole podcast series about service meshes, but did you have anything in particular about service meshes? I guess, in the interest of Cilium, that you can build... if you have multiple clusters, they could all be connected on a service mesh fairly simply then, or?

**Ville** \[16:57\]: Yeah. Yeah, you can do things like that. And it's a really interesting piece of tech. And we had a federated DNS in Kubernetes for quite some time, but it's kind of different than what you can get if you have the clusters interconnected on a network level, so it's a good piece of tech.

**Jacob** \[17:18\]: Well, that's cool. So let's say you have a team or many teams working on, let's say, a collected product built up of multiple services, and you're basically doing multitenancy in a way, even though they need to talk with each other in the cluster, you could now break it out and give them each their own Kubernetes cluster and have them connected on the service mesh rather than one mammoth, one cluster to rule them all and managing RBAC and everything very closely.

**Ville** \[17:43\]: That's possible. That's also already possible on pretty much any Kubernetes installation. And I would say that most companies have more than one Kubernetes cluster, and one way or another, they will have to talk to each other. Oftentimes having a sort of on a port-level ports talking across networks from one cluster to another doesn't make sense in, let's say, most cases. Oftentimes you just might want to expose the service from another cluster somehow to the other one, either through internal networks or public network, depending on the setup, of course.

But there are other good use cases for such, let's say, service mesh approaches even inside a single cluster, having more visibility or observability into the operations, what's going on and getting retries out of the box and offloading the port-level networking to a separate agent and getting mutual TLS. The list is endless. And I think, even though not related to this product at all, I think it's solving real world problems, so.

**Jacob** \[18:50\]: Yeah, but this is really interesting because one of the things I wanted to get out of this today as well is, I mean, if I had to build a Kubernetes service, all the extra things that you could do with minimal effort, I think are really interesting because of the community and the Open Source projects around it. And this sounds like exactly one of those things, with Cilium now kind of like... I don't think you can do this with many of the other CNIs? I know Calico has a lot of extra features that you can use. I've only used it for very simple use cases. But I know that we had this case recently where I needed to have clusters talking to each other with Calico and there isn't a built in way of doing it, so you end up needing something else on top.

**Ville** \[19:33\]: Yeah, other options include things like Istio, for example, or Linkerd that pretty much offer you the same thing, except it's on a different level. It's not a CNI network level thing, it's rather usually an agent that helps you out and sort of takes care of the networking part, and you offload it basically to a side car.

**Jacob** \[19:57\]: Yeah, which is great. And Cilium and eBPF... Here we go, eBPF again, makes that even more interesting because it's sidecarless, isn't it?

**Ville** \[20:10\]: Yes, indeed.

**Jacob** \[20:11\]: Yeah.

**Ville** \[20:12\]: Even though the side cars might be written in Rust and the memory footprint is negligible, but still it's a nicer way of doing things if you don't have to use sidecars.

**Jacob** \[20:22\]: It's an extra port, and for everybody who's been working with AWS EKS and the default AWS VPC CNI, might have been running out of IP addresses and knowing that every port needs an IP address. The fewer ports the better in a way, yeah.

**Ville** \[20:38\]: Yep. Yep.

**Jacob** \[20:39\]: So there's definitely a benefit there.

### Observability and monitoring in UKS

**Lauri** \[20:42\]: Speaking of observability, let's say I haven't had that much experience in UpCloud and monitoring and so on, so if I was to have a UKS cluster and I wanted to set up alerting and monitoring for it, what would be my first go-to then?

**Ville** \[21:03\]: Well, pretty much the same as you would have in any other cloud. So people have different needs or, let's say, requirements for having monitoring and alerting, and for a cloud provider it's a place of... hard to be sort of opinionated on. So that's why having a generic thing that anyone can deploy whatever they want in there makes a lot of sense. But this is another thing where the community comes in and helps because for, let's say, the basic set of having Kube-Prometheus-stack running on a Kubernetes cluster is really fast and easy to set up, and you get Prometheus and Grafana and Alertmanager and all that stuff, out of the box, and then you have really good visibility on the cluster.

Another thing that even though we're offering just the cluster for our users or customers, it's such a flexible thing that people can set up even their most sort of advanced set ups on top of it, and we make it possible.

**Jacob** \[22:07\]: Well, basically you give people a Kubernetes API, which is the, as they're saying now, the cloud operating system.

**Ville** \[22:15\]: Yeah, you could call it that. Yeah, sure.

**Jacob** \[22:16\]: Yeah. Yeah. I think it's an interesting way of phrasing it because it does... Yeah, it's not just Kubernetes anymore. And actually, I want to talk with you later about the load balancer Terraform discussion, because I think it's... maybe if we have enough time at the end of this, because I think it's an interesting topic and there are definitely use cases where you need to create the load balancer with Terraform and... Yeah.

**Ville** \[22:39\]: We can talk about that.

**Jacob** \[22:40\]: We can do. Yeah.

**Lauri** \[22:41\]: Speaking of Terraform, UpCloud has a very robust Terraform provider and I guess that's going to release with the UKS? It's going to have some modules for that too?

**Ville** \[22:56\]: Definitely, yeah. Yeah. For all the new products that we're working on or have released within the last year or so, each one of them will have Terraform support. Currently all of them are supported. For new products, of course, we will bring in the Terraform support as well. There's huge demand for it and it's something people expect from us already.

**Jacob** \[23:20\]: Yeah. Yeah. But it's not just Terraform that you cater for when you build a provider nowadays, because many of the other IAC tools, like Crossplane and Pulumi and so on, basically build off the Terraform providers, so.

**Lauri** \[23:31\]: But you could go a step back, that you nowadays can't build a product without a proper API to interact with.

**Ville** \[23:38\]: Sure, sure. And you also have Terraform CDK now, which is also something we support or have bindings for out of the box, so it's another cool automation built on top of the providers.

### What are the main benefits of using UpCloud and UKS?

**Jacob** \[23:51\]: Nice. Nice. And I have to praise UpCloud a little bit too, because having worked with the other cloud providers and the Terraform providers for them, when I work with UpCloud and I run Terraform plan or something, I think it's failed because it's like... I press plan and like, "There's the plan." And I'm like, "Okay, so something went wrong," and it's like, "Oh no, it actually planned in..." It's a matter of seconds and you get the feedback. So if your Kubernetes service is going to be on that speed as well, I think that's going to be nice. And maybe this is a perfect time to ask as well, quite a way into the podcast now, and we've gone quite deep into the discussion about UKS, but if you had to tell somebody why they should consider UKS, what would be the main points that you would-

**Lauri** \[24:43\]: And UpCloud in general.

**Ville** \[24:44\]: Yeah, maybe UpCloud in general is a good starting point because we're local but global and that's one thing that really sums up the situation quite well. Data locality and especially GDPR, Data Server NT in general, are things that are one of our strongest points. That combined with good performance. And it gives you a really robust cloud platform to run your infrastructure on, whether that's Kubernetes or databases, you get a good platform to run things on. So as the platform works for databases and virtual machines, then obviously the same platform will... we're building the product on the same platform as virtual machines are, or other products that we have. So nothing really changes, you get the same benefits as with any other products we have on our plan.

**Jacob** \[25:37\]: Great, cool.

### When will UKS be available, and what’s next?

**Lauri** \[25:40\]: So you guys are going to release at the end of this year, hopefully, and what then? What happens after?

**Ville** \[25:53\]: Maybe too premature to talk about what happens after, but-

**Jacob** \[25:57\]: Bug fixing.

**Ville** \[26:00\]: Bug fixing. Yeah, of course. Bug fixing. And we are improving things. We have plenty of things in our books. We are recruiting, there's two new positions open probably for... if you're listening to this in August 2022, then the position still might be open, but maybe not later on. So we are growing a lot, recruiting people, building, starting with Kubernetes but there's a lot of cool things coming afterwards as well. The fun doesn't stop here.

**Jacob** \[26:29\]: Yeah. No, that's really excellent to hear. And I think upkeeping a Kubernetes service is a challenge in itself because it's moving fast.

**Ville** \[26:38\]: It's very fast moving, yeah. And each release has some new cool feature that everyone wants, so us keeping up with the pace and allowing people to have the latest and greatest makes a lot of sense.

**Lauri** \[26:52\]: And if people want to get involved, like as a client or an Open Source contributor, that's probably going to happen in near future?

**Ville** \[26:59\]: Yes. Yes. As soon as we have the product out then yes, in an Open Source manner for sure.

**Jacob** \[27:05\]: Great.

**Lauri** \[27:08\]: Great to hear.

**Ville** \[27:09\]: Yes.

### Example use case: using UKS with Terraform for IaC

**Jacob** \[27:11\]: Is there anything that you want to talk about more, Ville, that you feel we haven't covered already?

**Ville** \[27:17\]: Maybe something about the load balancers. I'm curious to hear your story on load balancers from Terraform.

**Jacob** \[27:22\]: Yeah. Well, maybe instead of just asking the question outright, we can now cover a kind of like, "I have my product and I want to use UKS." Let's say it's some customer-facing SaaS product.

**Ville** \[27:39\]: SaaS app?

**Jacob** \[27:39\]: SaaS app, yeah.

**Ville** \[27:40\]: Dot fi.

**Jacob** \[27:41\]: Yeah. Is that domain taken? Because after this it's going to go viral.

**Ville** \[27:45\]: Yeah, it's going to go...

**Jacob** \[27:45\]: Yeah. Okay. So I buy my domain, configure my DNS to point to... well, I need to create the thing that is going to point to, which is going to be the load balancer. And I have my containers and I have my CI and everything ready and it pushes stuff. So what I want to do is I want to create... well, let's maybe say we have this use case where we want to have multiple clusters, maybe because we have different teams working in different... quite separate teams that need to talk with each other, the services need to talk with each other, so we create multiple Kubernetes clusters or UKS clusters.

**Jacob** \[28:19\]: Let's say we do that using Terraform. And now we want to have, well, maybe one or multiple load balancers to then hook up our DNS so when somebody goes to, was it SaaSapp.fi? It's going to push them towards the load balancers and the load balancers are going to send them to our, presumably, an ingress controller running on the correct cluster, because the load balancer would be responsible for forwarding the request to the correct cluster.

**Ville** \[28:48\]: Sounds like a very cool setup.

**Jacob** \[28:52\]: Yeah. Sounds like exactly the setup you need. Yeah. So I mean, I think this load balancer to cluster is a really interesting topic because there's a lot of discussions that I see online, a lot of GitHub issues that I end up scrolling through like 50 plus comments with everybody's interesting hacks and different solutions. So yeah, I guess the alternatives for doing it are that you create the load balancer with Terraform and then you create the cluster. And then when you deploy your ingress controller, you need to point the load balancer, basically target groups or that kind of concept.

So those need to point to the pod, to the service, and you can do that via NodePort, for example, declaring the service as type NodePort, and then you could point to the nodes and the ports on there. Or you could do it, I guess, like the cloud specific native way, like in AWS or Google Cloud and maybe now in UpCloud as well, you can define your service as type LoadBalancer, and it basically takes care of that.

But in the case that you have, let's say, one load balancer and you have multiple clusters, then that poses a problem now, because each time you define your ingress controller in... let's say you just have two clusters and you define it as type LoadBalancer, it's now going to create a load balancer per cluster. And we only wanted the one load balancer, and now we need to go and update our DNS records as well, and.

**Ville** \[30:26\]: But if you have multiple clusters and just one load balancer, what's the point there? Or why would you want to have just one load balancer?

**Jacob** \[30:35\]: For example, multitenancy of the Kubernetes clusters and maybe... Well, I mean, one case that I think many people have found themselves in, including me, is in a bigger company we might not have control of DNS records. And basically we have... we create a ticket, so point something, SaaSapp.fi, put in this CNAME or whatever, which will be the load balancer. And now we have our load balancer managed with Terraform and we can't really change the CNAME or the IP address. So that could be one use case, but also maybe cloud costs, you might want to keep your costs down if you're going to have many Kubernetes clusters for like multitenancy or this sort of thing, then maybe you don't want to have them spinning up load balances for every single cluster.

**Ville** \[31:25\]: Yeah. I mean, in that sense, it's a classic setup. So you might have sort of an ingress layer, you have always have a single entry point to any infrastructure basically, or any cluster you have running, and then for all of the clusters, you also might want to have an egress layer where each one of those clusters will show, for example, to the outside world through a single IP address. So that's totally doable. Or it's doable on any cluster, I would say.

So if you have needs for very specific, let's say, ingress routing, like you said, you might want to have multitenancy inside the clusters, then having a flexible ingress layer is something you need. Solutions include things like what comes to my mind is Envoy, for example, which works for both ingress and egress routing really well. So you might want to have, let's say, SaaSapp.fi/fu goes into cluster X and then bar goes into cluster Y and it's all managed in one layer basically.

**Jacob** \[32:33\]: But how is that communication handled? And let's say we've gone with the... I think you called it not best practice earlier. I'll let that one slip, maybe you can buy me a beer to make up for it. But let's say that I've created my load balancer using Terraform and it's configured so that SaaSapp.fi/fu will go to-

**Ville** \[32:55\]: Cluster X.

**Jacob** \[32:55\]: Yeah. But how old does it go, where does it go to cluster X? And that's maybe, if we're talking in UpCloud, then how could we do that? Would we do it as a node port, the ingress controller running as a node port, or would it be running as...

**Ville** \[33:11\]: I'm thinking, what's your sort of anticipated result? How many things do you want to configure?

**Jacob** \[33:19\]: I want my mind to be blown.

**Ville** \[33:20\]: One possibility is to, as we talked with having a CNI that spans cluster networks across, you have a unique cluster network range, for example, for each cluster, each cluster can talk to each other. Of course, you will have or hopefully you will have network policies to restrict or allow specific network access. But from your ingress layer, you could just directly call the ports, let's say the Kubernetes DNS name of your service that's running, for cluster X to then route your traffic.

**Jacob** \[33:53\]: But the load balancer isn't running in the cluster.

**Ville** \[33:56\]: I mean, it's running in your ingress layer somewhere. Does it have to run in the cluster? If you have a cross-cluster connectivity, then you can just have the layer running on some cluster or some of the clusters, and then eventually the CNI, the Kubernetes network player, will take care of routing the request into your cluster X or service X running on cluster X.

**Jacob** \[34:22\]: Okay.

**Ville** \[34:22\]: Does it make sense?

**Jacob** \[34:24\]: I think I'm still missing the part where we go... So just to make sure we're talking about the same thing here, the load balancer is like the managed load balancer in UpCloud, nothing to do with UKS or?

**Ville** \[34:35\]: Sure. It is part of the ingress layer too. So you will have one load balancer. This is not UpCloud specific even. I would say it's possible on any cloud. So you have a cluster running, ingress layer having a load balancer there, exposed over a single DNS name, for example. And because that ingress layer has knowledge on other clusters and services running on any cluster that's tied to it, then it knows where to route the request to, as long as you have those routing rules defined in that layer.

Not possible on any load balancer. So for example, in your case you would have a TCP load balancer, for example, that forwards it to some sort of smart proxy, such as Envoy, for example, and then from there on you route it to wherever you want. So not exactly a thing you would get as a service from any provider, but rather something you need to build by yourself, but that's the beauty of Kubernetes, you can do things like this on your own.

**Jacob** \[35:40\]: But I think this is why this always becomes such an interesting topic because there's so many different ways of doing this. And sometimes your hand is forced such as not being able to control DNS records. And other times it's like going for the holy grail of management can be... well, not always the best use of time and money.

**Ville** \[36:00\]: Yeah. And it can get really... like corner cases can come from, let's say, in your case, you have the clusters running on different providers, so you can't even expect on things working the same way across clusters, because they might be different ones.

**Jacob** \[36:15\]: Yeah. Yeah. But at least, I mean, what I've done most of this has been with AWS, and if you go from an AWS load balancer, at least the only way that I know about doing it, is then exposing the ingress controller as a node port on the nodes and then setting your security groups and things so that the load balancer can talk with the hosts. You see two instances where the ports will be running other services of type NodePort on that port, and then basically DNS points to the load balancer and the load balancer forwards you to a port on your worker nodes where the service of type NodePort is listening, and then that goes to your ingress pod, for example.

**Ville** \[36:55\]: The example I gave is that the beauty there is that you wouldn't have to use Terraform for it. And that's the kicker here. This sort of ingress layer would be just another service running on Kubernetes, exposed through a type LoadBalancer type of service, for example. So you wouldn't have to use Terraform for anything. Well, you might want to use it for the clusters, but for the routing itself you could do that in YAML.

**Jacob** \[37:23\]: Yeah, well, I was about to say that's the nicer way of doing it, but then you mentioned YAML. But I meant doing it via the Kubernetes APIs has certain benefits obviously, but it can also lead to these... when you don't have control over DNS records, or you don't want to create load balancers for every single cluster left, right and centre, then those kind of... controlling it on that level becomes more difficult when simply defining a type LoadBalancer creates your load balancer.

**Ville** \[37:53\]: Yeah. That's really important to know that the things you can control in these clusters, for example, in your scenario would be the service names. And because Kubernetes DNS exposes those over DNS names, which you can use in such setups, you don't necessarily need to have a public DNS for those things. You can use the cluster local ones.

**Jacob** \[38:16\]: Yeah. Yeah. And I think this space is evolving a little bit now as well, because, well, so AWS have the load balancer controller as well, which is a control you install on the cluster and you can then provision these load balancers and things. And I know there are other projects doing similar things, which is cool because that's like a Kubernetes native thing, but now they're also the... Well, there's Crossplane obviously, but there's a massive Terraform controllers coming out for Kubernetes now as well.

I think we need to search our #random channel on Slack because every time we find one, we share it and there's a mass of them coming out now as well. So I think this problem can be solved even with Terraform, but provisioning from the cluster, which is... Yeah, I guess maybe we should leave the topic to rest there though.

**Ville** \[39:11\]: I mean, we can talk about controllers all day, or operators, because I love them as well. It's a nice pattern to solve problems. What comes to your example? I mean, I think there's one cool project for Google Cloud where you can create Google Cloud projects by running an operator in Kubernetes. So you're defining projects as CRDs and then create... It's like a chicken and egg kind of thing, but some people apparently have needs for those.

**Jacob** \[39:44\]: Yeah. But yeah, and I remember when Crossplane was first becoming a thing and everybody was like, "Yeah, but if I need to have my cluster, how am I supposed to..." It was like, "Yeah, just create a cluster then." Create a cluster and kind of move on. And especially now with projects like Argo CD, I know Flux CD has been around for about just as long time, but Argo CD definitely seems to be the more popular or used one. I mean, now if you have a controller for even just vanilla Terraform and you have something like Argo CD there, then this is becoming easier and easier to do in a really nice, seamless way, so.

**Ville** \[40:28\]: Indeed, indeed. And the pattern works for any set up basically where you want to sort of offload some of the knowledge or internal sort of offloading of things into having an object. And then the only way to interact with between different controllers for example, is the spec, the object spec basically. So then you don't have to have those two things talking to each other, but rather one defining the object and the other reading the object and then working as expected. Yeah. It's a nice pattern.

**Jacob** \[41:02\]: It is a nice pattern. And I think emerging trend that I see now is having some kind of route cluster. So maybe this would be something really nice to do with UKS once we get into beta or whatever. Let's build our route UKS cluster and on there we provision the CRDs or whatever we're using, some Terraform controller through Argo CD, and it will start creating our other UKS clusters for each tenant. We can connect those with our single or our massive load balancers depending on how we want to do it, and then we have SaaSapp.fi running on UKS in a really, well, really nice way. Yeah. Cool.

**Ville** \[41:44\]: Sounds good to me.

### Wrapping up

**Jacob** \[41:48\]: Should we start to wrap up then?

**Lauri** \[41:49\]: I guess so. That last part might be a bit heavy in a way, though interesting.

**Jacob** \[41:57\]: Yeah. Yeah. Well, I did ask if we should discuss it in this podcast or not, and I got the green light, so. Yeah.

**Ville** \[42:05\]: What's the stuff? I don't remember any of it.

**Jacob** \[42:08\]: Ingresses and load balancers. It's all in history now, so.

**Ville** \[42:12\]: Oh yeah, yeah, yeah, yeah.

**Jacob** \[42:13\]: Yeah. So this has been our episode with Ville talking about UKS and-

**Ville** \[42:18\]: And Kubernetes in general it seems, not only UKS, which is good.

**Jacob** \[42:22\]: Yeah. Yeah. And it's never just Kubernetes. So we've been talking about other clouds and other interesting projects and CNIs and Cilium and eBPF a little bit, and service meshes and a whole wealth of interesting stuff. So what's your ending? What do you want to say? What's your wrap up about UKS to get people hyped?

**Ville** \[42:44\]: Stay tuned and we will see you by the end of this year at UpCloud. Well, I want to mention still about those two positions we have open, so please join. If you want to work as the Kubernetes engineer, we are hiring. So send us your stuff and we'll talk.

**Jacob** \[43:00\]: Yeah. And a little side story there. We asked Ville before the recording that we might ask you what you've struggled with, what have been your main challenges? And that's when this came up.

**Ville** \[43:11\]: Yes, we need more people. That's definitely the surprise we had.

**Jacob** \[43:16\]: Yeah. That's great. Any final comments from you, Lauri?

**Lauri** \[43:20\]: Not really. So thanks to Ville from UpCloud, and UKS is coming out in the end of this year, so stay tuned for that.

**Ville** \[43:33\]: Yes. And thank you for having me again. This was a nice chat.

Jacob \[43:36\]: Yeah. Awesome. You're very welcome to come back and discuss once UKS is out for real and we've had a chance to play with it.

**Lauri** \[43:43\]: Then we can chew on it.

**Jacob** \[43:44\]: Yeah.

**Ville** \[43:44\]: Yeah.

**Lauri** \[43:44\]: You promised us this, why isn't... it's not delivered.

**Ville** \[43:49\]: Yes. And maybe we can talk about the load balancer set up again.

**Jacob** \[43:54\]: Yeah, yeah. Maybe a whole new project has come out and the CLBI, the container load balancer interface.

**Ville** \[44:01\]: Yes. Something like that. Yeah. We can make it work.

**Jacob** \[44:04\]: Cool. All well, thanks everybody for listening and we'll see you next time.

**Ville** \[44:07\]: Yes. Thank you.