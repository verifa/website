<script>
	import { crewNameById } from '$lib/crew/crew';

	import { seo } from '$lib/seo/store';

	import PostBadges from './postBadges.svelte';
	import PostGrid from './postGrid.svelte';

	export let relatedBlogs = [];

	export let type;
	export let title;
	export let subheading;
	export let authors;
	export let tags;
	export let date;
	export let image;

	seo.reset();
	$seo.title = title;
	$seo.description = subheading;
	$seo.image.url = image;
	$seo.article = {
		authors: authors,
		tags: tags,
		published_time: new Date(date),
		modified_time: new Date(date)
	};
</script>

<svelte:head>
	<link rel="stylesheet" href="/prism-ghcolors.css" />
</svelte:head>

<div class="max-w-5xl mx-auto">
	<article id="blog-container">
		<div class="mb-8">
			<div class="mb-8">
				<img src={image} alt={title} />
			</div>
			<h2>{title}</h2>
			<h4>{subheading}</h4>
			<div class="mb-8">
				<PostBadges {type} {tags} />
			</div>
			<p class="mb-4">Published on {new Date(date).toDateString()}</p>
			<p class="mb-4">Authors: {authors.map((a) => crewNameById(a)).join(', ')}</p>
		</div>
		<slot />
	</article>
	<hr />
	<section>
		<h2>Comments</h2>
		<script
			src="https://utteranc.es/client.js"
			repo="verifa/website"
			issue-term="pathname"
			label="blog"
			theme="boxy-light"
			crossorigin="anonymous"
			async>
		</script>
	</section>
	{#if type != 'Job'}
		<section>
			<h2>Read similar posts</h2>
			<PostGrid posts={relatedBlogs} showBadges={false} />
		</section>
	{/if}
</div>
