---
type: Blog
title: "Getting started with Jenkins Config as Code"
subheading: "How to run JCasC without using an orchestrator."
authors:
- jlarfors
featured: true
tags:
  - Jenkins
  - Continuous Integration
date: 2020-10-01
image: /blog/2020-10-01/main.svg
---

**Navigating all the content around Jenkins can look like an impossible challenge, especially for newbies. In this post we'll explore one of its most successful and useful plugins - Jenkins Config as Code.**

If you are using Jenkins as your CI server you've probably heard about the Jenkins config as code plugin. It enables automation of the tool setup whilst providing better control and visibility of your configuration. The plugin (<https://plugins.jenkins.io/configuration-as-code/>) makes it possible to configure Jenkins instances using YAML files for the configuration. I would highly recommend that anyone administrating Jenkins master(s) should check it out, and this blog should help you to get up and running.

We've observed that for newcomers to Jenkins, or even experienced admins, adopting config as code can be a big challenge because most online examples are quite complicated and use orchestrators like Kubernetes, adding an extra layer of difficulty to the setup. So, we thought we'd strip back all the layers of difficulty that are not associated with Jenkins, and provide some basic examples using only Docker. After all, we still need some kind of environment to run Jenkins in :)

Our tutorial should give you a very basic initial setup for you to build on. We will take you through the following steps:

1. Build a Docker image based on the community Jenkins image with a minimal config-as-code setup

2. Extend the config-as-code setup

3. Talk about next steps and tips

## Prerequisite

The only prerequisite for this tutorial is that you have Docker installed and are able to run the community Jenkins image.

```bash
docker run --rm --name jenkins jenkins/jenkins:lts-jdk11
```

## 1. Minimal Config as Code

If you follow the [documentation](https://github.com/jenkinsci/configuration-as-code-plugin#getting-started) on the config-as-code plugin page you'll find concise information on getting started. To summarize, there are 2 things that we need to do:

1. Install the config-as-code plugin to Jenkins

2. Setup the environment variable `CASC_JENKINS_CONFIG` to tell Jenkins where to find the YAML configuration

How this looks in a `Dockerfile` is as follows, with some useful comments:

```bash
# Dockerfile

FROM jenkins/jenkins:lts-jdk11

# copy the list of plugins we want to install
COPY plugins.txt /usr/share/jenkins/plugins.txt
# run the install-plugins script to install the plugins
RUN /usr/local/bin/install-plugins.sh < /usr/share/jenkins/plugins.txt

# disable the setup wizard as we will set up jenkins as code :)
ENV JAVA_OPTS -Djenkins.install.runSetupWizard=false

# copy the config-as-code yaml file into the image
COPY jenkins-casc.yaml /usr/local/jenkins-casc.yaml
# tell the jenkins config-as-code plugin where to find the yaml file
ENV CASC_JENKINS_CONFIG /usr/local/jenkins-casc.yaml
```

And the list of `plugins` can be listed in a `plugins.txt` file which looks as follows:

```bash
# plugins.txt

configuration-as-code
```

And a basic `jenkins-casc.yaml` file which can contain something basic like:

```yaml
# jenkins-casc.yaml

jenkins:
  securityRealm:
    local:
      allowsSignup: false
      users:
        # create a user called admin
        - id: "admin"
          password: "admin"
  authorizationStrategy: loggedInUsersCanDoAnything
```

This creates a base Jenkins master with a username called admin who can login and do anything because of the `authorizationStrategy: loggedInUsersCanDoAnything`.

Let's test that this basic setup works.

```bash
# build the docker image
docker build -t jenkins-casc:0.1 .

# run the docker image and map ports 8080
docker run --rm --name jenkins-casc -p 8080:8080 jenkins-casc:0.1
# ...
# ...
# INFO hudson.WebAppMain$3#run: Jenkins is fully up and running
```

Once you see the message `Jenkins is fully up and running` you can go to your browser `<localhost:8080>` and you should see your Jenkins instance running. Success!

Now let's continue to develop the Jenkins instance further.

## 2. Extending config-as-code setup

The above minimal example should have given you a basic setup to develop further.
Now we are going to add two basic things to this setup:

1. Add another user `dev` that is not an administrator and setup permissions using the [matrix-auth plugin](https://plugins.jenkins.io/matrix-auth/)

2. Create some initial Jenkins pipelines

### 2.1 Add a dev user and permissions

First we need to install the [matrix-auth plugin](https://plugins.jenkins.io/matrix-auth/).
To add plugins simply add them to the `plugins.txt`

```bash
# plugins.txt

configuration-as-code

matrix-auth
```

Now that the plugin is available we can add the configurations to the config-as-code YAML file.

```yaml
# jenkins-casc.yaml

jenkins:
  securityRealm:
    local:
      allowsSignup: false
      users:
        - id: "admin"
          password: "admin"
        - id: "dev"
          password: "dev"
  authorizationStrategy:
    globalMatrix:
      permissions:
        - "Job/Build:dev"
        - "Job/Cancel:dev"
        - "Job/Read:dev"
        - "Job/Workspace:dev"
        - "Overall/Administer:admin"
        - "Overall/Read:authenticated"
        - "Run/Replay:dev"
        - "Run/Update:dev"
```

### 2.2 Create some initial pipeline jobs

The Jenkins config-as-code plugin uses the Jenkins [job-dsl plugin](https://plugins.jenkins.io/job-dsl/) plugin for handling jobs in Jenkins.
There is a nice interactive API for this plugin so you can check what's available: <https://jenkinsci.github.io/job-dsl-plugin/>

This is obviously specific to your Jenkins instance and the plugins that you have available, and the API viewer is available for your instance once you have installed the plugin at `<localhost:8080/plugin/job-dsl/api-viewer/index.html>`

Let's add this plugin to `plugins.txt` as well as the pipeline plugins that we need to create some pipeline jobs in Jenkins

```bash
# plugins.txt

configuration-as-code
job-dsl

blueocean
workflow-job
workflow-cps

matrix-auth
```

Now that we have the necessary plugins installed, let's add the configuration to the config-as-code YAML file.
I always like to split the job-dsl groovy scripts to create jobs away from the YAML, and this can be done by specifying a path to a file that will exist inside our Jenkins Docker image.

```yaml
# jenkins-casc.yaml
---
# specify the seedjob groovy script that will create the  pipelines for us
jobs:
  - file: /usr/local/seedjob.groovy
```

See the full config-as-code YAML file here:

```yaml
# jenkins-casc.yaml

jenkins:
  securityRealm:
    local:
      allowsSignup: false
      users:
        - id: "admin"
          password: "admin"
        - id: "dev"
          password: "dev"
  # authorizationStrategy: loggedInUsersCanDoAnything
  authorizationStrategy:
    globalMatrix:
      permissions:
        - "Job/Build:dev"
        - "Job/Cancel:dev"
        - "Job/Read:dev"
        - "Job/Workspace:dev"
        - "Overall/Administer:admin"
        - "Overall/Read:authenticated"
        - "Run/Replay:dev"
        - "Run/Update:dev"

  # make sure our jenkins master has 1 executor so that we can run our pipelines
  numExecutors: 1

# specify the seedjob groovy script that will create the  pipelines for us
jobs:
  - file: /usr/local/seedjob.groovy
```

Here is an example `seedjob.groovy` script to show you some of the things you can do:

```groovy
// seedjob.groovy

// create an array with our two pipelines
pipelines = ["first-pipeline", "another-pipeline"]

// iterate through the array and call the create_pipeline method
pipelines.each { pipeline ->
    println "Creating pipeline ${pipeline}"
    create_pipeline(pipeline)
}

// a method that creates a basic pipeline with the given parameter name
def create_pipeline(String name) {
    pipelineJob(name) {
        definition {
            cps {
                sandbox(true)
                script("""
// this is an example declarative pipeline that says hello and goodbye
pipeline {
    agent any
    stages {
        stage("Hello") {
            steps {
                echo "Hello from pipeline ${name}"
            }
        }
        stage("Goodbye") {
            steps {
                echo "Goodbye from pipeline ${name}"
            }
        }
    }
}

""")
            }
        }
    }
}
```

Finally we need to update our `Dockerfile` to copy the `seedjob.groovy` file into it:

```bash
# Dockerfile

FROM jenkins/jenkins:lts-jdk11

# copy the list of plugins we want to install
COPY plugins.txt /usr/share/jenkins/plugins.txt
# run the install-plugins script to install the plugins
RUN /usr/local/bin/install-plugins.sh < /usr/share/jenkins/plugins.txt

# disable the setup wizard as we will set up jenkins as code :)
ENV JAVA_OPTS -Djenkins.install.runSetupWizard=false

# copy the seedjob file into the image
COPY seedjob.groovy /usr/local/seedjob.groovy
# copy the config-as-code yaml file into the image
COPY jenkins-casc.yaml /usr/local/jenkins-casc.yaml
# tell the jenkins config-as-code plugin where to find the yaml file
ENV CASC_JENKINS_CONFIG /usr/local/jenkins-casc.yaml
```

Let's test this extra config as before.

```bash
# build the docker image, we can give it a new version
docker build -t jenkins-casc:0.2 .

# run the docker image and map ports 8080
docker run --rm --name jenkins-casc -p 8080:8080 jenkins-casc:0.2
# ...
# ...
# INFO hudson.WebAppMain$3#run: Jenkins is fully up and running
```

Once you see the message `Jenkins is fully up and running` you can go to your browser `<localhost:8080>` and you should see your Jenkins instance running.
Login with user `admin` and password `admin` and you should see your two pipelines.
You can also login as user `dev` with password `dev` if you want the experience of a developer.

## 3. Next steps & tips

Hopefully, this tutorial has given you a very basic introduction to how config-as-code works and will help you configure a Jenkins instance automatically.
Given that all the configurations are version controlled, and there's a process behind handling updates and deployments of Jenkins masters, config-as-code _should_ be the way people now maintain their Jenkins instances. It's the easiest way to avoid typical problems like _Jenkinstein_ and having to manage updates in a maintenance window and hope that nothing goes wrong :)

So, next steps.

### 3.1 Adding more configuration to your config-as-code YAML

My approach to configuring Jenkins and plugins/features that I have not used before is to always do things manually first.
Read the docs, spin up a test instance, play around with the UI and try to get things working. Once I have a working recipe, check the config-as-code configuration for my Jenkins instance at `<localhost:8080/configuration-as-code/viewExport>`

There are nice examples in the configuration-as-code github repository [demo folder](https://github.com/jenkinsci/configuration-as-code-plugin/tree/master/demos).
Also, if you search the repository for any snippets you are looking for there are likely some tests for that feature that you could copy bits of working YAML from.
But, don't expect this to work first time - you'll probably need to be patient and experiment a little.

Finally, you can check the "documentation" for your Jenkins instance. I've still not fully figured out how the syntax should look even after 2+ years working with the plugin.
The link is available on your running instance of Jenkins at `<localhost:8080/configuration-as-code/reference>`

### 3.2 Test, staging, production, etc environments

The types of environments you want to define are up to you, but you **should have at least one other environment separate from production**.
Test things before you deploy to production, and have a test environment for you to play around in.
Try not to experiment with new features or possible breaking changes in
the production instance.
However mild the changes are, things always go wrong :)

### 3.3 Build into image or make a configuration

A common question that comes up in our projects is whether the Jenkins
config-as-code YAML files and groovy scripts should be built into the
Docker image (as we did in our example), or if they should be mounted to
the container at runtime, using something like Kubernetes ConfigMaps, Nomad templates, or Docker configs.
My general rule-of-thumb is that maintaining Docker images is an extra overhead and most of these configs are runtime configs.
So, I would remove them from the Docker build stage and make them some kind of mountable config.

The exception to this rule is the `plugins.txt`.
Plugins can get messy and I have ended up in many debates about strictly versioning the plugins you install vs getting the latest. Let's take an example:

```bash
# plugins.txt

configuration-as-code

# vs

configuration-as-code:1.43
```

In the first case Jenkins will get the latest version of the plugin. In the second case it specifically installs version `1.43`.
The main reason for not getting the latest is to make things deterministic.
However, if you build the `plugins.txt` and the plugin installation process into the Docker image, then this process is deterministic because the plugins are installed into the image and are not installed at runtime.
So, no matter how or where you run the Docker image, you will always have the same plugin versions.
On the flip side, managing specific versions of plugins means you have to continuously check the versions and update them.
Personally, I go with the former method where we use the latest versions
of plugins and build them into the Docker image, mainly because it
feels like less maintenance whilst remaining deterministic.

### 3.4 Orchestrators

In a project with proper infrastructure you are likely using some kind of orchestrator like Kubernetes, Nomad, or Docker Swarm.
The purpose of this tutorial was to strip away the added complexity that comes with these tools.
My hope would be that you can now use some of the orchestrators' capabilities and community support to improve your deployment and configuration of Jenkins.

For example, if you are using Kubernetes, then there is a very powerful [Helm chart](https://github.com/helm/charts/tree/master/stable/jenkins) available for Jenkins.
We have been using [Flux](https://github.com/fluxcd/flux) to deploy Jenkins using [HelmReleases](https://github.com/fluxcd/helm-operator).
For newcomers to this environment, there are a lot of extra layers between the deployment and Jenkins config-as-code, so it is not exactly a gentle introduction. It does provide a much better experience from the operations side, though.
Consider this when you move your Jenkins config-as-code to production.

### 3.5 Plugins

One of the things that made Jenkins so popular and easy to use is its plugin
ecosystem. Conversely, the plugin ecosystem can become a bit of a nightmare
when ensuring cross-plugin compatibility, maintenance, upgrades, and so on.

So, the question arises: should you use plugins? Using Jenkins without
them would make it pretty useless because everything is implemented as a
plugin, including pipelines, config-as-code, job-dsl, docker, kubernetes, etc.
The most important thing is to ensure you have a minimal set of plugins to help you achieve your goals. Also, and this goes for any OSS tool, framework, library,
package, etc. - check the stability of the project and the community around it. Is
it being maintained? Are issues resolved? Are there any any other factors that can help
you reach a good judgement on whether to use a plugin or not.

Fortunately, if you have made it this far in the tutorial you will be jubilant
to hear that managing Jenkins using Config as Code is a **VERY GOOD** approach
to avoiding Jenkins plugin bloat and ending up with a _Jenkinstein_ managing
your CI process.

<img src="/blog/2020-10-01/img-01.png" alt="Halloween Jenkins" style="max-width: 300px;"/>

## Conclusion

Hopefully this tutorial has given you a basic introduction to Jenkins config-as-code and given you a working setup to develop. We will be publishing a tutorial on Jenkins pipelines as a natural follow-up to this.
