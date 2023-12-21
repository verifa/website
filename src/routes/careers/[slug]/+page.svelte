<script lang="ts">
	import LayoutJob from '$lib/posts/layout-job.svelte';
	import type { PageData } from './$types';
	export let data: PageData;
	$: ({ post, component } = data);
</script>

<LayoutJob {post}>
	<svelte:component this={component} />
</LayoutJob>
