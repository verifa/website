[![Contributors][contributors-shield]][contributors-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![Apache 2.0][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

# Verifa Website

This repository contains Verifa's website live at <https://verifa.io>

## Built with

### Development

* [Go](https://go.dev/)
* [Templ](https://templ.guide/)
* [Tailwind](https://tailwindcss.com/)

### Deployment

* [![google-cloud][google-cloud-shield]][google-cloud-url]
* [![github-actions][github-actions-shield]][github-actions-url]

### Other notable mentions

1. [giscus](https://giscus.app/) - adding comments to our blog using GitHub issues

## Design

The website, logos and brand were designed by [The Pizzolorusso Design Agency](https://pizzolorusso.com/about).
Highly recommended :)

## Developing

Things you need to install:

1. [Go](https://go.dev/), check the [go.mod](./go.mod) for the required version
2. [NodeJS](https://nodejs.org/en/download/package-manager)

To start a local development server:

```bash
make dev

# Open your browser at http://localhost:3000
```

## Building

Before pushing any code changes, make sure that it builds

```bash
make pr
```

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
[google-cloud-shield]: https://img.shields.io/badge/Google_Cloud-4285F4?style=for-the-badge&logo=google-cloud&logoColor=white
[google-cloud-url]: https://cloud.google.com/
[github-actions-shield]: https://img.shields.io/badge/GitHub_Actions-2088FF?style=for-the-badge&logo=github-actions&logoColor=white
[github-actions-url]: https://github.com/features/actions
