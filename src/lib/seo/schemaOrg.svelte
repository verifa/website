<script lang="ts">
	import { page } from '$app/stores';
	import { crewNameById } from '$lib/crew/crew';
	import { config, seo } from './store';

	let schemaOrgArray = [];
	// Make sure we regenerate the schema org for each page
	page.subscribe(() => {
		let url = `${config.siteUrl}${$page.url.pathname}`;

		let schemaOrgEntity = {
			'@type': 'Organization',
			'@id': `${config.siteUrl}/#entity`,
			name: config.siteTitle,
			url: config.siteUrl,
			sameAs: [config.sameAs]
		};

		let schemaOrgWebsite = {
			'@type': 'WebSite',
			'@id': `${config.siteUrl}/#website`,
			url: config.siteUrl,
			name: $seo.title,
			description: $seo.description,
			inLanguage: config.siteLanguage
		};

		let schemaOrgWebpage = {
			'@type': 'WebPage',
			'@id': `${config.siteUrl}/#webpage`,
			isPartOf: {
				'@id': `${config.siteUrl}/#website`
			},
			url: url,
			name: $seo.title,
			description: $seo.description,
			inLanguage: config.siteLanguage,
			potentialAction: [
				{
					'@type': 'ReadAction',
					target: [url]
				}
			]
		};

		let schemaOrgArticle = null;

		if ($seo.article) {
			schemaOrgArticle = {
				'@type': 'Article',
				'@id': `${url}#article`,
				isPartOf: {
					'@id': `${url}/#webpage`
				},
				author: crewNameById($seo.article.authors[0]),
				headline: $seo.title,
				datePublished: $seo.article.published_time.toISOString(),
				dateModified: $seo.article.modified_time.toISOString(),
				mainEntityOfPage: {
					'@id': `${url}#webpage`
				},
				publisher: config.entity,
				articleSection: ['blog'],
				inLanguage: config.siteLanguage
			};
		}

		schemaOrgArray = [schemaOrgEntity, schemaOrgWebsite, schemaOrgWebpage];

		if (schemaOrgArticle !== null) {
			schemaOrgArray.push(schemaOrgArticle);
		}
	});
</script>

<svelte:head>
	{@html `
        <script type="application/ld+json">
        ${JSON.stringify({
					'@context': 'http://schema.org',
					'@graph': schemaOrgArray
				})}
        ${'<'}/script>`}
</svelte:head>
