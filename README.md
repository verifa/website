# Verifa Website

This repository contains Verifa's website live at <https://verifa.io>

## Stack

This project uses [Svelte](https://svelte.dev/), [SvelteKit](https://kit.svelte.dev/) and [TailwindCSS](https://tailwindcss.com/).
GitHub Actions are used to build and deploy the website to a Google Cloud bucket, created with Terraform and hosted by Verifa.

### Other notable mentions

1. [mdsvex](https://mdsvex.pngwn.io/) - for converting Markdown into pages
2. [utterances](https://utteranc.es/) - adding comments to our blog using GitHub issues

## Design

The website, logos and brand were designed by [The Pizzolorusso Design Agency](https://pizzolorusso.com/about).
Highly recommended :)

## Developing

To start a local development server:

```bash
npm run dev

# or start the server and open the app in a new browser tab
npm run dev -- --open
```

## Building

Before pushing any code changes, make sure that it builds

```bash
npm run build
```

> You can preview the built app with `npm run preview`, regardless of whether you installed an adapter. This should _not_ be used to serve your app in production.

## Releasing

Pushing to the `main` branch deploys the changes to the staging environment <https://staging.verifa.io>

Tagging a version makes deploys it to the production environment <https://verifa.io>

## Contributing

Contributions and PRs welcome! More to follow on this topic.
