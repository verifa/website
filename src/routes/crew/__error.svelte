<script>
	import ButtonLink from '$lib/buttonLink.svelte';
	import { seo } from '$lib/seo/store';

	seo.reset();
</script>

<section>
	<h1>
		<span class="text-v-lilac">Woops...</span>
		<span class="text-v-pink">That crew member doesn't exist yet.</span>
	</h1>
	<h3>Want to check people who really do work at Verifa?</h3>
	<ButtonLink href={'/crew'}>Back to crew</ButtonLink>
</section>
