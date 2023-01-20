---
type: Blog
title: 'How to automate Artifactory clean up'
subheading: Hassle-free artifact retention policies through FileSpecs
authors:
- tlacour
tags:
- Artifactory
- Continuous Delivery
date: 2023-01-23
image: "/blogs/automate-artifactory-clean-up.png"
featured: true

---

JFrog’s Artifactory is used to store all kinds of artifacts, but its built-in support for deleting unnecessary files is surprisingly lackluster. I say surprisingly because it already has all the features to set up flexible retention policies without much of a hassle. It just hasn’t pieced them together. So let’s do it ourselves!

## Long story short

If you only want the gist of the setup, here’s a TL;DR:

- You use Artifactory’s [FileSpecs](https://www.jfrog.com/confluence/display/JFROG/Using+File+Specs) to *match* artifacts you wish to prune, e.g. *“Everything in foo-dev-local, with 0 downloads, older than 2 weeks”*
- You use the [JFrog CLI](https://jfrog.com/getcli/) to *delete* artifacts matching your FileSpecs, e.g. `jf rt delete --spec policies/foo-dev-local.json`
- All that remains is automating this using your scheduler of choice, be it a simple cron job, a GitHub Actions workflow, etc.

I’ve also written a simple JFrog CLI plugin, [rt-retention](https://rt-retention.verifa.io/), to help run policies in bulk and add some templating capabilities.

## FileSpecs as retention policies

FileSpecs are JSON files containing various properties and rules you use to match artifacts. They can contain [AQL (Artifactory Query Language)](https://www.jfrog.com/confluence/display/JFROG/Artifactory+Query+Language) queries, allowing you to filter artifacts on things such as age, last download, custom property values, etc. All of which are very relevant when finding artifacts you’d like to clean up.

Here’s a simple example that matches all artifacts from a specific repository that haven’t been downloaded in 14 days:

```json
// policies/nuget-dev-local.json
{
  "files": [
    {
      "aql": {
        "items.find": {
          "repo": "nuget-dev-local",
          "stat.downloaded": { "$before": "14d" }
        }
      }
    }
  ]
}
```

The idea is to write out your retention policies as FileSpecs that match the artifacts you want removed. Now, they don’t do anything on their own, that’s where the JFrog CLI comes in.

For more information on using FileSpecs, check the [official documentation](https://www.jfrog.com/confluence/display/JFROG/Using+File+Specs).

## Deleting artifacts using the JFrog CLI

The JFrog CLI is a command line tool used to interact with JFrog’s various products, including Artifactory. The noteworthy command here is `jf rt delete`, which can delete all artifacts that match a given FileSpec.

Matching and deleting happens server-side, which is the main reason you use the JFrog CLI and FileSpecs over the REST API and raw AQL. With the latter, you have to parse the search results yourself and send a `DELETE` request for each match, which is slower and more effort.

Below’s an example run of `jf rt delete`:

```
**$ jf rt delete --quiet --spec policies/pip-dev-local.json**
[Info] Searching artifacts...
[Info] Found 7 artifacts.
[Info] [Thread 2] Deleting pip-dev-local/archon/0.1.0/archon-0.1.0.whl
[Info] [Thread 1] Deleting pip-dev-local/archon/1.0.0/archon-1.0.0.whl
[Info] [Thread 0] Deleting pip-dev-local/archon/1.0.1/archon-1.0.1.whl
[Info] [Thread 1] Deleting pip-dev-local/archon/1.0.2/archon-1.0.2.whl
[Info] [Thread 2] Deleting pip-dev-local/crop/1.4.1/crop-1.4.1.whl
[Info] [Thread 1] Deleting pip-dev-local/crop/1.4.4/crop-1.4.4.whl
[Info] [Thread 0] Deleting pip-dev-local/crop/1.4.3/crop-1.4.3.whl
{
  "status": "success",
  "totals": {
    "success": 7,
    "failure": 0
  }
}
```

Since we define our retention policies as FileSpecs, it’s easy to enforce them by periodically running your `jf rt delete` commands. Preferably, you’d automate it. At my current customer, I keep the FileSpecs in a GitHub repository and have a simple GitHub Actions workflow run them every night.

For more information on interacting with Artifactory through the JFrog CLI, check the [official documentation](https://www.jfrog.com/confluence/display/CLI/CLI+for+JFrog+Artifactory).

## Optional: bulk retention with the rt-retention plugin

If, like me, you’re maintaining a sizeable Artifactory instance with hundreds of repositories, the thought of writing and maintaining hundreds of FileSpecs is not a fun one. 

To deal with that, I’ve written [rt-retention](rt-retention.verifa.io/), a simple JFrog CLI user plugin to make it all a little easier. It has only two features:

- Basic FileSpec templating — Generate FileSpecs from templates rather than maintaining large amounts of similar ones
- Delete artifacts using FileSpecs in a given directory structure — Replace hundreds of JFrog CLI calls with a single command

The templating allowed me to cleverly delegate defining retention policies to my users, without them having to get comfortable with the ins and outs of FileSpecs, while the bulk delete command is a small quality of life feature. 

I won’t dig into the technical details of the plugin here. If you’re interested, you’ll find all the info you need at [rt-retention.verifa.io](http://rt-retention.verifa.io/)

## Putting it all together

To summarise; define your retention policies as FileSpecs, enforce them through the JFrog CLI, and automate the process in whatever scheduler you’re comfortable with.

Preferably also version control your FileSpecs, be careful with your wildcards and dry-run before you start vaporising artifacts.

Good luck tidying up your instance, the sooner you start, the better!