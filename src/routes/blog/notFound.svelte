<script>
	import ButtonLink from '$lib/buttonLink.svelte';
</script>

<section>
	<h1>
		<span class="text-v-lilac">Woops...</span>
		<span class="text-v-pink">There's no blog for you here.</span>
	</h1>
	<h3>We recently updated our website, and some old links may have broken. Sorry about that.</h3>
	<ButtonLink href={'/blog'}>Back to blogs</ButtonLink>
</section>
