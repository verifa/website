<script lang="ts">
	import { page } from '$app/stores';

	import hash from 'object-hash';
	import { config, seo } from './store';

	const url = `${config.siteUrl}${$page.url.pathname}`;

	const schemaOrgEntity = {
		'@type': 'Organization',
		'@id': `${config.siteUrl}/#entity`,
		name: config.siteTitle,
		url: config.siteUrl,
		sameAs: [config.sameAs]
	};

	const schemaOrgWebsite = {
		'@type': 'WebSite',
		'@id': `${config.siteUrl}/#website`,
		url: config.siteUrl,
		name: $seo.title,
		description: $seo.description,
		inLanguage: config.siteLanguage
	};

	const schemaOrgWebpage = {
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
			author: $seo.article.authors[0],
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

	let schemaOrgArray = [schemaOrgEntity, schemaOrgWebsite, schemaOrgWebpage];

	if (schemaOrgArticle) {
		schemaOrgArray.push(schemaOrgArticle);
	}

	const jsonLdContents = `
    <script type="application/ld+json">
    ${JSON.stringify({
			'@context': 'http://schema.org',
			'@graph': schemaOrgArray
		})}
    ${'<'}/script>`;
</script>

<svelte:head>
	{@html jsonLdContents}
</svelte:head>
