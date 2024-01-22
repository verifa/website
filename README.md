[![Contributors][contributors-shield]][contributors-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![Apache 2.0][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

# Verifa Website

This repository contains Verifa's website live at <https://verifa.io>

## Built with

### Development

* [![svelte][svelte-shield]][svelte-url]
* [![sveltekit][sveltekit-shield]][sveltekit-url]
* [![tailwind][tailwind-shield]][tailwind-url]
* [![typescript][typescript-shield]][typescript-url]

### Deployment

* [![google-cloud][google-cloud-shield]][google-cloud-url]
* [![github-actions][github-actions-shield]][github-actions-url]

### Other notable mentions

1. [mdsvex](https://mdsvex.pngwn.io/) - for converting Markdown into pages
2. [giscus](https://giscus.app/) - adding comments to our blog using GitHub discussions

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

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/verifa/website.svg?style=for-the-badge
[contributors-url]: https://github.com/verifa/website/graphs/contributors
[stars-shield]: https://img.shields.io/github/stars/verifa/website.svg?style=for-the-badge
[stars-url]: https://github.com/verifa/website/stargazers
[issues-shield]: https://img.shields.io/github/issues/verifa/website.svg?style=for-the-badge
[issues-url]: https://github.com/verifa/website/issues
[license-shield]: https://img.shields.io/github/license/verifa/website.svg?style=for-the-badge
[license-url]: https://github.com/verifa/website/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://www.linkedin.com/company/verifa
<!-- STACK -->
[svelte-shield]: https://img.shields.io/badge/Svelte-4A4A55?style=for-the-badge&logo=svelte&logoColor=FF3E00
[svelte-url]: https://svelte.dev
[sveltekit-shield]: https://img.shields.io/badge/SvelteKit-FF3E00?style=for-the-badge&logo=Svelte&logoColor=white
[sveltekit-url]: https://kit.svelte.dev/
[tailwind-shield]: https://img.shields.io/badge/Tailwind_CSS-38B2AC?style=for-the-badge&logo=tailwind-css&logoColor=white
[tailwind-url]: https://tailwindcss.com/
[typescript-shield]: https://img.shields.io/badge/TypeScript-007ACC?style=for-the-badge&logo=typescript&logoColor=white
[typescript-url]: https://www.typescriptlang.org/
[google-cloud-shield]: https://img.shields.io/badge/Google_Cloud-4285F4?style=for-the-badge&logo=google-cloud&logoColor=white
[google-cloud-url]: https://cloud.google.com/
[github-actions-shield]: https://img.shields.io/badge/GitHub_Actions-2088FF?style=for-the-badge&logo=github-actions&logoColor=white
[github-actions-url]: https://github.com/features/actions
