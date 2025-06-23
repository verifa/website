---
type: Podcast
title: OSS Compliance and Tools with Lars Geyer-Blaumeiser
subheading: In this episode of The Verifa Podcast, Andreas chats with Lars Geyer-Blaumeiser,
  Senior Expert Open Source Program Office at Bosch, who shares his insights on open
  source compliance and tooling.
authors:
- alarfors
- avijayan
tags:
- Open Source
- OSS Compliance
date: 2021-06-27
image: "/blog/2021-06-28/ep03_-podcast_9201080_small.png"
featured: true
---

<iframe title="Embedded podcast player" src="https://anchor.fm/verifa/embed/episodes/OSS-Compliance-and-Tools-with-Lars-Geyer-Blaumeiser-e13bptn" height="151px" width="100%" frameborder="0" scrolling="no"></iframe>

<div class="flex gap-x-4">

[![Listen on Spotify](/blog/2021-03-30/listen-on-spotify.png)](https://open.spotify.com/show/12yStrneLdEsXn1Bjp6Myz)

[![Listen on Apple Podcasts](/blog/2021-03-30/listen-on-apple-podcasts.png)](https://podcasts.apple.com/gb/podcast/the-verifa-podcast/id1561051552)

[![Listen on Google Podcasts](/blog/2021-03-30/listen-on-google-podcasts.png)](https://www.google.com/podcasts?feed=aHR0cHM6Ly9hbmNob3IuZm0vcy81Mzg0NzE1Yy9wb2RjYXN0L3Jzcw==)

</div>

**Welcome to The Verifa Podcast, where we break down the complex world of Cloud and DevOps.**

In this episode of The Verifa Podcast, Andreas chats with Lars Geyer-Blaumeiser, Senior Expert Open Source Program Office at Bosch. Lars shares his insights on why open source compliance is vital for managing license risk and security risk, how you can effectively manage these risks and the tools required to do so.

## **During this episode we discuss**

* Why is open source compliance important? \[02:55\]
  * Managing OSS License risk \[03:23\]
  * Managing OSS Security risk \[09:30\]
* How do you manage risk when using open source software? \[14:31\]
  * Scan, Evaluate, Monitor
  * Challenges for IoT and embedded systems \[24:58\]
* What open source compliance tools are available? \[27:44\]
  * Evolution of open source compliance tooling
  * OSS Review Toolkit (ORT) \[34:21\]
  * How to manage OSS compliance technical debt \[41:29\]
  * How to build a database of OSS License data \[46:16\]
  * Standards for OSS Component Metadata \[48:11\]
* Lars’ ideal OSS management tool chain \[50:40\]

## **Mentioned in the podcast**

[Lars’ talk at EclipseCon2020 - Automated Open Source Compliance in Action](https://www.youtube.com/watch?v=_3r4XfMJBUA&list=PLy7t4z5SYNaSqqMEaZgVFyyF1Uoqa8Q3F&index=52)

[Lars’ talk at EclipseCon2019 - Automating Open Source Compliance with OSS Tooling](https://www.youtube.com/watch?v=aWwROyjTuCk&list=PLy7t4z5SYNaT_yo5Dhajb9i-Pf0LbQ3z8&index=69)

[OSS-Review-Toolkit](https://github.com/oss-review-toolkit/ort)

[SW360](https://www.eclipse.org/sw360/)

[Open Source Reference Tooling Work Group](https://oss-compliance-tooling.org)

[SPDX](https://spdx.org/licenses/)

[CycloneDX](https://cyclonedx.org/)

## **About Lars Geyer-Blaumeiser**

Lars is working in the central department for IoT and Digitalization at Bosch. He is responsible for strategic Open Source projects in the Open Source Program Office within Bosch, supporting Open Source projects throughout the company. In his previous role he was project lead for Eclipse SW360 and an active participant in the definition of an Open Source based methodology to automate Open Source Compliance activities within the Open Chain Reference Tooling Group. He got his PhD from the university of Kaiserslautern.

## **Connect with today’s podcast team on Linkedin**

[Andreas Lärfors](https://www.linkedin.com/in/andreas-l%C3%A4rfors-51253270/)

[Lars Geyer-Blaumeiser](https://www.linkedin.com/in/lars-geyer-blaumeiser-6002572b/)

[Anoop Vijayan](https://www.linkedin.com/in/anoopvijayan/)

## **Transcript**

### **Hello and welcome!**

**Andreas:** \[00:00:12\] Alright. Hi. Welcome to another episode of The Verifa Podcast. I'm your host, Andreas Lärfors, and with me I've got Anoop and Lars. Lars is our guest for this podcast. He's long term, long time OSS compliance leader from Bosch, so he's got a lot of interesting stuff to talk about. So I think we should get into it. Let's start with a brief introduction to say who we are and what our backgrounds are.

I come from a software development background but spent the last 10 years in CI and DevOps and more recently cloud. And I probably have three years or so of OSS compliance management as a service or as a consultancy service using both commercial and open source. So, Anoop, do you want to give a brief introduction? I know you're quite new to the whole OSS compliance thing, but you can just talk about briefly your knowledge area and so on.

**Anoop:** \[00:01:05\] Yes. I'm Anoop. I'm a cloud architect at Verifa. 15-plus years of experience doing DevOps and system administration things. Quite happy to be here meeting new friend, especially Lars, and quite looking forward to know more about OSS compliance. Thanks.

**Andreas:** \[00:01:26\] Yep. Nice to have you here. And then Lars, if you would please.

**Lars:** \[00:01:31\] Yeah, of course. Thank you for the invitation. So my name is Lars. I was project lead of the now archived Project Software 360 Antenna and also still project lead of the Project Software 360. I started working in this area 2017 after 10 years of extensively using open source, the Eclipse IDE, basically building proprietary in-house tools at Bosch on top of Eclipse, also later then commercial products.

And I started at another subsidiary of Bosch which was about open source services. And one of the service we provide is open source compliance tooling and also open source compliance methodology. And when I joined, it was just the time that the in-house tooling for open source compliance which was built on Sonatype IQ with an in-house tooling to put this to the open source.

So this tooling around IQ became Antenna. And when I joined, it was just the time when we had the initial contribution of the Eclipse project to put in place. Another thing that happened at that time was the foundation of the now called open chain reference tooling work group. At that time, it was a loosely coupled working group from different companies, quality tooling landscape, where we try to assess the situation of open source compliance, especially open source based tooling for open source compliance.

Yeah, that's me.

### **Why is open source compliance important?**

**Andreas:** \[00:03:05\] Great. Thanks. We will dive into a lot of that a bit later on. But I think just for clarity's sake, we're going to recap why we need to do OSS compliance, why we need to manage our usage of OSS, and then how we manage the usage of OSS. So I think the two main points as to why come down to licensing when it comes down to risk, and one of those risks is licensed risk and the other type of risk is the security risk.

### **Managing OSS License risk**

**Andreas:** And I think the licensing one is the one where most companies or most people get involved with to begin with because it can be a big legal matter. Security is more reserved for cybersecurity teams, but also developers who knows the software on a technical level and can figure that out. But the licensing is a ... it's a legal matter.

There is a question as to how you use the OSS component, as to how the license applies. But if we can just reach an understanding here of licensing, my understanding is that copyright is a necessary part of an OSS component. Each OSS component has a copyright, and that identifies the owner of that component, if that's people or an organization. And then the license tells you how you can use that OSS component.

How you're allowed to use it, what you must do if you use it, and so on. The conditions, the terms of using that OSS component. Have I got that roughly right?

**Lars:** \[00:04:38\] In principal, I think you're totally right. So the thing with copyright is that you simply cannot use any copyrighted material until you get an allowance, and this is exactly what the license is all about. And that's something that has to be understood when using open source because you can share some code on GitHub that's easy, but in a legal way, to use it you need simply the license.

And that's why companies, especially big companies, make a fuss about it. That they clearly understand what the license situation is so that they can fulfill all the obligations that come with those licenses.

**Andreas:** \[00:05:17\] Because a very naive way of looking at OSS is that it's a free software.

**Lars:** \[00:05:23\] Yes.

**Andreas:** \[00:05:24\] Yeah, but it comes with conditions. I guess as they say, it's not free, like free beer. And there are two main categories of licenses too. There's the copyleft, which is the one where people are actually concerned about. So that typically states that if you can use this piece of open source software, but if you do, the software that you make must also be open source.

It must be released under the same license. So that's like the GPL and so on.

**Lars:** \[00:05:58\] Yeah. Even copyleft comes with two flavours. The weak and the strong copyleft. What you've mentioned is actually the strong copyleft. The legal term is derivative work. So it's a piece of art in the copyright sense, which is licensed, and derivative work is what you make out of this piece of art that you get, that you use. And copyleft licenses claim that derivative work has to be licensed under the same license.

But they restrict what derivative is. Strong copyleft, basically all the software that is using your software. But there are also legal things when this applies. So for example, is dynamic linking is still the same as static linking and stuff like that. There is also the weak copyleft which only applies to the software itself. So if you use a piece of software and you change it, you have to publish the changes under the same license.

This does not mean you have to put them upstream into the open source project, but your customer to whom you give the software, you have to give the source code of this piece of software you're using with your changes so that he can rebuild the software.

**Andreas:** \[00:07:14\] Okay. And I guess for things like build tools and test frameworks, let's say, just as an example, if you use something to build your software more as a utility, it's not actually embedded within the software product than ... Not even strong copyleft counts there.

**Lars:** \[00:07:30\] The licenses typically set obligations when you redistribute the software. So when you put it to someone else, or when you're only using the software, you can do and typically without any applications. There are modern licenses like the AGPL who also cover, for example, the software as a service case, execute the software. So you're not giving the software out but you run it in your computing center.

But even then, the license obligations apply, which is not the case for many of the licenses that put in place prior to software as a service, as a technology.

**Andreas:** \[00:08:09\] Yeah. I guess a lot of these licenses are quite dated by now and they haven't kept up with software as a service, or software as a service didn't exist on the scale that it does now.

**Lars:** \[00:08:20\] Yeah.

**Andreas:** \[00:08:21\] The other type of license categories, typically, which essentially says, "Do what you want with it." I think there's even a, "Do what you want with it," license or it might not be-

**Lars:** \[00:08:29\] There's I think even as a swear word.

**Andreas:** \[00:08:32\] Yes.

**Lars:** \[00:08:33\] Yeah. Not necessarily do what you like with it, but because typically the copyright owners still want to be mentioned and the copyright still stays in place. But you can use this software on a free range. So you can put it into your products, you can sell your product. And as long as you give credits to the copyright holders, this is typically the standard application that comes.

**Andreas:** \[00:09:04\] Yeah. Yeah. So the thing that both permissive and copyleft have in common is that if you use open source components in your product, if you distribute open source components as part of your product, you must at minimum state which open source components those are and include the original licenses. It's the thing that it all has in common. Yeah.

When we talk about OSS compliance for companies distributing software or hardware containing software, this is the baseline. They need to be able to produce a list of OSS components in their products and include the original licenses.

**Lars:** \[00:09:40\] Right.

### **Managing OSS Security risk**

**Andreas:** \[00:09:41\] Yeah. That's a baseline requirement. So that's a little bit about license risk. And then there's the security risk. And this is ... It's a double-edged sword, because if you have ... You basically have a team of developers working ... If you use an OSS component, you have almost a team of developers working for you. If somebody identifies a security vulnerability, maybe you don't have to fix it yourself.

There's already an OSS community out there who are maintaining this OSS component. So you can just wait for them to fix it for you. But this throws up a whole new host of issues as well. Not only can you rely upon that community, that depends. But then also, how do you tell the good actors from the bad actors? Is it possible to insert malicious code in OSS components?

So it's like you've outsourced and taken in OSS. In some way, you've outsourced the management to a third-party that you can't verify the intentions of. There's a small risk of that, right?

**Lars:** \[00:10:44\] Yeah. But here we come into an area where we talk about the selection of open source components, and open source as a right range. So you have some code dumps on GitHub on the one hand side, and you have foundation-based projects with good governance and a clear process on who has right access to the project, and so on. On the other hand side, you have company-driven open source projects, where the copyright is basically held by one company.

So there's lots of different dimensions which decide on the quality of an open source component. But I think it's common sense that security ... that open source favours the handling of security issues because it's open. Vulnerabilities can be easily found. They can of course be used, and also easily if you're faster than the good guys in finding them.

But in principle, there is a bigger interest in solving those. With proprietary code, it's harder to find, but the bad guys they know how to find those stuff as well. And I also saw this from my experience that in such a case companies think about how to deal with the vulnerability. If the vulnerability is open, it's clear, it's there, there is no discussion.

But if you are responsible in how your customers update, for example, their components, you're very reluctant in publishing the information. So then it becomes really difficult to have a clear strategy. I think that the strategy with open source is much easier to deal with. And for example, there are nowadays also tooling that allows you to check whether vulnerability is really affecting you.

Because using a big open source component does not mean that you'll run the affected code because you'll need user part of the open source component, potentially a part that is not affected. I think, yeah, of course there are ... It's not the bright new world without any issues, but in total I think, the advantages are much higher than the disadvantages when it comes to open source.

**Andreas:** \[00:13:08\] Yeah. And I guess that there have been headlines. I have seen reports of malicious code and bad actors. But I think seeing those is a good thing. It means they're being discovered, they're being identified and then they're being reacted to. So that's also ... When you have so many ... Essentially with OSS, especially a widely used OSS component, you essentially have a lot more stakeholders than you would with a proprietary piece of code.

**Lars:** \[00:13:34\] And look at what happened lately in the Linux kernel. It was a massive reaction from the community and from the response from people when ... Was it some university who tried to check what happens when they enter a malicious code?

**Andreas:** \[00:13:52\] Yeah.

**Lars:** \[00:13:52\] And yeah, I think ... And especially in such projects like Linux, there are so many people who are really doing this by heart and with a lot of effort and emotions and they are really interested in delivering good software. And that's something that you can't expect always when someone is paid to do something he doesn't like perhaps.

**Andreas:** \[00:14:18\] Yeah.

**Lars:** \[00:14:19\] So from a quality perspective, I think open source is at the top. Not all, of course, and not every project, but I think from software quality, when you put range, put the best projects into a sequence, I think the top projects would be open source. That's what I'm convinced on.

**Andreas:** \[00:14:38\] Yeah. And I think I would be inclined to agree with you. There's such a wide range too though. You have open source projects that are used by millions each day, and then you have tiny ones that are used by maybe 10 or 100 people. So it's a wide range. The Linux kernel case that you brought up, that's a really interesting story for anyone listening who wants to go and have a read. I believe it's the University of Minnesota.

**Lars:** \[00:15:03\] Yeah, I think Minnesota.

**Andreas:** \[00:15:05\] And they ran a research project to study how easy or difficult it would be to insert malicious code into a project like the Linux kernel. And so they actually did that research by inserting malicious code into the Linux kernel. I didn't follow the fallout too long afterwards, but the initial fallout was that they banned all contributors from Minnesota University and began back dating, removing their contributions too.

That's what I saw from that. So a swift response from the community.

**Lars:** \[00:15:41\] There was also this other case of this developer in the Linux kernel who made a business model out of it by suing companies who didn't follow the GPL applications completely, and he got also thrown out. So it was again a very swift and massive reaction by the community by removing the code because of course they are interested in companies keeping the license obligations.

But they're not interested in people creating business models out of that or exploiting the idea. So that's, I think not really a big danger in the big values.

**Andreas:** \[00:16:23\] Yeah. There's also the concept of ... Coming back to licenses briefly, there's the concept of license trolls, right?

**Lars:** \[00:16:30\] Yeah.

**Andreas:** \[00:16:31\] Who I believe, and correct me if I'm wrong, I believe their mode of operation is that they will take a product with containing software. They will somehow identify an OSS component being part of that product with a license breach by the company. And then they will attempt to blackmail the company into paying them off. "If you don't pay me this amount of money, I will expose you as breaching these license terms." Is that about right?

**Lars:** \[00:16:57\] Yeah, something like that. I don't know if I can explain that in English properly. But what they do is they point out that a company has done something wrong concerning the obligations to the company. And there is a mechanism in law where you can do that as an owner of the copyright. Yeah. You can do this and make a fine or define a fine if they do this again, that they have to pay the fine.

So that's a legal act, point that out, with the fine. And the company then has to accept that they have not fulfilled the obligations. If not, they get sued. So they typically do, but with big companies, it's this problem that you can't guarantee that this will not happen again because those guys who were affected by this, they know now what the obligations are.

But the other parts of the company, even they could have done this already and put in place. And so then the second case, that's where the money comes from. They find the second case and then tell the company, "Hey, you guys. You have signed that you will never do this again, but here you do it again. So you have to pay the fine."

**Andreas:** \[00:18:18\] So it can be a completely different group of people who've breached the license agreement?

**Lars:** \[00:18:22\] Yep, but it's the same company.

**Andreas:** \[00:18:23\] But because it's the same company ... Yeah. Okay.

**Lars:** \[00:18:25\] And this works with big companies because there are so many business units. If some get sued, the rest will not recognise.

**Andreas:** \[00:18:32\] This highlights as well how important or how effective OSS management can be on an organizational level-

**Lars:** \[00:18:40\] Right.

**Andreas:** \[00:18:40\] ... instead of just doing it on a team level or a project level. And that's one of the mission statements of Software 360, right?

**Lars:** \[00:18:48\] Yes. The thing is to ensure traceability. To know which components are used within the company and to know where they are used. And that's one of the core cases, the core use cases of open source management, to really get the overview.

**Andreas:** \[00:19:05\] Yeah, because one of the ... In big companies, especially, one of the challenges is to just know what's right and what's wrong. You may have legal advisors or something available to you or maybe not, but to collect everything in one place and be able to see that, as an organization, this is how we operate. These licenses are essentially white listed, they're okay.

These ones are gray listed, they must be investigated on a case-by-case basis. And then these are blacklisted, you must basically never use or distribute anything containing an OSS component with this license.

**Lars:** \[00:19:39\] Yep.

### **How do you manage risk when using open source software?**

**Andreas:** \[00:19:40\] So we've been through why we need to manage OSS briefly on how we do it. It's really three steps. We scan, somehow, to identify the OSS components. Then we evaluate the risk, the license risk, the security risk, and then post-release, we have to make sure we monitor, because the security ... The license data is usually static. There's been a case recently where a license was changed.

I can't remember which one that was ... Was it-

**Lars:** \[00:20:08\] The Ruby case.

**Andreas:** \[00:20:09\] It was the Ruby, yeah. Do you want to-

**Lars:** \[00:20:12\] They used ... Yeah, but that's always the fear of every open source officer, that you use a component and there is an undetected GPL license piece in there. And the GPL simply has this rival effect because all that is using this component is derivative work. So, why are the open source component? Your code gets into the situation that it has to be licensed under GPL.

That's a rare case, but this is one of the fears of open source officers, that something like this gets undetected into the products. But, I mean, in principle from a law perspective of course, you have used this license, but in a normal case there will be solutions to get around this when you react fast and swift.

**Andreas:** \[00:21:03\] Yeah. In the Ruby case, I think they just ... It was actually a dependency of Ruby with the GPL license. So they just changed the dependency, issued a new version and then everybody could-

**Lars:** \[00:21:15\] Exactly.

**Andreas:** \[00:21:15\] ... continue. Yep.

**Lars:** \[00:21:16\] So I think that also the reaction was very swift, very clear, and pointed out exactly the issues that make the usage of open source from, "Hey, let's download something and use it to some work that has to be done within the company," especially when you want to sell products on top of the open source.

**Anoop:** \[00:21:34\] So what about the components which don't define any licenses? So there are some GitHub repositories and other things which actually do not define any licenses. What's your thought of them?

**Lars:** \[00:21:44\] The law is pretty simple. So there's actually a difference between the US and the European law, because in the US you really get something like public domain. So you can make source code public domain, which was for a long time not accepted in Europe because this possibility is simply not there in European law. And the copyright law in principle is very clear.

You can't use copyrighted material until you have a license. So if someone simply pushes some stuff to GitHub and it looks nice, you're not allowed to use. Nowadays, there is the interpretation that if someone clearly states that he is giving this to the public domain, he is issuing something like a license. He agrees that the software is used as you like.

It's really a statement. But if there's nothing like that, only the source code, keep your fingers away from that. You're legally not allowed to use it. It's very simple.

**Andreas:** \[00:22:47\] And not just you're not allowed to redistribute it, you're not even allowed to use it.

**Lars:** \[00:22:53\] Yeah, it's copyright. It's like as I said, it's a piece of art. If you don't buy a piece of music, the right to listen to a piece of music, you're not allowed to listen to this music. And the same applies to software. It's simply a concept that has been introduced. I always ask myself why this has been introduced to software, because it was the closest relationship to what software is.

Because copyright actually was done for art, for works of art. But like also a textbook and stuff like that, and since books are some printed letters and source code are some printed letters, perhaps that was the relationship that has decided that this is copyrighted material.

**Andreas:** \[00:23:40\] Because I don't think I've ever seen a license in a book I've read. You don't usually open a book and see, "This book is provided without warranty or guarantees."

**Lars:** \[00:23:51\] That's right, but it's also typically not open source. So there is warranty from a legal perspective and there's a copyright in there. And by buying the book, you get the license basically.

**Andreas:** \[00:24:03\] There is typically, all of this work is fiction and a new representation similar to other characters.

**Lars:** \[00:24:10\] \[Yeah\] .

**Andreas:** \[00:24:10\] Yeah.

**Lars:** \[00:24:11\] That's another case.

**Andreas:** \[00:24:13\] Yeah. Okay. Let's try to bring this slightly back on topic. Yeah. You do have to monitor your releases. So you've released a piece of software, it's out there, people are using it. That doesn't mean you can just say that the work with OSS is done. So as we've discussed, the license can change. It usually does not, but it can. But I guess the data that moves more is the security vulnerability.

**Lars:** \[00:24:33\] Right.

**Andreas:** \[00:24:33\] So security vulnerabilities are identified daily. They're kept in the NVD, the National Vulnerability Database. So there's a community drive to identify security vulnerabilities. And then of course, within each OSS component, there's a drive to fix those. So that means that using OSS is a repeating process, a continuous process of, are there any new security vulnerabilities introduced that are now possibly in my product? And do I need to push out a new piece of software?

**Lars:** \[00:25:05\] Yeah.

**Andreas:** \[00:25:06\] And this ... Sorry.

**Lars:** \[00:25:08\] No. Go on.

### **Challenges for IoT and embedded systems**

**Andreas:** \[00:25:09\] Yeah. So this is obviously if you were running a web service or software as a service, great. You can probably push out a release the next day. But within embedded software where I spent a lot of my time it can be a different matter. You've got the challenge of vehicles, for example, which, okay, now we're getting over the air updates. Companies like Tesla are actually pushing updates daily to their vehicles.

But if you go down into low-level embedded ... I spoke recently with a company that makes the cranes in harbours around the world and they told me that an hour of downtime in a harbour, that's millions of dollars gone. It's in millions of dollars per hour, so they just cannot afford any downtime. They have to thoroughly test all the software that goes out. So they have a much harder time.

Even if they want to push out an upgrade or an upgraded version, a new version of their software, it can take months for the customers to actually begin using it.

**Lars:** \[00:26:06\] Yeah, absolutely. And the thing there is that times are changing. And for those companies, this is really a game changer. Because you mentioned the car, I recently talked to another business unit here who does fire alarm systems, and those systems are isolated. They're not in the network. A car never was in the network in the past. Even if there is a vulnerability, you need physical access to the device to make use of it, which is quite hard.

But nowadays with IoT and everything connected, this becomes a totally different threat and even parts of a car, which are even still not connected but connected via the car network to those parts who are connected can be attacked by outsiders over the network. And that's simply a game changer which requires that also the whole development methodology has to be reworked because people were used to, "Okay, we built this product, we make the software, we put it on stocks, we produce it."

But there is not a big maintenance pressure on the software, because as long as the product does what it has to do, everything is fine. But nowadays, you really have to support the devices, 10, 15 years. And with such embedded things, it's typically 10, 15 years, and not only two years or three years like in the smartphone case. And we all know it from the smartphone, from the Android case, how bad it is to get updates on even only two years when you have the wrong phone.

**Andreas:** \[00:27:43\] And now of course we have Android Auto is built into many modern vehicles. The whole instrument cluster, the whole entertainment system can be based on Android, just as an example there.

**Lars:** \[00:27:55\] Yep.

### **What open source compliance tools are available?**

**Andreas:** \[00:27:55\] Yep. It's going to be interesting to follow developments there, but that's okay. So we've covered why we need to manage OSS and we've covered how we need to do it. We need to scan, we need to evaluate the risk and then we need to monitor post-release. Of course, this would be a lot more difficult if there weren't tools available. And I think that's where we're getting into the stuff that we're most passionate about now is the tooling around OSS.

I guess from my view, this has long been a space dominated by commercial tools. It's been companies basically with money making tools, selling those tools to manage OSS. And the concept of open source tools that manage OSS has come more recently in the last four or five years. Is that about right?

**Lars:** \[00:28:44\] Roughly, yes.

**Andreas:** \[00:28:45\] Yeah? I guess the first question then, why did it take so long? Why did it take so long for the community that works on OSS to make tools ... If we talk about the OSS community as one community, why did it take so long to go from making OSS components to then making tools that helps people manage the usage of those? Is it because it's just a matter of when commercial companies are using the OSS components?

**Lars:** \[00:29:13\] Yeah. So I think it's a long story, to be honest. And in compliance cases, typically happening was that it was the fear that would drive this whole business. The fear of being non-compliant, the fear of using some GPL software and thus forcing you to open up your own proprietary code and stuff like that. The fear of being sued and have a reputation loss.

And that was what the companies ... and actually at that time, for example, in embedded software, open source was not an option. Open source management was about denying open source and ensuring that you're not using open source. And that was where all the snippet scanning companies came up and basically made the market for open source compliance.

That's the logical thing to do. If you want to ensure that you don't use open source, you use such a snippet scanner who tries to figure out if you're using some known public code and you're fine if there isn't any finding. And this has been established in many companies, and they were quite successful in pushing that fear and pushing that solution.

And big companies simply have this tendency, if they can buy ... If this is a solution, that's fine, especially in such a case, because there is no guarantee. But what you need in the end is not a guarantee but you need the approval that you're using state-of-the-art technology, because when you get sued, you have to show that you did state-of-the-art methodology to find out that you did something wrong.

Only then, you don't get fines too hard.

**Andreas:** \[00:30:57\] Oh. So, I guess, if you breach a license agreement and you didn't do anything to prevent that, that's worse on a legal perspective than if you breach it but you did everything you could to not breach it?

**Lars:** \[00:31:12\] Definitely.

**Andreas:** \[00:31:12\] Oh. I thought it would either be a breach or no breach.

**Lars:** \[00:31:16\] Yeah, but there are really strange cases. So what do you do in this Ruby case? So the guy who are using this dependency, it's his best to his knowledge to make a clear statement for his component. But he didn't understand that using this GPA component was an issue. And there are really weird cases. People who don't state their copyright in a way that could be easily found.

You can't go through all the files and really check every line, whether this is something. So you use tooling to detect copyright statements. And some things are so weird that they are not found. So, errors happen. And that's what you have to understand in this area. Of course, it's a breach of the license or not, but it's always also the situation, the case under which this has happened.

And if you're simply naïve and think, "Oh, everything is fine," that's definitely a different case than when you can prove that you have an established process in place, that you did all the things that, in this case, this has also been executed and for some reason nothing was found. That's one point, and that was why a commercial solution is really attractive for a company.

Because as long as this is state-of-the-art, they can claim, "Hey, we have bought ... We have spent so many money on this tool. We have the state-of--the-art."

**Andreas:** \[00:32:40\] We have the best tools.

**Lars:** \[00:32:41\] That's it. Yeah. But on the other hand side, what you saw is a lot of tools came up. in-house in companies. I mentioned, we had this tool that later became Software 360 Antenna. Software 360 was an in-house tool at Siemens, toolkit was developed here for the internal purposes first. So a lot of companies detected that those tools, the commercial tools had gaps and they built tooling around the commercial tools.

### **OSS Review Toolkit (ORT)**

And this was the jumpstart for all this open source activities basically, that people detected that they have gaps and they have to fill the gaps. Another thing that happened was that when open source usage became more and more popular, people simply saw analysing an open source component with a snippet scanner simply does not make sense, because what it detects is that you're using open source.

So in the first place. And that it's much better to check what you're really using. If you look at modern languages; Javascript, whatever, they all have package managers. So you don't copy some open source into your repository and use it, but you simply state, "I'm using component X, Y, Z in version four." And the built system drags that into your build and makes use of that.

And in the same way, you can get the list out of your package manager, what he will take in. And with this usage, you also get this problem of transitive usage. So, dependencies of dependencies. But that's a simple call. In Maven, you make a Maven dependency list and you get a list of all the dependencies. And this was a change in how the methodology was about identifying the components.

So you simply ask your system, "What am I using?" And then for each of the components you're using, you check what is the license situation. But even there, you can automate a lot. That's exactly what the OSS Review Toolkit is doing. For many package managers, that it simply tracks in all the information you can get and even runs a copyright scan.

So not a snippet scan but a tool like FOSSology or ScanCode, which checks for copyright statements and license statements in the source files to get an overview of whether what they declare as their license matches to the findings in the files, which is often a deviation. And with simply using the OSS Review Toolkit in the end, you get good observation of which open source you're using.

What the status is in terms of, can this be found? Is the source repository available easily? What is the state ... What would they claim from a license? What is found in the code? And so on. And there are other companies who also do commercial tooling in this way, but I think they're really in the open source world now as state-of-the-art and delivers the tools you need to do this kind of thing.

But the difficulty here comes with the technology. So you have to support each and every package manager, and you have this unfortunate world of the C and C++ developers who typically do not use any package management but still copy the third-party code into their repositories. In the OSS Review Toolkit, we actually lately added a possibility to do the detection manually because this sounds difficult and you have this manual step.

But in the end, if you have a repository and this is all your code that you built, you don't have this transitive dependency issue because your code is complete and you copied it in. So, in principle, the knowledge, what was copied in is there. So it's not a big step to make the development team responsible for claiming what they have entered.

The rest of the processing is again done automatically, but they need the starting information to understand what is in. And there is no good solution. So I know from a commercial tool, what they do, they have an archive of known files. But then in the C/C++ case, you again have this issue that you not only copy the open source into your repository but you're slightly adapted, and then this mechanism fails.

So you thought the best solution is really to make them claim what they are using for this case. It's possible, it's not too difficult, and do the automation after the step and still have a good situation when it comes to identifying the dependencies and processing it throughout the whole work life cycle of the compliance management.

**Anoop:** \[00:37:40\] So, one question. Now we tumbled upon very interesting things. You mentioned about transitive dependencies. And if you look from a developer point of view, especially when running modern tools like PIP or NPM, we typically put some version or put the latest. And what's in your opinion would be a good phase in the software development or during the pipeline to test this part?

So if we look at, towards the left, extreme end where the developer does it, he changes things a lot. So probably that might not be a good spot. But then by the time of the release, if you do the scanning checking, then it might be too late. So what's in your opinion would be a best-

**Lars:** \[00:38:19\] Yeah. Yeah. We're in the phase of continuous delivery, aren't we. So the time of your commit is not that far away from the time of the release of that stuff. Actually, the OSS Review Toolkit allows you to process this on a daily basis, or even on a pull request base, for example. Yeah, you can check every pull request, whether there has been some additions or changes concerning dependencies.

And this doesn't cost too much time. So perhaps not for every pull request, but at least on a daily basis, I would run my tool chain, detect if there is a deviation from the previous thing. And I mentioned open source compliance as a service, the whole methodology is about identifying what you're using, mapping what you've found to the knowledge you already have.

So in the end, if you want to work in a clean way, you have to approve the metadata of every component. In a sense that, as I said, you have the licenses that are claimed by the project, you have the licenses that are found in the source code of the component, and if this deviates, you have to decide what is the situation? You have components. My good best example is always, you have a component which has a BSD 3-Clause License in the root folder.

And then in every file you have, in the header, this file is BSD License. BSD is no license. There are three different BSD licenses with two, three, and four clause. The files actually have not a correct license statement, but there is this BSD 3-Clause License in the root core folder, so you can determine that they mean with this ... We have this file, this BSD License, that they mean this 3-Clause License.

But this is a deduction that has to be done by someone who's trained in such things. This is not an automation you can do. And the methodology also gets the metadata into a state that we say, "Okay, this is approved by the open source back office." They check each and every component used. In most cases, this is easy. They state Apache 2.0 and they only find Apache 2.0, so no big deal.

The copyrights are also extracted well by the tools. You can simply automate this, but in the end, you have approved information of components. And then in your CI pipeline, you run, you identify the components, and then you can simply find out, "Do I know all metadata for all those components or other new ones I haven't seen so far?" And for them, you can start this process automatically.

So from a project perspective, they would add a new component that has not been used in this company. The tooling would run through and in the end they would see, "Okay, you're not ready to release because there is one open compliance topic." This is already triggered. The back office would take over this information, process it, approve the meta information in the database.

And then on the next day when they run this again, they would in the end see, "Okay, everything clean. Metadata is available. You can go with a release, everything fine." And this methodology, this is simply modern ... from my point of view, modern open source compliance management.

### **How to manage OSS compliance technical debt**

**Andreas:** \[00:41:39\] Wow. Yeah, that sounds fantastic, having seen how some companies do this. That is the way I would recommend them to do it. What a great way to get people to work, because it's always the ... Yeah. I think in a perfect world, they would only be ... you would only detect the new changes. I think one of the challenges is that there's usually a huge technical debt here.

A huge backlog to get through. People are actively right now adopting an OSS management process. So if they haven't scanned before, if they haven't managed it well before, they could already have hundreds of thousands of lines of code. They could have hundreds of OSS components in many of their projects. And so they face this massive ... They start by, "Okay, let's do a scan and see what we've got," and then, "Oh, wow.

We have 100 OSS components here." And then you have maybe 50 software projects that all look similar. So you have this massive amount of work. You go from not doing anything at all thinking, "Hey, let's start managing our OSS," to suddenly, "Wow, this is more work than we ever expected." And then you've got to run just to stay where you are. You can't run to catch up.

You're just running to stay where you are, to manage what you've got. But I guess that's just a question of an efficient workflow and the right resource. Have you seen them being solved in some other good way?

**Lars:** \[00:42:56\] Yeah. So what I just described as how we think from a big company's perspective, and there's always this thing that for a big company, it's always more critical to be compliant because you're so attractive to get money out. Not meaning that not everyone has to be compliant, but the risk is simply on a different level. So for example, the Open Source Review Toolkit, what it does is it gives you quite a good overview of ...

It takes a while to run because they check every component with ScanCode, or in the worst case, I had the project running for a week for the first analysis. Yeah, but if you use ... And for example, if you use JavaScript, you don't end up with hundreds of dependencies. You end up with potentially thousands of dependencies, because this is an ecosystem which has lots and lots of components, and verifying components.

And it simply takes time to identify the component, to download it, to run ScanCode on the sources and stuff. So this simply takes a while. But in the end, what you have is really an overview of which components you have, which licenses that they have, which licenses have been found in the files. And in some ecosystems, like in JavaScript, this is typically not diverging too much, because the components are small and they have ingrained mechanisms to state a license, at least.

And so they don't have this culture of copying stuff in. They simply use it as a dependency. You get all the information, and typically you get a clear picture on what you have to work on for this project. Having an approved metadata makes sense if you're in a company who's really reusing the same open source component again and again. Because then it makes sense to synchronise the analysis of all the component because it's really saving effort.

But if you basically have one or two projects, three projects, the synergies between those projects might be too less to really have this approval of workflow. But then again, you get the list of components and you can see, does it make sense to have a deeper look? So where, for example, is a divergence between the found license and the claimed license, and you can have a look at that.

So even if you're starting, and even if you have a legacy, if you have a modern language where you can detect your dependencies, this is not too difficult. It becomes hard if you have grown a C/C++ repository with lots of developers long gone who have entered stuff. And again, that's also a training issue. For example, if you're coming from university, you learn Stack Overflow is your friend.

That's a normal thing, to go to Stack Overflow. Read and copy the stuff into. That's copyright infringement, unfortunately. So you have to train people. But from my point of view, if you train the people, if you spend the effort in training them, the risk is mitigated a lot. So then you talk about accidental issues and not about things done on purpose because someone didn't simply know that he's doing something wrong.

**Andreas:** \[00:46:27\] Yeah. So it's not just an issue ... or it's not just a problem with tooling and process, it's also a cultural issue within a company. You've brought up this license conflicts quite a lot, which is surprising to me. It's something that I haven't seen that much of. It may be because I work with a lot of mostly commercial tools on the project where I work, who maintain their own database and they tell you, "This is the license."

### **How to build a database of OSS License data**

**Andreas:** But can we just touch upon ... There's obviously the NVD, which is National Vulnerability Database. It's where we gather the known security vulnerabilities of OSS components. What's the equivalent for OSS licenses? Is there a central repository where we can go and say, "I have this OSS component. Tell me what license."

**Lars:** \[00:47:12\] Not on this quality level, because that's a national initiative, the NVD and there's a whole process around that. There is actually something called ClearlyDefined. That's a project that was pushed by Microsoft a lot. What they basically do is they run FOSSology and ScanCode scans on every open source component they can grab on, and store the raw results.

And then they have a mechanism, a process to curate the data. So you can, for example, if there is no link to the source repository, you can fix that. You can fix ... or you can harmonize or normalize the license statements. There is the SPDX license identifiers, which is a clearly defined way of expressing the license of an open source component. But people tend to use some strange variations of this.

So you can normalize this so that you have a machine-readable license or statement.

**Andreas:** \[00:48:21\] Right. Because SPDX, that's software package data exchange. And initially, it's just a data format to uniquely identify a software package.

**Lars:** \[00:48:32\] Exactly. It's an exchange format for bill of materials, for components and their metadata, especially from the license compliance world. So you can trail down unto the files and even snippets for license information, but in general, it's a good tool to share the information you have.

So since you have to fulfill the obligations, I think we haven't touched that point, which as you mentioned earlier, typically means that you have to give the list of components you're using together with the copyright and the license information, as well as sometimes the source code. So you have to gather this, and if you want to process this in a supply chain, SPDX is a good format to transfer the information to your customer who then can make use of the information. Yeah, but that's SPDX.

**Andreas:** \[00:49:23\] How close would you say that is to being a standard? How-

**Lars:** \[00:49:27\] It's standard.

**Andreas:** \[00:49:28\] But it's a standard, how many people use it? How widely adopted is it? Because ideally, everyone would be using it, then the problem would be solved.

**Lars:** \[00:49:35\] Yeah. There is quite a competition in the area of Software Bill of Material formats. There is one coming out of the world of security area. How is it called? I forgot.

**Andreas:** \[00:49:46\] We can jump over the name, but-

**Lars:** \[00:49:49\] Yeah.

**Andreas:** \[00:49:49\] Okay. So you have competing standards?

**Lars:** \[00:49:51\] Yeah, and this other one seems to be more successful currently, simply because license compliance nowadays is typically done by pushing a PDF to your customer. And so this ... this supply chain management of open source is not well established as of now. That's what open chain is all about, but there are things to go. And moving the data through the supply chain is ... for some reason hasn't been established.

And that's why the security world, the Bill of Materials simply is much more needed when it comes to your own management over the whole life cycle because of late findings of vulnerability, which are not known at development time. They had more to come up with a solution and that's why they have had their Bill of Material technology in place earlier, and which seems to be more currently successful.

### **Lars' ideal OSS management tool chain**

**Andreas:** \[00:50:50\] So what would your perfect OSS management tool chain look like if we're talking about running a scan, evaluating the risk, monitoring post-release and also providing the Bill of Materials to end customers? Could you put that together?

**Lars:** \[00:51:06\] Yeah. Clear. CycloneDX is the name of the-

**Andreas:** \[00:51:10\] CycloneDX, okay, yeah.

**Lars:** \[00:51:11\] Yeah. So it's basically what I presented on the last EclipseCon. Whoever's interested can watch my talk there.

**Andreas:** \[00:51:19\] Yep.

**Lars:** \[00:51:20\] It's open source compliance as a service. And this is basically centered around the OSS Review Toolkit, because this is really a good solution for this continuous integration stuff. So you can simply add this to your continuous integration, get a report at the end, get the documents in the end of the processing, and get a list of issues that you have to solve.

And by the way, I would to mention that there's also a policy check component in there. So you can even automate the policy checking like you mentioned, blacklists, licenses and stuff. So you can ensure that nothing is found which has a license that is on your blacklist and stuff like that.

**Andreas:** \[00:52:02\] Is ORT still essentially command line-based. Is it more-

**Lars:** \[00:52:05\] It's command line-based.

**Andreas:** \[00:52:06\] Yeah. So it's an automated ... You set up an automated system, and then you have to honor and configure how you want those alerts to be displayed. So it would tie in with your existing infrastructure, yes.

**Lars:** \[00:52:18\] So what we did is basically we ... So there is already a tool in a Docker container. So you can simply use ORT in the Docker container. And we added that too and built infrastructure. So we started with Jenkins. Lately the guys are using Azure DevOps, which provides simply a build and you can trigger this build out of your build with the links to the component to be checked.

And this is run in a Docker environment, completely automatic. And that's how you run it in DevOps environment. Yeah, so that's the core and what you get there is for whenever you run it, you get information. What is in your project? What is the license situation of those? And what is the policy state of those? And you get the documents so you can ship your product with the compliance information. And that's simply a report.

So you can generate a SPDX cycle on the XPDF HTML. So for apps, typically, you need HTML because you want to add it to the product about page and stuff like that. That's the backbone of the whole thing to run this constantly. The second important thing is the database, and we recommend Software 360 at that point, which stores the whole data. So which stores the component information, which stores the approved state so that you know which components you're using and whether you have checked in the metadata.

**Andreas:** \[00:53:48\] So ORT will ... It will scan for you, it will produce a report, it will handle policy violations. And one of those reports will be to push data into Software 360, which persists the data. Acts like a database, like a software catalog. Is that the right term?

**Lars:** \[00:54:03\] Yeah.

**Andreas:** \[00:54:04\] Yep. And then that supports ... The Software 360 then supports the workflow, the lifecycle of the issues found?

**Lars:** \[00:54:11\] Exactly. You store the project information that traces to the components there. Also, there's a license database so that you can store your interpretation of licenses, because that's also a legal deduction that you have to read the license and you have to conclude what are the obligations, and how do we want to fulfill them? And yeah, that's then the entry point for later phases.

When you detect a vulnerability, there's also connection to a CVE database in Software 360, but I'm not sure if this is still working, because I think this service was shut down that was used there. That's actually one of the big issues, getting the data appropriately into the system. And that's typically something where you would like to pay some money for a good data provider.

And you enter the data, you identify the components involved, and you can trace to the projects in which you have used those components. And then the process starts that the responsibilities for those projects have to decide on, but they are informed. So it's moving from a polling-based mechanism to an event-based mechanism. You really get an automation and the right people get informed that there's something to do.

**Andreas:** \[00:55:33\] Great. Imagine if we could get these commercial actors who are maintaining their databases of licenses and additional security data. And if they could all just pull their information together, and then if you wanted to use that, maybe you had to pay a subscription fee or something. It could be offered like that data as a service, could be offered to you.

And it would just link up with all these open source tools that do the actual scanning work.

**Lars:** \[00:56:00\] Yeah. So, from a vulnerability perspective, I would pay money. That's a good research office who checks for vulnerabilities and does this mapping to, for example, unique IDs, there's the package url mechanism for identifying components and things like that. And that's really worth something, and something you don't want to do every day in your company. That simply does not make sense.

Then it comes to compliance information. That's one of the big selling points for commercial tools, but I'm a bit more reluctant there, because what we experienced basically is that they have when you start up to get all the information. But when you run the system, you're faster than they built up the database and you need the information when you add the component, basically not when ... two weeks later.

The turnaround time with commercial databases typically is too long.

**Andreas:** \[00:56:54\] Right. Okay.

**Lars:** \[00:56:55\] And that's why I would prefer to use open source tooling and building up your own known database of components and component metadata, simply because when you ... There is this upfront invest to fill the database with the stuff you're using at a certain point in time when you start. But then when you run it, it's much more smoother than having a commercial vendor in the picture that has to update his database.

And then they don't scan this repository. And then you have to fill in the data from the side. So it becomes really ugly if you are not in control of the data there.

**Andreas:** \[00:57:34\] Right. So in short then, to get that license data once, you've got ORT in place which is managing your scans and pushing data to Software 360 where you store the information and you have your workflows. And in order to fill the gaps with that license data, to get the license data, you scan the OSS components that are found to identify the licenses. You have to work out license conflicts and harmonize the license statements.

**Lars:** \[00:57:58\] Yeah.

**Andreas:** \[00:57:59\] And then that basically builds your internal, or your own managed database of license data.

**Lars:** \[00:58:05\] Yeah. And together with ClearlyDefined. So until now, unfortunately, ClearlyDefined seems to be not really working out, because the willingness to share the information there seems to be not as high as it should be, or that it has to be. But this would be a cool solution, that companies share their information in this database and make use of this database.

But this is definitely the more likely approach than building on top of a commercial data provider, because yeah, as I said, this makes things ugly.

**Andreas:** \[00:58:45\] And do you see it on the horizon? Do you think it's ... Is there any work being done there to start there?

**Lars:** \[00:58:50\] Yeah, ClearlyDefined is there. The problem is that ... Yeah, so the Eclipse Foundation, for example, shared that ClearlyDefined Microsoft is working with that. But there are still things to do to make this working in this. But this is the solution, I think that could be successful. And I don't see the need for another methodology. It's more that either we get this running or we don't have this, because there is no reason to build up something besides that.

**Andreas:** \[00:59:20\] This has been excellent. I have enjoyed this very much. We have taken up much more of your time than we said we would, Lars. But time really flew. It was a lot of fun. Thanks for taking your time and being an excellent guest. Would you like to run down a list of projects you're involved with that you'd like to pitch to people or give them an idea where ... to see the things that are being done, to see the work that you're doing?

**Lars:** \[00:59:45\] I could, but to be honest, I would basically refer to my talk on the EclipseCon, which is publicly available because there I summarised basically the whole methodology.

**Andreas:** \[00:59:55\] Okay.

**Lars:** \[00:59:56\] So this would not add much value if I would, basically-

**Andreas:** \[01:00:00\] Right. Right.

**Lars:** \[01:00:02\] But I can send you the link to screen cast there.

**Andreas:** \[01:00:06\] Sure, sure. We'll include the links to the things that we've talked about, the projects that we've talked about mentioned in this podcast. We'll include those in the post of this podcast and the description. I can tell that you're a man of efficiency, Lars. Very good. If you're interested in finding out more about Lars and his work that he's working on, go look at the EclipseCon talk, which was Compliance as a Service, OCAS, O-C-A-S. Is that right?

**Lars:** \[01:00:30\] Yeah. I'll pitch you the link.

**Andreas:** \[01:00:32\] Yeah. So we'll put the link in the description as well. So yeah, I've been Andreas this whole time. Anoop, would you like to say good-bye?

**Anoop:** \[01:00:40\] Yes. Yeah, just very much excellent. Even from a developer perspective, it was really good. I'm going to, right away after this, going to go to the OSS Compliance as a Service doc and get to know more. It's really interesting now. Thanks. Thanks so much, Lars.

**Lars:** \[01:00:54\] Okay. Yeah. Again, thank you for inviting me. It was really a pleasure to talk about this topic. I still like it, although actually I basically moved out. I'm now doing other open source stuff here. But this is simply a topic that I was very involved in and I like to share.

**Andreas:** Yep. Great. Good. Thanks a lot, and bye-Bye.

**Lars:** Bye-Bye.

**Anoop:** \[01:01:18\] Bye-bye.
