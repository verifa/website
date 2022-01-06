<script context="module" lang="ts">
	import { getPost } from '$lib/posts/posts';

	import NotFound from './notFound.svelte';

	export function load({ params }) {
		const post = getPost(params.slug);
		if (post !== null) {
			return {
				props: {
					post: post
				}
			};
		}
		return {
			props: {
				post: NotFound
			}
		};
	}
</script>

<script lang="ts">
	export let post;
</script>

<svelte:component this={post} />
