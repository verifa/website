<script context="module" lang="ts">
	export const load = async ({ params, fetch }) => {
		try {
			const post = await import(`../../posts/${params.slug}.md`);
			const { title, tags } = post.metadata;
			const relatedBlogs = await getRelatedBlogs(fetch, title, tags);
			return {
				props: {
					post: post.default,
					relatedBlogs: relatedBlogs.blogs
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
	import { getRelatedBlogs, type Post } from '$lib/posts/posts';
	import NotFound from '../careers/notFound.svelte';

	export let post = null;
	export let relatedBlogs: Post[];
</script>

<svelte:component this={post} {relatedBlogs} />
