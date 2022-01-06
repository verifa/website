<script lang="ts">
	import { seo, config } from './store';
</script>

<svelte:head>
	<meta name="twitter:card" content="summary_large_image" />
	<meta name="twitter:title" content={$seo.title} />
	<meta name="twitter:description" content={$seo.description} />
	<meta name="twitter:site" content="@{config.twitterUsername}" />
	<meta name="twitter:creator" content={`@${config.twitterUsername}`} />
	{#if $seo.image}
		<meta name="twitter:image" content={$seo.image.url} />
		<meta name="twitter:image:alt" content={$seo.image.alt} />
	{/if}
</svelte:head>
