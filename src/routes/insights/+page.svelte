<script lang="ts">
	import { goto } from '$app/navigation';
	import { seo } from '$lib/seo/store';
	import { onMount } from 'svelte';

	onMount(() => {
		setTimeout(() => goto('/blog'), 1000 * 3);
	});

	seo.reset();
</script>

<section>
	<h1>
		<span class="text-v-lilac">New website.</span>
		<span class="text-v-pink">New routes.</span>
	</h1>
	<h3>Redirecting you to where you want to be...</h3>
</section>
