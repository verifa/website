---
type: Blog
title: How to Debug Failing Build Agent Pods in Kubernetes-enabled Jenkins
subheading: A step-by-step process for debugging failing Jenkins Build Agent Pods
  in Kubernetes.
authors:
- alarfors
tags:
- Jenkins
- Kubernetes
- Continuous Integration
- Containers
date: 2021-07-05
image: "/static/blog/2021-07-05/how-to-debug-failing-build-agent-pods-kubernetes-jenkins.png"
featured: true
---
**Running Jenkins in a Kubernetes cluster is a great way to enable auto-scaling of the infrastructure hosting your Jenkins Build Agents. However, when the Build Agent Pods fail to start correctly, it can be difficult to troubleshoot. In this short article we look at some simple things you can do to figure out what's going wrong.**

Our Jenkins system is deployed on Kubernetes, with both Jenkins Master and the Jenkins Build Agents running as Pods. We recently had an issue where Build Agent Pods were not starting correctly. Because of our auto-scaling setup, this led to hundreds or thousands of Build Agent Pods being created in the cluster, and the infrastructure (Nodes) scaling up accordingly. This was first discovered when we reached our Cost Cap limit on our EKS cluster. Cleaning up required the deletion of around 7,000 Pods.

It was clear that we had to identify the cause of this issue. The cause itself is not that interesting in itself, but the method of debugging is, and might help you if you are facing similar issues.

### Step 0: Inspect your Jenkins Build Logs

You've probably already done this, but you should of course start by looking at the build logs in Jenkins. Here's what ours were saying:

```
15:58:40 Started by user andreasverifa

15:58:40 Running in Durability level: MAX_SURVIVABILITY

15:58:40 [Pipeline] Start of Pipeline

15:58:41 [Pipeline] podTemplate

15:58:41 [Pipeline] {

15:58:41 [Pipeline] node

15:58:47 Created Pod: kubernetes staging/example-projects-andreas-test-1-x34f1-qzgqg-3cggn

15:58:56 Still waiting to schedule task

15:58:56 All nodes of label 'Example-Projects_andreas-test_1-x34f1' are offline

15:58:57 Created Pod: kubernetes staging/example-projects-andreas-test-1-x34f1-qzgqg-1yvt8

15:59:07 Created Pod: kubernetes staging/example-projects-andreas-test-1-x34f1-qzgqg-l3t8l

15:59:17 Created Pod: kubernetes staging/example-projects-andreas-test-1-x34f1-qzgqg-w784y
```

Oh dear, a new Pod created every 10 seconds. It's easy to see why our cluster was filling up with failed Pods.

### Step 1: Inspect your Pods

So why are the Pods failing to start? Let's (kubectl) describe them:

```bash
$ kubectl describe pod example-projects-andreas-test-...

Containers:
  jnlp:
    Container ID:   ...
    Image:          jenkins/inbound-agent:4.3-4-jdk11
    Image ID:       ...
    Port:           <none>
    Host Port:      <none>
    State:          Failed
      Started:      ...
    Ready:          False
    Restart Count:  0
    Requests:
      cpu:     100m
      memory:  256Mi
    Environment:
      ...
    Mounts:
      ...
  build:
    Container ID:   ...
    Image:          ...
    Image ID:       ...
    Port:           <none>
    Host Port:      <none>
    State:          Running
      Started:      ...
    Ready:          True
    Restart Count:  0
    Requests:
      cpu:        600m
      memory:     768Mi
    Environment:  <none>
    Mounts:
      ...
```

Our Pipeline defines a Pod that contains two Containers:

* "jnlp" - the container which runs the Jenkins Agent
* "build" - the container in which we build our software and run the pipeline steps

From the description above we can see that the jnlp container has Failed, so let's get the logs and investigate why:

```bash
$ kubectl logs example-projects-andreas-test-... jnlp

INFO: Protocol JNLP4-connect encountered an unexpected exception
java.util.concurrent.ExecutionException: org.jenkinsci.remoting.protocol.impl.ConnectionRefusalException: Unknown client name:
...
Jun 18, 2021 12:04:41 PM hudson.remoting.jnlp.Main$CuiListener error
SEVERE: The server rejected the connection: None of the protocols were accepted
java.lang.Exception: The server rejected the connection: None of the protocols were accepted
```

From these logs we can see that the connection from the jnlp Agent Container to the Jenkins Master failed. And the reason for the failure is that the server (Master) rejected the connection.

The next step in this process is in hindsight unnecessary, but it documents our troubleshooting process and might be useful for you.

### Step 2: Manually Override Pods (Optional)

We've established that the jnlp Container fails to connect to the Jenkins Master. One of the issues with containerized systems is that we often cannot debug what is going wrong on startup without modifying the container. That's what we're going to do now.

Our Jenkins build containers are defined as inline-YAML in our pipeline scripts. Let's take an example pipeline and make the jnlp container sleep for 10000 seconds instead of completing the normal startup routine (where the connection would usually fail):

```groovy
pipeline {
    agent {
        kubernetes {
            yaml '''
apiVersion: v1
kind: Pod
metadata:
  name: andreas-test
spec:
  containers:
  - name: jnlp
    image: jenkins/inbound-agent:4.3-4-jdk11
    command: ["sleep", "10000"]
'''

        }
    }

    stages {
        stage('Say Hello World') {
            steps {
                container("jnlp") {
                    echo "Hello World!"
                }
            }
        }
    }
}
```

On startup, the jnlp container does not register with the Jenkins Master (as we've overridden the default startup/entrypoint with the **command** value). So the Jenkins Job finds no available Nodes and continues spawning Pods... So let's cancel the Jenkins Job.

But we are now left with at least one Pod which is in state "Running" instead of "Failed". This allows us to debug the connection process.
Let's exec into the running container and try running the entrypoint. We can determine the entrypoint by looking at the Dockerfile for this image, for example:

`ENTRYPOINT ["/usr/local/bin/jenkins-agent"]`

```bash
$ kubectl exec -it example-projects-andreas-test-... -- /bin/sh

$ sh -x /usr/local/bin/jenkins-agent
...
exec /usr/local/openjdk-11/bin/java -cp /usr/share/jenkins/agent.jar hudson.remoting.jnlp.Main -headless -tunnel jenkins-agent:50000 -url http://jenkins:8080/ -workDir /home/jenkins/agent ... example-projects-andreas-test-...
...
```

Ok, there's our connect command! We can now run this manually to test the connection.

We have seen that the server is refusing the connection because of "Unknown client name", which suggests the server is not expecting the client (jnlp Agent) to connect.

We manually added a Node in the Jenkins GUI with the name of our Pod. Once this was done, running the entrypoint/Java connect command succeeded in connecting the jnlp Agent to the Jenkins Master!

So we have now deduced that the Agent is failing to connect to the Master because the Master is not expecting the Agent. This suggests that the Kubernetes plugin is failing to "register" the Pod with the Master.

### Step 3: Inspect Kubernetes Plugin Logs

We probably should have jumped straight from Step 1 to Step 3. However, Step 2 helped confirm that there are no external causes for why the Agent is failing to connect to the Master. It is just conditions in Jenkins which are causing the issue.

Suspecting the Kubernetes plugin, we want to see the logs of the Kubernetes plugin. The Jenkins plugin page for the Kubernetes plugin contains the following information:

> For more detail, configure a new Jenkins log recorder for **org.csanchez.jenkins.plugins.kubernetes** at **ALL** level.

Once done, we could see the following output in the new log recorder:

```
Created Pod: kubernetes staging/example-projects-andreas-test-...
... WARNING org.csanchez.jenkins.plugins.kubernetes.KubernetesLauncher launch
Error in provisioning; agent=KubernetesSlave name: example-projects-andreas-test-...
java.lang.NoSuchMethodError: 'java.lang.Object io.fabric8.kubernetes.client.dsl.PodResource.watch(java.lang.Object)'
  at org.csanchez.jenkins.plugins.kubernetes.KubernetesLauncher.launch(KubernetesLauncher.java:170)
  at hudson.slaves.SlaveComputer.lambda$_connect$0(SlaveComputer.java:294)
  at jenkins.util.ContextResettingExecutorService$2.call(ContextResettingExecutorService.java:46)
  at jenkins.security.ImpersonatingExecutorService$2.call(ImpersonatingExecutorService.java:80)
  at java.base/java.util.concurrent.FutureTask.run(FutureTask.java:264)
  at java.base/java.util.concurrent.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1128)
  at java.base/java.util.concurrent.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:628)
  at java.base/java.lang.Thread.run(Thread.java:829)
```

Aha! "**NoSuchMethodError**" is typical of version incompatibilities between plugin dependencies!

### Conclusion

Skipping over some details, we concluded that the error is caused by a version conflict; the ["kubernetes" plugin](https://plugins.jenkins.io/kubernetes)has a dependency on the ["kubernetes-client-api" plugin](https://plugins.jenkins.io/kubernetes-client-api/). We are using fixed version numbers for our plugins, and our kubernetes plugin version was fixed at 1.29.2. Looking at the pom.xml for the source of this version, we could see that the dependency of the kubernetes-client-api plugin was for **version 4.13.2-1**.

Our Jenkins image had recently been updated and this seemed to have updated all the non-fixed versions of plugins, i.e. the transitive dependencies. Upon inspection, we were running **version 5.4.1 of the kubernetes-client-api** plugin!

The quick solution was to upgrade the kubernetes plugin to the latest version, 1.30.0. This immediately resolved the issue and jobs began running as normal again.

The long-term solution to this problem, to prevent it from occurring again, is to build our own tagged & versioned Jenkins image. Each change then means incrementing the tag/version, and so we can more easily roll back changes. This is, in other words, a "complete" Jenkins image, which is bundled with the specific plugin versions we want to use.

We hope this breakdown of our problem and troubleshooting helped you. Feel free to [comment](mailto:info@verifa.io) with any questions or tips on making this information better!
