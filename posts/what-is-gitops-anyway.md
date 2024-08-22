---
type: Podcast
title: What is GitOps anyway?
subheading: In this episode of The Verifa Podcast, Jacob and Andreas get into a deep
  discussion about GitOps, how it works, and its applications and benefits.
authors:
- jlarfors
- alarfors
tags:
- Continuous Delivery
- Git
- GitOps
- Continuous Integration
date: 2021-05-04
image: "/static/blog/2021-05-05/ep02_podcast_9201080.png"
featured: true
---

<iframe title="Embedded podcast player" src="https://anchor.fm/verifa/embed/episodes/What-is-GitOps-anyway-e1051nv" height="151px" width="100%" frameborder="0" scrolling="no"></iframe>

<div class="flex gap-x-4">

[![Listen on Spotify](/static/blog/2021-03-30/listen-on-spotify.png)](https://open.spotify.com/show/12yStrneLdEsXn1Bjp6Myz)

[![Listen on Apple Podcasts](/static/blog/2021-03-30/listen-on-apple-podcasts.png)](https://podcasts.apple.com/gb/podcast/the-verifa-podcast/id1561051552)

[![Listen on Google Podcasts](/static/blog/2021-03-30/listen-on-google-podcasts.png)](https://www.google.com/podcasts?feed=aHR0cHM6Ly9hbmNob3IuZm0vcy81Mzg0NzE1Yy9wb2RjYXN0L3Jzcw==)

</div>

**Welcome to The Verifa Podcast, where we break down the complex world of Cloud and DevOps.**

In this episode of The Verifa Podcast, Jacob and Andreas get into a deep discussion about GitOps, its applications and benefits, and best practices.

## **During this episode we discuss**

* What exactly is GitOps \[00:00:52\]
* What are the tools and processes involved \[00:03:29\]
* Sealed Secrets \[00:09:13\]
* What are the benefits and problems solved by GitOps \[00:12:44\]
* Where and when to start your GitOps journey \[00:19:00\]

## **Mentioned in the podcast**

* [FluxCD](https://fluxcd.io/)
* [WeaveWorks](https://www.weave.works/)
* [ArgoCD](https://argoproj.github.io/argo-cd/)
* [Sealed Secrets](https://github.com/bitnami-labs/sealed-secrets)
* [HashiCorp Vault](https://www.vaultproject.io/)

Connect with today's podcast team on Linkedin:

* [Jacob on Linkedin](https://www.linkedin.com/in/jlarfors/)
* [Andreas on Linkedin](https://www.linkedin.com/in/andreas-l%C3%A4rfors-51253270/)

## Transcript

**Jacob:** \[00:00:00\] So hello everyone, to episode two of Verifa's podcast. And today we're going to talk about GitOps, which is a topic I personally like very much because it combines two of my favourite things, which is Git workflows, which obviously tie in very closely with CI, and then continuous deployment, which is automating releases and deployments. And we decided to talk about GitOps because we feel like there's a lot of confusion out there. And we want to try and clarify things a little bit, talk about what GitOps is and what GitOps isn't. And to help me with that, I've got Andreas with me today.

**Andreas:** \[00:00:37\] Hello, Jacob.

**Jacob:** \[00:00:39\] How's it going?

**Andreas:** \[00:00:40\] It's going very well. How you doing?

**Jacob:** \[00:00:43\] Yeah. I'm all good. I'm hyped up to talk about GitOps.

**Andreas:** \[00:00:48\] That makes two of us.

**Jacob:** \[00:00:49\] Awesome. So let's Git to it.

**Andreas:** \[00:00:52\] Should we talk about what GitOps is?

**Jacob:** \[00:00:54\] I can give my view on what GitOps is, and to me it's continuous deployment. So about being able to automate deployments, whether it be infrastructure or application deployment, it's that coupled with a workflow around Git. So how do you take, for example, a Git commit and turn that into a release could be in emerging to master branch, or it could be creating a tag and then yeah. Having a continuous deployment kick in and take that into the target environment. And that to me is what GitOps is all about.

And the main goal there is that what you have in version control as like a snapshot, a commit, is synchronised with the target environment. And that is obviously the job of continuous deployment is to make sure that if you have a specific version and this is the latest release, let's say, so this is the one that should be in the target environment. Then it's the job of continuous deployment, whichever method or tool you use for that to make sure that's the case.

**Andreas:** \[00:02:00\] Yep. So it's, in several parts it's using Git or something that offers very similar features, some other SCM tool. But it's using Git and the workflow that it offers to manage your versions, manage your releases. And then automating the deployment of all of your definitions into kind of a working system of some sort.

**Jacob:** \[00:02:25\] Absolutely. Yeah. Yeah.

**Andreas:** \[00:02:28\] I tried to make that as short as I possibly could. The shortest possible definition of GitOps while not excluding too much.

**Jacob:** \[00:02:36\] Yeah. I think that's a good level though. It's not too abstract and it's not too specific.

**Andreas:** \[00:02:40\] Yep. So I take my infrastructure as code, my configuration as code, I version those in Git. I define a workflow in Git that allows me to systematically or programmatically understand when I have a release. Whether you do branching and pull request into a master branch or something like that. Whenever something occurs on the master branch, then we have a release and that's when our automatic deployment is triggered. So that's kind of a solid example of Git is infrastructure as code, configuration as code or versioned in Git when a change occurs on a master branch that triggers our deployment and our system is deployed from those things defined as code.

**Jacob:** \[00:03:29\] Yeah. And what I find quite amazing about that is it's not, this isn't some crazy revolutionary thing that's just come out in the last couple of years. This is to a lot of people, a basic instinct. I know it has been for me quite a ... I almost don't consider it a practice because it comes quite naturally. Of course, I'm going to use Git to version things. Of course, I'm going to have a workflow. Of course, I'm going to have some kind of pipeline to automate tasks that I need to do repeatedly. So it's not GitOps is this brand new thing. But I think it's just a slightly clearer definition maybe to doing things.

The fact that you use Git, and there's now pretty standard Git branching strategies that you can use. And you can find a ton online when you search, if you don't already have a favorite already. I think what is quite challenging though, is nowadays on the continuous deployment side, choosing a tool or even knowing which strategy to use. And this actually leads back to quite an interesting topic around the history of GitOps. Because when GitOps first came out in 2017, it was coined by a company called Weaveworks who are the creators of a tool called Flux CD, which is an operator that runs inside Kubernetes and basically watches Git repositories for changes and then automatically deploys those changes. So it's a pool based operator.

**Andreas:** \[00:04:51\] And that's F-L-U-X, right? Flux CD.

**Jacob:** \[00:04:54\] Yeah.

**Andreas:** \[00:04:55\] Yeah.

**Jacob:** \[00:04:55\] F-L-U-X CD. We can drop the links to this in the release notes. Of course.

**Andreas:** \[00:05:00\] So I guess what was not, I don't know if revolutionary is taking it too far, but what was different with Flux was that previously people had been observing these changes to our repositories, where we have our infrastructure as code, configuration as code. They'd be watching those repositories and doing a push change. So when they see a change, they would have something that triggers a push to the production environment or whatever this is being deployed. The difference between that, which people had been doing and Flux when it came along, is that Flux sat in the production environment already and pulled the change into the production environment. Is that a good way of explaining it?

**Jacob:** \[00:05:44\] Yeah. Yeah. I think that summarises the difference between push and pull based deployment. And I think this was a source of some confusion to begin with because it was almost as if GitOps wasn't so much a practice, but more of a specific tool or a specific method for doing something.

And so I think when the term was initially coined, it was like, okay, if you use Flux or if you use Argo CD, which is an alternative to Flux that does very similar things. If you weren't using those and doing this kind of pull based deployment, then you weren't doing GitOps.

And I think just like every term that comes out starts to manifest and starts to transform. And I think in the end it should just be a useful thing to find something useful. And it gives people something to practice, something to help them get better. And I think GitOps has evolved to become this more of a generic deployment. So I would say that now it doesn't matter if you're doing a pull based or a push-based deployment, as long as you follow the core principles of GitOps then I guess you could say you're practicing GitOps.

**Andreas:** \[00:06:47\] Yeah. I would say nothing in the definition of GitOps restricts you to which specific tools to use. It's more about what you do than how you do it.

**Jacob:** \[00:06:58\] Except Git. Yeah.

**Andreas:** \[00:06:59\] Yeah I guess Git, people argue that you can use other SCM tools that offer some kind of features like a release flow or a management flow. But yeah, I don't see why you wouldn't use Git these days. It's just so defacto, when it comes to SCM tools.

Would you say that when we talk about GitOps, that it in general, it encompasses more than just the application delivery? Because that's my impression of GitOps, is that it's more than just the application.

**Jacob:** \[00:07:32\] Yeah. I would say just like we saw that the rise of infrastructure as code and these other config as code before that as well. The idea of applying all the awesome practices and things we've built up around software development should be applied to the things that ops people would have traditionally done with infra as code and config as code. So I don't see any reason to draw any kind of lines between using GitOps for application code, as opposed to using it for infrastructure or whatever you need to get your environment running for your customers or for the end user.

**Andreas:** \[00:08:09\] I guess when the term was coined, it was in relationship to Flux CD. So it was about a little bit more than just application deployment.

**Jacob:** \[00:08:21\] I suppose so, yeah. And even if you were just doing things in Flux in Kubernetes now, which doesn't make you just GitOps, because that's what you're doing. But let's just pretend for instance, you are doing that, which is actually something we do quite often. We're actually avid fans and users of Flux. But there's now things like Crossplane coming out, which is infra as code. It's an alternative to Terraform, which is a Kubernetes operator. So you can now deploy changes to Kubernetes and Crossplane will then be like a CD tool for infrastructure as code based on the Kubernetes manifests you've made. So world's going crazy. You can now GitOps your way into GitOps almost.

**Andreas:** \[00:09:03\] Yeah. We are going to be continuously deploying our continuous deployment systems.

**Jacob:** \[00:09:10\] Exactly.

**Andreas:** \[00:09:10\] It's like the chicken and egg problem.

**Jacob:** \[00:09:13\] Yeah. That's actually a very interesting topic to talk about because I think at least I suffered with this. I think mentally more than anything, because when I started to do these kinds of practices, years back, it was always like, okay, there has to be a starting point. Where do you like ... Okay, I can create all the tools that I need and I want to do all this amazing stuff. But if I don't have a CI tool or if I don't have, even if you go as far back as not using a SaaS tool for version control, it's like if I use GitOps practices to create and host all my development tools, what do I do before I have those?

And there's quite a funny topic, I guess nowadays it's much easier because you have cloud providers and you have all these SaaS services, you don't really need to host all your own stuff. So, that gives you a really good starting point to build off. But there are still things like secrets, which will probably never, however amazing you try and be, however great you are with GitOps. Everybody has secrets and everybody has to take great care when working with secrets and everybody has manual work around secrets.

**Andreas:** \[00:10:20\] Yep. I was sat here itching to start talking about secrets, and then you beat me to it. Yeah. I promise. We're not going to talk about Kubernetes too much, but like sealed secrets, such a cool solution. It doesn't get rid of the egg problem. You still have to create them at some point. You will never get away from having to authenticate. So there is still the egg problem of authentication or secrets management to solve, but we're getting really close now to it. Once the initial setup is done, it can be securely managed automatically.

**Jacob:** \[00:10:54\] Yeah. I don't want to burst your bubble there, but if we're talking about, GitOps for GitOps, and you need a Kubernetes cluster before you can create sealed secrets, because if people don't know what sealed secrets are, they are basically a custom resource in Kubernetes that it creates a key in your cluster. A private key that's specific to your cluster. And then you can sign secrets against that cluster. So they can only be decrypted in that cluster. So that you can check them into version control, because unless somebody has access to your Kubernetes cluster and can apply your sealed secrets there, then they can't decrypt your sealed secrets. But that requires that you have a Kubernetes cluster created. And if that's being created automatically, then yeah, you still got to go in there and do things manually. Yeah.

**Andreas:** \[00:11:44\] My point was still more that technology keeps coming out that makes things better and better and better. So yeah, we are moving more and more towards automated deployment of everything. Yeah.

**Jacob:** \[00:11:58\] Yeah, absolutely. And with things, like with HashiCorp vault as well, and the integration that you get with Kubernetes now using basically service accounts to access vault and get the secrets into Kubernetes pods and things that's more and more moving in a nice dynamic. I guess, zero trust kind of environment where there's less need for manual work. So there's less risk of doing things stupidly and less secrets are now really long live static things that get copied from place to place in a much more dynamic. So the access you get to these secrets is dynamic. We are moving in the right direction, but I guess secrets will never go away. And there probably will never be a perfect way of handling secrets.

**Andreas:** \[00:12:44\] Yeah. Good. That was an excellent eight minute tangent in this GitOps oriented podcast. Should we talk briefly about the advantages, benefits, the problems solved by GitOps and related to practices? Because I think that's interesting for people.

**Jacob:** \[00:13:00\] Yeah, I think so. Going back to the point earlier that we said, GitOps isn't anything revolutionary, and it's basically all the benefits of having a solid Git workflow, plus all the benefits of having a solid continuous delivery or continuous deployment workflow. I'm not sure what more there is to add really then. If you're not doing it, you probably should be, unless you're a one man team or a small team where it just doesn't make sense to put in the initial work to get something up and running. But things like traceability through Git. So basically every change you make, you can roll back or see who did what where exactly. Why as well. Yeah. And if it's gone through a PR then who approved it and then obviously continuous deployment, well, if you want to do things well, then you should do things often. And if you can automate those things, then that's really good thing to get into as well. It means you can just deliver faster and leaves less room for human error as well in the deployment process. So.

**Andreas:** \[00:14:06\] Yeah, I really like that part because verifying changes to an environment or to an application where maybe you put them first in a staging environment before you applied those changes to your production environment. Having everything defined as code just allows you to. And you can verify your changes in staging. And if you just apply the same code in production, then theoretically it should be identical. What works in staging should work in your production. There's no, oh, I missed a step or this is different between the two environments. I need to keep this in mind. That whole part of it goes away. It's so nice from a development and a collaboration standpoint too. Somebody can pick up your work and go with it.

**Jacob:** \[00:14:53\] Yeah. Yeah, absolutely. So I guess that leaves a question for us then. Andreas, would you recommend GitOps as a practice for people?

**Andreas:** \[00:15:06\] Yes. Definitely. I believe you mentioned just a few minutes ago that the only case, I agree that the only case where you wouldn't want to use it, is just whether it's too much of an overhead or the ROI is too low for it. If it takes too long to set it up then the time saved once it's up and running. Yeah. Prototyping and things like that, don't use it unless you think it's fun. But in general, yeah. It almost only has benefits. There's so few drawbacks. And you?

**Jacob:** \[00:15:48\] Yeah, I agree as well. And like I said earlier, it's an instinct for me. If I was coming into a project or starting a new project, Git is obviously, just shouldn't even have to be a thought. Of course, you're going to version control your source code and choosing Git for that. And then the CD part. Yeah. I'm in a early stage startup right now and we have Flux set up and we have the capability to just enable it. But right now we are doing manual deployments using scaffolding, Kubernetes. And we just don't have the real benefit for enabling Flux yet, like right now. In a couple of months time, we will be starting to do it. So it's not like I, first thing I do is like, all right, production infrastructure GitOps, pipeline or Flux or whatever. And get everything ready and then start writing the code. It is an incremental thing that evolves over time. But there's no case in a long-term project or a real project. If you get what I mean, that I wouldn't be doing these things.

**Andreas:** \[00:16:52\] So I think from what you said there, it sounds like your pace of work and how many people you are collaborating. They sound like the two factors why GitOps just hasn't been applied yet. You're still moving along slowly. And you're so few people that it's okay to have that manual control.

**Jacob:** \[00:17:09\] Yeah. I think it would slow us down right now. We're in the stage where we're on a feature branch and I want to get it into the environment before it's even merged in a PR. Did I say that out loud publicly? No, but-

**Andreas:** \[00:17:23\] We can cut that bit.

**Jacob:** \[00:17:24\] Yeah, we can cut that bit. This has just been happening. We don't have real customers using our production environment yet. So we're in that golden stage where we can afford to do these things. And if we were doing GitOps formerly and properly, obviously it would be to prevent people doing stupid stuff. And this is a classic case of somebody doing stupid stuff. But it's stupid stuff that I mean to do right now, because it just means that we can move that little bit faster.

It does add overhead, not just the initial setup, but there is a formal process behind it. Git is supposed to act almost like an audit trail for every change that you make and every change that gets deployed. And that doesn't come for free. Especially if you have maybe compliance requirements, regulatory compliance requirements, not just internal ones. Where everything has to be code reviewed and that doesn't come for free. PRs and code reviews don't really cost you anything in terms of money or even that much time to do them. But they do add another stage in your process. And ...

**Andreas:** \[00:18:31\] Yeah, but when you're prototyping and just trying to prove a concept, let's get stuff up and running. We'll build this today, but we might tear it down tomorrow and change direction. That's just not a good situation for GitOps.

**Jacob:** \[00:18:43\] Yeah, I think so. Yeah. If you, yeah, if you really want that, if you really want to be on edge with everything and building constantly and prototype and experimenting, then don't go for GitOps. But as soon as you know what you're actually doing and you want to do it properly and formally, then do do GitOps.

**Andreas:** \[00:19:00\] Yeah. I would say that the last time, or the latest point at which you should enable GitOps in an ideal world would be when something is released into production. Whether you're developing some kind of embedded device or more so an app or a web service or whatever it is. When that goes into production for the first time, you should definitely have GitOps and preferably a long time before that. So you've had time to develop the system and work in that workflow.

**Jacob:** \[00:19:31\] Yeah, I would say before that, definitely because it's not once you set these tools up and tick a box that like yeah, we're doing GitOps, that it's going to run smoothly from there on in. It's going to take a bit of a cultural shift, it is going to take a change in the way you work. And it is not going to be perfect first time around, it's going to take continuous improvement and incremental progress. It shouldn't be left to the last minute, but then much like a lot of the things like testing things do get left to later on.

**Andreas:** \[00:20:04\] What do you think about applying GitOps to something that already exists in production? Let's say a company with 50 developers has a web service that's in production. It's in use, let's say, I don't know, 20,000 users. How would you apply GitOps to that? They have an existing release process. They have an existing way of working.

**Jacob:** \[00:20:23\] I think in incremental change and incremental progress is the best way to go. And there's no point trying to go for a big bang. They're like, okay, we're going to rewrite everything we've done in Terraform and have all our infra done. And we're going to get Vault in to handle secrets. And then we're going to bring a CI tool to build containers. And we use Flux or Argo CD for the deployment and move everything to Kubernetes in that process too. That's just crazy. So I guess just starting small.

**Andreas:** \[00:20:47\] Maybe you could start with the internal parts. Okay. Let's make sure that our internal CI process or the task process, whatever. That conforms to GitOps, and then ...

**Jacob:** \[00:20:58\] That's actually a really good point that you raised, because having these continuous deployment requires you to be much more rigorous with your automated testing. Basically your automated process. If your process involves a lot of manual intervention, and it's not just putting artifacts in places or doing this, but it's actually testing or some quality assurance type of activity, then there's no point trying to move towards an automated deployment because your real problem is no matter how good your deployment is, if your QA activities involve lots of manual work, then that's really the first place you should start.

So maybe go back to CI and get CI working properly before you start trying to do GitOps and continuous deployment. Because I would argue, and this is something I really like to talk about when a lot of people tell us, they just say that, yeah CI, we have automated builds. We're doing CI. And it's like, yeah. One could argue that CI stretches a lot further into the release than just running your build.

**Andreas:** \[00:22:07\] It's a matter of perspective, isn't it?

**Jacob:** \[00:22:09\] Yeah, exactly. And it's not just running a tool and getting output, but it's okay. Isn't it about the feedback of that as well? So it's yeah. I guess it's going back to fundamentals.

**Andreas:** \[00:22:17\] Yeah. It's like the key question there is. Okay, what do you consider integrated? If you make a change to the code base and it still builds, do you consider that integrated? Or does it need to be tested as well? Integrated should be, it's now part of our software. It's now part of our product. That's integrated to me. It's now part of it. And if you've got an untested piece of code in there, is that integrated? Anyway, but that's a rhetorical question. Of course it's not, let's avoid that tangent. That's a one for another chat, I think.

**Jacob:** \[00:22:49\] Yeah, I think so. I think now we've moved away from, GitOps to just talk about CI.

**Andreas:** \[00:22:55\] But I think that's a symptom of GitOps not actually being, like we've said, not very revolutionary. It's more like a collection of best practices. It's CI, CD, configuration as code, infrastructure as code, Git workflow, those things.

**Jacob:** \[00:23:12\] Yeah. And I think actually what we did just now is a very good job of not just saying to everybody that, hey of course, you should do GitOps. But it's actually that yeah, in theory, it would be best if you were practicing GitOps, but maybe it's not perfect for everybody. I just told you my example right now, it's more of a choice. It's not like we can't do it, but it's just I choose not to do it for now. But then also it might be a stage where you are in your project and it's not you're doing things badly. It might just be that you have an older project, you have a code base that's been around for a lot longer. You might be developing embedded devices where you can't just spin up a Kubernetes cluster and run your tests there. And you might be talking real hardware and real physical devices. Everything gets suddenly a lot more challenging. So I think GitOps is about finding a good balance. Finding a good balance in your development delivery deployment process.

**Andreas:** \[00:24:12\] Yep. Not all of us are afforded a perfect development environment.

**Jacob:** \[00:24:19\] Yeah, exactly. All right.

**Andreas:** \[00:24:21\] Do what you can and then aim for GitOps.

**Jacob:** \[00:24:25\] Yeah. Is that the official summary now?

**Andreas:** \[00:24:29\] Should I do GitOps? Well do what you can and aim for it if possible.

**Jacob:** \[00:24:33\] Yeah. Good luck and GitOps. Keep calm and GitOps. Yeah.

**Andreas:** \[00:24:39\] Keep calm and GitOps. Yeah. So I've been Andreas.

**Jacob:** \[00:24:42\] And I've been Jacob, pretending to be anyway. Yeah.

**Andreas:** \[00:24:46\] I'm still Andreas and-

**Jacob:** \[00:24:47\] Yeah, I'm still Jacob. I probably will be by the time we get to next episode as well. So.

**Andreas:** \[00:24:52\] Yeah, I guess it's good time for a pitch as well. If you do want some help with GitOps, come to our website, book a free consultancy verifa.io, V-E-R-I-F-A dot I-O.

**Jacob:** \[00:25:08\] Or if you just want to talk about GitOps, good and bad. Then feel free to reach out to us too. We'd be happy...

**Andreas:** \[00:25:15\] Yeah, we're on Twitter and LinkedIn and ...

**Jacob:** \[00:25:17\] Yeah. Happy to have a rant.

**Andreas:** \[00:25:20\] Yeah.

**Jacob:** \[00:25:20\] Yeah. Okay. Thanks everybody for listening. Thanks Andreas. Thanks Jacob. And catch you next time.

**Andreas:** \[00:25:30\] Yep. Bye-bye.
