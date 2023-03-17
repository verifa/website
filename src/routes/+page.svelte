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
	import HeaderLine from '$lib/headerLine.svelte';
	import ClientLogos from '$lib/clients/clientLogos.svelte';

	export let data: PageData;
	const posts: Post[] = data.posts.posts;

	seo.reset();
</script>

<section>
	<div class="flex">
		<div>
			<h1>Accelerate your software delivery</h1>
			<h3>
				We help design, implement and evolve platforms that accelerate the delivery of software
			</h3>
			<ButtonLink href="/work#continuous-delivery">Learn More</ButtonLink>
		</div>
		<div>
			<img
				class="object-contain h-full w-full"
				src="/continuous-delivery.svg"
				alt="continuous delivery"
			/>
		</div>
	</div>
</section>

<section>
	<h3>Assess. Coach. Build. Evolve.</h3>
	<p>
		We’re an expert crew of experienced software delivery consultants based in the Nordics. We help
		our clients optimise their software delivery flow and empower teams to deliver xxxx. We believe
		in working together with our clients to assess, coach and build ‘solutions’ (???), so that they
		ultimately retain ownership.
	</p>
</section>

<section>
	<Grid>
		<div>
			<h3>Value Stream Assessment</h3>
			<p>
				Visualize your value streams to identify waste and accelerate the flow and delivery of
				value.
			</p>
			<ButtonLink href="/work#continuous-delivery">Learn More</ButtonLink>
		</div>
		<div>
			<h3>Software Delivery Platform</h3>
			<p>
				Visualize your value streams to identify waste and accelerate the flow and delivery of
				value.
			</p>
			<ButtonLink href="/work#continuous-delivery">Learn More</ButtonLink>
		</div>
		<div>
			<h3>Team Topologies</h3>
			<p>
				Visualize your value streams to identify waste and accelerate the flow and delivery of
				value.
			</p>
			<ButtonLink href="/work#continuous-delivery">Learn More</ButtonLink>
		</div>
		<div>
			<h3>Cloud stuff</h3>
			<p>
				Visualize your value streams to identify waste and accelerate the flow and delivery of
				value.
			</p>
			<ButtonLink href="/work#continuous-delivery">Learn More</ButtonLink>
		</div>
	</Grid>
</section>

<section>
	<ClientLogos />
</section>

<section>
	<Columns reverse={true}>
		<Column class="self-start">
			<div class="flex flex-col gap-y-4">
				<h2>
					‘A key and crucial part of our Continuous Integration journey was having Verifa's
					involvement’.
				</h2>
				<h4>— David Hoslett, Siemens Mobility.</h4>
			</div>
		</Column>
		<Column class="self-start">
			<div class="flex flex-col gap-y-0">
				<div>
					<h4>Predictable software deliveries.</h4>
					<p>
						We are a small and dedicated crew who care about the value we deliver. We look to
						tomorrow's problems and solve them for you ahead of time.
					</p>
				</div>
				<div>
					<h4>Less firefighting, more improvements.</h4>
					<p>
						We don't just solve problems but teach you how to solve them, and we do it as a team.
						That's our goal and what sets us apart from other consultancy houses.
					</p>
				</div>
				<div>
					<h4>Reduce lead times and handovers.</h4>
					<p>
						We have helped international teams of all sizes, across many industries like Automotive,
						Medical, Finance, Mobile and Web.
					</p>
				</div>
				<div>
					<ButtonLink href="/clients">Our Clients</ButtonLink>
				</div>
			</div>
		</Column>
	</Columns>
</section>

<!-- <section>
	<Grid>
		<div>
			<img
				class="object-contain h-full w-full"
				src="/continuous-delivery.svg"
				alt="continuous delivery"
			/>
		</div>
		<div>
			<h3>Continuous Delivery</h3>
			<p>
				We help teams unlock their continuous release potential through workshops and coaching. We
				use our knowledge and experience, guiding teams to create processes that deliver value.
				Through self-discovery we align everyone to a common goal and help transform the way
				software is delivered.
			</p>
			<ButtonLink href="/work#continuous-delivery">Learn More</ButtonLink>
		</div>
	</Grid>
</section>

<section>
	<Grid>
		<div class="lg:order-last">
			<img
				class="object-contain h-full w-full"
				src="/cloud-architecture.svg"
				alt="cloud architecture"
			/>
		</div>
		<div>
			<h3>Cloud Architecture</h3>
			<p>
				Designing and building scalable, reliable and cost-effective cloud architectures is our
				passion. We want to make our knowledge and experience available to help teams accelerate
				their projects. The cloud ecosystem is like a jungle, and we can help you navigate through
				it.
			</p>
			<ButtonLink href="/work#cloud-architecture">Learn more</ButtonLink>
		</div>
	</Grid>
</section>

<section>
	<h1>Workshops, coaching and implementation.</h1>
</section> -->

<!-- <section>
	<Columns reverse={true}>
		<Column>
			<img
				class="object-contains h-full w-full"
				src="/everything-is-connected.svg"
				alt="everything is connected"
			/>
		</Column>
		<Column>
			<h4>
				We help teams unlock their continuous release potential through workshops and coaching. We
				have worked with many teams, who struggle with many of the same challenges.
			</h4>
			<ButtonLink href="/work#workshops">Learn More</ButtonLink>
		</Column>
	</Columns>
</section>

<MainReference />

<section>
	<Columns>
		<Column class="flex self-stretch items-center bg-v-black">
			<div class="p-20">
				<img
					class="object-contain h-full w-full"
					src="/partners/hashicorp-horizontal-white.svg"
					alt="hashicorp"
				/>
			</div>
		</Column>
		<Column class="flex self-stretch items-center">
			<h3 class="mb-0">
				We partner with those that help us deliver the best possible cloud experience to our
				customers.
			</h3>
		</Column>
	</Columns>
</section>

<section>
	<h1>Great teamwork is more than just great tools.</h1>
</section>

<section>
	<Columns reverse={true}>
		<Column>
			<img class="object-contain h-full w-full" src="/round-table.svg" alt="round table" />
		</Column>
		<Column>
			<div class="flex flex-col gap-y-0">
				<h4>
					We all share a passion for teamwork and continuous learning. We allocate 20% of our work
					time for internal projects and personal development, which helps us stay ahead of the
					curve with technology.
				</h4>
				<div>
					<ButtonLink href="/company">Learn More</ButtonLink>
				</div>
			</div>
		</Column>
	</Columns>
</section> -->

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
