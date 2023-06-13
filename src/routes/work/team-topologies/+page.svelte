<script lang="ts">
	import Column from '$lib/column.svelte';
	import Columns from '$lib/columns.svelte';
	import { seo } from '$lib/seo/store';
	import ButtonLink from '$lib/buttonLink.svelte';
	import HeaderLine from '$lib/headerLine.svelte';
	import CtaButton from '$lib/ctaButton.svelte';

	seo.reset();
	$seo.title = 'What we do: Team Topologies';
</script>

<section>
	<Columns reverse={false}>
		<Column>
			<h1 style="margin-bottom: 0px">Team Topologies.</h1>
			<h4 style="margin-top: 1rem; margin-bottom: 2.5rem; ">
				We help organise happy, high performing teams that deliver better value to customers.
			</h4>
		</Column>
		<Column>
			<img
				class="object-contain h-full w-full"
				src="/round-table.svg"
				alt="An illustration of a group of people seated at a round table."
			/>
		</Column>
	</Columns>
</section>

<section>
	<HeaderLine />
	<h2>Faster flow, more value</h2>
	<Columns reverse={false} class="lg:items-baseline">
		<Column>
			Even with automation, building software is an inherently human undertaking and the ways of
			working play a huge part in its success. The composition of teams, relationships and
			interactions is an often overlooked practice when improving a continuous delivery process.
		</Column>
		<Column>
			Verifa’s holistic approach, based on <a href="https://teamtopologies.com/" target="_blank"
				>Team Topologies</a
			>, is about organising technical teams in a such a way that they are able to deliver value
			continuously and sustainably.
		</Column>
	</Columns>
</section>
<section id="how-we-help">
	<HeaderLine />
	<h2 class="text-v-lilac">How we can help</h2>
	<h2 class="mb-0">
		Every team is different, and we need to understand your unique situation. The starting point for
		this is working together with your team to conduct a <span class="text-v-pink"
			>Value Stream Assessment</span
		> to uncover your current way of working, interactions and software delivery flow.
	</h2>
	<ButtonLink href="/work/value-stream-assessment">Learn More</ButtonLink>

	<Columns reverse={false}>
		<Column>
			<img
				class="object-contains w-full h-full"
				src="/everything-is-connected.svg"
				alt="everything is connected"
			/>
		</Column>
		<Column>
			<CtaButton
				text={'Want a happier, more productive team?'}
				contact={'Let’s chat'}
				fullScreen={false}
			/>
		</Column>
	</Columns>
</section>
