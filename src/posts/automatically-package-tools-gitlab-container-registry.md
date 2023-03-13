# Automatically package tools into the GitLab Container Registry

Stakeholders: Zach Laster
Status: In Review
Type: blog - tech

Subtitle: How to optimize GitLab CI runtime environments using custom Docker images

Kicker: **Often when running CI/CD jobs we need to use custom built tools and applications. While we could download the things we need each time we run a relevant job, it would be more efficient to have them already available on the images we are using. Fortunately, we can build our own Docker images, and we can have GitLab manage them for us.**

---

# Introduction

GitLab CI allows us to run jobs on our repos. These can be tests, compilation, deployment, or anything else you’d like to do with your code. GitLab CI uses Docker images to provide the environments in which it runs jobs, and we can specify what image to use for each job. While images often come from Docker Hub, we can use GitLab’s Container Registry to have a private registry of our own.

# Process

For this test, I wanted a GitLab project that, on push to repository, would build, tag, and push an image to the container registry for that project.

As an extension, I tested using the resulting image.

## Itch.io Butler

For the purposes of this test, I wanted something relatively simple to dockerize that still reflected a typical use case. [Itch.io](http://Itch.io) is a distribution platform commonly used by “indie developers” to share digital games, assets, printables, and other products. It has a CLI tool named Butler which is used to automate the deployment of products to created pages on the platform.

The following bash script downloads and unpacks butler inside our docker image.

```bash
#!/usr/bin/bash

mkdir -p /opt/butler/bin 
cd /opt/butler/bin

curl -L -o butler.zip https://broth.itch.ovh/butler/linux-amd64/LATEST/archive/default
unzip butler.zip

# GNU unzip often fails to set the executable bit, regardless of the state in the .zip
# In other contexts, this command might require sudo, but inside an image that usually doesn't work!
chmod +x butler
```

## Dockerfile

Docker images are created by building from a Dockerfile. A Dockerfile specifies a base image to build from and then applies layers to that image in order to produce the desired filesystem/installation state. Often, those layers are a series of annotated bash script lines.

Ubuntu makes a good base in our case, as it is very small (less than 30MB). From there we only need to ensure we have the tools used in the script and to run the script itself.

```docker
FROM ubuntu:23.04
Label Author="Zach Laster <zlaster@verifa.io>"

RUN apt-get update && apt-get install -y --no-install-recommends \
	ca-certificates \
	unzip \
	zip \
	curl \
	&& rm -rf /var/lib/apt/lists/*

ADD get_butler.sh /opt/butler/get_butler.sh
RUN bash /opt/butler/get_butler.sh
RUN /opt/butler/bin/butler -V

ENV PATH="/opt/butler/bin:${PATH}"
```

Note that we also add `butler` to the PATH.

## Building

With the above files, we are able to create our image (assuming we have docker or some other image building tool installed). There are other options, but to build our image in GitLab CI we’ll use “[Docker-In-Docker](https://docs.gitlab.com/ee/ci/docker/using_docker_build.html#use-docker-in-docker)”.

Since all we want to do is create a Docker image and push it to our private registry, our CI pipeline is quite straightforward. The simplest CI file could look like this:

```yaml
image: docker:23.0.1 # Specify the Docker image
services:
  - docker:23.0.1-dind # We are using Docker-In-Docker

# We define a variable to based on GitLab-provided variables.
# This will be used as the name of the tag of our image.
# Note that we use the REF_SLUG and not REF_NAME.
variables:
  DOCKER_IMAGE_BUILD_TAG: $CI_REGISTRY_IMAGE:$CI_COMMIT_REF_SLUG

# Build job in build stage which runs the Docker commands.
build:
  stage: build
  script:
		# Our login is provided to us by GitLab.
    - docker login -u $CI_REGISTRY_USER -p $CI_REGISTRY_PASSWORD $CI_REGISTRY
		# Build and tag image
    - docker build --pull -t $DOCKER_IMAGE_BUILD_TAG .
		# Push image with tag. This puts it in the repo's registry.
    - docker push $DOCKER_IMAGE_BUILD_TAG
```

If we push these files to a repo then the CI pipeline will make us an image. Simple!

<aside>
💡 If you are running a self-hosted GitLab instance, you might need to ensure that the container registry is enabled. At the time of this writing, SaaS GitLab seems to have it enabled for every group/project by default, with no way to disable it.

</aside>

## Fleshing out and Testing

My full CI file looks like this:

```yaml
image: docker:23.0.1
services:
  - docker:23.0.1-dind

stages:
  - build
  - test
  - release

variables:
  DOCKER_IMAGE_BUILD_TAG: $CI_REGISTRY_IMAGE:$CI_COMMIT_REF_SLUG
  DOCKER_IMAGE_LATEST_TAG: $CI_REGISTRY_IMAGE:latest

before_script:
  - docker login -u $CI_REGISTRY_USER -p $CI_REGISTRY_PASSWORD $CI_REGISTRY

build:
  stage: build
  script:
    - docker build --pull -t $DOCKER_IMAGE_BUILD_TAG .
    - docker push $DOCKER_IMAGE_BUILD_TAG

test-version:
  image: $DOCKER_IMAGE_BUILD_TAG
  stage: test
  before_script: []
  script:
    - ./tests/check_version.sh

# When pushing the the main branch of the repo, we also tag the image as "latest"
release-image:
  stage: release
  script:
    - docker pull $DOCKER_IMAGE_BUILD_TAG
    - docker tag $DOCKER_IMAGE_BUILD_TAG $DOCKER_IMAGE_LATEST_TAG
    - docker push $DOCKER_IMAGE_LATEST_TAG
  # This particular job should only run when we push to the branch named 'main'
  only:
    - main
```

And the `check_version.sh` file looks like this:

```bash
#!/usr/bin/bash
butler -V # This just outputs the version identifier of butler
```

This gives a quick verification that butler is correctly set up and that the image is working.

The CI file also now tags the image as `latest` if it passes the test and this was a push to the main branch.

<aside>
💡 Note that we’d need to do something smarter to not push images that are untested/fail the test, possibly using temporary tags and cleanup rules.

</aside>

<aside>
💡 When using git on Windows, executable files such as the version test script may not be flagged with the executable flag. You should do this explicitly when committing the file.
`git update-index --chmod=+x tests/check_version.sh`

</aside>

# Testing it Out

Now we have an image in our project’s container registry and we’ve even made sure it works within the project. Let’s try using it in a different project.

Let’s spin up a repo that contains only a `.gitlab-ci.yml` file:

```yaml
# .gitlab-ci.yml
test-version:
  image: registry.gitlab.com/verifa/examples/docker-butler:latest
  stage: test
  script:
    - butler -V
```

Let’s initialize a repo locally, commit the file, and push it to a new GitLab repo.

```bash
git init
git add .
git commit -m"Test the butler version in the latest image"
git remote add origin git@gitlab.com:verifa/examples/butler-image-test.git
git push --set-upstream origin main
```

If we check the pipeline job, we see this:

![Failed](/blogs/automatically-package-tools-gitlab-container-registry/failed_pipeline.png)

Well that’s not good!

Turns out, we can’t access this image outside of our initial project! That’s because the image is in a private project.

There’s several ways to address this. A very simple option is to go to the GitLab project page, **Settings→CI/CD→Token Access** and add the name of our test project (including the group name) to the `CI_JOB_TOKEN` access list.

Testing after that yields this:

![Passed](/blogs/automatically-package-tools-gitlab-container-registry/passed_pipeline.png)

Job done!

That’s great for a few projects, but what if we have 100 projects that all use this image? We probably don’t want to add the name to the list every time we spin up a new project.

The description for Token Access is confusing at best, but let’s try disabling the `CI_JOB_TOKEN` management in the image project completely.

![Passed Again](/blogs/automatically-package-tools-gitlab-container-registry/passed_pipeline2.png)

Well that’s surprising. That system is actually a whitelist, and disabling it permits anyone with access to the repo to use the images.

The [documentation for this feature](https://docs.gitlab.com/ee/ci/jobs/ci_job_token.html#allow-access-to-your-project-with-a-job-token) doesn't mention how access works when it is disabled, but the documentation for the original, deprecated, outbound behavior says that when it is disabled then the user's access permissions are used, so it is probably that.

# Conclusion

It’s very easy to set up a pipeline whereby our own custom tools and processes can be containerized for future use, which is a very valuable way of making our pipelines faster!

Here we’ve seen how to do this with a simple tool we download from a public URL, but it’s also practical to do this with artifacts from other projects or installations from other sources.

Accessing the image afterward is somewhat more difficult. I’m still exploring how to make the container registry available to other projects within the same group and beyond. Adding projects to the `CI_JOB_TOKEN` access list works well enough, but is a bit impractical in a larger organization with many projects. Meanwhile turning off Token Access works in the same way blowing a hole in something lets water through; it works, but it’s not very managed and might be unsafe. It’s also likely that disabling the Token Access flag only permits pipelines started by users with access to the project, which might have other issues. This will be looked at more in the future.

I'm working on some related test projects to have code compiled using a custom image and then to Dockerize that into its own image. Stay tuned for more!

* [Image Repo](https://gitlab.com/verifa/docker-itchio-butler)
* [Test Repo](https://gitlab.com/verifa/docker-itchio-butler-test)
