<script context="module" lang="ts">
	export const load = async ({ params, fetch }) => {
		const member = crewByID(params.slug);
		if (!member) {
			return {
				status: 404,
				props: {
					member: null
				}
			};
		}
		const postsUrl = `/posts/blogs-by-${member.id}.json`;
		const res = await fetch(postsUrl);
		if (res.ok) {
			return {
				props: {
					member: member,
					blogs: await res.json()
				}
			};
		}
		return {
			status: res.status,
			error: new Error(`could not load ${postsUrl}`)
		};
	};
</script>

<script lang="ts">
	import { crewByID, type Member } from '$lib/crew/crew';
	import PostGrid from '$lib/posts/postGrid.svelte';
	import type { Blogs } from '$lib/posts/posts';
	import { seo } from '$lib/seo/store';

	export let member: Member;
	export let blogs: Blogs;

	seo.reset();
</script>

{#if member.active}
	<section>
		<div class="flex flex-col gap-y-8 sm:flex-row sm:gap-x-4">
			<div class="">
				<img src={member.image} alt={member.id} class="h-48" />
			</div>
			<div>
				<h2>{member.name}</h2>
				<h4>{member.position}</h4>
			</div>
		</div>
	</section>
{:else}
	<section>
		<div class="flex flex-col gap-y-8 sm:flex-row sm:gap-x-4">
			<div class="">
				<img src={member.image} alt={member.id} class="h-48" />
			</div>
			<div>
				<h2>{member.name}</h2>
			</div>
		</div>
	</section>
{/if}
<section>
	{#if blogs.blogs.length == 0}
		<h2>No posts by author</h2>
	{:else}
		<h2>Posts by author</h2>
		<PostGrid posts={blogs.blogs} />
	{/if}
</section>
