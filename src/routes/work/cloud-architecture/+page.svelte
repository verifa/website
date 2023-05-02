<script lang="ts">
	import Column from '$lib/column.svelte';
	import Columns from '$lib/columns.svelte';
	import { seo } from '$lib/seo/store';
	import ButtonLink from '$lib/buttonLink.svelte';
	import HeaderLine from '$lib/headerLine.svelte';

	seo.reset();
	$seo.title = 'What we do: Cloud Architecture';
</script>

<section>
	<Columns reverse={false}>
		<Column>
			<h1>Cloud Architecture</h1>
			<h3>
				We help organise happy, high performing teams that deliver better value to customers.
			</h3>
		</Column>
		<Column>
			<img
				class="object-contain h-full w-full"
				src="/cloud-architecture.svg"
				alt="An illustration of a group of people designing a complex 3-dimensional object"
			/>
		</Column>
	</Columns>
</section>

<section>
	<HeaderLine/>
	<h4>What</h4>
	<Columns reverse={false}>
		<Column>
			A Software Delivery Platform serves to reduce the cognitive load for software delivery teams, 
		enabling them to focus on delivering value instead of firefighting.
		</Column>
		<Column>
			Developing software is increasingly complex, requiring many moving parts. A well designed Platform 
			seeks to standardise as much as possible, to create a human friendly interface together with a 
			stable and reliable environment to develop, test and deploy software. It should also be designed 
			to scale, and tools to be replaceable in order to avoid repeating pain. 
		</Column>
	</Columns>
	<h3 class="mb-0">
		Platform Engineering helps to define the ‘Golden Path’ for software teams to deliver value. 
		It all begins with research on how teams work in order to create a proper "Platform Abstraction". 
		This drives what the platform will do, who will use it, and what workflows it needs to fulfil.
	</h3>
</section>
<section id="how-we-help">
	<h2>How we can help</h2>
	<p class="mb-0">
		Every team is different, and we need to understand your unique situation. The starting 
		point for this is working together with your team to conduct a Value Stream Assessment 
		(link to VSA page) to uncover your current way of working, interactions and software delivery 
		flow. Each VSA is customised but often includes:
	</p>
	<ul>
		<li>Item 1</li>
		<li>Item 2</li>
		<li>Item 3</li>
		<li>Item 4</li>
	</ul>
<ButtonLink href="/contact">Get started</ButtonLink>
</section>
