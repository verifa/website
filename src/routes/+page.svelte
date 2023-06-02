<script lang="ts">
	import PostGrid from '$lib/posts/postGrid.svelte';
	import ButtonLink from '$lib/buttonLink.svelte';
	import Columns from '$lib/columns.svelte';
	import Column from '$lib/column.svelte';
	import Grid from '$lib/grid.svelte';
	import { seo } from '$lib/seo/store';
	import MainReference from '$lib/mainReference.svelte';
	import type { PageData } from '.svelte-kit/types/src/routes/$types';
	import type { Post } from '$lib/posts/posts';
	import ClientLogos from '$lib/clients/clientLogos.svelte';
	import TechnologyLogos from '$lib/technologies/technologyLogos.svelte';

	export let data: PageData;
	const posts: Post[] = data.posts.posts;

	seo.reset();
</script>

<section>
	<Columns reverse={false}>
		<Column>
			<h1>Accelerate your software delivery.</h1>
			<h3>We help optimise your software delivery flow and build high performing teams.</h3>
			<ButtonLink href="/contact">Chat with an expert</ButtonLink>
		</Column>
		<Column>
			<img
				class="object-contain h-full w-full"
				src="/continuous-delivery.svg"
				alt="continuous delivery"
			/>
		</Column>
	</Columns>
</section>

<section>
	<h2>Assess. Coach. Build. Evolve.</h2>
	<h4>
		We’re an expert crew of experienced software delivery consultants based in the Nordics. We help
		our clients make informed decisions about their software delivery processes and platforms, team
		structures, open source technologies, and cloud services. We believe in working together with
		our clients to assess their current situation, coach their teams, and design and implement
		sustainable solutions that they own.
	</h4>
</section>

<section>
	<Grid>
		<div>
			<h3>Value Stream Assessment</h3>
			<p>
				Visualise your value streams to identify waste and highlight opportunities for faster flow.
			</p>
			<ButtonLink href="/work/value-stream-assessment">Learn More</ButtonLink>
		</div>
		<div>
			<h3>Software Delivery Platform</h3>
			<p>
				Find the right platform abstraction to enable delivery teams in accelerating software
				delivery.
			</p>
			<ButtonLink href="/work/software-delivery-platforms">Learn More</ButtonLink>
		</div>
		<div>
			<h3>Team Topologies</h3>
			<p>We help organise happy, high performing teams that deliver better value to customers.</p>
			<ButtonLink href="/work/team-topologies">Learn More</ButtonLink>
		</div>
		<div>
			<h3>Cloud Architecture</h3>
			<p>Designing and implementing cloud-native architectures that are robust yet evolvable.</p>
			<ButtonLink href="/work/cloud-architecture">Learn More</ButtonLink>
		</div>
	</Grid>
</section>

<section>
	<h2>Trusted by.</h2>
	<ClientLogos />
</section>

<section>
	<Columns reverse={true}>
		<Column class="self-start">
			<div class="flex flex-col gap-y-4">
				<h1 style="margin-bottom: 0">
					‘A key and crucial part of our Continuous Integration journey was having Verifa's
					involvement’.
				</h1>
				<h4 style="margin-top: 0">— David Hoslett, Siemens Mobility.</h4>
			</div>
			<div>
				<ButtonLink href="/clients">Our clients</ButtonLink>
			</div>
		</Column>
		<Column class="self-start">
			<div class="flex flex-col gap-y-0">
				<div>
					<h3 style="margin-top: 0">Why work with Verifa.</h3>
					<p>
						We are a small and dedicated crew who care about the value we deliver. We look to
						tomorrow's problems and solve them for you ahead of time.
					</p>
				</div>
				<div>
					<h3>What makes us different.</h3>
					<p>
						We don't just solve problems but teach you how to solve them, and we do it as a team.
						That's our goal and what sets us apart from other consultancy houses.
					</p>
				</div>
				<div>
					<h3>Who we have helped.</h3>
					<p>
						We have worked with regulated and embedded environments in industries like Automotive,
						Medical, Finance, Mobile and Web.
					</p>
				</div>
			</div>
		</Column>
	</Columns>
</section>

<section>
	<h2>Technologies.</h2>
	<Columns reverse={true}>
		<Column>
			<div class="flex flex-col gap-y-2">
				<h2 style="margin: 0px">
					We work with the world's leading technologies to help us deliver the best possible
					experience for our clients.
				</h2>
			</div>
		</Column>
		<Column>
			<TechnologyLogos />
		</Column>
	</Columns>
</section>

<section>
	<h1>Learn more on our blog.</h1>
</section>
<section>
	<h3 class="text-v-lilac">Search by popular keyword</h3>
	<div class="flex flex-col gap-y-12 mb-8">
		<div class="-my-2 flex flex-wrap gap-x-4">
			{#each data.posts.keywords as tag}
				<a href="/blog?keywords={tag}" class="inline-block ">
					<span class="inline-flex items-center my-2 px-3 py-0.5 bg-v-gray">
						<p class="m-0 capitalize text-v-white">{tag}</p>
					</span>
				</a>
			{/each}
		</div>
		<PostGrid showBadges={false} {posts} />
	</div>
	<ButtonLink href="/blog">All posts</ButtonLink>
</section>
