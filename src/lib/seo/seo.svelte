<script lang="ts">
	import OpenGraph from './openGraph.svelte';
	import SchemaOrg from './schemaOrg.svelte';
	import { config, seo } from './store';
	import Twitter from './twitter.svelte';

	const website = {
		entity: 'Verifa',
		ogLanguage: 'en_GB',
		siteLanguage: 'en-GB'
	};
</script>

<svelte:head>
	<title>{$seo.title}</title>
	<meta name="description" content={$seo.description} />
	<meta name="robots" content={`${$seo.robots.noindex ? 'noindex' : 'all'}`} />
	{#if $seo.canonical}
		<link rel="canonical" href={$seo.canonical} />
	{/if}
</svelte:head>

<OpenGraph />
<SchemaOrg />
<Twitter />
