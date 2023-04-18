<script lang="ts">
	import { seo } from '$lib/seo/store';
	import ButtonLink from '$lib/buttonLink.svelte';

	seo.reset();
	$seo.title = 'What we do: Team Topologies';
</script>

<section>
	<div class="flex">
		<div>
			<h1>Team Topologies</h1>
			<h3>
				We help organise happy, high performing teams that deliver better value to customers.
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
			Even with automation, building software is an inherently human undertaking and the ways 
			of working play a huge part in its success. The composition of teams, relationships and 
			interactions is an often overlooked practice when improving a continuous delivery process.

			Verifa’s holistic approach, based on Team Topologies (link to TT site??), is about organising 
			technical teams in a such a way that they are able to deliver value continuously and sustainably.
		</p>
	</div>
</section>
<section id="how-we-help">
	<div>
		<h2>How we can help</h2>
		<p class="mb-0">
			Every team is different, and we need to understand your unique situation. The starting 
			point for this is working together with your team to conduct a Value Stream Assessment 
			(link to VSA page) to uncover your current way of working, interactions and software delivery 
			flow. Each VSA is customised but often includes:

			- item 1
			- item 2
			- item 3
			- item 4
		</p>
		<ButtonLink href="/contact">Get started</ButtonLink>
	</div>
</section>
