<script lang="ts">
	import Column from '$lib/column.svelte';
	import Columns from '$lib/columns.svelte';
	import CtaButton from '$lib/ctaButton.svelte';
	import { seo } from '$lib/seo/store';
	import ButtonLink from '$lib/buttonLink.svelte';

	seo.reset();
	$seo.title = 'Continuous Delivery';
	// TODO: description
	$seo.description = 'TODO: something about continuous delivery';
	$seo.image.url = '/continuous-delivery.svg';
	$seo.image.alt = 'continuous-delivery';
</script>

<section id="continuous-delivery">
	<h1>Continuous Delivery.</h1>
</section>
<section>
	<Columns>
		<Column class="self-start">
			<h3>Continuous Delivery</h3>
			<p class="mb-0">
				Continuous Delivery is about building a process where software is always releasable whilst
				ensuring short release cycles. Often the pains and struggles to help achieve this are not
				related to technology but stem from a fragmented delivery process with multiple handovers
				causing friction and delays.
			</p>
		</Column>
		<Column class="self-start">
			<h3>What we believe</h3>
			<p class="mb-0">
				Every software team is slightly different and best practices are not a one-size-fits all.
				Our workshops and coaching help you unlock new potential by guiding you through your own
				Continuous Delivery journey. We will save you time by guard-railing you from common
				pitfalls, whilst learning from each other as we go.
			</p>
		</Column>
	</Columns>
</section>
<section>
	<h3>Improve your productivity, quality and sustainability with Continuous Delivery.</h3>
</section>
<section>
	<Columns fullWidth={true} reverse={true}>
		<Column>
			<img
				class="object-contain h-full w-full"
				src="/work/cd-metrics/four-key-metrics.png"
				alt="four-key-metrics"
			/>
		</Column>
		<Column>
			<h3>Measuring performance</h3>
			<p>
				To understand if changes to your software delivery have a positive effect you have to
				measure your performance. We use the four key DORA metrics as a foundation for measuring
				software delivery performance.
			</p>
			<ButtonLink href={'/work/measuring-performance'}>Learn more</ButtonLink>
		</Column>
	</Columns>
</section>
<section>
	<Columns fullWidth={true}>
		<Column>
			<img
				class="object-contain w-full h-full"
				src="/work/cd-workshop/mock-vsm-sample.png"
				alt="mock-vsm-sample"
			/>
		</Column>
		<Column>
			<h3>Identifying improvements</h3>
			<p>
				Identifying what to change in your software delivery is not easy and should not be driven by
				metrics. We use Value Stream Mapping as a method to help teams make informed decisions.
			</p>
			<ButtonLink href={'/work/continuous-delivery-workshop'}>Learn more</ButtonLink>
		</Column>
	</Columns>
</section>
<section>
	<h3>
		We use the <span class="text-v-pink">Build-Measure-Learn</span> feedback cycle to implement change.
		Althought our approach has an extra learn phase where we use Value Stream Mapping to map out your
		development process, discuss team topologies and identify possible changes.
	</h3>
	<div>
		<img
			class="object-contain h-full w-full"
			src="/work/cd-workshop/build-measure-learn.png"
			alt="build-measure-learn"
		/>
	</div>
</section>

<section>
	<Columns>
		<Column>
			<div class="flex flex-col gap-y-6">
				<div class="flex space-x-6 items-center">
					<div>
						<img class="h-4 w-4 object-contain" src="/logo-element.png" alt="logo-ement" />
					</div>
					<div>
						<h4 class="mb-0">Visualize your release process</h4>
					</div>
				</div>
				<div class="flex space-x-6 items-center">
					<div>
						<img class="h-4 w-4 object-contain" src="/logo-element.png" alt="logo-ement" />
					</div>
					<div>
						<h4 class="mb-0">Pinpoint opportunities and waste</h4>
					</div>
				</div>
				<div class="flex space-x-6 items-center">
					<div>
						<img class="h-4 w-4 object-contain" src="/logo-element.png" alt="logo-ement" />
					</div>
					<div>
						<h4 class="mb-0">Understand team topologies</h4>
					</div>
				</div>
				<div class="flex space-x-6 items-center">
					<div>
						<img class="h-4 w-4 object-contain" src="/logo-element.png" alt="logo-ement" />
					</div>
					<div>
						<h4 class="mb-0">Justify investment</h4>
					</div>
				</div>
			</div>
		</Column>
		<Column>
			<CtaButton
				fullScreen={false}
				text={'Ready to unlock your CD potential?'}
				contact={'Get started.'}
			/>
		</Column>
	</Columns>
</section>
<section>
	<h2>‘Best practices are a good solution to someone else's problem’</h2>
	<h4>— Bosse Nyström, Chief of Wisdom, Verifa</h4>
</section>
