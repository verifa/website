<script context="module" lang="ts">
	export async function load({ fetch }) {
		const postsUrl = `/posts/jobs.json`;
		const res = await fetch(postsUrl);
		if (res.ok) {
			return {
				props: {
					data: await res.json()
				}
			};
		}

		return {
			status: res.status,
			error: new Error(`Could not load ${postsUrl}`)
		};
	}
</script>

<script lang="ts">
	import Column from '$lib/column.svelte';
	import Columns from '$lib/columns.svelte';
	import CompanyValues from '$lib/company/companyValues.svelte';

	import type { Jobs, Post } from '$lib/posts/posts';
	import PostGrid from '$lib/posts/postGrid.svelte';
	import HeaderLine from '../_headerLine.svelte';
	import Form from './_form.svelte';
	import { seo } from '$lib/seo/store';

	export let data: Jobs;

	$seo.title = 'Join us: we care about our work and the impact we have';
	$seo.image.url = '/round-table.svg';

	const jobs: Post[] = data.jobs;
</script>

<section>
	<HeaderLine />
	<h4>Join us</h4>
	<h1>We care about our work and the impact we have.</h1>
</section>
<section>
	<h3>
		We are an experienced crew of DevOps and Cloud professionals dedicated to helping our customers
		with Continuous practices and Cloud adoption. We each bring our own specialities, experiences
		and know-how when collaborating on customer projects, tailoring services to our customers'
		individual needs.
	</h3>
</section>

<section>
	<HeaderLine />
	<h4>Our Values</h4>
	<CompanyValues />
</section>

<section>
	<Columns>
		<Column>
			<h2>What we look for.</h2>
			<p>
				We look for people passionate about continuous practices, cloud architecture and all the
				wonderful things that will help our customers deliver greatness. If you like to be
				challenged and want to keep developing, then you should get in touch.
			</p>
		</Column>
		<Column>
			<h2>What we offer.</h2>
			<p>
				Our goal is to provide a great place to work and continue to learn whilst delivering value
				to our customers. You'll get to work with cool tech and discuss with others who share a
				passion to deliver.
			</p>
		</Column>
	</Columns>
</section>
<section id="open-positions">
	<HeaderLine />
	<h4>Open positions</h4>
	<h1>Check out our open positions.</h1>
</section>
<section>
	<PostGrid posts={jobs} showBadges={false} />
</section>
<section>
	<HeaderLine />
	<h4>Apply now</h4>
	<h1>Join our crew.</h1>
</section>
<section>
	<Form />
</section>
