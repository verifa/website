<script>
	import { crewByID } from '$lib/crew/crew';

	import { seo } from '$lib/seo/store';

	import PostBadges from './postBadges.svelte';
	import PostGrid from './postGrid.svelte';
	import { PostType } from './posts';

	export let relatedBlogs = [];

	export let type;
	export let title;
	export let subheading;
	export let authors;
	export let tags;
	export let date;
	export let image;

	const crewAuthors = authors.map((id) => crewByID(id));

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
	<article>
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
			<h4>Authors</h4>
			<div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-y-4 sm:gap-y-8">
				{#each crewAuthors as author}
					<a href="/crew/{author.id}" class="group">
						<div class="flex gap-x-4 items-center">
							<img src={author.image} alt={author.id} class="h-16 w-16" />
							<div>
								<h5 class="mb-0 group-hover:text-v-lilac">{author.name}</h5>
							</div>
						</div>
					</a>
				{/each}
			</div>
		</div>
		<div id="blog-container">
			<slot />
		</div>
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
	{#if type != PostType.Case}
		<section>
			<h2>Read similar posts</h2>
			<PostGrid posts={relatedBlogs} showBadges={false} />
		</section>
	{/if}
</div>
