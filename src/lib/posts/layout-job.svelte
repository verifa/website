<script lang="ts">
	import './posts.css';

	import CareersForm from '$lib/careers/careersForm.svelte';

	import { crewNameById } from '$lib/crew/crew';

	import { seo } from '$lib/seo/store';

	import type { Post } from './posts';
	import PostBadges from './postBadges.svelte';

	export let post: Post;

	seo.reset();
	$seo.title = post.title;
	$seo.description = post.subheading;
	$seo.image.url = post.image;
	$seo.article = {
		authors: post.authors,
		tags: post.tags,
		published_time: new Date(post.date),
		modified_time: new Date(post.date)
	};
</script>

<svelte:head>
	<link rel="stylesheet" href="/prism-ghcolors.css" />
</svelte:head>

<div class="max-w-5xl mx-auto">
	<article class="prose lg:prose-lg xl:prose-xl 2xl:prose-2xl">
		<div class="mb-8">
			<div class="mb-8">
				<img class="w-full h-full" src={post.image} alt={post.title} />
			</div>
			<h1>{post.title}</h1>
			<h3>{post.subheading}</h3>
			<div class="not-prose mb-8">
				<PostBadges type={post.type} tags={post.tags} />
			</div>
			<p class="mb-4">Published on {new Date(post.date).toDateString()}</p>
			<p class="mb-4">Authors: {post.authors.map((a) => crewNameById(a)).join(', ')}</p>
		</div>
		<div id="blog-container">
			<slot />
		</div>
	</article>
	<section>
		<h2>Apply now!</h2>
		<CareersForm />
	</section>
</div>
