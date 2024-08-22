---
type: Blog
title: "How to automatically deploy applications to a private GKE Cluster using CloudBuild"
subheading: "This article describes the problem of automatically deploying code to a private GKE cluster and my workaround solution to it."
authors:
- lsuomalainen
featured: true
tags:
  - Kubernetes
  - GCP
  - Continuous Delivery
date: 2020-11-05
image: /static/blog/2020-11-05/main.png
---

I have been using Google Cloud Platform (GCP) and Google Kubernetes Engine (GKE) for quite some time and I like them a lot. Another GCP product I have become very fond of is CloudBuild which I use to automatically build and deploy code. Compared to previous Jenkins-based CI/CD pipelines I've used, CloudBuild offers similar customizability without the need to manage a build server, and with a significantly smaller configuration code base.

However, the more I use these tools the more I encounter shortcomings and corner cases that are not obvious at first glance. Sometimes these are minor and can be disregarded, but on other occasions they need some workarounds.

This article describes the problem of automatically deploying code to a private GKE cluster and my workaround solution to it. Understanding this article will require some knowledge on GKE or Kubernetes in general.

## The Problem

CloudBuild is a fantastic tool. A single CloudBuild run takes a configuration file, spins up a VM, and runs commands in order in a series of Docker containers which all have a common volume to write on. Typically, you pull up your source code from a source code repository, and build one or more Docker images (and perhaps other artifacts) which are then stored in the confusingly named Container Registry. My typical cloudbuild.yaml configuration file looks like this:

```bash
steps:
- name: 'gcr.io/cloud-builders/docker'
  args: [
    'build',
    '-t', 'eu.gcr.io/$PROJECT_ID/$REPO_NAME:$COMMIT_SHA', '.', '-f${_DOCKERFILE}',
  ]

- name: 'gcr.io/cloud-builders/docker'
  args: ["tag", "eu.gcr.io/$PROJECT_ID/$REPO_NAME:$COMMIT_SHA", "eu.gcr.io/$PROJECT_ID/$REPO_NAME:latest"]

- name: 'gcr.io/cloud-builders/kubectl'
  args: ['set', 'image', 'deployment/${_KUBERNETES_DEPLOYMENT}', '${_CONTAINER_NAME}=eu.gcr.io/$PROJECT_ID/$REPO_NAME:$COMMIT_SHA']
  env:
  - CLOUDSDK_CONTAINER_CLUSTER=${_KUBERNETES_CLUSTER}
  - CLOUDSDK_COMPUTE_ZONE=${_COMPUTE_ZONE}

images: ['eu.gcr.io/$PROJECT_ID/$REPO_NAME:$COMMIT_SHA', "eu.gcr.io/$PROJECT_ID/$REPO_NAME:latest"]
```

Each one of the steps is a distinct Docker container. First, CloudBuild builds a Docker image and tags it with the commit SHA. The image is then tagged as 'latest'. In the next step the newly built image is used to replace the previous container in a Kubernetes deployment. Finally, the image with both tags is stored in the Container Registry.
This works fine when you have a GKE cluster without IP restricted access. However, if you have a private GKE cluster i.e. you are using Authorized Networks to restrict IP addresses from which your cluster can be accessed, the third step 'kubectl set' will fail to connect to the cluster.

This is because CloudBuild is a managed service, running on an arbitrary managed server with an arbitrary IP address, that you probably haven't allowed in your Authorized Networks configuration.

Now, obviously, you could just add the whole IP range CloudBuild uses and it would work out fine. But, the bad news here is that the range in question is the whole public GCE IP address range, so allowing it will mean your Authorized Networks configuration is practically useless.

## The Solution

GCP doesn't offer any smart or neat solutions to this problem, so I decided to hack my way around it. The main idea is to surround the 'kubectl set' statement with steps containing scripts that will add the CloudBuild VM's IP address to the cluster's Authorized Networks, and then remove the address once the deployment is updated. To perform such operations in my scripts, I need to include Service Account credentials which have permissions to modify the authorized networks list.

To make matters worse, CloudBuild doesn't allow for injecting the keys from outside of the repo from which the source code is pulled (although nowadays this probably could be circumvented by using GCP's Secret Manager), so we have to include the key to the Service Account within the repo. Of course, storing keys in the source code repo is something we should not do by default, at least not in the plaintext, so I alleviate the issue somewhat by using GCP's Key Management Service (KMS) to encrypt the key in the repo.

## Step 1

For this scheme to work, we need to make a new service account which has permissions to modify the Authorized Networks configuration in the private GKE cluster. We also need to give the CloudBuild service account permissions to decrypt the encrypted key for the former account. Cloudbuild service account permissions should look like this:

![a screenshot](/static/blog/2020-11-05/img-01.png)

For the service account used to manage the Authorized Networks configuration, only Kubernetes Engine Admin role is needed. Create that account and download it's key as a json. Rename your key to something less complicated: I named mine kubernetes-admin.json

Create a keyring with KMS and a symmetric key to it that will be used to encrypt the key. I named mine cloud-build and cloud-build-encrypt respectively, but you can name yours whatever you want. Now, in your repo, encrypt the key by running:

```
gcloud kms encrypt --key=cloudbuild-encrypt
--ciphertext-file=kubernetes-admin.json.enc
--plaintext-file=kubernetes-admin.json --location=global --keyring=cloud-build
```

Now you have an encrypted key `kubernetes-admin.json.enc` which you can push back to your repo. Remember to not push the plaintext. Better yet, straight up delete it after you're done encrypting.

## Step 2. Add the decryption step to your pipeline

Next, before getting to authorized network modifications, let's decrypt the key! In our cloudbuild.yaml add the following step:

```
- name: gcr.io/cloud-builders/gcloud
  args:
  - kms
  - decrypt
  - --ciphertext-file=${_KMS_KEYNAME}.enc
  - --plaintext-file=${_KMS_KEYNAME}
  - --location=global
  - --keyring=cloud-build
  - --key=cloudbuild-encrypt
  - --project=${_PROJECT}
```

`_KMS_KEYNAME` is my substitution for 'kubernetes-admin.json' and `_PROJECT` is, you guessed it, the project in which the encryption key is located. This decrypts the key and makes it available in /workspace/ in your current Cloudbuild VM.

## Step 3. Add the script that finds the Cloudbuild VM's IP address and adds it to the Authorized Networks configuration

My next CloudBuild step looks like this:

```
-name: 'gcr.io/secret-fleuri-project/cloud_build_add'
  env:
    - CREDENTIALS=/workspace/${_KMS_KEYNAME}
    - PROJECT=${_PROJECT}
    - CLUSTER=${_KUBERNETES_CLUSTER}
    - REGION=${_COMPUTE_ZONE}
```

It runs the custom image called `cloud_build_add` I've uploaded to the Container Registry. A simple python script runs in it. Now, for proprietary reasons, I won't provide the script in full, but I'll walk you through what it does so you can build it yourself.

First, it retrieves the IP address of the VM by querying the Compute Engine API:

<https://metadata.google.internal/computeMetadata/v1/instance/network-interfaces/0/access-configs/0/external-ip>

Second, the script assumes the service account to modify the k8s cluster using the key we decrypted earlier.

Then, using [this API Method](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters/list), it gets the details of the cluster we're deploying to, especially the Authorized Networks configuration:

```
master_authorized_networks_config =
response['clusters'][0]['masterAuthorizedNetworksConfig']
#This assumes your cluster was first on the list
```

Penultimately, we construct the new configuration by appending the VM's IP to the list like this:

```
cloud_build_address = [{'displayName': 'CloudBuild', 'cidrBlock': ip}]
master_authorized_networks_config['cidrBlocks'] += cloud_build_address
```

Now, using the updated `master_authorized_networks_config`, we construct a [ClusterUpdate](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/ClusterUpdate) and use it to update the cluster configuration using [this API call](https://googleapis.dev/python/container/latest/gapic/v1/api.html#google.cloud.container_v1.ClusterManagerClient.update_cluster).

## Step 4. The clean up

Now we can run the `kubectl set image` step in the pipeline. The whole configuration looks like this:

```
steps:
- name: 'gcr.io/cloud-builders/docker'
  args: ['build', '-t', 'eu.gcr.io/${_PROJECT}/$REPO_NAME:$COMMIT_SHA', '.', '-f${_DOCKERFILE}' ]

- name: 'gcr.io/cloud-builders/docker'
  args: ["tag", "eu.gcr.io/${_PROJECT}/$REPO_NAME:$COMMIT_SHA", "eu.gcr.io/${_PROJECT}/$REPO_NAME:latest"]

- name: gcr.io/cloud-builders/gcloud
  args:
  - kms
  - decrypt
  - --ciphertext-file=${_KMS_KEYNAME}.enc
  - --plaintext-file=${_KMS_KEYNAME}
  - --location=global
  - --keyring=cloud-build
  - --key=cloudbuild-encrypt
  - --project=${_PROJECT}

- name: 'gcr.io/secret-fleuri-project/cloud_build_add'
  env:
    - CREDENTIALS=/workspace/${_KMS_KEYNAME}
    - PROJECT=${_PROJECT}
    - CLUSTER=${_KUBERNETES_CLUSTER}
    - REGION=${_COMPUTE_ZONE}

- name: 'gcr.io/cloud-builders/kubectl'
  args: ['set', 'image', 'deployment/${_KUBERNETES_DEPLOYMENT}', '${_CONTAINER_NAM                                 E}=eu.gcr.io/${_PROJECT}/$REPO_NAME:$COMMIT_SHA']
  env:
   - CLOUDSDK_CONTAINER_CLUSTER=${_KUBERNETES_CLUSTER}
   - CLOUDSDK_COMPUTE_ZONE=${_COMPUTE_ZONE}
   - CLOUDSDK_CORE_PROJECT=${_PROJECT}

- name: 'gcr.io/cloud-builders/kubectl'
  args: ['set', 'image', 'deployment/${_KUBERNETES_CRON_JOB_POD}', '${_CRON_JOB_CONTAINER}=eu.gcr.io/${_PROJECT}/$REPO_NAME:$COMMIT_SHA']
  env:
    - CLOUDSDK_CONTAINER_CLUSTER=${_KUBERNETES_CLUSTER}
    - CLOUDSDK_COMPUTE_ZONE=${_COMPUTE_ZONE}
    - CLOUDSDK_CORE_PROJECT=${_PROJECT}

- name: 'gcr.io/secret-fleuri-project/cloud_build_remove'
  env:
    - CREDENTIALS=/workspace/${_KMS_KEYNAME}
    - PROJECT=${_PROJECT}
    - CLUSTER=${_KUBERNETES_CLUSTER}
    - REGION=${_COMPUTE_ZONE}

images: ['eu.gcr.io/${_PROJECT}/$REPO_NAME:$COMMIT_SHA', "eu.gcr.io/${_PROJECT}/$REPO_NAME:latest"]
```

The last step `cloud_build_remove` is mostly similar to `cloud_build_add`, but the script does not need to find out the VM´s IP address because it's not adding it. Instead, it goes through `master_authorized_networks_config['cidrBlocks']` and removes every entry with the displayName `CloudBuild` (There could be multiple if your 'set image' step fails for some reason).

## Summary

This was my pretty hacky way of deploying image versions to a private GKE cluster. Apparently, there aren't very many texts concerning the topic on the internet apart from this StackOverflow question. Other possible solutions to [this](https://stackoverflow.com/questions/51944817/google-cloud-build-deploy-to-gke-private-cluster) problem could be configuring a [proxy](https://cloud.google.com/solutions/creating-kubernetes-engine-private-clusters-with-net-proxies) to the private cluster, or there could be a solution that would pull new images from an outside repository instead of trying to push them to the cluster. If you have a better idea of how to solve it let us know!
