---
type: Blog
title: "The C and C++ Package Manager"
subheading: "Conan is an open source project for managing software components (or packages) for C and
C++. Supporting C and C++ packages requires a flexible model, especially to allow packages to
be pre-compiled into binaries, and Conan really hits the mark. This blog gives a short
introduction to Conan and how it could help your project."
authors:
  - jlarfors
tags:
  - Conan
  - Open Source
  - Continuous Integration
date: 2020-09-04
image: /static/blog/2020-09-04/main.svg
featured: true
---

**Managing dependencies in C and C++ projects is a common challenge and lots of different ways have been invented to deal with these issues, each with varying results. Conan aims to solve this challenge once and for all by providing a universal approach to managing dependencies with all the benefits of a package manager.**

Managing dependencies in C and C++ projects has always presented a challenge. As such I
have seen a variety of inventive ways to share C and C++ packages. A very non-exhaustive list
of such approaches would cover the main strategies, such as:

- Checking in library and header files into version control;
- Copying files from a shared drive;
- Using Git submodules or using APT packages (and sometimes a combination of these).
  None of these approaches were designed for managing C and C++ packages, thus all have
  drawbacks. This is where Conan comes in.

## What is Conan and how does it work?

Conan is a package manager designed for C and C++, meaning it supports multiple platforms
and build tools, stores source code, as well as pre-built libraries/binaries in a remote repository.
Conan, much like any other package manager, provides the simple concepts of creating,
publishing and resolving packages, so there are no surprises there. However, most package
managers do not need to deal with the C and C++ landscape - which means supporting
multiple compilers, linkers and a vast array of build tools.

Conan has quite a simple yet powerful model comprising settings and options to deal with the
different build environments (such as compilers and linkers). Settings typically define project-
wide and environment related settings, such as the OS and compilers, whereas options define
package-specific options that can have defaults (such as whether a package should be statically
or dynamically linked).

If you build a package with a different setting or option, Conan will treat this as a different binary.
This, on a very high level, is how Conan enables distributing pre-built binaries across multiple
platforms and environments.

![conan diagram](/static/blog/2020-09-04/conan-diagram.svg)

Where Conan also packs a punch is when it comes to integrating with multiple build tools. I like
to think of it as Conan generates the "glue" for your build system to understand which packages
Conan has resolved and where to find the relevant files (e.g. where to find header and library
files).

With CMake, for example, Conan will generate a .cmake file that defines CMake targets for
each of the required packages that can be used in your [CMakeLists.txt file](https://docs.conan.io/en/latest/integrations/build_system/cmake/cmake_generator.html#targets-approach).

One interesting thing to note is that if your desired combination of settings and
options has not yet been pre-compiled into a Conan package, Conan also stores the source
code for creating the package. You can then build Conan packages as part of "resolving" them.
And this does not just include direct dependencies, but also indirect ("transitive") dependencies.
That's right: Conan of course supports transitive dependencies too!

## The Conan community

For any framework to become successfully adopted and embraced in today's market there
needs to be a thriving community, so let's talk a bit about the one surrounding Conan.
For starters, all the source code is available on GitHub for you to explore.

[The documentation is very nice](https://docs.conan.io/en/latest/introduction.html) and there is an active Slack channel
in the Cpplang Slack workspace. [JFrog](https://jfrog.com/) are actively supporting the project with conferences [Conan Days -](https://conandays.conan.io/) and free online training is available [in partnership with Verifa](https://verifa.io).

Last but not least, [Conan Center](https://conan.io/center/) is a publicly available Conan
repository with a vast number of open source projects that have already been "Conanized". In
other words, it's the process of taking an ordinary piece of source code and creating a Conan
package out of it.

The Conan Center is populated by the conan-center-index repository on GitHub which contains
all the Conan recipes for [creating the packages](https://github.com/conan-io/conan-center-index).

## How Conan can cure your project woes

Still not convinced Conan is worth taking a closer look at? Then let's talk about some of the
pains that I have experienced in the past without it, and how Conan solves those problems.

### Dependency tree (or Bill of Materials)

Understanding what has gone into a final product in C and C++ can be really, really difficult, and
often requires a lot of undesired process and documentation.

But what if there was a convenient file that just specified the dependencies (or requirements)
and whether each dependency was used only for the build/test phase, or if they actually got
shipped with the product?

Hmmm, it sounds like a package manager would solve this problem! Bonus points for tools like
[oss-review-toolkit](https://github.com/oss-review-toolkit/ort) that support Conan (thanks to a
contribution from [Verifa](/) to help build the clearing process into Continuous Integration.

### Slow build times

It is quite common for teams to be building a set of libraries or an entire platform, which
application teams then develop their logic on top of. Often the application teams need to rebuild
the libraries or platform because what they get delivered are not pre-built libraries, but the
source code with instructions for how to include it in their projects.

What if the application team could simply specify the libraries or platform they want to use with a
semantic version range, and then only had to compile their code and not the whole product?
A versatile package manager could solve this by providing pre-built binaries.

### Development environment dependencies

This is perhaps a slightly more obscure use case for Conan, but I wanted to mention it. Lots of
teams create a README or some instructions for setting up their development environment
which includes things like the compiler and linter. What if you could specify these build
requirements in a file, with versions, so that developers could set up their environments
automatically?

As Conan is built on Python it can make use of Python's virtualenv to configure an environment
(e.g. setting PATH) to point to binaries that could be included in Conan packages - your linter,
for example.

## What are the challenges with using Conan?

It would not be fair to talk so highly about Conan without also considering some of the
challenges I have observed, so let's cover that.

Conan is not the most simple tool; it has to provide a lot of features and flexibility and nothing in
the C and C++ world is usually plug and play. That said, the Slack thread is available and the
documentation for Conan is pretty good. Plus you can always reach out to us at [Verifa](/contact/) about organising some training.

The second point to mention is that if you are developing lots of small components and need to
make continuous changes to several of these components at the same time before build and
test, then you need to consider the overhead of making each of these components a Conan
package. It can be counterproductive and add overhead.

So the granularity of Conan packages should be considered, which is more of a design or
architectural decision to make than necessarily being a drawback of Conan. Excitingly, this is a
known topic and work is underway to help, such as [Conan's workspace feature](https://docs.conan.io/en/latest/developing_packages/workspaces.html).

## What are the alternatives to Conan?

There are other package managers in the C and C++ space but as we do not have as much
experience with them (added to the fact they are constantly changing), it is not easy to make
any kind of fair comparison.

To briefly mention a few alternatives by name, the most well-known is [vcpkg](https://github.com/Microsoft/vcpkg). Then there are tools like [build2](https://build2.org/) and [Hunter](https://github.com/cpp-pm/hunter), which seem to be less adopted in the market.

We have decided to settle for Conan as our choice of package manager because of the great
support we have received from both the community and from industry tools like Artifactory to
store pre-built binaries. This provides a great option for internal Conan repositories for large
software projects.

## Conan summary

In conclusion, if you are looking to solve some of the common challenges mentioned above I
would strongly consider looking at Conan and running a proof-of-concept.

At Verifa we have helped several companies adopt Conan and in running proof-of-concepts on
customer codebases to get an understanding of its benefits and how it affects your workflow.
If you are interested in learning more, we suggest reading our [What Is Conan? post](https://bincrafters.github.io/2018/07/14/What-Is-Conan/). Alternatively, you can [get in touch with us](/contact/) and we can provide all the necessary services and training to your team.

We look forward to hearing from you!
