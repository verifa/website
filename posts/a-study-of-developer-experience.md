---
type: Blog
title: A study of Developer Experience
subheading: Developer Experience promises improved productivity, satisfaction, engagement and retention of development teams. We conducted a study of existing material and in this post we share the summary together with our real-world experience.
authors:
- jlarfors
- tlacour
tags:
- Developer Experience
- DevOps
- Platform Engineering
date: 2024-08-21
image: "/static/blog/a-study-of-developer-experience/devex-a-study.png"
featured: true
toc: true
---

In recent years Developer Experience has gained significant interest for teams and organisations wanting to improve their software delivery. To better understand Developer Experience we conducted a short study of existing material and summarised it together with our real-world experience.

The goal of this post is to give the reader an understanding of Developer Experience, why it is important and how you can improve it.

## What is developer experience?

In 2012, Fagerholm and Münch proposed a definition of Developer Experience that is influenced by User Experience (UX) and is comprised of three dimensions: **Cognition**, **Affect** and **Conation** \cite{fagerholm2012developer}.

![devex-concept-and-definition-2012](/static/blog/a-study-of-developer-experience/devex-concenpt-def-2012.png)

The **cognitive** dimension consists of factors that affect how the developers perceive their development infrastructure, such as interactions with tools and processes. The **affective** dimension is about developers’ feelings towards work, and is largely related to social factors (respect, belonging, attachments to individuals and teams). The **conative** dimension is about a developer’s perceived value of their contribution. Visibility into how work transforms into value, and spending time on meaningful tasks positively impacts developer experience. In their definition, “developer” refers to anyone who is engaged in the activity of developing software and therefore does not limit this to only people with “developer” as their title or role.

This definition has been a base for more recent work that summarises it succinctly as “How developers think about, feel about, and value their work” \cite{greiler2022actionable}. A practical model for understanding Developer Experience was later proposed, also with three dimensions: feedback loops, cognitive load and flow state \cite{noda2023devex}. The dimensions in this model closely relate to the dimensions in the earlier definition \cite{fagerholm2012developer}: cognitive load is a measure of cognition, flow state helps individuals focus and feel productive (affect), and feedback loops help individuals value their work (conation).

The definition of “developer” is something that still lacks clarity. We would argue that it is not just anyone involved in *developing* software, but anyone involved in *developing and delivering* software.

## What’s the background/evolution of Developer Experience?

The first formal definition of Developer Experience we can find is from 2012 \cite{fagerholm2012developer}, although there are online publications from 2011 \cite{effectivedx2011}. While we did not have a formal definition of Developer Experience before then, improvements to the software development process have positively impacted the Developer Experience as defined today.

For example, those following the Agile Manifesto Principles \cite{agileprinciples} published in 2001 will see an improved Developer Experience. It touches **cognition** (enabling and support individuals), **affect** (promoting sustainable development, trust for development teams and collaboration with business people) and **conation** (promoting continuous delivery and close interaction with customers). And of course, it promotes continuous improvement which, if done correctly, will help with all three dimensions.

DevOps as a methodology (and not as a team or a set of tools/technologies) will also improve the Developer Experience.  One could argue that DevOps has added to the cognitive load placed on developers, though let us consider that a consequence of poor implementation. DevOps promotes a collaborative environment, shared responsibility and improved feedback; all positively impacting the Developer Experience.

The goal of Platform Engineering is to abstract away the low level details of a “software development platform” and provide a bespoke interface (also known as an Internal Developer Platform) to support teams with the development, deployment and/or operations of their software. For a platform team, User Experience is a large part of the overall Developer Experience (depending on how far a platform spreads across the software development lifecycle). The notion of treating a “platform as a product” is largely about collaborating with your users (the developers) and iteratively delivering something that improves the User (or Developer) Experience.

So what is significant about Developer Experience if we have had methodologies and practices to improve the software development process all along? The purpose of Developer Experience is to help practitioners better understand, analyse, design and improve software development environments from the perspective of a user (or developer).

## Why is Developer Experience important?

A better Developer Experience is proven to improve individuals satisfaction, engagement and retention \cite{noda2023devex} as well increasing performance outcomes for individuals, teams and organisations \cite{devexinaction}. A happy and well-performing team will naturally lead to better recruitment because people will talk positively about your workplace and even recommend it to their connections. Thus, any company that delivers software should optimise their work environment from the perspective of those expected to participate in it. The emphasis here should be on the experience of the users of the software delivery process; an overemphasis on productivity is the best way to lose it \cite{endres2003handbook}.

## What makes a good Developer Experience?

What defines a good Developer Experience is difficult to capture. However, the presence of a bad Developer Experience is more easily identified. For example, interruptions, unrealistic deadlines, and friction in development tools negatively affect Developer Experience \cite{noda2023devex}.

A common misconception is that Developer Experience is primarily affected by tools \cite{noda2023devex}, but it encompasses less tangible factors such as feelings about work (e.g. respect, attachment, belonging) and the value of one’s own contribution (e.g. alignment of one’s own goals with those of the project, plans, intentions, commitment and feedback) \cite{fagerholm2012developer}.

It stands to reason that a good developer experience comes from a process uniquely well-fit for its organisation. It cannot be bought or borrowed but be developed internally. As such, “good” Developer Experience should not be a goal but an undertaking that requires Continuous Improvement.

## How do you measure Developer Experience?

Measuring Developer Experience holistically is difficult as it requires collecting quantitive and qualitative data from a variety of sources. The SPACE framework \cite{forsgren2021space} is arguably the closest practical research-based framework for measuring Developer Experience. Though the SPACE framework was created to measure developer *productivity*, multiple studies observed correlation between developer productivity and satisfaction \cite{graziotin2017unhappiness}\cite{storey2019towards}\cite{bellet2024does} and the “S” in SPACE stands for “satisfaction and well-being”.

Other metrics, like those defined by DORA \cite{forsgren2018accelerate}, focus on performance which is an outcome of Developer Experience and can therefore be complementary, but not all encompassing as a Developer Experience measure.

User Experience metrics can also be used to uncover usability issues in a development environment, such as an “Internal Development Platform” or self-service tools. These metrics are not related to the software development process directly, but related to the environment in which the developers work. For example, knowing that users spend an extra long time on a specific page, or that a self-service capability often fails, can help to improve the Developer Experience, the result of which should be reflected in productivity metrics.

The first practical step you can take when measuring Developer Experience is to study the SPACE framework \cite{forsgren2021space}. It does not tell you explicitly which metrics you need, but provides a framework for helping you select and fine tune your metrics. It is not holistic, but it is a very good start. At Verifa, we like to start our investigations with [Value Stream Mapping](https://verifa.io/services/assessments/value-streams/). It is an excellent way to understand the development flow, dependencies across teams and get a general feel for the Developer Experience within an organisation.

## How do you improve Developer Experience?

Like any improvement process it needs to be done iteratively with measures in place to validate the effect of changes. In order to make a long-lasting change you will need support from those paying the bills and those who will be affected. Hence, the first step will be to conduct preliminary research about the current state of Developer Experience to identify pain points (opportunities) and show just cause for implementing improvements. We recommend mapping out your Value Streams and (manually) collecting some metrics (e.g. from the SPACE framework \cite{forsgren2021space}). Do not spend time implementing automated data collection tools and/or dashboards at the beginning.

Once the initial data has been collected and you have a rough understanding of your Developer Experience it is time to identify improvements you want to make. Consider your business (goals and challenges) and optionally set goals for your improvement. Remember that your proposed changes should be implemented gradually, letting the dust settle a bit in order to re-evaluate your course using updated metrics.

## Conclusion

The parallel between User Experience (UX) and Developer Experience is interesting; observing your software development lifecycle from a UX perspective rather than measuring outcomes will help with developer satisfaction and retention. This is also particularly interesting for DevOps, Platform and/or Enablement teams (as per Team Topologies \cite{skelton2019team}) that support development teams because it emphasises the very important “human” aspect that is often forgotten or overlooked.

While few practical frameworks or methods exist today for measuring or improving Developer Experience we hope more practical work and shared experiences will  improve the Developer Experience to everyone’s benefit.

If you are looking for an experienced team to help guide you with Developer Experience, do get in touch!
