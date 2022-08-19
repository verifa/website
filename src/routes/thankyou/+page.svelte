<script lang="ts">
	import { goto } from '$app/navigation';
	import { seo } from '$lib/seo/store';
	import { onMount } from 'svelte';

	onMount(() => {
		setTimeout(() => goto('/'), 1000 * 3);
	});

	seo.reset();
	$seo.title = 'Thank you for contacting us';
</script>

<section>
	<h1>Thank you for contacting us</h1>
	<h3>We will get back to you as soon as possible.</h3>
</section>
