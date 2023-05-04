<script lang="ts">
	import Column from '$lib/column.svelte';
	import Columns from '$lib/columns.svelte';
	import { seo } from '$lib/seo/store';
	import ButtonLink from '$lib/buttonLink.svelte';
	import HeaderLine from '$lib/headerLine.svelte';

	seo.reset();
	$seo.title = 'What we do: Software Delivery Platforms';
</script>

<section>
	<Columns reverse={false}>
		<Column>
			<h1>Software Delivery Platforms.</h1>
			<h3>
				Find the right platform abstraction to enable delivery teams in accelerating software
				delivery.
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
	<HeaderLine />
	<div>
		<h4>Human-friendly platforms</h4>
		<Columns reverse={false}>
			<Column>
				Agile and DevOps have enabled cross-functional and autonomous teams, but has also led to
				additional responsibilities for the product teams. As well as development of the software,
				they also need to support the infrastructure and daily operations. This causes additional
				cognitive load and takes time away from providing actual value.
			</Column>
			<Column>
				A Software Delivery Platform seeks to reduce the cognitive load for software delivery teams,
				allowing them to focus on delivering value instead of firefighting. It’s a set of tools,
				processes and technologies that enable teams to build, test and deploy software more
				efficiently and reliably.
			</Column>
		</Columns>
		<h3 class="mb-0">
			A well-designed platform removes the complexity of modern infrastructures by providing a
			human-friendly interface and a stable, self-serve environment, so that teams can focus on
			delivering value faster.
		</h3>
	</div>
</section>

<section id="how-we-help">
	<HeaderLine />
	<h4>How we can help</h4>
	<Columns reverse={false}>
		<Column>
			Often, a ‘platform’ is a collection of tools maintained by a team, with little thought into
			how they can help delivery teams. Platform Engineering helps to define the ‘Golden Path’ for
			software teams to deliver value. It all begins with research on how teams work in order to
			create a proper "Platform Abstraction". This drives what the platform will do, who will use
			it, and what workflows it needs to fulfil.
		</Column>
		<Column>
			Verifa helps teams to design and implement Software Delivery Platforms that decouple teams,
			allowing them to work independently and deliver value faster. Value Stream Assessments (link
			to VSA page) are our method for understanding team interactions and flow of delivery.
			<ButtonLink href="/contact">Chat with an expert</ButtonLink>
		</Column>
	</Columns>
</section>
