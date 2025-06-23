---
type: Blog
title: 'Exploring ORAS: OCI Registries as Storage'
subheading: This post is a short overview of my adventures and my first impressions of ORAS.
authors:
- tlacour
tags:
- Developer Experience
- Open Source
date: 2023-04-27
image: "/blog/exploring-oras-oci-registries-as-storage.png"
featured: true
---

**People always find ways to abuse new tech when it comes out. Back when I first discovered Docker, I didn't take long to also discover the people trying to use it for binary distribution - ramming their artifacts into images.**

**I recently discovered ORAS: *"an attempt to define an opinionated way to leverage OCI Registries for arbitrary artifacts without masquerading them as container images."* This piqued my interest, so I took it for a spin.**

**This post is a short recording of my humble adventures and first impressions.**

## Setup

### Installing ORAS CLI

*Official docs: [CLI/Installation](https://oras.land/docs/cli/installation/)*

First up is installing the ORAS CLI.

Installation instructions are available for all platforms, and there's a Docker image available, so follow the relevant steps and you'll be up and running.

```
$ oras version
Version:        1.0.0
Go version:     go1.20.2
Git commit:     b58e7b910ca556973d111e9bd734a71baef03db2
Git tree state: clean
```

### Setting up an OCI registry (Optional)

I spun up a registry to isolate my mucking-about.
It's fairly simple and makes cleaning up a lot easier.
(I used podman, but Docker should work just fine.)

```
REGISTRY_DIR=~/oras-adventures/registry
mkdir -p $REGISTRY_DIR
podman run -d --name oras_registry -p 5000:5000 -v $REGISTRY_DIR:/var/lib/registry --restart=always registry:2
```

We'll want to add our registry to `/etc/containers/registries.conf` as well:

```
[[registry]]
location = 'localhost:5000'
insecure = true
```

A quick restart of the podman service and we should be good to go:

```
systemctl restart podman
```

To test, let's try pushing an image to our registry.

```
$ podman pull hello-world:latest
Trying to pull docker.io/library/hello-world:latest...
Getting image source signatures
Copying blob e07ee1baac5f done
Copying config feb5d9fea6 done
Writing manifest to image destination
Storing signatures

$ podman tag hello-world:latest localhost:5000/hello-world:latest

$ podman push localhost:5000/hello-world:latest
Getting image source signatures
Copying blob e07ee1baac5f done
Copying config feb5d9fea6 done
Writing manifest to image destination
Storing signatures
```

Alright, looks like our playground is good to go!

## Playing with ORAS

### Pushing an artifact

*Official documentation: [CLI/Pushing](https://oras.land/docs/CLI/pushing)*

Let's push our first artifact.
To do so, we'll need to specify the artifact's type and our file.
I have Verifa's lovely logo at hand, so let's use that.
It's an SVG file, so I'll be using `image/svg+xml` as the artifact type.

```
$ oras push localhost:5000/verifa-logo:0.1.0 --artifact-type="image/svg+xml" ./verifa-logo.svg
Uploading 157f230b1d99 verifa-logo.svg
Uploaded  157f230b1d99 verifa-logo.svg
Pushed [registry] localhost:5000/verifa-logo:0.1.0
Digest: sha256:a1a4a2a322c8989464ebdb8d4af375f7610c5f1c28b8c9dc6e8a1b975b8456a0
```

If you're unsure what type to use, you can check [IANA's media type list](https://www.iana.org/assignments/media-types/media-types.xhtml).

### Finding an artifact

*Official documentation:* *[CLI reference/oras_repo_ls](https://oras.land/docs/cli_reference/oras_repo_ls/)* / *[CLI reference/oras_repo_tags](https://oras.land/docs/cli_reference/oras_repo_tags/)*

A look around the repo reveals both our `verifa-logo` artifact and the `hello-world` image we pushed earlier. Everything is as it should be.

```
$ oras repo ls localhost:5000
hello-world
verifa-logo

$ oras repo tags localhost:5000/verifa-logo
0.1.0
```

### Pulling an artifact

*Official documentation: [CLI/Pulling](https://oras.land/docs/cli/pulling/)*

Let's try pulling our artifact back down again.

Not the most realistic test case, but it's what we've got.

```
$ oras pull localhost:5000/verifa-logo:0.1.0
Downloading 157f230b1d99 verifa-logo.svg
Downloaded  157f230b1d99 verifa-logo.svg
Pulled [registry] localhost:5000/verifa-assets:0.1.0
Digest: sha256:a1a4a2a322c8989464ebdb8d4af375f7610c5f1c28b8c9dc6e8a1b975b8456a0

$ ls
verifa-logo.svg
```

Looks like pulling an artifact dumps its contents in your working directory. Fair enough, I can work with that.

### Artifacts with multiple files

*Official documentation: [CLI/Pushing](https://oras.land/docs/CLI/pushing#pushing-artifacts-with-multiple-files)*

You can push an artifact with multiple files, where each file ends up as a its own layer.
Note that each layer should have a media type assigned.
I'll be pushing the following as an artifact:

- `README.md`
  - A simple markdown file. I'll use `text/markdown`
- `images/`
  - A directory containing two images. ORAS automatically tars up directories, so this layer will contain a single tar file. There's no media type for tar files, I'll just use the common `application/x-tar`

```
$ tree
.
├── images
│   ├── verifa-logo-small.svg
│   └── verifa-logo.svg
└── README.md

$ oras push localhost:5000/verifa-assets:0.1.0 \
  README.md:text/markdown \
  ./images/:application/x-tar
Uploading 21e7f9fe06dd images
Uploading b22b00913462 README.md
Uploaded  21e7f9fe06dd images
Uploaded  b22b00913462 README.md
Pushed [registry] localhost:5000/verifa-assets:0.1.0
Digest: sha256:de6be28cb314cf63e6f7af0d6a3f9ada2ad316e0892a9a99241ff903e49ae6d6
```

Pulling the artifact is the same as before.
It dumps its contents in your working directory, exploding the tar along the way.

```
$ oras pull localhost:5000/verifa-assets:0.1.0
Downloading 21e7f9fe06dd images
Downloading b22b00913462 README.md
Downloaded  21e7f9fe06dd images
Downloaded  b22b00913462 README.md
Pulled [registry] localhost:5000/verifa-assets:0.1.0
Digest: sha256:de6be28cb314cf63e6f7af0d6a3f9ada2ad316e0892a9a99241ff903e49ae6d6

$ tree
.
├── images
│   ├── verifa-logo-small.svg
│   └── verifa-logo.svg
└── README.mdw
```

## Closing thoughts

Having only barely scratched ORAS' surface, I'm quite fond of it.
I'm always on the lookout for generic dependency management tools, and while I wouldn't use ORAS as such (yet), it holds quite a bit of potential. You can build some nice tooling around this, especially with the ability to define/handle custom media types.

I'll be keeping an eye on it for sure.
