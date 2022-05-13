<script context="module" lang="ts">
	export async function load({ fetch }) {
		const postsUrl = `/posts/cases.json`;
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
	import HeaderLine from '$lib/headerLine.svelte';
	import Testimonial from './_testimonial.svelte';

	import ClientLogos from '$lib/clients/clientLogos.svelte';
	import CtaButton from '$lib/ctaButton.svelte';
	import Grid from '$lib/grid.svelte';
	import { seo } from '$lib/seo/store';
	import type { Cases } from '$lib/posts/posts';
	import PostGrid from '$lib/posts/postGrid.svelte';

	export let data: Cases;

	seo.reset();
	$seo.title = 'Our clients: we care and we deliver';
	$seo.description =
		'We are proud to work with some of the biggest companies in the world across a range of industries. With every customer our goal is always the same: to solve problems and share knowledge. Our business is based on building trust and long-lasting relationships.';
	$seo.image.url = '/clients.svg';

	const quotes: {
		text: string;
		person: string;
		logo: string;
	}[] = [
		{
			text: 'Verifa helped lead discussions around many topics which was useful in gaining a common understanding of what we can do to improve. We are very happy with the report Verifa prepared for us.',
			person: 'Olli Suihkonen, Visma',
			logo: '/clients/visma.svg'
		},
		{
			text: "A key and crucial part of our Continuous Integration journey was having Verifa's involvement in holding-our-hands and giving us confidence while we moved our active projects to a CI methodology.",
			person: 'David Hoslett, Siemens Mobility',
			logo: '/clients/siemens.svg'
		},
		{
			text: 'We were looking for help with professionalising our Continuous Integration and software delivery platform whilst moving the setup to the cloud. Verifa was a very reliable partner in implementing a modern setup using their years of experience.',
			person: 'Laurent Muller, Vyaire Medical',
			logo: '/clients/vyaire.png'
		},
		{
			text: "Verifa helped us visualize, assess and define metrics for our software architecture using Lattix and integrated this into our CI pipelines using an 'Architecture as Code' approach to monitoring the architecture, which was a big help for our project goals.",
			person: 'Juha Mailisto, ABB Drives',
			logo: '/clients/abb.svg'
		},
		{
			text: 'Working with Verifa for software services was very easy, and the staff we were working with were very technically minded. We had no problems communicating what we needed, or with getting the required feedback.',
			person: 'Pierre-Henri Staneck, QA Systems',
			logo: '/clients/qa-systems.png'
		}
	];
</script>

<section>
	<HeaderLine />
	<h4>Our Clients</h4>
	<h1>Continuously delivering for our clients.</h1>
</section>
<section>
	<Columns>
		<Column>
			<img class="object-contain" src="/clients.svg" alt="client" />
		</Column>
		<Column>
			<h3>We care and we deliver.</h3>
			<p>
				We are proud to work with some of the biggest companies in the world across a range of
				industries. With every customer our goal is always the same: to solve problems and share
				knowledge. Our business is based on building trust and long-lasting relationships.
			</p>
		</Column>
	</Columns>
</section>
<section>
	<HeaderLine />
	<h4>Cases</h4>
	<h1>How we have helped our clients.</h1>
</section>
<section>
	<PostGrid showBadges={false} posts={data.cases} />
</section>
<section>
	<HeaderLine />
	<h4>Quotes</h4>
	<h1>What our clients say about us.</h1>
</section>
<section>
	<Grid>
		{#each quotes as quote}
			<div class="self-start"><Testimonial {quote} /></div>
		{/each}
	</Grid>
</section>
<section>
	<ClientLogos />
</section>
<section>
	<CtaButton />
</section>
