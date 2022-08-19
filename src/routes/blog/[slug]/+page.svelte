<script lang="ts">
	import type { PageData } from './$types';
	export let data: PageData;
	$: ({ post, relatedPosts } = data);
</script>

<svelte:component this={post} {relatedPosts} />
