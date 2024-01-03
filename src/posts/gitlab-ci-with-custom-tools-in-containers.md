---
type: Blog
title: Supporting GitLab CI with custom tools in containers
subheading: How to Compile and Deploy Custom CI Tools via Docker images in GitLab
authors:
- zlaster
tags:
- Containers
- Continuous Integration
date: 2023-04-28
image: "/static/blog/gitlab-ci-with-custom-tools-in-containers/supporting-gitlab-ci-with-custom-tools-in-containers.png"
featured: true

---

**[In a previous post](https://verifa.io/blog/automatically-package-tools-gitlab-container-registry/) we showed how to build and store Docker images in GitLab. Building on that, we will demonstrate a workflow to Dockerise a custom tool which first needs to be compiled while minimising time to deployment by storing artifacts in the Package Registry.**

***

The repo for this post is [here](https://gitlab.com/verifa/godot_images/-/tree/headless), if you wish to jump straight to the final results.

## Introduction

We often need custom tools in our pipelines. Whether they are modified versions of larger open-source projects or tools developed entirely in-house, simple packagers or complex compilers, these custom tools need to be built and made available to our CI pipelines. Deploying these tools as images leverages the ease of deployment and reliability of Docker.

For this demonstration, we will build and deploy a compilation-like tool, the Godot engine. The Godot engine is a game development editor and compiler tool that is growing in popularity. One of the key reasons for this is that it is free and open-source. This means that, while we could use it as provided, we are also able to create customised variants of it in order to better suit our projects.

We will compile and deploy the Godot engine in a way which allows us to automatically run tests on a project created in Godot. We will also minimise the time it takes to deploy the image, which is valuable when improving or fixing the tool.

## Compiling the Tool

> [!WARNING]
> Godot compiles to a few different artifacts. The only one we care about for this demonstration is the “headless server” build. However to deploy the game we would also need some export templates, and to work on the project we would need an editor build for every platform we wish to use. We’ll address those in a later post.

In order to compile the Godot engine we’ll need a slightly more elaborate than usual compilation environment. Therefore, first we will create a Docker image. The headless server artifact is compiled on a Linux environment, which is convenient for our purposes.
The compilation environment image itself is fairly straightforward; we install `scons`, which is needed to compile Godot, and the libraries needed to compile for the target platform.

```docker
#Linux
FROM ubuntu:20.04
LABEL Author="Zach Laster <zlaster@verifa.io"

RUN apt-get update && apt-get install -y --no-install-recommends \
 ca-certificates \
 build-essential \
 scons \
 pkg-config \
 libx11-dev \
 libxcursor-dev \
 libxinerama-dev \
 libgl1-mesa-dev \
 libglu-dev \
 libasound2-dev \
 libpulse-dev \
 libfreetype6-dev \
 libudev-dev \
 libxi-dev \
 libxrandr-dev \
 git \
 && rm -rf /var/lib/apt/lists/*

# Specify the python version in the scons script file
RUN sed -i "/\#\! \/usr\/bin\/python/c\\#\! \/usr\/bin\/python3" $(which scons)
```

> [!WARNING]
> Linux binaries of Godot often will not run on distributions that are older than the distribution they were built on, so we use Ubuntu 20.04 LTS. It won’t matter for this demonstration, but it would for exporting the game.

We’ve already gone over [how to build this Docker image and store it in a GitLab registry](https://verifa.io/blog/automatically-package-tools-gitlab-container-registry/), so we’ll skip over that part here.

Now that we have our environment image, we can compile Godot! For our CI file we’ll utilise a few techniques which will make the CI jobs easier to work with.

First, we’ll specify a `workflow` block, which allows us to specify high level rules for when any jobs are run.

Second, we’ll move most of the logic for the artifact job into a hidden job. Hidden jobs are denoted by a prefix of `.`, to let GitLab know that this isn’t a job to execute. Instead, they are very useful for definitions blocks. We can use GitLab’s `extends` keyword to base our actual job on hidden jobs, thereby reusing definitions and separating concerns.
The hidden job specifies our stage, where the artifacts will be found and how long to keep them, and two extra script steps. This separation already helps to make responsibilities more clear, as we can see at a glance what is common, boilerplate logic and what is the specific instruction and configuration for building. This will be a significant benefit when we have more artifacts.

As to the actual compilation, our instruction is a single line which we run on the image we built previously. We’ll use `before_script` to clone the official Godot repository from GitHub, and `after_script` to move the artifacts to a bin folder in the root of the workspace.

```yaml
workflow:
  rules:
    - if: $CI_MERGE_REQUEST_ID
      when: never
    - when: always

stages:
  - artifacts

variables:
  GODOT_REPO_URL: https://github.com/godotengine/godot.git
  GODOT_REPO_BRANCH: 3.x

.artifact:
  stage: artifacts
  artifacts:
    expire_in: 1 day
    paths:
      - bin/
  before_script:
    - git clone --depth 1 --branch "$GODOT_REPO_BRANCH" "$GODOT_REPO_URL"
  after_script:
    - if [ -d godot/bin/ ]; then mkdir bin && mv godot/bin/* bin/; fi #Relocate artifacts folder to root

headless:
  extends: .artifact
  image: $CI_REGISTRY_IMAGE/environments/linux:latest
  script:
    - scons --directory=godot  profile=../editor.py platform=server tools=yes target=release_debug
```

> [!WARNING]
> Compiling Godot can take upward of an hour. Be prepared. We’ll address this further down this post.

Pushing this file triggers the pipeline which will build our headless artifact.

> [!NOTE]
> GitLab CI configuration allows us to import files. To make my CI files easier to read, I usually move the definitions and hidden blocks to another file, such as `ci/.definitions.yml`. Larger pipeline specifications will, of course, require more files and possibly techniques such as templating.

## Deploying the Image

Now that our artifact can be built, it’s time to put it into a Docker image for external use.

The following GitLab CI configuration adds a `headless-image` job that depends on the headless build job and artifact. Again, we use hidden jobs to simplify the structure and move the boilerplate elsewhere. This can be very useful in extending the pipeline.

```yaml
.dind:
  image: docker:23.0.1
  services:
    - docker:23.0.1-dind

.image:
  extends: .dind
  stage: images
  variables:
    DOCKER_IMAGE_BUILD_TAG: $CI_REGISTRY_IMAGE/$CI_JOB_NAME_SLUG:$CI_COMMIT_REF_SLUG
    DOCKER_FILE: Dockerfile
  script:
    - docker login -u $CI_REGISTRY_USER -p $CI_REGISTRY_PASSWORD $CI_REGISTRY
    - docker build --pull -t $DOCKER_IMAGE_BUILD_TAG -f $DOCKER_FILE .
    - docker push $DOCKER_IMAGE_BUILD_TAG

headless-image:
  extends: .image
  variables:
    DOCKER_IMAGE_BUILD_TAG: $CI_REGISTRY_IMAGE/headless:$CI_COMMIT_REF_SLUG
    DOCKER_FILE: headless.dockerfile
  needs:
    - job: headless
      artifacts: true
```

The image is created using the specification in `headless.dockerfile`, which is very simple: We use the same version of Ubuntu as before, ensure we have python and git, and copy in our artifact to the user binary folder.

```docker
FROM ubuntu:20.04
LABEL Author="Zach Laster <zlaster@verifa.io>"

RUN apt-get update && apt-get install -y --no-install-recommends \
 ca-certificates \
 git \
 python \
 python-openssl \
 && rm -rf /var/lib/apt/lists/*

COPY bin/godot_server.x11.opt.tools.64 /usr/local/bin/godot

RUN mkdir ~/.cache && mkdir -p ~/.config/godot
```

This should work, but the compilation of the artifact takes much too long to be practical here. We need to improve on this.

## Improving Godot Compilation Time

While this is a complex and advanced subject, and how to improve the compilation time of an artifact tends to depend a lot on the artifact, the usual answers are to either reduce what we are building or to cache something.

Godot provides numerous flags for controlling the build output and process. These will affect both the artifact size and build time. There is a [file in the repo](https://gitlab.com/verifa/godot_images/-/blob/main/editor.py) which specifies what we want to build.

More impactfully, we can make use of GitLab caching to significantly improve build times after the first time. Since `scons` can cache its work to speed up compilation on subsequent attempts, we can inform it to do that and have GitLab track that cache.

```yaml
.scons_cache:
  variables: &scons_cache_variables
    SCONS_CACHE: ../.scons-cache/ #Up a directory since we are working in ./godot/
    # Limit to 7 GiB to avoid having the extracted cache fill the disk.
    SCONS_CACHE_LIMIT: 7168
  cache:
    - key: "$CI_COMMIT_REF_SLUG-scons-$CI_JOB_NAME_SLUG"
      paths:
        - .scons-cache/
      policy: pull-push
```

We’ll use a mixture of YAML anchors and `!reference` for the `scons_cache` hidden block. Using `!reference` gives us more control and makes it easier to merge with other logic, but it doesn’t work well with variables blocks if we have more than one imported set of variables.
Adding these lines to `.artifact` adds the cache.

```yaml
variables:
    <<: *scons_cache_variables
  cache:
    - !reference [.scons_cache, cache]
```

`scons` caching drastically reduces our compilation times, even when some of the code changes. However, the linking portion can still take ten minutes or more. A long time to wait if we don’t actually need to do anything.

## Reusing Artifacts

Rather than compile Godot every time we run the job, we could first check if we have already compiled the current version. The easiest check for this is if we already have a compiled artifact of the current git commit SHA which we are about to clone.

To accomplish this, we need to store our compiled artifacts in a way than enables us to retrieve them. We could utilise GitLab job artifacts, but there’s some limitations with retrieving specific versions and the API is not streamlined for this use case.
Instead, we can utilise another GitLab feature, the Package Registry.

To use the GitLab package registry, the only tool we’ll need is `wget`. Pushing to and retrieving from the registry is handled via a REST API.

`wget --header="JOB-TOKEN:$CI_JOB_TOKEN" -O ${ARTIFACT_PATH}${ARTIFACT_FILE} $ARTIFACT_URL`

The above line will use our current job token to download an artifact file from a given URL.

`ARTIFACT_URL=${CI_API_V4_URL}/projects/${CI_PROJECT_ID}/packages/generic/$GENERIC_PACKAGE_NAME/$SRC_REMOTE_SHA/$ARTIFACT_FILE`

The artifact URL is fairly predefined; we only decide the last three values. The structure of `<package name>/<version>/<filename>` is fixed, but we can use almost anything for them. For our case, we’ll use the job name slug as the package name, the Godot repo SHA we are compiling as the version, and the name of the artifact as the actual file name.

Uploading works the same way as retrieval: `wget --header="JOB-TOKEN:$CI_JOB_TOKEN" --method=PUT --body-file="${ARTIFACT_PATH}$ARTIFACT_FILE" "$ARTIFACT_URL"`. With very little around these we can push and pull on the registry.

```yaml
.artifact_package:
  variables: &artifact_package_variables
    ARTIFACT_FILE: example_file
    ARTIFACT_PATH: bin/
    GENERIC_PACKAGE_NAME: $CI_JOB_NAME_SLUG
  download: #Bash commands we will execute to download the artifact based on variables
    - if [ -z $SRC_REMOTE_SHA ]; then SRC_REMOTE_SHA=`git ls-remote $GODOT_REPO_URL $GODOT_REPO_BRANCH | head -1 | sed "s/\t.*//" | tee godot.sha`; fi; echo "Compiling Godot repo:"" $SRC_REMOTE_SHA"
    - ARTIFACT_URL=${CI_API_V4_URL}/projects/${CI_PROJECT_ID}/packages/generic/$GENERIC_PACKAGE_NAME/$SRC_REMOTE_SHA/$ARTIFACT_FILE; echo $ARTIFACT_URL
    - (wget --header="JOB-TOKEN:$CI_JOB_TOKEN" -q --spider $ARTIFACT_URL && mkdir -p ${ARTIFACT_PATH} && wget --header="JOB-TOKEN:$CI_JOB_TOKEN" -O ${ARTIFACT_PATH}${ARTIFACT_FILE} $ARTIFACT_URL) || true
  check: #Bash command to check if the artifact exists and exit if so
    - if [ -f ${ARTIFACT_PATH}$ARTIFACT_FILE ]; then echo Retrieved":" $ARTIFACT_FILE from package registry; exit 0; fi
  upload: #Upload the artifact
    - ls "${ARTIFACT_PATH}$ARTIFACT_FILE"
    - SRC_REMOTE_SHA=`cat godot.sha`
    - ARTIFACT_URL=${CI_API_V4_URL}/projects/${CI_PROJECT_ID}/packages/generic/$GENERIC_PACKAGE_NAME/$SRC_REMOTE_SHA/$ARTIFACT_FILE; echo $ARTIFACT_URL
    - wget --header="JOB-TOKEN:$CI_JOB_TOKEN" --method=PUT --body-file="${ARTIFACT_PATH}$ARTIFACT_FILE" "$ARTIFACT_URL"
```

With this last piece, our .artifact block looks like this

```yaml
.artifact:
  stage: artifact
  tags:
    - offload
  artifacts:
    expire_in: 1 day
    paths:
      - bin/
  variables:
    <<: *scons_cache_variables
    <<: *artifact_package_variables
  cache:
    - !reference [.scons_cache, cache]
  before_script:
    - !reference [.artifact_package, download] # We import the bash commands
    - !reference [.artifact_package, check]
    - git clone --depth 1 --branch "$GODOT_REPO_BRANCH" "$GODOT_REPO_URL"
  after_script:
    - if [ ! -d godot/bin/ ]; then exit 0; fi
    - mkdir -p bin && mv godot/bin/* bin/ #Relocate artifacts folder to root
    - !reference [.artifact_package, upload]
```

Now when we next run the compilation job, it will first check if we’ve already pushed a compiled artifact to the registry and, if so, use that. If there isn’t a package for the current SHA, then the job runs as before and then pushes the resulting artifact to the registry. This cuts our compilation job time down from 15 minutes to **50 seconds 🎉** in the case of having previously compiled the artifact for this version of Godot.

## Summary

We now have a Godot image which we can use to run tests on our Godot projects. We’ve verified that the image works and the contained Godot executable runs and knows its name and version.

This process and workflow is applicable to any number of tools we might wish to compile and deploy using docker images. The ability to run containers containing custom-built tooling can have a massive impact on our CI jobs.

In future blog posts, we will explore how to handle compiling for multiple architectures, as well as handling images which contain multiple artifacts.
<!--TODO: Add link to next post here-->

[GitLab Repo](https://gitlab.com/verifa/godot_images/-/tree/headless)

> [!NOTE]
> A note on the repo: The compilation job is tagged in the repo as needing to run on the “offload” runner. We have a high-powered self-hosted runner in use, which reduces used CI quota and does the job much faster. To use this repo’s CI configuration in your own project you would need to remove that tag or provide a self-hosted runner with the same tag.
