<script lang="ts">
	import Column from '$lib/column.svelte';
	import Columns from '$lib/columns.svelte';
	import CtaButton from '$lib/ctaButton.svelte';
	import Grid from '$lib/grid.svelte';
	import { seo } from '$lib/seo/store';
	import ButtonLink from '$lib/buttonLink.svelte';
	import CDWorkshop from '../_cdWorkshop.svelte';

	seo.reset();
	$seo.title = 'What we do: Value Stream Mapping';
</script>

<section>
	<div class="flex">
		<div>
			<h1>Value Stream Assessment</h1>
			<h3>
				Visualize yous value streams to identify waste and highlight opportunities for faster flow
			</h3>
		</div>
		<div>
			<img
				class="object-contain h-full w-full"
				src="/continuous-delivery.svg"
				alt="continuous delivery"
			/>
		</div>
	</div>
</section>

<section>
	<div>
		<p class="mb-0">
		A Value Stream Map is a visual tool showing the series of events that take a product or
		service from the beginning of a process until it reaches its end (i.e. a Value Stream). 
		Value Stream Mapping (VSM) is a method for defining a Value Stream and analysing the 
		current flow in order to remove waste and to derive a more lean and continuous workflow.

		The idea of implementing change in your software delivery is not too dissimilar from 
		implementing change in your software itself, and should be done in a *continuous* way. 
		View it as using Continuous Delivery for improving your Continuous Delivery.
		</p>
	</div>
</section>
<section id="how-we-help">
	<h2>How we can help</h2>
	<CDWorkshop />
</section>
