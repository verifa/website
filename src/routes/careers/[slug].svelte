<script context="module">
	import NotFound from './notFound.svelte';
	export const load = async ({ params }) => {
		try {
			const post = await import(`../../posts/${params.slug}.md`);
			return {
				props: {
					post: post.default
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
	export let post;
</script>

<svelte:component this={post} />
