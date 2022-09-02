<script lang="ts">
	import Column from '$lib/column.svelte';
	import Columns from '$lib/columns.svelte';
	import Grid from '$lib/grid.svelte';
	import { seo } from '$lib/seo/store';
	import type { PageData } from '.svelte-kit/types/src/routes/$types';
	import CdWorkshop from '../_cdWorkshop.svelte';

	seo.reset();
	$seo.title = 'Continuous Delivery Metrics';
	// $seo.description =
	// 	'Our Continuous Delivery Workshops help teams unlock their Continuous Delivery potential. We use Value Stream Mapping as our method for helping teams understand their current release process and create a backlog of improvements.';

	const metrics = [
		{
			title: 'Deployment frequency',
			description: 'How often a software team pushes changes to production'
		},
		{
			title: 'Change lead time',
			description: 'The time it takes to get committed code to run in production'
		},
		{
			title: 'Change failure rate',
			description: 'The share of incidents, rollbacks, and failures out of all deployments'
		},
		{
			title: 'Mean time to recovery',
			description: 'The time it takes to restore service in production after an incident'
		}
	];
</script>

<section>
	<h1>Continuous Delivery Metrics.</h1>
</section>
<section>
	<h3>Our approach to ...</h3>
</section>
<section>
	<Grid>
		{#each metrics as metric, index}
			<div>
				<h3>{index + 1}. {metric.title}</h3>
				<p class="mb-0">{metric.description}</p>
			</div>
		{/each}
	</Grid>
</section>
<section>
	<CdWorkshop />
</section>
