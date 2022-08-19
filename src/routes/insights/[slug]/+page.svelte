<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { seo } from '$lib/seo/store';
	import { onMount } from 'svelte';
	const { slug } = $page.params;

	onMount(() => {
		setTimeout(() => goto(`/blog/${slug}`), 1000 * 2);
	});

	seo.reset();
</script>

<section>
	<h1>New webiste. New routes. Redirecting...</h1>
</section>
