---
type: Blog
title: Bag of Git tricks
subheading: Neat git tricks in bite-sized chunks.
authors:
- tlacour
tags:
- Git
date: 2023-05-15
image: "/blogs/bag-of-git-tricks/bag-of-git-tricks.png"
featured: true
---

<script>
    import Admonition from '$lib/posts/admonition.svelte'
</script>

---

Git has a lot of bells and whistles, some of which I find myself recommending to people often.
To save myself some time, I compiled a list of some of my favourite commands and options.
Take a look around, you might find something useful.

Content:

- Aliases (git config)
- Patch mode (git add)
- Interactive mode (git rebase)
- git reflog]
- Searching with git log
- git bisect


## Aliases (git config)

[Documentation](https://git-scm.com/docs/git-config#Documentation/git-config.txt-alias)

You can configure aliases for commands in git, allowing you to do more with fewer keystrokes:

`git config alias.<shortname> "<command>"`

In the following example, I configure `st` to be an alias for `status.`

```
$ git config --global alias.st "status"

$ git st
On branch main
Your branch is up to date with 'origin/main'.

nothing to commit, working tree clean
```

My personal favourite is my alias for a pretty log:

`git config --global alias.graph "log --graph --pretty=tformat:'%C(bold blue)%h%Creset %s %C(bold green)%d%Creset %C(blue)<%an>%Creset %C(dim cyan)%cr' --abbrev-commit --decorate"`

Netting me the following with little effort:

![bag-of-git-tricks-example](/blogs/bag-of-git-tricks/bag-of-git-tricks-example.png)

## Patch mode (git add)

[Documentation](https://git-scm.com/docs/git-add#Documentation/git-add.txt--p)

Patch mode allows you to stage changes interactively, allowing for fine-grained control of what goes into your commit.

`git add -p <paths>`

In the following example, I have two changes. I stage the former, but choose not to stage the latter.

```
$ git add -p .
diff --git a/README.md b/README.md
index a5af44b..fbce946 100644
--- a/README.md
+++ b/README.md
@@ -57,7 +57,7 @@

-Tagging a version makes deploys it to the production environment <https://verifa.io>
+Tagging a version deploys it to the production environment <https://verifa.io>
 
(1/2) Stage this hunk [y,n,q,a,d,e,?]? y

diff --git a/README.md b/README.md
index a5af44b..fbce946 100644
--- a/README.md
+++ b/README.md
@@ -104,6 +104,8 @@ 

+    TODO: add a section on local testing here
+

(2/2) Stage this hunk [y,n,q,a,d,e,?]? n
```

I love patch mode for two reasons; it allows me to stage on a finer granularity than all changes in a file, and it’s great for reviewing what I’m actually committing. It allows me to craft clean atomic commits, and keeps those debug printouts I forgot to clean up out of my commits.

## Interactive mode (git rebase)

[Documentation](https://git-scm.com/docs/git-rebase#Documentation/git-rebase.txt--i)

Interactive rebase is a powerful tool for rewriting local history.

`git rebase -i <target>`

In the following example, I rebase my `my-feature` branch onto `main`.

```
$ git switch my-feature
$ git rebase -i main
```

My text editor pops up, asking me what I wish to do with the commits. Here’s what I asked git to do for me:

```
# Argh, typos! Please let me rewrite the commit message
reword aa82449 Updaet order interfce
# Hmm, let's just add these small changes to the previous commit
fixup  a32c293 small tweak to interface change
# Uhh, I don't need this commit anymore, drop it.
drop   fc560aa add some debug printout
# This commit's fine as is.
pick   a3c6428 Update Order interface documentation
```

After closing my editor, git does what I specified, netting me a clean history:

```
* 0b4c83e Update Order interface documentation (HEAD -> my-feature)
* d3f750e Update Order interface
* d3c7e07 Fix off-by-one in item paging (main, tag: v0.2.4)
```

I find interactive rebase to be an amazing tool for tidying up local changes before pushing them upstream. It’s without a doubt one of my favourite tools in the git arsenal.

## git reflog

[Documentation](https://git-scm.com/docs/git-reflog)

Reflog lists the local history and movements of a ref (branch/HEAD/etc.)

`git reflog <ref>`

In the following example, you see the short history of where my HEAD’s been, including the commands that moved it:shows you

```
$ git reflog HEAD
4c89d17 (HEAD) HEAD@{0}: commit: Add Service and Project to Request OpenAPI schema
34fd162 HEAD@{1}: commit: Move logout to NavBar
23e58c7 HEAD@{2}: commit: Vaporize dead code
e50695b HEAD@{3}: pull -r: Fast-forward
cc4c5fd HEAD@{4}: commit: Started work on request form generation
d6e9edb HEAD@{5}: pull -r: Fast-forward
888fcb5 HEAD@{6}: commit: Donut session - Start working on request form
fe65c7e HEAD@{7}: clone: from github.com:verifa/coastline
```

`reflog` is an invaluable tool when troubleshooting a colleague or student’s git issues. It lets me see what they actually did to get tangled up, rather than what they remember they did.

## Searching with git log

[Documentation](https://git-scm.com/docs/git-log#Documentation/git-log.txt--Sltstringgt)

Git log has some powerful search options, helping you track down changes.
`git log -S <pattern>`

`git log -G <regex>`

In the example below, I look for any commits with changes including `filterPosts`, netting me three commits:

```
$ git log --oneline -S filterPosts
75b9ce3 Upgrade SvelteKit with major design overhaul
79169d2 Cleanup of posts logic and fix missing related blogs
207461e Create single endpoint for all post types
```

`git log -G` functions the same way, but matches regular expressions instead of a plain string. Note that these are just options to `log`, you can combine them with others as you wish.

## git bisect

[Documentation](https://git-scm.com/docs/git-bisect)

Git bisect facilitates searching through your history for the commit that introduced a bug.
`git bisect <subcommand> <options>`
When using bisect, you mark two commits as start- and endpoints. Git will then perform a binary search through your history, checking out commits as it goes. All you do is check if the bug is present or not, then mark the commit as `good` or `bad`, until git pinpoints the commit that introduced the problem.

In the below example, I supply `git bisect run` with a script that tests if the bug is present. This automates the testing process, quickly finding me the culprit commit:

```
$ git bisect start      # start a new bisect run
$ git bisect bad HEAD   # current commit is BAD
$ git bisect good 0.1.0 # bug wasn't present in 0.1.0, it was GOOD
Bisecting: 8 revisions left to test after this (roughly 3 steps)
[7423a9476e708d766c8373c434951a0d74b76810] Added image caption checking

$ git bisect run ../test-filters.sh
running ../test-filters.sh
Bisecting: 4 revisions left to test after this (roughly 2 steps)
[ab5d5a375ca8edf3d24f2c1be6ca36990feb0cb3] Small fix to post filtering
running ../test-filters.sh
Bisecting: 1 revision left to test after this (roughly 1 step)
[c4405540d56f16699d32baf791ede0ab3d23a1b8] Update K8S + Vault blog
running ../test-filters.sh
Bisecting: 0 revisions left to test after this (roughly 0 steps)
[125f83d85742c93d2b34f3574ccdefdfc5c612f6] Update main page styling
running ../test-filters.sh
db5d5a375ca8edf3d24f2c1be6ca36990feb0cb3 is the first bad commit

commit ab5d5a375ca8edf3d24f2c1be6ca36990feb0cb3
Author: grinch bugsly <gbugsly@verifa.io>
Date:   Wed Nov 20 09:58:45 2022 +0200

    Improve sorting after post filtering

 src/lib/posts.ts | 4 +-
 1 file changed, 2 insertion(+), 2 deletion(-)
```

I absolutely love `bisect`, but almost never get a chance to use it. Still, it’s a nice tool to have in your back pocket .

## Far from the end

Git has a truly enormous amount of interesting features, and while the bells and whistles listed above might be some of my favourites, there’s plenty more that went unmentioned.

Why not run `git help -a` and explore some commands yourself?