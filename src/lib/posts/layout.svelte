<script>
	import './posts.css';

	import { crewByID } from '$lib/crew/crew';

	import { seo } from '$lib/seo/store';

	import PostBadges from './postBadges.svelte';
	import PostGrid from './postGrid.svelte';
	import { PostType } from './posts';
	import  * as readtime from './readtime';

	export let relatedPosts = [];

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

	let container;
	let readtimeDisplay;
	$: if (container && readtimeDisplay) {
		const result = readtime.processText(container.innerText);
		readtimeDisplay.innerText = "Read time: " + result.humanizedTime;
	}
</script>

<svelte:head>
	<link rel="stylesheet" href="/prism-ghcolors.css" />
</svelte:head>

<div class="max-w-5xl mx-auto">
	<article class="prose lg:prose-lg xl:prose-xl 2xl:prose-2xl">
		<div class="mb-8">
			<div class="mb-8">
				<img class="w-full h-full" src={image} alt={title} />
			</div>
			<h1>{title}</h1>
			<h3>{subheading}</h3>
			{#if type !== PostType.Event}
				<div class="not-prose mb-8">
					<PostBadges {type} {tags} />
				</div>
				<p class="mb-4">Published on {new Date(date).toDateString()}</p>
				<p bind:this={readtimeDisplay} class="mb-4">Read time: </p>
				<h4>Authors</h4>
				<div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-y-2 sm:gap-y-8">
					{#each crewAuthors as author}
						<a href="/crew/{author.id}" class="group">
							<div class="flex gap-x-4 items-center">
								<img src={author.avatar} alt={author.id} class="h-16 w-16 m-0" />
								<div>
									<h5 class="mb-0 group-hover:text-v-lilac">{author.name}</h5>
								</div>
							</div>
						</a>
					{/each}
				</div>
			{/if}
		</div>
		<div id="blog-container" bind:this={container}>
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
		</section>
		<section>
			<PostGrid posts={relatedPosts} showBadges={true} />
		</section>
	{/if}
</div>
