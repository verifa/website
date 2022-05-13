<script context="module" lang="ts">
	export async function load({ fetch }) {
		const postsUrl = `/posts/blogs.json`;
		const res = await fetch(postsUrl);
		if (res.ok) {
			return {
				props: {
					data: await res.json()
				}
			};
		}

		return {
			status: res.status,
			error: new Error(`Could not load ${postsUrl}`)
		};
	}
</script>

<script lang="ts">
	import { page } from '$app/stores';
	import type { Blogs, Post } from '$lib/posts/posts';
	import PostGrid from '$lib/posts/postGrid.svelte';
	import { writable } from 'svelte/store';
	import { onMount } from 'svelte';
	import { seo } from '$lib/seo/store';

	export let data: Blogs;

	seo.reset();
	$seo.title = 'Verifa blog and news';
	$seo.description =
		'We write about all the great things happening in Cloud, DevOps, Continuous Delivery and our culture';

	const allBlogs: Post[] = data.blogs;
	const allKeywords: string[] = data.keywords;
	let showBlogs: Post[] = allBlogs;

	const selectedKeywords = writable<boolean[]>(Array(allKeywords.length).fill(false));
	onMount(() => {
		// Get the query params and create initial search
		const paramKeywords = $page.url.searchParams.get('keywords') || '';

		paramKeywords.split(',').forEach((kw) => {
			const kwIndex = allKeywords.indexOf(kw);
			if (kwIndex != -1) {
				$selectedKeywords[kwIndex] = true;
			}
		});
	});

	// filterBlogs runs the filtering of blogs based on the selected keywords
	const filterBlogs = (keywords: string[]): Post[] => {
		if (keywords.length == 0) {
			return allBlogs;
		}
		// If there are keywords, filter the blogs
		return allBlogs.filter((blog) => {
			for (let index = 0; index < keywords.length; index++) {
				const keyword = keywords[index];
				if (!blog.tags.includes(keyword)) {
					return false;
				}
			}
			return true;
		});
	};

	// Subscribe to the filter store to refresh list of blog posts
	// when the query params change
	selectedKeywords.subscribe((selectedKeywords) => {
		let keywords: string[] = [];
		selectedKeywords.forEach((kw, index) => {
			if (kw) {
				keywords.push(allKeywords[index]);
			}
		});
		// Update the list of blogs
		showBlogs = filterBlogs(keywords);
	});
</script>

<section>
	<div class="flex space-x-4">
		<div>
			<img class="h-16 w-16 object-contain" src="/logo-element.png" alt="logo-ement" />
		</div>
		<div>
			<h4>Our blog and news.</h4>
			<h1>Out and about.</h1>
		</div>
	</div>
	<div>
		<h3>Filter by keyword</h3>
		<div class="-my-2 flex flex-wrap gap-x-4">
			{#each allKeywords as keyword, index}
				<!-- 
					If keyword is selected show a different button.
					Simpler than lots of conditional logic to style the elements.
					Sometimes copy+pasta is better, maybe.
				 -->
				{#if $selectedKeywords[index]}
					<button
						class="flex items-center gap-x-1 my-2 border-0 px-3 py-0.5 bg-v-pink hover:bg-v-pink focus:bg-v-pink"
						on:click={() => {
							$selectedKeywords[index] = !$selectedKeywords[index];
						}}
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							class="h-6 w-6"
							fill="none"
							viewBox="0 0 24 24"
							stroke="currentColor"
						>
							<path
								stroke-linecap="round"
								stroke-linejoin="round"
								stroke-width="3"
								d="M5 13l4 4L19 7"
							/>
						</svg>
						<p class="m-0 text-v-white font-semibold border-v-lilac">{keyword}</p>
					</button>
				{:else}
					<button
						class="inline-block my-2 border-0 px-3 py-0.5 bg-v-gray hover:bg-v-gray focus:bg-v-gray"
						on:click={() => {
							$selectedKeywords[index] = !$selectedKeywords[index];
						}}
					>
						<p class="m-0 text-v-white">{keyword}</p>
					</button>
				{/if}
			{/each}
		</div>
	</div>
</section>
<section>
	<PostGrid posts={showBlogs} />
</section>
