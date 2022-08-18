<script context="module" lang="ts">
	export const load = async ({ params, fetch }) => {
		try {
			const post = await import(`../../posts/${params.slug}.md`);
			try {
				const { title, tags } = post.metadata;
				const res = await fetch(
					'/posts/posts.json?' +
						new URLSearchParams({
							limit: '3',
							skip_title: title,
							keywords: tags.join(',')
						})
				);

				if (res.ok) {
					const data: PostsData = await res.json();
					return {
						props: {
							post: post.default,
							relatedPosts: data.posts
						}
					};
				} else {
					const error = await res.text();
					return {
						status: res.status,
						error: new Error(error)
					};
				}
			} catch (error) {
				return {
					status: 500,
					error: error
				};
			}
		} catch (error) {
			return {
				props: {
					post: NotFound
				}
			};
		}
	};
</script>

<script lang="ts">
	import type { Post, PostsData } from '$lib/posts/posts';
	import NotFound from './_notFound.svelte';

	export let post = null;
	export let relatedPosts: Post[];
</script>

<svelte:component this={post} {relatedPosts} />
