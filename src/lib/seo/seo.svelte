<script lang="ts">
	import { page } from '$app/stores';

	import OpenGraph from './openGraph.svelte';
	import SchemaOrg from './schemaOrg.svelte';
	import { config, seo } from './store';
	import Twitter from './twitter.svelte';

	const website = {
		entity: 'Verifa',
		ogLanguage: 'en_GB',
		siteLanguage: 'en-GB'
	};

	let url = `${config.siteUrl}${$page.url.pathname}`;
</script>

<svelte:head>
	<title>{$seo.title}</title>
	<meta name="description" content={$seo.description} />
	<meta name="robots" content={`${$seo.robots.noindex ? 'noindex' : 'all'}`} />
	<link rel="canonical" href={url} />
</svelte:head>

<OpenGraph />
<SchemaOrg />
<Twitter />
