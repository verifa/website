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
			<h1 style="margin-bottom: 0px">Accelerate your software delivery.</h1>
			<h4 style="margin-top: 1rem; margin-bottom: 2.5rem; ">We help optimise your software delivery flow and build high performing teams.</h4>
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
	<h3 class="text-v-lilac">Assess. Coach. Build. Evolve.</h3>
	<h3>
		We’re an expert crew of experienced software delivery consultants based in the Nordics. We help
		our clients make informed decisions about their software delivery processes and platforms, team
		structures, open source technologies, and cloud services. We believe in working together with
		our clients to assess their current situation, coach their teams, and design and implement
		sustainable solutions that they own.
	</h3>
</section>

<section>
	<Grid>
		<div>
			<h3>Value Stream Assessment</h3>
			<span>
				Visualise your value streams to identify waste and highlight opportunities for faster flow.
			</span>
			<ButtonLink href="/work/value-stream-assessment" class="text-base lg:text-xl">Learn More</ButtonLink>
		</div>
		<div>
			<h3>Software Delivery Platform</h3>
			<span>
				Find the right platform abstraction to enable delivery teams in accelerating software
				delivery.
			</span>
			<ButtonLink href="/work/software-delivery-platforms" class="text-base lg:text-xl">Learn More</ButtonLink>
		</div>
		<div>
			<h3>Team Topologies</h3>
			<span>We help organise happy, high performing teams that deliver better value to customers.</span>
			<ButtonLink href="/work/team-topologies" class="text-base lg:text-xl">Learn More</ButtonLink>
		</div>
		<div>
			<h3>Cloud Architecture</h3>
			<span>Designing and implementing cloud-native architectures that are robust yet evolvable.</span>
			<ButtonLink href="/work/cloud-architecture" class="text-base lg:text-xl">Learn More</ButtonLink>
		</div>
	</Grid>
</section>

<section>
	<h3>Trusted by</h3>
	<ClientLogos />
</section>

<section>
	<Columns reverse={true}>
		<Column class="self-start">
			<div class="flex flex-col gap-y-4">
				<h1 class="leading-tight" style="margin-bottom: 0">
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
					<span>
						We are a small and dedicated crew who care about the value we deliver. We look to
						tomorrow's problems and solve them for you ahead of time.
					</span>
				</div>
				<div>
					<h3>What makes us different.</h3>
					<span>
						We don't just solve problems but teach you how to solve them, and we do it as a team.
						That's our goal and what sets us apart from other consultancy houses.
					</span>
				</div>
				<div>
					<h3>Who we have helped.</h3>
					<span>
						We have worked with regulated and embedded environments in industries like Automotive,
						Medical, Finance, Mobile and Web.
					</span>
				</div>
			</div>
		</Column>
	</Columns>
</section>

<section>
	<Columns reverse={true}>
		<Column>
			<div class="flex flex-col gap-y-2">
				<h4 style="margin: 0px">Technologies</h4>
				<h2 style="margin: 0px">
					We work with the world's leading providers to help us deliver the best possible
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
	<h4 style="margin: 0px">Learn more on our blog</h4>
	<h2>Search by popular keyword</h2>
	<div class="mb-5 md:mb-8 lg:mb-12"></div>
	<div class="flex flex-col gap-y-12 mb-8">
			<div class="-my-2 flex flex-wrap gap-x-4">
			{#each data.posts.keywords as tag}
				<a href="/blog?keywords={tag}" class="inline-block ">
					<span class="inline-flex items-center my-2 px-3 py-0.5 bg-v-gray">
						<span class="m-0 capitalize text-v-white">{tag}</span>
					</span>
				</a>
			{/each}
		</div>
		<PostGrid showBadges={false} {posts} />
	</div>
	<ButtonLink href="/blog">All posts</ButtonLink>
</section>
