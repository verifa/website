# How to write blogs

1. Create a new file, or copy an existing post, under [src/posts](../src/posts/)
2. Be sure add the required [frontmatter](https://jekyllrb.com/docs/front-matter/):

    - `type`: one of `Blog`, `Podcast`, `Case`[study], `Job`, `Event`
    - `title`: title of the blog
    - `subheading`: subtitle of the blog
    - `authors`: string list of author IDs (e.g. `jlarfors`, your Verifa email prefix)
    - `tags`: string list of tags (do not create any new ones), check list here: <https://verifa.io/blog/index.html>
    - `date`: date the blog was written/published
    - `image`: display image for the blog
    - `featured`: whether it should be featured in the top/latest blog posts

3. Write the content and create a Pull Request

## Beyond Markdown

Our blogs are written in Markdown, but sometimes you want a little more than Markdown.

We use [MDSvex](https://mdsvex.pngwn.io/) to compile our Markdown files to Svelte.
We can also just embed Svelte into them. Which we have done.

### Figure captions

Example:

```html
<figure>
  <img src="/static/blog/nodeless-aws-eks-clusters-with-karpenter/karpenter-how-it-works-diagram.png" alt="karpenter-how-it-works">
  <figcaption>Source: https://aws.amazon.com/static/blog/aws/introducing-karpenter-an-open-source-high-performance-kubernetes-cluster-autoscaler/</figcaption>
</figure>
```

### Citation

Example:

```md
> Love the problem, not the solution. There is no endgame with Continuous Delivery and time is finite.
>
> <footer>Jacob Lärfors, Verifa</footer>
```

### Admonitions

We follow GitHub's markdown format: <https://github.com/orgs/community/discussions/16925>

For example:

```md
> [!NOTE]
> Highlights information that users should take into account, even when skimming.

> [!TIP]
> Optional information to help a user be more successful.

> [!IMPORTANT]
> Crucial information necessary for users to succeed.

> [!WARNING]
> Critical content demanding immediate user attention due to potential risks.

> [!CAUTION]
> Negative potential consequences of an action.
```
