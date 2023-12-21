<script lang="ts">
	import Layout from '$lib/posts/layout.svelte';
	import type { PageData } from './$types';
	export let data: PageData;
	$: ({ post, component, relatedPosts } = data);
</script>

<Layout {post} {relatedPosts}>
	<svelte:component this={component} {relatedPosts} />
</Layout>
