<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	const { slug } = $page.params;

	onMount(() => {
		setTimeout(() => goto(`/blog/${slug}`), 1000 * 2);
	});
</script>

<section>
	<h1>New webiste. New routes. Redirecting...</h1>
</section>
