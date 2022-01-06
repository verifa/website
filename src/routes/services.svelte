<script lang="ts">
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';

	onMount(() => {
		setTimeout(() => goto('/work'), 1000 * 3);
	});
</script>

<section>
	<h1>
		<span class="text-v-lilac">New website.</span>
		<span class="text-v-pink">New routes.</span>
	</h1>
	<h3>Redirecting you to where you want to be...</h3>
</section>
