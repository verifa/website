<script context="module" lang="ts">
	export const load = async ({ params, fetch }) => {
		try {
			const post = await import(`../../posts/cases/${params.slug}.md`);
			const { title, tags } = post.metadata;
			return {
				props: {
					post: post.default,
				}
			};
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
	import NotFound from './_notFound.svelte';

	export let post = null;
</script>

<svelte:component this={post} />
