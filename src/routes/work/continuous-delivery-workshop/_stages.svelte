<script lang="ts">
	let selected: number = 0;

	const stages: string[] = ['Preparation', 'Workshop', 'Implementation'];

	const stageDescriptions: { title: string; description: string }[][] = [
		// Stage 1
		[
			{
				title: 'Identify owner',
				description:
					'The workshop owner has the mandate to drive the workshop and own the results from your side.'
			},
			{
				title: 'Define scope',
				description:
					'Define the start and end of the Value Streams. Commonly we conduct planning interviews to set a structure for the workshop.'
			},
			{
				title: 'Schedule workshop(s)',
				description:
					'Schedule the workshop(s), make it clear whether it is remote or in-person and all the equipment and software is available.'
			}
		],
		// Stage 2
		[
			{
				title: 'Motivation',
				description:
					"An inspirational presentation that will present the idea of Continuous Delivery so that it's clear why we are doing this."
			},
			{
				title: 'Value Streams',
				description:
					"For the majority of the workshop we will facilitate the Value Stream's organic growth on the whiteboard as we go over various topics."
			},
			{
				title: 'Waste deep dives',
				description:
					'We dive further into specific topics, e.g. particularly painful processes, requirements or initiatives. These may be planned beforehand or come up during the workshop.'
			}
		],
		// Stage 3
		[
			{
				title: 'Report',
				description:
					'We prepare a report summarizing the workshop, the pain points (i.e. waste) and our own suggestions that we can think of. This is the first version of the backlog.'
			},
			{
				title: 'Feedback',
				description:
					'A discussion about the workshop and report. It is important that you understand the report and are comfortable with the implementation.'
			},
			{
				title: 'Build - Measure - Learn',
				description:
					'This is where the real work of implementing change begins. And it is so important to do it gradually with feedback.'
			}
		]
	];
</script>

<nav>
	<ol class="gap-y-4 md:flex md:gap-y-0 md:gap-x-8 list-none mb-0">
		{#each stages as stage, index}
			<li class="md:flex-1">
				<button
					on:click={() => (selected = index)}
					class="group w-full pl-4 py-2 flex flex-col border-l-4 md:pl-0 md:pt-4 md:pb-0 md:border-l-0 md:border-t-4 hover:border-v-lilac {selected ===
					index
						? 'border-v-lilac'
						: 'border-v-gray'}"
				>
					<p
						class="{selected === index
							? 'text-v-lilac'
							: 'text-v-gray'} mb-0 group-hover:text-v-lilac"
					>
						Stage {index + 1}
					</p>
					<h4 class="mb-0">{stage}</h4>
				</button>
			</li>
		{/each}
	</ol>
</nav>
<div class="my-8">
	<dl class="flex flex-col">
		{#each stageDescriptions[selected] as item}
			<div class="flex flex-col gap-y-2">
				<dt class="flex items-center gap-x-4">
					<img class="h-4 w-4 object-contain" src="/logo-element.png" alt="logo-ement" />
					<h4 class="mb-0 group-hover:text-v-lilac">{item.title}</h4>
				</dt>
				<dd>
					<p class="pl-8">
						{item.description}
					</p>
				</dd>
			</div>
		{/each}
	</dl>
</div>
