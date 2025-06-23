---
type: Blog
title: "Getting started with Jenkins Pipelines"
subheading: "How to develop an effective CI process"
authors:
- jlarfors
featured: true
tags:
  - Jenkins
  - Continuous Integration
date: 2020-09-21
image: /blog/2020-09-21/main.jpg
---

**New to Jenkins pipelines? It can be difficult to figure out what a minimal working setup is, and whether to use normal or multibranch pipelines. Let's take a look at what you need to do to get up and running.**

If you are already using Jenkins as your CI server you will probably have heard about Jenkins pipelines. And those of you who have not already migrated from Freestyle jobs might want to start planning the migration!

Jenkins pipelines provide a very rich way of developing an automated process for releasing software. Pipelines are based on groovy and provide a high-level Domain Specific Language (DSL) that enable you to do lots of things with very few characters. I'll try to give you an overview of all the key elements so that you can start developing in line with best practices while avoiding the pitfalls.

A "pipeline" in modern continuous integration tools is a process that typically allows you to break your automated workflow into "stages".
Typically, a pipeline is triggered by changes to version control, such as a `push` in git, and proceeds to run stages like building the software, running a linter, static analysis, and a series of tests.
With continuous deployment, such a pipeline might be extended to also `deploy`
the software automatically to some sort of environment.

With Jenkins, a pipeline follows this concept and is implemented using a suite of
plugins in the Jenkins ecosystem.

Let's get into what you need to know to start developing Jenkins pipelines. And for those of you already using Jenkins, I'll touch upon some neat tips and tricks to help you improve your CI process.

## 1. What is a Jenkins pipeline?

### 1.1 Jenkins pipelines and Jenkinsfiles

Let's take a basic example of a pipeline to start with:

```groovy
pipeline {
    agent { docker { image "golang:1.15.0" } }
    stages {
        stage("build") {
            steps {
                sh "go version"
            }
        }
    }
}
```

This pipeline will use the docker agent, which means the `golang:1.15.0` docker
image will be started and the pipeline will run inside a running container based
on that image.
The docker container will run on a Jenkins "agent" which is just a node (or VM)
which is configured as a Jenkins agent.
The pipeline has one stage, `build`, which will execute `go version` in a shell.
Easy, right?

When you create a basic pipeline job in Jenkins you will need to provide the
pipeline code itself, and you have two options for doing this:

1. Provide the text directly inside the Jenkins job
2. Provide a remote repository (e.g. git repository) containing the Jenkins
   pipeline, typically in a filled named `Jenkinsfile`

Option 1 is mostly for testing and should not be used for production. Jenkins pipelines should always be version controlled and stored in a separate repository in a `Jenkinsfile` (although you can give it any name you like). For instance, in some projects we have created "utility" pipelines
and put them in one Git repository with names that make sense, e.g. `artifactory-daily-cleanup` which will perform the daily cleanup of artifacts in Artifactory.

### 1.2 Pipeline vs Multibranch Pipeline

A basic pipeline job in Jenkins can fetch a `Jenkinsfile` from a **specific
branch** in a Git repository.
However, what if we want the `Jenkinsfile` (pipeline) to run on every branch in
a Git repository?
Do we create a pipeline job for each branch? What about new branches?

This situation gave rise to "Multibranch Pipelines" in Jenkins. Multibranch
pipelines are just another Jenkins job type that scans Git repositories for
branches and creates children pipeline jobs dynamically for each branch (based
on some specified conditions).
By using multibranch pipelines we can now store our `Jenkinsfile` in the same
repository as our code base and have it run over all the branches we want!

### 1.3 Pipeline Syntax

There is a very useful page that I refer to frequently:
<https://www.jenkins.io/doc/book/pipeline/syntax/>

It is worth noting straight away that there are two flavours of Jenkins
pipelines (this is **VERY IMPORTANT** to remember):

1. [Declarative
   pipelines](https://www.jenkins.io/doc/book/pipeline/syntax/#declarative-pipeline)
2. [Scripted
   pipelines](https://www.jenkins.io/doc/book/pipeline/syntax/#scripted-pipeline)

In nearly all cases you should use `declarative pipelines` because they are much easier to read and maintain, and provide a stricter syntax which reduces the chance for mistakes to occur.
However, if you have a need for heavily dynamic pipelines, where stages get generated on-the-fly with nested levels of parallelism, then it is often cleaner to write pure `scripted pipelines`, rather than trying to force your dynamic logic within declarative pipelines.
You can embed scripted syntax inside declarative pipelines using a `script {...}` block.

## 2. Developing Jenkins pipelines

Let's use the official Jenkins pipeline syntax [documentation](https://www.jenkins.io/doc/book/pipeline/syntax/) as a reference for this section and pick out some of the most useful things you need to know.
We will only cover `declarative syntax` of pipelines.

### 2.1 Agent

The [agent](https://www.jenkins.io/doc/book/pipeline/syntax/#agent) tells Jenkins where to run the pipeline, and therefore usually comes first in the pipeline because it is pretty important.
The `agent` can be specified either at the top-level of a pipeline, or for each `stage` block in the pipeline.

```groovy
// top-level
pipeline {
    agent {
        ...
    }
}

// or stage-specific
pipeline {
    // no global agent
    agent none
    stages {
        stage("Build") {
            agent {
                ...
            }
        }
    }
}
```

There are different types of agents that you can define and the approach to use depends on how you have configured your Jenkins agents.
If you configure static nodes (e.g. VMs) in your Jenkins master you need to tell the pipeline which of these nodes you would like to run on.
Typically, this is done by labelling your nodes and then using `label` in your pipeline and Jenkins will match the label in your pipeline to the labels on your nodes.

Another approach, the more "cloud-native" approach to Jenkins agents, is to configure a `cloud`.
A cloud can be configured for something like Kubernetes so that Jenkins will dynamically provision an agent in a Kubernetes cluster.
What this means in the case of Kubernetes is that a pod specification is needed and the Jenkins pipeline will run inside this pod.
Other examples of supported clouds: Docker, Docker Swarm, Mesos, Amazon EC2.

### 2.2 Stages, Parallel and Matrix

[Stages](https://www.jenkins.io/doc/book/pipeline/syntax/#stages) are quite intuitive, so the only point worth mentioning here is how you can nest stages and also run them in [parallel](https://www.jenkins.io/doc/book/pipeline/syntax/#parallel) blocks.
Let's take an example:

```groovy
pipeline {
    agent any
    stages {
        stage("Parallel Hello") {
            // run the stages in the parallel block... in parallel, natürlich
            parallel {
                stage("Hello 1") {
                    steps {
                        echo "Hello 1"
                    }
                }
                stage("Hello 2") {
                    steps {
                        echo "Hello 2"
                    }
                }
            }
        }
    }
}
```

A recent addition to this group is the use of a [matrix](https://www.jenkins.io/doc/book/pipeline/syntax/#declarative-matrix).
Matrices allow you to provide axes with combinations of configurations that a pipeline block will run with.

Here is a simple example of a matrix pipeline for different platforms:

```groovy
pipeline {
    agent any
    stages {
        stage("Parallel Hello") {
            matrix {
                axes {
                    axis {
                        name "PLATFORM"
                        values "linux", "mac", "windows"
                    }
                }
                // all the stages here will run in parallel
                stages {
                    stage("Hello Axes") {
                        steps {
                            echo "Hello on ${PLATFORM}"
                        }
                    }
                }
            }
        }
    }
}
```

### 2.3 Other directives (environment, options)

There are some other directives that can be useful to use in your pipelines.
A significant one I'd like to mention is [environment](https://www.jenkins.io/doc/book/pipeline/syntax/#environment) which lets you define environment variables that will be available to your steps.
Like `agent`, `environment` can be declared on the top-level of the pipeline declaration or be stage-specific.

Let's take an example of building a golang application where we specify the operating system as a top level environment variable. We'll then use a matrix to define some different architectures and set the GOARCH environment variable at the stage level with each of the different GOARCH axis values.

```groovy
pipeline {
    agent { docker { image "golang:1.15.0" } }
    environment {
        GOOS = "linux"
    }
    stages {
        stage("Build") {
            matrix {
                axes {
                    axis {
                        name "GOARCH"
                        values "amd64", "arm64", "ppc64"
                    }
                }
                // all the stages here will run in parallel
                stages {
                    stage("Build with GOARCH") {
                        // set the GOARCH environment variable based on the GOARCH
                        // axis values
                        environment {
                            GOARCH = "${GOARCH}"
                        }
                        steps {
                            echo "GOARCH = ${env.GOARCH}"
                            sh "go version"
                        }
                    }
                }
            }
        }
    }
}
```

[Options](https://www.jenkins.io/doc/book/pipeline/syntax/#options) is another very useful directive that allows you to configure some pipeline options.
Probably the most used one is `timeout` which sets a timeout for the pipeline.
It should almost be a necessity to add a `timeout` to a pipeline of reasonable complexity, especially in a large environment where you want to make sure jobs do not hang and occupy resources for other builds.
Other useful options are `skipStagesAfterUnstable` and things like `quietPeriod` to tell Jenkins how long a job should sit in a queue or wait for the SCM checkout operation to complete before killing the job.
Check the official documentation for the complete list.

### 2.4 Steps (actually doing stuff)

So, maybe the moment you've all been waiting for, the [steps](https://www.jenkins.io/doc/book/pipeline/syntax/#steps) section.
Within the `steps` block you define all the steps you want to execute, so this is where you put the things that you actually want the pipeline to do.
Steps are required because without them your pipeline is close to useless :)

#### 2.4.1 Different steps

In the examples above we looked at steps such as `echo` to echo some string to the executing pipeline and `sh` to execute a shell command.
Pipelines come with these generic steps and also others such as `sleep`, `build`, `error`, `readFile`, etc.

Jenkins plugins can also provide their own steps.
If, for example, there is a Jenkins plugin from a tool vendor they might very likely define a step that will serve as a convenience wrapper for some logic.
E.g. `runMyToolAnalysis` might be a step for MyTool to run the analysis, and it will take some arguments and underneath it might just invoke some things from the command line.
So, before going off and writing everything as a shell script, consider checking the available steps.

#### 2.4.2 Where things are executed (master vs agent)

The important thing to note here is that when a pipeline is executed the logic is
being executed on the Jenkins master, and only when you invoke some DSL steps,
like `sh` or `bat` does the work get executed on the agent. As such, if you try
to run code like `new File("my_file.txt").text` to read a file it will execute
on the master.

So, just make sure you use the built-in DSL steps available and most likely
you will want to install the [Pipeline Utility
Steps](https://www.jenkins.io/doc/pipeline/steps/pipeline-utility-steps/) plugin
for extra functionality like reading/writing files.

### 2.5 Credentials

There is not much need to duplicate documentation, so check the official
documentation on [using
credentials](https://www.jenkins.io/doc/book/using/using-credentials/).

### 2.6 Script block

The [script block](https://www.jenkins.io/doc/book/pipeline/syntax/#script)
provides a way to include bits of groovy scripts inside a `declarative`
pipeline. The official documentation covers most of what you need to know about
this.

### 2.7 Snippet Generator

Sometimes getting the syntax for steps in a pipeline can be challenging. To aid
this process, I strongly suggest using the [Snippet
Generator](https://www.jenkins.io/doc/book/pipeline/getting-started/#snippet-generator)
to generate snippets of pipeline code.

It is easily found by going to the URL: `${YOUR_JENKINS_URL}/pipeline-syntax`

### 2.8 Replaying pipelines

When you are developing pipelines and fine tuning some settings sometimes it
doesn't work right away and you need to keep trying. If you are fetching your
pipeline code from a repository (which you _should_ be doing), that might mean
continuously checking new changes into version control before running another
test of your pipeline.

The ability to [replay
pipeline](https://www.jenkins.io/doc/book/pipeline/development/#replay) is here
to help!

This feature allows you to replay a pipeline with some modifications that you
can make directly in the Jenkins UI. This can save a ton of extra work creating
and cleaning up commits.

### 2.9 Shared Libraries

Jenkins shared libraries is a topic that warrants its own blog post, but let's
keep it brief here.

If you want to make sure your pipelines remain DRY (Don't Repeat Yourself) and
want to keep certain logic centralized so that it can be reused in multiple
pipelines across your entire organisation, then creating a shared library is
probably what you want to do.

We have seen projects where the declarative `Jenkinsfile` in each repository is
just this:

```groovy
@Library('awesome-library')_

// this method returns a common declarative pipeline {...}
getCommonDeclarativePipeline()
```

That's DRYer than the sahara!

Other times shared libraries are used for simple things like logging, uploading
artifacts, error handling, etc.

## Conclusion

Jenkins Pipelines are very powerful and there are lots of ways to use them. We
hope that this blog post has given you some knowledge that will help you in
developing an effective CI process. Find out more about our work with Jenkins here.
