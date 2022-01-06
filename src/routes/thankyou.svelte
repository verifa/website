<script lang="ts">
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';

	onMount(() => {
		setTimeout(() => goto('/'), 1000 * 3);
	});
</script>

<section>
	<h1>Thank you for contacting us</h1>
	<h3>We will get back to you as soon as possible.</h3>
</section>
