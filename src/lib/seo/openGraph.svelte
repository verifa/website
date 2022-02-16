<script lang="ts">
	import { crewNameById } from '$lib/crew/crew';
	import { seo, config } from './store';
</script>

<svelte:head>
	<meta property="og:site_name" content={config.siteShortTitle} />
	{#if $seo.canonical}
		<meta property="og:url" content={$seo.canonical} />
	{/if}
	<meta property="og:type" content={$seo.article ? 'article' : 'website'} />
	<meta property="og:title" content={$seo.title} />
	<meta property="og:description" content={$seo.description} />
	{#if $seo.image}
		<meta property="og:image" content={`${config.siteUrl}${$seo.image.url}`} />
		<meta property="og:image:alt" content={$seo.image.alt} />
	{/if}
	{#if $seo.moreImages}
		{#each $seo.moreImages as image}
			<meta property="og:image" content={`${config.siteUrl}${image.url}`} />
			<meta property="og:image:alt" content={image.alt} />
		{/each}
	{/if}
	{#if $seo.article}
		{#each $seo.article.authors as author}
			<meta property="article:author" content={crewNameById(author)} />
		{/each}
		{#each $seo.article.tags as tag}
			<meta property="article:tag" content={tag} />
		{/each}
		{#if $seo.article.published_time}
			<meta property="article:published_time" content={$seo.article.published_time.toISOString()} />
		{/if}
		{#if $seo.article.modified_time}
			<meta property="article:modified_time" content={$seo.article.modified_time.toISOString()} />
		{/if}
	{/if}
</svelte:head>
