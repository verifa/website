<script lang="ts">
	import Column from '$lib/column.svelte';
	import Columns from '$lib/columns.svelte';
	import CtaButton from '$lib/ctaButton.svelte';
	import { seo } from '$lib/seo/store';
	import ButtonLink from '$lib/buttonLink.svelte';

	seo.reset();
	$seo.title = 'Measuring software delivery performance';
	// TODO: description
	$seo.description = 'TODO: something about continuous delivery';
	$seo.image.url = '/continuous-delivery.svg';
	$seo.image.alt = 'continuous-delivery';
</script>

<h1>Measuring software delivery performance.</h1>

<section>
	<h3>TODO</h3>
</section>
