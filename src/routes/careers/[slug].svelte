<script context="module" lang="ts">
	import { getJobsBySlug } from '$lib/posts/posts';

	import NotFound from './notFound.svelte';

	export function load({ params }) {
		const post = getJobsBySlug(params.slug);
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
