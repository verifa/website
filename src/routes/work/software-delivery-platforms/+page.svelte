<script lang="ts">
	import Column from '$lib/column.svelte';
	import Columns from '$lib/columns.svelte';
	import { seo } from '$lib/seo/store';
	import ButtonLink from '$lib/buttonLink.svelte';

	seo.reset();
	$seo.title = 'What we do: Software Delivery Platforms';
</script>

<section>
	<Columns reverse={false}>
		<Column>
			<h1>Software Delivery Platforms</h1>
			<h3>
				Find the right platform abstraction to enable delivery teams in accelerating software delivery.
			</h3>
		</Column>
		<Column>
			<img
				class="object-contain h-full w-full"
				src="/company-verifa_colors.svg"
				alt="An illustration of people engaging with a wall-sized dashboard of charts and information"
			/>
		</Column>
	</Columns>
</section>


<section>
	<div>
		<h4 class="mb-0">
			A Software Delivery Platform serves to reduce the cognitive load for software delivery teams, 
			enabling them to focus on delivering value instead of firefighting.
			
			Developing software is increasingly complex, requiring many moving parts. A well designed Platform 
			seeks to standardise as much as possible, to create a human friendly interface together with a 
			stable and reliable environment to develop, test and deploy software. It should also be designed 
			to scale, and tools to be replaceable in order to avoid repeating pain. 
			
			Platform Engineering helps to define the ‘Golden Path’ for software teams to deliver value. 
			It all begins with research on how teams work in order to create a proper "Platform Abstraction". 
			This drives what the platform will do, who will use it, and what workflows it needs to fulfil.
		</h4>
	</div>
</section>
<section id="how-we-help">
	<div>
		<h2>How we can help</h2>
		<p class="mb-0">
			Often, a ‘platform’ is a collection of tools maintained by a team, with little thought into how 
			they can help delivery teams. Verifa helps teams to design and implement Software Delivery 
			Platforms that decouples teams, allowing them to work independently and deliver value faster.

			Value Stream Assessments (link to VSA page) are our method for understanding team 
			interactions and flow of delivery.
		</p>
		<ButtonLink href="/contact">Get started</ButtonLink>
	</div>
</section>
