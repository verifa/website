---
type: Blog
title: 'GitOps: All you need is Git (and Continuous Deployment)'
subheading: GitOps is a modern approach to automate deployments using Git as the single
  source of truth to define the deployment.
authors:
- avijayan
tags:
- Git
- GitOps
- Continuous Delivery
date: 2021-06-13
image: "/blogs/2021-06-14/all-you-need-is-git-image-01.png"
featured: true
jobActive: true

---
**Automating deployments (A.K.A. Continuous Deployment) is not just running your _“deploy”_ command in a pipeline, but it requires a lot more. How do we manage our releases? What if something goes wrong? There is much more to automated deployments... enter GitOps!**

Kelsey Hightower, the father of Kubernetes put it nicely: “GitOps is versioned CI/CD on top of declarative infrastructure” \[1\]. Using version control for software development is nothing new, and neither is version controlled CI/CD pipelines nowadays. So what's all the fuss about GitOps? Essentially, it's about bringing all of these things together into a streamlined process that's going to enable teams to perform better, whether the goal is higher release velocity or more controlled and confident releases and upgrades. Let's briefly explore GitOps to guide our understanding of what it entails in order to help us improve our GitOps practices.

### **What does GitOps really mean?**

There is some confusion about GitOps and what it means. When we talk about utilising GitOps, what we are really saying is:

* _We define our infrastructure and configuration as code_
* _We define a pathway to deploy that infrastructure and configuration (e.g. pipelines)_
* _We manage releases through Git mechanics (tags, pull-requests etc.)_
* _When a new release occurs, an automatic deployment occurs_
* _What you see in Git is what you see in the target environment (reconciliation)_

GitOps is a way of taking infrastructure / configuration / everything-as-code, stored in Git, into a release through Git mechanics and having that reflected in target environments. For example, an app deployment or cloud infrastructure. It can also be said to be a single source-of-truth for what goes into a system and the pathway to deploy it.

In reality, GitOps empowers developers to take on operations in a way that says “You own it, you ship it!" and graduates from being a slogan to a reality that can be executed day in and day out. \[2\]

### **GitOps - a modern approach to manage automated deployments**

GitOps is a modern approach to manage automated deployments which has many benefits. Before practices like Infra as Code and Continuous Deployment became mainstream, there was little interest to save or version our infrastructure (as code). But with Infrastructure as Code, just like we version Software, we also save our Infrastructure (as code) in version control. Added to that, we also manage the release of our Infrastructure versions in the same way we manage the release of our Software versions. Typically this means employing Git workflows like branching, pull-requests and tagging code commits.

Sometimes there is confusion between Continuous Deployment (CD) and GitOps. GitOps is a subset of Continuous Deployment (CD) and defines **what** we deploy and **how** we deploy a particular release.

### **Achieving GitOps**

The two main ingredients to achieve GitOps are using a Git workflow to manage the release of our versions and a continuous deployment process to automatically deploy those releases.

#### Git Workflow

A Git Workflow is a recipe or recommendation for how to use Git to accomplish work in a consistent and productive manner. Git becomes the interface for operations, which means any change to a system is introduced only by a git commit. Some of the Git workflow types include Centralized, Trunkbased \[3\], Feature branching, GitFlow, Forking etc. \[4\]. The user is free to choose the type of workflow but software teams should agree on which type to follow. Git workflows also define the process to mark releases. For example, one could publish a release based on a chosen git commit or git tag from a git repository.

#### Continuous deployment

As soon as a version is released, the Continuous Deployment (or CD) system ensures that deployment occurs in the correct environment. Tests and other quality control measures are typically part of the CD process. Apart from this, a GitOps CD system should be convergent in nature. This means the system is always in synchronous state with Git with convergent loops, for achieving continuous software delivery and deployment. This can be achieved automatically using strategies like:

1. Retry - Keep trying to achieve the desired state from an actual state.
2. Modify - Update every change performed in the system back to Git

A simple GitOps system is based on having these two key concepts.

### **How to deploy GitOps**

There are two ways of deploying in GitOps. They are:

#### a) Push based - Continuous Deployment

![Push based GitOps](/blogs/2021-06-14/gitops-blog-diagram-1-white-04.png)

In a push based approach, the changes are pushed to the target environment as a part of continuous deployment. In a Kubernetes world, a tool performs _kubectl apply_, or _helm upgrade_ at a Deploy stage in the Continuous Deployment pipeline. One should remember that it is still controlled by changes in version control, but just an extension of Continuous integration. Needless to say, this is the traditional approach to Continuous Deployment. E.g. Pipelines implemented with Jenkins, GitLab, GitHubActions etc.

#### b) Pull based - Continuous deployment

![Pull based GitOps](/blogs/2021-06-14/gitops-blog-diagram-2-white-05.png)

In a pull based approach, there is an agent, proxy or operator running inside your target environment which pulls and deploys changes to itself, from a version control system. It's still from version control, but not an extension of your Continuous Integration (CI) pipeline. Also it eliminates the need to have a broker in between.

One of the clear benefits of pull approach is, one does not have to manage secrets and access control for deploying changes to the target environment. In the case of push-based approach, there is a clear need for credential handling and access control because the changes are being pushed from outside. Some examples of frameworks using pull based approach are Git operators like FluxCD, ArgoCD.

GitOps itself does not favour one approach over the other, but in the end, how changes are deployed are a matter of choice and feasibility.

### **10 key benefits of GitOps**

 1. **Automated deployments**

    Though automated deployment is a key component of GitOps, this can be considered as a benefit for organizations as well. By making changes in the version control, the update makes it all the way into the production system. This reduces lead time and increases release frequency.
 2. **Auditability**

    Due to the nature of GitOps workflow starting from version control, one gets an built-in audit log of the changes directly from the version control system. Therefore, there is no need to maintain a separate system for the audit trail of the deployments.
 3. **Traceability & rollback**

    This is closely related to Auditability. In a working GitOps system, upon a failure, one could traceback the changes which caused the failure and decide to rollback or rollover to resolve the problem caused due to failure. Traceability is key to rolling back for problem resolution.
 4. **Automated rollbacks**

    There are GitOps systems which perform automated rollback upon failure with auditability preserved. This means that a failure was detected upon a new change and the system rolled back to the previous version, all done automatically.
 5. **Improved collaboration & development**

    With the GitOps approach, the development, test and deployment teams are not working in silos anymore, but designed for optimum collaboration. All changes to the system are transparent across teams. This also promotes contribution across functional teams.
 6. **Stage and Verify**

    GitOps due to its cohesive nature, it's simpler to stage a change or update to a system and verify correctness. After which, the changes can be released all the way to live production environments. This is also called phased deployments.
 7. **Quality control**

    This is slightly related to Staging and verifying changes. It's easier to perform a Quality Control (QC) ranging from light-weight to every-extensive in this system due to the possibility of having phased deployments. This also includes running security vulnerabilities check.
 8. **Automated tests**

    Automated deployments force automated tests. It's almost impossible for a software using GitOps practices to scale without automated testing. This can be considered as a positive side-effect of GitOps approach.
 9. **Disaster recovery**

    Thanks to the infrastructure as code (IaaC) feature of GitOps, Disaster recovery is more easily managed. Now it's possible to restore infrastructure partly or fully in times of disaster.
10. **Drift detection and remediation**

    In a GitOps environment, due to reconciling happening all the time, there is less possibility of configuration drift. If it happens due to some manual intervention, remediation kicks-in and gets the system back to the desired state from git.

### **Drawbacks to GitOps**

With all the benefits that GitOps brings, there could also be some drawbacks to GitOps.

#### Cost of setup

GitOps is not recommended for cases when the Return On Investment (ROI) is low and cases whenever overhead outweighs the benefits. This can be in scenarios like simple applications, very small teams, perhaps in the early days of a project etc.

#### On-Boarding

As simple as it sounds, implementing GitOps for the first time in an organisation can be a cumbersome process. Challenges can be both from a technical and non-technical front.

### **Is GitOps the right approach for me?**

Every approach has pros and cons. For GitOps, we can clearly see there are more benefits than drawbacks. This does not mean that just by using GitOps you will be successful. Moving into GitOps also needs proper planning and most importantly the right mind-set. But one thing that we know for sure is that GitOps can be a path to success and a key ingredient in successful products and companies around the world.

### **References**

1. GitOps definition: [https://www.gitops.tech/](https://www.gitops.tech/ "https://www.gitops.tech/")
2. What is GitOps: [https://www.cloudbees.com/gitops/what-is-gitops](https://www.cloudbees.com/gitops/what-is-gitops "https://www.cloudbees.com/gitops/what-is-gitops")
3. Trunk based development: [https://trunkbaseddevelopment.com/](https://trunkbaseddevelopment.com/ "https://trunkbaseddevelopment.com/")
4. Git workflows: [https://www.atlassian.com/git/tutorials/comparing-workflows](https://www.atlassian.com/git/tutorials/comparing-workflows "https://www.atlassian.com/git/tutorials/comparing-workflows")
