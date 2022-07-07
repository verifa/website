---
type: Blog
title: Keep your pipelines simple, stupid
subheading: Build pipelines like GitHub Actions, GitLab CI/CD and ol’ Jenkins have become ubiquitous these last few years. While they are certainly useful, it’s surprisingly easy to burn your fingers on them. Here’s my three rules to keep pipelines more manageable.
authors:
- tlacour
tags:
- Continuous Integration
- Continuous Delivery
date: 2022-07-07
image: "/blogs/keep-your-pipelines-simple.png"
featured: true

---

Build pipelines like GitHub Actions, GitLab CI/CD and ol’ Jenkins have become ubiquitous these last few years. While they are certainly useful, it’s surprisingly easy to burn your fingers on them. Indulging yourself too much can leave you with a monstrous build that is difficult to maintain or debug, and impossible to run locally or migrate.

After having burnt my own fingers plenty, I’ve come to live by three simple rules that keep my build pipelines more manageable, and after inheriting my fair share of searing hot builds, it’s time I wrote these down and put them out there.

## The three rules of build pipelines

Here’s my three rules to keep my pipelines more manageable and my colleagues sane:

1. Minimise dependencies on plugins and libraries
2. Keep build logic out of the pipeline
3. Limit layers of abstraction

### 1. Minimise dependencies on plugins and libraries

Many platforms offer plugins or tool integrations that make it seductively easy to configure and run your tools for you. However, the flip side of that particular coin is often forgotten. For each plugin or custom action you use, you increase platform lock-in, introduce a new dependency, introduce a new attack vector, and add another build step you can’t run locally.

Do the dirty work, rather than delegating it. 

***Don’t:***

```yaml
- name: build
  uses: gradle/gradle-build-action@latest
  with:
    arguments: build

- name: tag
  uses: random-person-in-nebraska/useful-tagging-plugin@latest
  with:
    tag: $VERSION
```

***Do:***

```yaml
- name: build
	run: ./gradlew build

- name: tag
	run: |
		VERSION=$(git describe)
		git tag -a $VERSION -m "Tagged by build pipeline"
    git push origin $VERSION
```

### 2. Keep logic out of the pipeline

From invoking tools to actual scripting, it’s easy to jam it all into your pipeline, but this makes it difficult to run locally. Move such logic to build scripts and have your pipelines call those instead. This allows you to run and debug your build steps locally, and keeps your pipelines focused on the build *flow*. 

Keep your pipeline stupid. The less it does, the better.

***Don’t:***

```yaml
- name: test
	env:
		SECRET: ${{ secrets.UPLOAD_KEY }}
		MODULES_FOR_TEST: ${{ steps.check.outputs.* }}
	run: |
    for mod in $MODULES_FOR_TEST; 
		do
			pytest modules/$mod/* -v --junitxml="$mod.xml"
			./upload.py -name="$mod" -key="$SECRET" -file="$mod.xml"
		done
```

***Do:***

```yaml
- name: test
	env:
		SECRET: ${{ secrets.UPLOAD_KEY }}
	run: |
		./scripts/test.sh --upload-key "$SECRET"
```

### 3. Limit layers of abstraction

There is value in encapsulating build steps in a script or build tool of sorts. However, adding tools is easier than removing them. Think twice before introducing yet another tool or abstraction layer. Keeping your build low to the ground makes it easier to understand, troubleshoot and modify.

No one wants to inherit a build that runs Make to call a bash script to spin up a Docker container which runs a Python script which…

***Don’t:***

```bash
$ ls build/
build.sh  call-test-scripts.sh  conanfile.py
config.yml Dockerfile  env/  __init__.py  modules/
run_build.py  TestReport.rb  test-scripts/
```

***Do:***

```bash
$ ls build/
Dockerfile Makefile README.md
```

*Note: The author does not imply you should only use Docker/Make for everything, they are merely exemplifying flattening your build’s tool stack.*

## Summary

Your build pipeline is part of your project. Try not to cut corners and be lazy with it, but keep it simple and down-to-earth. You’ll thank yourself in the long run, as will I when I don’t inherit yet another 20k line YAML behemoth.