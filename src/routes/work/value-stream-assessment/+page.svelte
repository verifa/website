<script lang="ts">
	import Column from '$lib/column.svelte';
	import Columns from '$lib/columns.svelte';
	import CtaButton from '$lib/ctaButton.svelte';
	import HeaderLine from '$lib/headerLine.svelte';
	import { seo } from '$lib/seo/store';
	import ButtonLink from '$lib/buttonLink.svelte';
	import CDWorkshop from '../_cdWorkshop.svelte';

	seo.reset();
	$seo.title = 'What we do: Value Stream Mapping';
</script>

<section>
	<Columns reverse={false}>
		<Column>
			<h1>Value Stream Assessment.</h1>
			<h3>
				Visualise your value streams to identify waste and highlight opportunities for faster flow.
			</h3>
		</Column>
		<Column>
			<img
				class="object-contain h-full w-full"
				src="/visibility.svg"
				alt="An illustration of woman holding a magnifying glass up to a screen displaying source code."
			/>
		</Column>
	</Columns>
</section>

<section>
	<HeaderLine />
	<h4>Lean and continuous workflows</h4>
	<Columns reverse={false} class="lg:items-baseline">
		<Column>
			A Value Stream Map is a visual tool showing the series of events that take a product or
			service from the beginning of a process until it reaches its end, i.e. a Value Stream. Value
			Stream Mapping (VSM) is a method for defining a Value Stream and analysing the current flow in
			order to remove waste and to derive a more lean and continuous workflow.
		</Column>
		<Column>
			The idea of implementing change in your software delivery is not too dissimilar from
			implementing change in your software itself, and should be done in a <i>continuous</i> way. View
			it as using Continuous Delivery for improving your Continuous Delivery.
		</Column>
	</Columns>
</section>
<section id="how-we-help">
	<HeaderLine />
	<h4>How we can help</h4>
	<CDWorkshop />
	<div class="flex">
		<ul>
			<li>Visualise your release process</li>
			<li>Pinpoint opportunities and waste</li>
			<li>Get your whole team onboard</li>
			<li>Justify investment</li>
		</ul>
		<Columns>
			<Column>
				<CtaButton
					text={'Ready to unlock your CD potential?'}
					contact={'Get Started'}
					fullScreen={false}
				/>
			</Column>
		</Columns>
	</div>
</section>
<section>
	<h2>
		‘Verifa helped lead discussions around many topics which was useful in gaining a common
		understanding of what we can do to improve’.
	</h2>
	<h4>— Olli Suihkonen, Visma Software Oy.</h4>
	<ButtonLink href="/clients/value-stream-mapping-with-visma/">Read the case study</ButtonLink>
</section>
