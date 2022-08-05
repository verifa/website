---
type: Blog
title: A first look at GKE Autopilot
subheading: In Verifa's first podcast we discuss Google's new Autopilot for Kubernetes,
  and what it means to have a managed Kubernetes service.
authors:
  - lsuomalainen
  - jlarfors
  - alarfors
  - slaitinen
tags:
- DevOps
- Continuous Delivery
date: 2021-03-30
image: "/blogs/2021-03-30/podcast_ep01_920-1080.png"
featured: true
jobActive: true

---
<iframe title="Embedded podcast player" src="https://anchor.fm/verifa/embed/episodes/A-first-look-at-GKE-Autopilot-ett6kl" height="151px" width="100%" frameborder="0" scrolling="no"></iframe>

<div class="blog-flex">

[![Listen on Spotify](/blogs/2021-03-30/listen-on-spotify.png)](https://open.spotify.com/show/12yStrneLdEsXn1Bjp6Myz)

[![Listen on Apple Podcasts](/blogs/2021-03-30/listen-on-apple-podcasts.png)](https://podcasts.apple.com/gb/podcast/the-verifa-podcast/id1561051552)

[![Listen on Google Podcasts](/blogs/2021-03-30/listen-on-google-podcasts.png)](https://www.google.com/podcasts?feed=aHR0cHM6Ly9hbmNob3IuZm0vcy81Mzg0NzE1Yy9wb2RjYXN0L3Jzcw==)

</div>

**Welcome to The Verifa Podcast, where we break down the complex world of Cloud and DevOps.**

In the very first episode of The Verifa Podcast, we take a first look at Google's new Autopilot for Kubernetes - GKE Autopilot. GKE Autopilot is a new mode of operations for managing Kubernetes, enabling the user to focus on software, while GKE Autopilot manages the infrastructure.

#### **During this episode we discuss**

* What do we mean by a 'managed' Kubernetes service? \[03:15\]
* What are the differences between GKE, AKS, EKS and Fargate? \[05:16\]
* A fully managed Kubernetes vs Serverless approach \[13:07 & 26:52\]
  * What does 'severless' really mean?
  * Is Autopilot serverless?
* Billing: GKE Standard vs. GKE Autopilot \[14:03\]
* When to use Autopilot and when not? \[36:34\]
* Are we going to use GKE Autopilot? \[41:12\]

#### **Mentioned in the podcast**

* [GKE Autopilot](https://cloud.google.com/kubernetes-engine/docs/concepts/autopilot-overview)

Connect with today's podcast team:

* [Jacob on Linkedin](https://www.linkedin.com/in/jlarfors/)
* [Andreas on Linkedin](https://www.linkedin.com/in/andreas-l%C3%A4rfors-51253270/)
* [Sakari on Linkedin](https://www.linkedin.com/in/sakarilaitinen/)
* [Lauri on Linkedin](https://www.linkedin.com/in/lauri-suomalainen/)

#### **Transcript**

##### Hello and welcome

**Jacob:** \[00:00:00\] Hello and welcome to The Verifa Podcast. This is the podcast where we take questions our customers have asked us or questions that we really wish our customers would ask us. And we share our view and try to break down the complex world of cloud and DevOps. I am Jacob, your host, and I have my co-host me today And we're joined by two very special guests. The first of which is Lauri.

**Lauri:** \[00:00:43\] Hello, I'm Lauri. I'm a consultant at Verifa, and I like to dabble with Google cloud mostly, but also have some experiences with AWS.

**Jacob:** \[00:00:57\] Cool. Thanks for being with us. And we also have Sakari.

**Sakari:** \[00:01:01\] hello, I'm Sakari. I'm also liking mostly Google cloud platform. Also have been doing a lot of stuff with AWS and Azure. So like pretty cloud agnostic, don't really care about which cloud it is. But yes, Google cloud is my favourite.

**Jacob:** \[00:01:30\] okay. Thanks guys. For telling me your favourite clouds. That's all the information I needed to know Which is your favourite cloud Andreas if we're going to be this?

**Andreas:** \[00:01:42\] My mine's missing. I'll ask what your favourite is in a minute too, but yeah, mine's Google too. Have worked with all three and Azure the most. But yeah, Google, Google's my favourite. It's the only one I haven't used commercially yet. I don't think that's got anything to do with it, but yeah.

**Jacob:** \[00:02:01\] I think Lauri wants to say something already.

**Lauri:** \[00:02:03\] I was hoping that someone would say Alibaba at some point, but

**Jacob:** \[00:02:08\] Yeah. Maybe one day, one day. Cool. I don't think I need to tell, which is my favourite public cloud, because I'm already outnumbered here. My opinion is obsolete. Anyway I guess the reason for everybody talking about clouds today is because we are going to talk about a cloud feature or a new cloud capability.

&nbsp;  

##### What do we mean by a managed Kubernetes service?

**Jacob:** We're going to talk about GKE's recently announced Autopilot feature, which is their new flavour, their new type of managed Kubernetes service. So that's going to be the topic for today, but of course we'll cover other things to do with that, and similar technologies and stuff too. And yeah, let's get to it.

So we're going to safely assume everybody listening knows what Kubernetes is and has some some kind of knowledge and understanding of containers in this world. Sorry if you don't but we want to get right to the meat of it. I'm gonna kick off by asking Lauri to tell us what does it actually mean for a Kubernetes service to be managed?

What is it that makes a service managed?

**Lauri:** \[00:03:26\] Kubernetes is itself quite a complicated beast to handle and. It consists of the workloads you want to run on it, but also the infrastructure which usually consists of nodes and networks and all that jazz. And traditionally that has been quite hard to administer, and let's say building your own Kubernetes cluster from scratch is a rite of passage for an aspiring Kubernetes administrator and full disclosure. I haven't done that yet, but I'm getting there. But when you talk about managed Kubernetes as a service the idea is that the Kubernetes provider provides you the control plane and also the physical hardware for your Kubernetes cluster.

So you only have to focus on handling the workloads you want to run on the cluster.

**Jacob:** \[00:04:41\] So basically if I wanted to have a Kubernetes API with some nodes and control plane and stuff, so I could send my workloads to be run, I could go to a managed Kubernetes service and say, Hey, create me this, like with these specs. And then I would basically have a, yeah, like a Kubernetes cluster.

**Lauri:** \[00:04:59\] you don't have to worry about which kind of service they are running on or what kind of network topology, there is, what sort of other considerations like multi-tenancy, is there in a cluster and so on because that's all managed for you.

&nbsp;  

##### What are the differences between GKE, AKS, EKS?

**Jacob:** \[00:05:16\] Yeah. Okay, nice. And in the world today with the three biggest cloud providers or the most common ones, they all offer their own managed Kubernetes service. So we've got Azure's AKS and AWS, there's EKS and Google's GKE. Are there any significant differences between those or are they on a high level,

are they quite the same, is it the same sort of transaction? I want this kind of cluster. Give me a cluster.

**Lauri:** \[00:05:51\] For a lot of what is my understanding, they are pretty much the same thing. Things of course, GKE being a Google products and Kubernetes being also a Google product is the Gold standard. As Sakari put it, but Sakari do you know, what about the more subtle differences between AKS EKS, GKE?

**Sakari:** \[00:06:17\] Not really, the biggest ones are between some kind of like storage. Otherwise, usually, because I would say all of those are nowadays like certified Kubernetes by CNCF. So they should all handle all the end to end tests in similar way. So you run one Kubernetes test in some place, and you can have another one in another cloud, it should handle the same.

So they should be very same.

**Lauri:** \[00:07:07\] like from the workload operations side of operating the services is similar.

**Jacob:** \[00:07:15\] And this might be a bit of a dumb question, but that's my job is to be here and ask dumb questions. What kind of tests do people run on these? If you want to have a certified cluster, what are the types of tests that get run on that?

**Sakari:** \[00:07:33\] Yeah, they, I don't even know the whole load of tests they run there. It's probably something like 200 tests or that kind. There are some tools you can run it in your cluster. But it's kinda create a pod, create a DaemonSet, create a StatefulSet create CronJobs, create everything, like everything should work the same.

**Jacob:** \[00:08:10\] So that testing the resources then. In that cluster. So basically talking the Kubernetes API language with that cluster and just making sure that it behaves as it's supposed to behave,

**Sakari:** \[00:08:23\] Yes.

**Jacob:** \[00:08:24\] which I guess in a nutshell is the definition of testing

&nbsp;  

##### What is GKE Autopilot?

**Jacob:** so we talked now a little bit about the different cloud providers and the different Kubernetes managed services that they have. And the topic for today was specifically to talk a little bit about Autopilot, which is GKE. So Google new flavour of it's the new flavour of GKE. So maybe we should introduce that topic briefly and I'll give that responsibility to Lauri.

Lauri, please tell us what is Autopilot.

**Lauri:** \[00:08:59\] that's a brand new thing. By the time of this recording it's been out for a week. So previously I talked about how to manage Kubernetes, you leave the Infrastructure to be managed by the service provider, but you still have to define it. You can you tell the provider that I want to have a six node cluster and it has to have this sort of a network and these additional features there. And after that you get to manage your own workloads.

But now Autopilot is the next logical step in managed Kubernetes, because you don't have to do that even. You just tell the service that I want to have this completely managed for me. I don't want to know anything about it. It's okay. What kind of workflow do you want to run? And then you just define the workloads and the service matches the infrastructure to your workloads.

So you don't have to know anything about the infrastructure now anymore.

**Jacob:** \[00:10:18\] Is it still like a managed service or is it more like a Kubernetes as a service now? It sounds a bit more like Kubernetes as a service, like you get the API end point where you could send your workloads and then it will just do it for you.

It is managed because they manage it for you. But yeah I guess the two worlds are getting closer and closer.

**Lauri:** \[00:10:36\] And that's more like a semantics

**Jacob:** \[00:10:39\] Yeah.

So this being quite a nice progression in the hosted and managed services that are provided. Is there anything already like Autopilot that exists, or have Google been the first to do this? Maybe Sakari, you want to take that one.

**Sakari:** \[00:11:00\] Yeah. AWS has been having this service called Fargate or is it EKS Fargate?. Anyway, it's, their like similar kind of thing where you pay by the pods. So you don't actually create cluster or node pool Okay, I want three machines or instances, but you just create your workloads and you pay by the resources they are using.

One thing with Google is they might come a bit late with their stuff, but they might be usually better.

**Lauri:** \[00:11:41\] Yeah, considering that Fargate was released in 2017, if I remember correctly. So Google had this on the back-burner for some time

**Jacob:** \[00:11:51\] But then Fargate is not just for Kubernetes. Because AWS have had their own container service. ECS, is it the elastic container service? And I think Fargate works for that too. So it hasn't been specifically built to manage Kubernetes workloads. But it can be used for that.

Whereas GKE AutoPilot is like specifically Kubernetes.

That's

**Lauri:** \[00:12:17\] correct. So would you say Fargate at least in its first iteration is more akin to CloudRun or Kubernetes instead of Autopilot?

**Jacob:** \[00:12:30\] That sounded almost like a riddle, but I guess you can use you can use Fargate for either. If you want to use ECS then you can use ECS Fargate or whatever that combo would be called. I guess Fargate is more the idea of managing the workload rather than the infrastructure nodes.

Right? Somebody please, correct me as I'm speaking out of my depths here. I'd be the one reading a bit of documentation and doing my homework and asking.

**Lauri:** \[00:13:05\] Yeah, you are correct.

&nbsp;  

##### Fully managed Kubernetes vs. Serverless approach

**Jacob:** \[00:13:07\] Okay, thanks. I guess this idea of managing infrastructure, managing nodes versus having a service available to you Is almost like going from a managed service almost to to server less compute resources. So I think that would be an interesting thing to talk about now. And compare like how does this new world of managed Kubernetes with Fargate, and when I say new, if it's been around since 2017, but yeah, at least the Autopilot side of things... how does that compare with serverless compute resources? I think to do that comparison effectively and nicely, we first need maybe an understanding, one that we agree on of what serverless really means because there are quite a few services that exist today that say they are serverless.

So what does that actually mean? And I don't know, Lauri or Sakari either one of you want to have a crack at helping us define serverless.

**Lauri:** \[00:14:14\] Traditionally, if you can say about such recent thing as serverless. The billing has been one of the main focus points. So let's say for Lambda or CloudRun when it's on a Google cloud point of view, they are built by requests made to the service.

So let's say you have a web server. In serverless that web server is idle. It doesn't accumulate your bill, unless it's called like with Autopilot, you don't have to manage the underlying infrastructure there. But unlike AutoPilot, when not in use, the resources are turned off. So for example, if you had the aforementioned web server before it gets his first call it will take some time to boot up.

Whereas Autopilot,in Autopilot, you have continuously running services. They don't need to be specifically called. And you're going to pay for the resources they consume. So in a way they are similar, but their modus operandi is a little bit different.

**Jacob:** \[00:15:53\] With serverless, it's not that you don't have to care about the server, which is almost what the name implies. There is no server. There is an underlying server, but you don't need to care about it, but you would put the focus more on the use of the service rather than the fact that you don't have to care about there being a server.

So billing being a specific topic there that you pay per requests. So you pay by the amount you actually use it rather than paying for the underlying resources.

**Lauri:** \[00:16:26\] In that respect the serverless Lambda functions and so on, let's say in traditional or standard GKE, as they say now, you pay for the nodes. So let's, I'm going to allocate free nodes. And then I have one pod running there and it wouldn't consume, but only one node's, resources.

But I would pay for the nodes themselves, but with Autopilot the service itself would allocate enough resources for my one pod to run. So in that way, I wouldn't be paying for things I don't use.

**Jacob:** \[00:17:12\] it's a bit of a, like a grey area then, because it's serverless in the sense that, you don't care about the underlying servers and you're not paying for not caring about the service you are paying for. Like actually using the service, you're actually paying to run Kubernetes workloads.

So it's kinda like server-less from that perspective. But if we look at the other ideas or benefits of serverless that you mentioned, which is things on all the time, that does not exist because I'm running a workload in Autopilot. So basically a pod running a web service or something like this is going to be on all the time.

It's not going to be scheduled based on a request or based on something like that.

**Lauri:** \[00:18:01\] Yeah, in comparison most of the serverless stuff is event based. So you'll get a call from say a HTTP end point or some sort of a message queue. And then you run a function and the function goes to sleep.

**Jacob:** \[00:18:17\] really interesting. Really interesting.

How about Andreas and Sakari? How'd you feel about that? I don't know.

**Andreas:** \[00:18:29\] It's really interesting. So it's almost like another level of abstraction added on top of what we already have. And by level, I mean a level away from the hardware. A long time ago we had server racks in offices. A lot of people still do have server racks in offices, but you would run your software on your hardware and you were in control of both of those.

And then along came virtualisation. We started spinning up a virtual machine on the hardware. And then on top of that, you've got cloud computing where the hardware could be shared amongst many different users. So I don't know, let's say a hundred users on a single server rack or with their own individual virtual machines.

And this sounds like another level of abstraction on top of that. So we're getting even further away from being involved with the hardware. For the last 10 years or more, I guess you've been able to go into one of these cloud providers, spin up a virtual machine or a a Kubernetes cluster with a node pool. But you were still somehow aware of like virtual machines, even if they were virtual. So you're like one, one abstraction away from the underlying hardware. This sounds like another abstraction away on top of that. It's like the services running, it's managing the virtualisation part for you.

Is that kind of a good way to look at this?

**Lauri:** \[00:20:09\] Got it. That's a good summary. I think. Yeah. What would be the next step, then you're going to a cloud provider and telling your business case and they take care of the rest

**Andreas:** \[00:20:19\] Yeah. The cloud provider creates the software for you. Yeah. They tell you what to run.

**Jacob:** \[00:20:27\] you like record that: I want this app, which is going to do this and this. And then five minutes later, the code has been auto-generated , auto- deployed in your cloud provider. And.

**Lauri:** \[00:20:38\] one. Yeah, Yeah, one thing I've always had a distaste for severless because I felt that it requires you to go to a vendor lock. So if you're going to go full serverless and you're going to write functions that defines your software, that you are stuck with, that particular provider's seller stack.

But I think with services like Autopilot which use the common denominator. There is containers. And let's say if you have a container, you have a, let's say Kubernetes deployment event going one level further up. You can use the same Kubernetes manifest in manage standard GKE AutoPilot and any other competitors and so on. So in that way, it's more possible when compared to serverless.

**Jacob:** \[00:21:43\] Yeah, it's much more cloud agnostic. So that's definitely brings some benefits. And not everything can be run in a serverless fashion either. So it's also the ability with Kubernetes you can basically run any workload in there. At least, I don't think you can run everything in serverless. Is that just a feeling I have? What about like databases and

**Andreas:** \[00:22:06\] Hmm, like, high availability services. And those can't possibly be suitable for serverless

**Jacob:** \[00:22:15\] Yeah. And I guess there's some costs with a startup on server-less resources too. If you need really instant real-time things, then.

Yeah I guess it depends a bit.

**Sakari:** \[00:22:29\] Yeah there are those serverless databases for example, Google Spanner or AWS Aurora, which are, well by my definition, the serverless is where you don't really choose instance size or something. The billing goes by requests and something else. Whatever the service is. So in Amazon Aurora, the service costs by requests and something, I don't know, similar kind of things in Google Spanner.

But anyway, you will not go to the console and click,. I want this size machine for it. So unless you click that one, it's serverless.

**Andreas:** \[00:23:27\] I guess in, practice there are actually quite a lot of things, depending on your definition, there could be quite a lot of serverless applications. If you define serverless as I don't want to be aware of, and I don't want to be in control of the underlying server. Then something like a managed database or some versions of that could be called serverless.

My understanding of serverless was always there is not guaranteed to be a server running your software at all times. There may be times when your software goes idle. That was my understanding of serverless. Which one of those is correct here.

**Sakari:** \[00:24:12\] I would say that Cloud SQL or Amazon RDS are not server less because you have to create the instance for your database.

**Andreas:** \[00:24:24\] And does that instance remain running, whether it is in use or not? Is it always on,

**Sakari:** \[00:24:31\] Yes, it's always on.

**Lauri:** \[00:24:35\] Yeah , in that way your definition is quite good because that's if you want to draw a line between managed and serverless, I think the point of not having a guarantee that your application is already on some server when it's called is one of the better ones I've heard.

**Andreas:** \[00:25:00\] Like a managed service would be your application is always on a server somewhere at all times. And serverless would be it isn't, but it will be.

**Lauri:** \[00:25:13\] Yeah. So , for example serverless where you wouldn't, you could run... it suits batch jobs very well, but not continuous jobs. Or streaming or that kind of stuff.

**Andreas:** \[00:25:24\] Yeah. So things that can be run on demand that don't need instantaneous, something where you can, it can accept like a queue of jobs that are managed In a timely manner. So probably not like a website or a web server where somebody is directly interacting with that service. But something where you can you can basically Oh, what's the word dispatch jobs or delegate jobs, anything where you delegate or dispatch jobs to something, and you're not necessarily waiting for those to come back.

So some kind of asynchronous delegation of jobs that would be very suitable for a serverless. Is that like the optimal use case for a serverless application?.

**Lauri:** \[00:26:07\] Yeah, why not? And in that vein, Cloud Build, for example, would qualify for a serverless. Even though it's not marketed as a such. Cloud Build being a CI/CD tool for which you can spin up multiple Docker containers and run jobs on them, usually for deploying code.

**Andreas:** \[00:26:31\] Yeah, I think maybe that's just with Cloud Build in mind, maybe that's the specifics of that. Like it's intended for building stuff, deploying, rather than a general IT service. So at least I have a good understanding of serverless and managed services.

**Jacob:** \[00:26:52\] Yeah, this has been a really interesting discussion to, challenge these terminologies and words that get used really often. But sometimes we don't step back and really think what they mean. And I think now that we have reached this somewhat common definition of serverless back to the original question then. Autopilot, is it serverless?

And if not, why not?

**Lauri:** \[00:27:21\] The answer is no. It's still managed, but this is a little bit further managed than standard GKE because you are paying by the resources. You are guaranteed to have those resources at your disposal where you define your workloads and run them in Autopilot.

**Andreas:** \[00:27:45\] So I'm interested in what's the level of detail, what's the level of granularity of defining resources. Cause I'm familiar with requesting a Kubernetes cluster and telling it how many nodes I want and what size those nodes should be, storage classes and things like that.

But if I'm not telling, if I'm not telling Autopilot any of that, what do I have to tell it? Is it done on resource limits and requests of pods or what's the, how does it figure everything out? Because it can't just run everything up, wait until it crashes and then say, Oh, that was wrong.

I'll make some adjustments and okay. We need a bigger pod here. We need more memory. We need more CPU. How do, how does what do I tell it? And how does it figure it out?

**Sakari:** \[00:28:37\] Yes, you create those requests, like resource requests and limits. That's how it works. And about the topic of managed service and serverless things. Like they are not always the same thing. So there is managed services and there are server less things. Probably you cannot have any serverless thing, which isn't managed.

**Jacob:** \[00:29:09\] So a service being managed is like a superset of serverless, then. I was continuing on, the question Andreas asked. Thanks for your views there, Sakari on the serverless and managed. That's nice contribution. Yeah. So if I send a Kubernetes manifest, declaring a pod, And I set the resource requests to whatever, some number and the Kubernetes algorithm for figuring out where to put that pod realises that there aren't enough compute resources available. Enough, like raw power on the nodes it has running. Won't that takes some time for Autopilot to create new nodes for us and then make sure my pod gets scheduled? Or is that something that Google have solved and take care of for us. And we make sure we get our resources on time. And without too much delay

**Sakari:** \[00:30:11\] It could take some time because it's pretty much like how the normal autoscaler works. Oka,. My requests are requesting something that in my cluster, there isn't like space to put it in. So then it needs to spin up new node and the pod is just pending until it can have some resources to run it.

**Lauri:** \[00:30:39\] Being quite new, most people don't have much experience with it. I'm also a bit concerned, consider a case where you have a faulty application, that leaks memory. Would it keep doing that in Autopilot? And then increase your, the capacity it requires, at the end of the month you get presented with that huge bill, because you're using all the memory available in the data centre,

**Andreas:** \[00:31:09\] Is there some upper limit you can set? Surely there's some kind of control like that for it.

**Lauri:** \[00:31:16\] it uses Quotas defined in the compute engine. But then again, in my opinion, it goes a little bit against the ethos of AutoPilot. If you're not supposed to manage infrastructure why you're supposed to make quotas, then too.

**Jacob:** \[00:31:30\] Yeah, I wish Autopilot could just manage my billing as well. I think that somewhat makes sense that you put your billing there, because yeah, that sounds like a very legitimate use case and your billing going sky-high. Although I haven't heard that many stories. This is just my experience.

I haven't had that many horror stories with people in GCP being, getting to the end of the month and having these crazy high bills like you have with other car providers. Not naming any names. So I would assume that you could use the quotas or billing alerts with Autopilot to make sure that doesn't happen.

**Lauri:** \[00:32:12\] With standard, it was easy because you could define the upper limit to which the nodes could , those could auto scale up too. And when the ceiling is reached, your purchase keeps failing, you probably get notified by that. If you have proper auditing system in place.

**Sakari:** \[00:32:31\] It is the same with pods. You can set the requests limits. Or resource limits. If you're hitting that limit, you probably get some basic out of memory error or something. So it's really not about that.

**Lauri:** \[00:32:49\] Yeah, one always should instantiate the limits.

**Andreas:** \[00:32:53\] What else apart from node size and the number of nodes, what else is Autopilot going to do for us? What else is there to do in the setup of a complete cluster? Obviously we don't need to talk about the normal managed Kubernetes service, like providing a control plane and things, but what additional things does Autopilot take care of? Apart from node size.

**Lauri:** \[00:33:22\] Apart from node size, of course, going to be node type. And then there is the auto-scaling part. What else? Networking?

**Andreas:** \[00:33:31\] Yeah. So I guess I've been running most of my clusters with a cube net and that's just worked for me. What other considerations are there when building a cluster in terms of networking because things like ingresses and things, that's stuff you define in your deployment, right?

Would you class an ingress as something you'd define in the cluster or something you'd define in the deployment?

**Lauri:** \[00:33:56\] In a deployment, definitely. At least when it comes to GKE.

**Jacob:** \[00:34:00\] I guess it depends like there is a connection, right? The type of ingress controller that you have and the type of the, the type of, for example, service. Like a load balancer would need some hardware to spin up that load balancer. So again, when you have the managed services by cloud providers quite easy, right?

Because everything's there. So if you declare a service of type load balancer, and so there's going to be a load balancer in the cloud that's going to be created for you. But if you're running your own Kubernetes cluster or running one for local development, then you probably need a way to change that. Or then provide some capability for a service of type load balancer to actually be able to fulfil that request.

**Andreas:** \[00:34:44\] Yeah, but in terms of what Autopilot brings, in addition to a normal managed Kubernetes service, what kind of network.. I'm trying to figure out, what does Autopilot need to take into account in terms of network that a managed Kubernetes cluster would, you know, you'd have to manage as a Of note here is that I haven't set up many Kubernetes clusters and the Kubernetes clusters I have set up have been quite simple. So I'm just trying to see what are like the networking tasks of a complex Kubernetes cluster that AutoPilot would take care of for you.

**Lauri:** \[00:35:26\] Quite similar to standard GKE. So if you're using GKE, you don't really have to think about that much of the networking except maybe for routing can be either based on pods or the nodes themselves. On AutoPilot, the pod based driving is the only available option.

**Jacob:** \[00:35:47\] Autopilot is just a flavour of GKE. When you create a GKE cluster, you either say whether you want standard, which is well, the standard one, or if you want Autopilot, which then removes the resources, like the compute engine nodes and stuff like this. They don't appear in your project and your Google cloud project because you don't manage them.

And you shouldn't know, you shouldn't care about them. So I think from the networking perspective, it's the same, right?

**Lauri:** \[00:36:14\] And also in a way, Autopilot is a subset of features of GKE standard. Of course you lose some freedom when giving giving up control to Google cloud platform in regards of your cloud service, that's the price you have to pay.

&nbsp;  

##### When to use Autopilot and when not?

**Jacob:** \[00:36:34\] We're now really touching on some of the next topics we were going to talk about which is very nice. And those were going to be when to use Autopilot and when not to use Autopilot. And leading on from that discussion where you're not in charge of your nodes, I guess that plays somewhat of a part in choosing when to use Autopilot and when not to use Autopilot. At least my understanding would be that if I'm going to run very vanilla containers and workloads in Kubernetes and I just want a Kubernetes cluster, then Autopilot would make perfect sense because I don't have to care. There's less stuff for me to care about.

But if I need control over some of the runtime OS, if the nodes, or I'm using some special container runtimes, or some special things that I need to care about, or the nodes need to be in some special network with some special things set up. Basically customisations. Then probably I'd want to roll with GKE where I have a bit of control over the nodes.

Would that be a fair summary?

**Lauri:** \[00:37:36\] You would be right. There are other restrictions too. Let's say you're going to do some heavy customisation, and you want to do stuff on the cube system side, in a cube system, namespace. Then you can't go with Autopilot. because, that is roughly restricted now.

Also, if you want to run, let's say hundred, about a hundred Watts per node. That's impossible in Autopilot. The quota is much lower

**Jacob:** \[00:38:07\] So first they take away our nodes and now they take away our cube system

**Lauri:** \[00:38:12\] Only if you let them. Considering the most workloads I've run on GKE. Hindsight, if I had Autopilot back then I would have gone with that.

**Jacob:** \[00:38:26\] Sakari, are there any times when you would choose standard GKE or if you have a choice now are you going to go with Autopilot as the default. If it works as described, cause we've only a week in.

**Sakari:** \[00:38:40\] It's bit like Lauri said that it's bit like subset of the GKE standard. So if you would want to run something like service mesh it still, it doesn't work in Autopilot. That's just one example. And if you really need some privileged stuff, you probably cannot run it in AutoPilot. So there are bit of limitations, but I think they try to be as, as close as the normal standard one. Supporting entire Kubernetes API and something like, pretty much everything should work, but there are like just little bit of limitations.

**Lauri:** \[00:39:30\] The most notable limitations are that if you would want to run a huge cluster with lots of nodes, Autopilot limits you to 400 by default, whereas GKE standard, you can have 15,000. That's the most notable thing. Also going to very small scale, if you want to run zonal cluster, you can't do that with Autopilot because the Autopilot clusters are by definition regional.

**Sakari:** \[00:39:59\] Yeah. And it, anyway, it's not like big things, but some things. So if you have your simple, I don't know, web server running in Autopilot cluster, it will probably work.

**Jacob:** \[00:40:13\] And I assume if you went down the path of Autopilot and then you grow out of your shoes, and you realise that you needed some more of these power features and power capabilities, then you can always switch to GKE. It's not exactly like you have to make the choice for the long run.

**Sakari:** \[00:40:31\] No, you cannot convert your cluster. So you should create a new one.

**Lauri:** \[00:40:37\] Yeah, but then again, you have your workloads, which should work as well in standard as they do in Autopilot. Sakari is right. You can just click a button and covert an Autopilot cluster to a standard GKE cluster.

**Jacob:** \[00:40:52\] Yeah, I guess that, that wasn't as easy as I was referring to, or that would be nice, but you'd have your stateful. Everything's stored in GCS for example, and you'd have PVs that you can delete without actually deleting the underlying storage and then create your new cluster. It could be quite a seamless transition.

**Andreas:** \[00:41:12\] Are you going to use it? Sakari, Lauri, are you going to use AutoPilot?

**Lauri:** \[00:41:18\] Most likely because first of all, I'm curious. And second of all I still haven't come across a workload that, or a use case where GKE standard would be the only option. I have come across a few that a managed Kubernetes server was not an option. But the choices hasn't so far been that if we had an even more managed service, rather than using a GKE standard. I guess time will tell, I predict that Autopilot is probably going to take off because of the ease of use.

**Sakari:** \[00:41:57\] I would also say most likely, like we used to have our internal deamon stuff and it used to be a cluster that was just pretty much doing nothing, and we need to pay for those instances. So compared, we would only pay for the parts that were running. That could've made quite a bit of change in the billing.

**Lauri:** \[00:42:23\] Yeah. And from a business, point of view people don't want to run a Kubernetes cluster. Kubernetes cluster is a tool. It doesn't contribute to your payroll. It's the applications you run on it that do. So this also shifts the focus more to. where the actual money comes from,

**Andreas:** \[00:42:49\] Do you think this simplifies auto scaling?

**Lauri:** \[00:42:53\] In what regard?

**Andreas:** \[00:42:55\] Is it going to make auto scaling more accessible for people?

**Lauri:** \[00:43:00\] Oh, I see, like in general. Possible, but then again, there's also, of course it's Kubernetes. You have to know how to run a Kubernetes and you have to know how Kubernetes operates. The threshold is slightly lower, but it's not entirely gone.

**Andreas:** \[00:43:17\] Yeah, because I've spent some time in spreadsheets with node sizes and the pods that I need to run and their limits and requests. Basically just trying to find like the optimal node pool for my deployment. So I guess this just takes care of that for me. That's the idea.

**Lauri:** \[00:43:35\] Yeah, but you still have to manage your workloads, horizontal auto-scaling and also vertical.

**Andreas:** \[00:43:42\] Yeah. So number of replicas.

**Jacob:** \[00:43:45\] But I guess the point is if you have a well-defined workload. So you have worked on defining the resource requests and the resource limits in your pods. These sorts of things, then Autopilot will optimally scale, the necessary nodes to run that. And then you, because you're paying per pod, you're paying for the amount of workload you're running.

So I think if that's the context that you have optimised your workloads and you've optimised your pods, Then, I guess your billing will be optimised for you by Autopilot, right? So that's nice. It puts more emphasis on your workloads rather than the infrastructure, which I think Lauri was to your point as well, that nobody cares how many servers you have. Nobody cares how complex your infrastructure is, what they care about is does it scale and is it maintainable and these sorts of things and I guess Autopilot gives you less stuff to care about so you can focus more on the business value in that sense.

**Lauri:** \[00:44:50\] Yeah from the customer, and let's say shareholder point of view, that is what you get. And it of course pains me as a DevOps guy and a giant nerd that when I do elegant things under the hood, when it goes unnoticed unless it breaks, but that's how it goes.

**Jacob:** \[00:45:11\] And when it breaks, people know about it for the wrong reasons.

**Lauri:** \[00:45:15\] Exactly..

**Jacob:** \[00:45:16\] Okay. I guess we're nearing the end of our session together. I think it's been a really wonderful discussion. I feel like we've been on a bit of a spiritual journey together here and discussing serverless and serverful and, just about everything else to do with that.

Are there any closing remarks anybody would like to make? What would you say to our viewers about GKE AutoPilot and what we've discussed today?

**Lauri:** \[00:45:44\] Considering that AutoPilot is so new. So it would be nice to revisit the topic at a later date and see how much people have adopted it. Maybe try a real use cases also by myself. Anyways, I've had a lot of fun.

**Sakari:** \[00:46:04\] I would say it would be worth a try. So if you can run your workloads in Autopilot, just do it. But if you have something, like special things you need to do, then GKE standard.

**Andreas:** \[00:46:21\] And what if I only have access to Azure?

**Lauri:** \[00:46:24\] Dead silence!

**Jacob:** \[00:46:26\] Do we need to answer that? I'm pretty sure Microsoft will be releasing their own version of this at some point. So it's do you want to wait or do you switch? No, I'm kidding.

**Lauri:** \[00:46:40\] There's going to be a Terraform provider for Autopilot, right?

**Jacob:** \[00:46:44\] I checked the Github. Actually I posted it in our Slack and I checked the update yesterday. So they are planning to make, it's going to be the same resource kind in Terraform, but with the parameter to say, whether it's standard or whether it's Autopilot. Following the way it works in the UI and the way it works from the command line, with the G-Cloud command line..

**Lauri:** \[00:47:08\] Yeah. So if you Terraform your current environment don't jump to cloud Autopilot just yet because Terraform provider is in the making.

**Jacob:** \[00:47:20\] The actual provider is coded. I think they're waiting on some go libraries for the GCP SDK to be updated, to support it. So actually the Terraform provider is there ready, but it's the Golang SDK is, from I guess Google maintains those, that's lagging behind.

So to be soon,

**Lauri:** \[00:47:44\] Looking forward to that.

**Jacob:** \[00:47:46\] Yeah, me too. In fact, I think Lauri, you wrote the deployment I'm using at the moment. So as soon as AutoPilot is there, I'm going to add that parameter in and away we go.

**Andreas:** \[00:47:57\] Nice. Yeah, thanks a lot, Lauri and Sakari.. Good hosting Jacob.

**Jacob:** \[00:48:05\] Good co-hosting Andreas! And to all the listeners hope you've enjoyed it. Hope it was educational. I hope it was fun. And we'll see you next time for another episode of The Verifa Podcast.
