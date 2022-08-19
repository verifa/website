<script lang="ts">
	import type { PageData } from './$types';
	export let data: PageData;
	$: ({ post } = data);
</script>

<svelte:component this={post} />
