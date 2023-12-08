<script lang="ts">
	import ButtonLink from '$lib/buttonLink.svelte';

	import Column from '$lib/column.svelte';

	import Columns from '$lib/columns.svelte';

	import CtaButton from '$lib/ctaButton.svelte';
	import Faq from '$lib/faq.svelte';
	import FaqItem from '$lib/faqItem.svelte';
	import PostGrid from '$lib/posts/postGrid.svelte';
	import type { Post, PostsData } from '$lib/posts/posts';

	import { seo } from '$lib/seo/store';
	import type { PageData } from '.svelte-kit/types/src/routes/$types';
	import Stages from './_stages.svelte';
	import Waste from './_waste.svelte';

	seo.reset();
	$seo.title = 'Continuous Delivery Workshop';
	$seo.description =
		'Our Continuous Delivery Workshops help teams unlock their Continuous Delivery potential. We use Value Stream Mapping as our method for helping teams understand their current release process and create a backlog of improvements.';

	export let data: PageData;
	const posts: Post[] = data.posts.posts;
</script>

<section>
	<h1>Continuous Delivery Workshop.</h1>
</section>
<section>
	<h3>
		Our Continuous Delivery workshops help teams unlock their Continuous Delivery potential. We will
		help you understand your release process and create a backlog of improvements.
	</h3>
</section>
<section>
	<Columns>
		<Column>
			<img
				class="object-contain w-full h-full"
				src="/work/cd-workshop/mock-vsm-sample.png"
				alt="mock-vsm-sample"
			/>
		</Column>
		<Column class="self-start">
			<div class="flex flex-col gap-y-16">
				<div>
					<h4>Value Stream Map</h4>
					<span class="mb-0">
						A Value Stream Map is a visual tool showing the series of events that take a product or
						service from the beginning of a process until it reaches its end (i.e. a Value Stream).
					</span>
				</div>
				<div>
					<h4>Value Stream Mapping</h4>
					<span class="mb-0">
						Value Stream Mapping (VSM) is a method for defining a Value Stream and analysing the
						current flow in order to remove waste and to derive a more lean and continuous workflow.
					</span>
				</div>
			</div>
		</Column>
	</Columns>
</section>
<section>
	<h2>
		‘Verifa helped lead discussions around many topics which was useful in gaining a common
		understanding of what we can do to improve’.
	</h2>
	<h4>— Olli Suihkonen, Visma Software Oy.</h4>
	<ButtonLink href="/clients/value-stream-mapping-with-visma/">Case study</ButtonLink>
</section>
<section>
	<Columns reverse={true}>
		<Column>
			<img
				class="object-contains h-full w-full"
				src="/everything-is-connected.svg"
				alt="everything is connected"
			/>
		</Column>
		<Column>
			<div class="flex flex-col gap-y-16">
				<div>
					<h4>The Goal</h4>
					<span class="mb-0">
						The goal of our Value Stream Mapping workshop is to build a common understanding of your
						release process and highlight waste as effectively as possible.
					</span>
				</div>
				<div>
					<h4>Finding Waste</h4>
					<span class="mb-0">
						Waste is any activity that consumes resources, but does not add value to the product or
						service a customer receives. Less waste means more flow.
					</span>
				</div>
			</div>
		</Column>
	</Columns>
</section>
<section>
	<h2>Types of waste.</h2>
</section>
<section>
	<Waste />
</section>

<section>
	<h2>Workshop stages.</h2>
</section>
<section>
	<Stages />
</section>

<section>
	<h2>Continuous Delivery for Continuous Delivery.</h2>
</section>
<section>
	<h4>
		The idea of implementing change in your software delivery is not too dissimilar from
		implementing change in your software itself, and should be done in a <span class="italic"
			>continuous</span
		> way. View it as using Continuous Delivery for improving your Continuous Delivery.
	</h4>
</section>

<section>
	<Faq title="Continuous Delivery principles">
		<FaqItem question="Infinite Game"
			><span class="mb-0">
				When you fix your biggest problem, you still have a biggest problem.
			</span></FaqItem
		>
		<FaqItem question="Plan vs Do"
			><span class="mb-0">
				Set goals, but don’t spend too much time planning the details. You learn as you go, allowing
				you to adjust course better than any plan ever would. Don’t plan for more than a couple of
				weeks.
			</span></FaqItem
		>
		<FaqItem question="What vs How"
			><span class="mb-0">
				It’s not about how to do it in the tools, it’s about what you want to achieve. Don’t let the
				tools control your process.
			</span></FaqItem
		>
		<FaqItem question="Problem vs Symptom"
			><span class="mb-0">
				Symptoms are easy to see and often easy to patch, but workarounds don’t solve the underlying
				problem, they only serve to obfuscate it. Look for the underlying problem.
			</span></FaqItem
		>
		<FaqItem question="Comfort Zone"
			><span class="mb-0">
				Technology will always change and old experts will become novices. It is better to learn
				something new and fail than to stay in the comfort zone risking the business.
			</span></FaqItem
		>
		<FaqItem question="KISS"
			><span class="mb-0">
				Keep it simple stupid; a good solution removes waste. Use simple tools, reduce complexity
				and maintenance costs. A simple solution is better than a clever solution.
			</span></FaqItem
		>
		<FaqItem question="Build - Measure - Learn"
			><span class="mb-0">
				Avoid big bangs as much as possible. Test, experiment, release gradually and be ready to
				revert. Build this in as part of the process.
			</span></FaqItem
		>
	</Faq>
</section>

<section>
	<h2>Want to find out for yourself?</h2>
</section>
<section>
	<h4>
		We want to help you unlock your Continuous Delivery potential and we believe in using Value
		Stream Mapping as a method for doing so. We are also a very experienced team of engineers which
		means we can also help you after the workshop with coaching and implementation.
	</h4>
</section>
<section>
	<CtaButton text={'Want to get started?'} />
</section>
<section>
	<h2>Suggested reading.</h2>
</section>
<section>
	<PostGrid {posts} showBadges={true} showPreview={false} />
</section>
