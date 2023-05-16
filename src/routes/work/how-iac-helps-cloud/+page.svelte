<script lang="ts">
	import Column from '$lib/column.svelte';
	import Columns from '$lib/columns.svelte';
	import CtaButton from '$lib/ctaButton.svelte';
	import Grid from '$lib/grid.svelte';
	import { seo } from '$lib/seo/store';
	import HeaderLine from '$lib/headerLine.svelte';
	import MainReference from '$lib/mainReference.svelte';

	seo.reset();
	$seo.title = 'How Infrastructure as Code can help you with Cloud';
</script>

<svelte:head>
	<script
		type="text/javascript"
		src="https://assets.calendly.com/assets/external/widget.js"
		async></script>
</svelte:head>

<section>
	<h1>How Infrastructure as Code can help you with Cloud</h1>
</section>
<section>
	<Grid>
		<div class="self-start flex flex-col gap-y-12">
			<img class="w-full h-full" src="/continuous-delivery.svg" alt="continuous-delivery" />
			<div>
				<h3 class="">Benefits of Infrastructure as Code</h3>
				<div>
					<p class="mb-4">
						Infrastructure as Code provides many benefits to your cloud usage. Most notably are:
					</p>
					<ul class="verifa-bullet">
						<li>Scalability by Automation</li>
						<li>Auditability through Version Control</li>
						<li>Cost optimization</li>
					</ul>
				</div>
				<p class="text-v-black">
					Schedule a meeting with our CEO and learn how you can benefit from IaC.
				</p>
			</div>
		</div>
		<div class="self-start h-full">
			<!-- Calendly inline widget begin -->
			<div
				class="calendly-inline-widget h-full"
				data-url="https://calendly.com/jlarfors/infrastructure-as-code?hide_gdpr_banner=1&background_color=f9fafb&text_color=0d0e12&primary_color=ad9ce3"
			/>

			<!-- Calendly inline widget end -->
		</div>
	</Grid>
</section>
<section>
	<h1>Why Infrastructure as Code?</h1>
</section>
<section>
	<Columns>
		<Column>
			<h3>What is IaC?</h3>
			<p>Some text here</p>
		</Column>
		<Column>
			<h3>Tooling</h3>
			<p>Some text here</p>
		</Column>
		<Column>
			<h3>Use cases</h3>
			<p>Some text here</p>
		</Column>
	</Columns>
</section>

<MainReference />
