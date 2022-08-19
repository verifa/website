<script>
	import ButtonLink from '$lib/buttonLink.svelte';
	import { seo } from '$lib/seo/store';

	seo.reset();
</script>

<section>
	<h1>
		<span class="text-v-lilac">Woops...</span>
		<span class="text-v-pink">There's no case for you here.</span>
	</h1>
	<h3>We couldn't find the case you are looking for. Sorry about that.</h3>
	<ButtonLink href={'/clients'}>Back to clients</ButtonLink>
</section>
