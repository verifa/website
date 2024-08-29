---
type: Blog
title: Unlock your Continuous Delivery potential with Value Stream Mapping
subheading: In this blog I cover challenges teams face while improving their release process, and how Value Stream Mapping can be used to overcome them.
authors:
- jlarfors
tags:
- Developer Experience
- Continuous Delivery
- Value Streams
date: 2022-09-07
image: "/static/blog/unlock-continuous-delivery-potential-vsm/unlock-your-continuous-delivery-vsm.png"
featured: true
lastMod: 2024-08-29
toc: true
---

**Every software team is slightly different, as are their problems, which makes practicing Continuous Delivery challenging. Value Stream Mapping is a method for helping teams make more informed decisions about how to improve their Continuous Delivery practices.**

Plenty of material exists for teams to read and learn about improving their Continuous Delivery practices. However, best practices are not a one-size-fits-all. Do you know what should change and where to invest? Value Stream Mapping is one method you should consider to help answer that question.

## What is Continuous Delivery and why is it hard?

Continuous Delivery is about building a process where software is always releasable while ensuring short release cycles. The pains and struggles to achieve this are often not related to technology but stem from a fragmented release process with handovers causing friction and delays.

Practicing Continuous Delivery is challenging because every team is different and best practices are not a one-size-fits all. Let’s cover some of the common challenges teams encounter when looking to improve their release process.

### Challenge: Understanding your current release process

Before you can improve your release process you must first understand its current state in order to identify what to improve. Even if you have the fortune of being able to define a completely new process (e.g. for a new project), it could be costly and challenging to implement, and it would be a shame if you did not learn from previous mistakes.

It is difficult to build a realistic and common understanding of a release process; one that truly represents the daily tasks, does not simply focus on the “happy path”, and one that highlights where the pain and struggles are. If only a few individuals sketch their own understanding of the release process it is likely going to miss important details and your analysis will thus be inaccurate.

### Challenge: Lack of involvement from the team

Every software team is different because of the people in it and not just the technologies they are working with. It is imperative to not lose sight of the human factor when implementing change and make sure you get people along for the ride.

Enforcing new practices without being able to properly explain ***why*** changes are being made can cause friction and internal opposition against the change. If it’s been working fine up until now, then what’s the problem? People need motivation for change.

Team members also have varying comfort zones and it would be foolish to assume everyone to be okay with change. Having people involved from the beginning will make it much easier to understand and support people in expanding their comfort zones.

Ownership is pivotal when making change. If those responsible for an activity simply get told how to do it “better” without the opportunity to provide input on the change, it can lead to repeating previous mistakes and lack of motivation.

### Challenge: Lack of feedback

Continuous Delivery is all about (continuous) feedback to the development team, yet ironically, it is common to struggle getting feedback on the improvement process itself. This is not just about metrics, such as lead time or mean time to recovery, but also about feedback from the people involved in the improvements.

Applying the Build-Measure-Learn process to Continuous Delivery will help in making a successful transformation and ensuring the Developer Experience is at the forefront of change and not an afterthought.

Implementing change to your release process should not be done in big chunks because big chunks are more difficult to adopt, measure and revert. At the same time, it is difficult to decompose problems and create an incremental plan for improving your Continuous Delivery.

### Challenge: Lack of support from management

Getting support from management to take time away from product development is a common challenge with Continuous Delivery. This is not because management does not want to invest into improvement, but the challenge is being able to justify the return on investment.

As mentioned earlier, it is important to involve the team. You cannot simply have someone improve your release process on their own. There will be a cost of practicing Continuous Delivery and the benefits need to be quite clear.

### Challenge: Following best practices

It is common to hear success stories with “tool X” that helped a team greatly improve their release process. While research is important, it is easy to mistakenly focus on the workflows and tools of successful teams, rather than ***why*** it works for them.

One of my wise colleagues told me once that “Best practices are a good solution to someone else’s problems”, and I have not found a better way to summarise it. Of course best practices should generally be followed, but not without understanding ***why*** you are doing something and the problem you are trying to solve.

Goals should always be tied to the business. For example, aiming for the shortest release cycle possible comes at a price (often a compromise) and if it does not add clear value to the business, you are likely sacrificing something else that might add value (e.g. stability, scalability, monitoring). There’s only so much you can do at a time, so choosing your battles should be done with some insight into your specific team, product and business requirements.

## Value Stream Mapping

If the above challenges with implementing Continuous Delivery have resonated with you then you will want to learn about Value Stream Mapping and how we at Verifa run our workshops.

### What is Value Stream Mapping?

A Value Stream Map is a visual tool showing the series of events that take a product or service from the beginning of a process until it reaches its end (i.e. a Value Stream). It displays all the critical steps and quantifies the time and volume taken at each event. The main purpose is to be able to identify and remove waste.

Value Stream Mapping (VSM) is a method for defining a Value Stream and analysing the current flow in order to remove waste and to derive a more lean and continuous workflow. When applied to Continuous Delivery, it documents the current delivery process and highlights waste as well as opportunities for improvement.

<figure>
  <img src="/static/blog/unlock-continuous-delivery-potential-vsm/mock-vsm.png" alt="simple-mock-value-stream-map">
  <figcaption>Example of a simplified Value Stream Map.</figcaption>
</figure>

Let’s go over the aforementioned challenges again and discuss how Value Stream Mapping can help address them.

### Solution: Understanding your current release process

A Value Stream Map can accurately describe the current release process and the waste within. Waste makes it clear where to improve while justifying ***why*** changes should be made. If we know the problem, it makes it easier for us to propose solutions and be able to measure the success of the solutions.

Visualising your release process brings many benefits, but it’s not just about the visual diagram you create. It’s about the process of creating it during which you share and discuss things that might never be discussed otherwise. This is the main reason why Value Stream Mapping should be a team activity; not something only a few individuals are involved in.

### Solution: Lack of involvement from the team

Value Stream Mapping is an interactive process and one objective should be to engage everyone involved to contribute. This will get people onboard right from the beginning.

The waste discovered while Value Stream Mapping will often originate from pain described in the release process. That relationship creates a clear mapping between any proposed change (i.e. solution) to the problem (i.e. waste) that originated from some pain. It is a great motivator to understand ***why*** changes are being made and the pain it will alleviate.

While Value Stream Mapping does not directly extend people’s comfort zones, it will create a backlog of improvements that is easy to communicate and discuss. The benefit is that people can prepare themselves for what’s to come and openly express concerns during the process or get some training on the side before being thrown in the deep end. It is a more comfortable approach.

The interactive nature of Value Stream Mapping is great for getting those responsible for activities involved in discussing waste and possible solutions right at the beginning. This allows people to take ownership of their work and be involved in shaping the future process.

### Solution: Lack of feedback

Value Stream Maps do not provide feedback about changes you implement directly but do provide guidance for what to measure. Asking the right question can be very difficult and Value Stream Maps make asking questions such as “did change X help solve waste Y?” a lot easier to write, communicate and measure.

Remember that metrics are only part of the feedback. The Developer Experience is not easy to measure with metrics (although they can be correlated, e.g. better performance typically leads to happier teams and vice versa). By referring back to the Value Stream Map and the waste that was identified, it makes it easier to discuss feedback from the team.

It is important to decompose problems into small incremental changes. This is much easier to do when the waste we have identified is related to steps in the release process. For example, sometimes waste is not a problem but a symptom of other waste. We can consider this waste to be transitive waste. This is valuable information when making changes because it can help us to order and plan changes incrementally.

### Solution: Lack of support from management

Waste identified in the Value Stream Map can represent the cost of inefficiencies in your current release process. Quite simply; if you remove this waste then that is your return on investment. This makes it much easier to justify investment into improvements.

It is recommended to include management in the Value Stream Mapping process; it’s a great way for them to learn more about how the teams work together. Once you start to tackle the waste it’s also much easier to plan the work into the ongoing project schedules. This is great because it means you can be quite flexible with how much time you invest into improvements based on how much you think it will benefit you.

### Solution: Following best practices

If you use Value Stream Mapping to help guide your Continuous Delivery journey, you will better understand what your team’s individual challenges are and better reason about possible solutions. Once you know ***why*** you are doing something, it’s much easier to start looking for best practices that try to solve the same thing. Love the problem, not the solution. There is no endgame with Continuous Delivery and time is finite.

> Love the problem, not the solution. There is no endgame with Continuous Delivery and time is finite.
>
> <footer>Jacob Lärfors, Verifa</footer>

## Conclusion

Continuous Delivery is hard because teams are different, their problems are different, and therefore the same set of solutions might not work for everyone. Learning about your own release process is not an obvious thing to do. Books and the internet can highlight problems others have suffered from (which might well be your problems) and suggest methods that you can try out. But it is important that you can relate problems back to your own release process; not just second-guess your problems and apply solutions.

The observant reader will have noticed that the word “why” has been highlighted throughout this post. There was intent behind it; I believe practices like Continuous Delivery should very much lead with ***why*** we do things, followed by ***how*** things will be done. The ***why*** is the problem statement and the ***how*** is an implementation detail. A better understanding of the problem will greatly help with the implementation.

Value Stream Mapping is one method for learning about your release process and identifying the problems with it (i.e. waste) and hopefully this post has given you some insight into how it can help. Do you want to run your own Value Stream Mapping workshop or have any questions about it? You can read about our [Value Stream Assessment](/services/assessments/value-streams/).
