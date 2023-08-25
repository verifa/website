<script lang="ts">
	import { fly } from 'svelte/transition';
	import ChevronRight from '$lib/icons/chevronRight.svelte';

	const wastes: {
		name: string;
		description: string;
		img: string;
	}[] = [
		{
			name: 'Conflict',
			description:
				'Conflicts slows down or even hinders the value flow. These can both be high or low level conflicts, ranging from conflicting interests and priorities, all the way down to reappearing merge conflicts.',
			img: '/work/cd-workshop/waste/conflict.png'
		},
		{
			name: 'Handover',
			description:
				'The drawback of handovers is loss of information and breaks in the flow. Handovers can happen between teams, tools and even team members.',
			img: '/work/cd-workshop/waste/handover.png'
		},
		{
			name: 'Manual work',
			description:
				'Manual work is prone to inconsistency and human error. Plus, a machine probably does it a lot faster anyway.',
			img: '/work/cd-workshop/waste/manual-work.png'
		},
		{
			name: 'Legacy',
			description:
				'Previously developed parts - processes, scripts or tools - that are no longer as compatible may need to be refactored, replaced or remade.',
			img: '/work/cd-workshop/waste/legacy.png'
		},
		{
			name: 'Late discovery',
			description:
				'Discovery of a flaw or fault in the process that forces you to return to a previous step. The later in the process, the higher the impact.',
			img: '/work/cd-workshop/waste/late-discovery.png'
		},
		{
			name: 'Unplanned work',
			description:
				'This is any work that comes as a surprise for the team and needs to be done “ASAP”. This could be urgent bugs, scope creep or new requirements.',
			img: '/work/cd-workshop/waste/unplanned.png'
		},
		{
			name: 'Queue',
			description:
				'A queue is a break in the flow that is predictable and can be planned around. Examples of these are events that happen at a certain time or process steps where “it’s your turn”.',
			img: '/work/cd-workshop/waste/queue.png'
		},
		{
			name: 'Waiting',
			description:
				'Waiting is when the break in the flow isn’t predictable; waiting for other processes, teams, team members, resources to be available, etc. ',
			img: '/work/cd-workshop/waste/waiting.png'
		}
	];

	let showWastes = Array(wastes.length).fill(false);
</script>

<dl class="grid grid-cols-1 gap-y-2 md:grid-cols-2 md:gap-y-6 md:gap-x-8">
	{#each wastes as waste, index}
		<div class="flex flex-col gap-y-2">
			<button on:click={() => (showWastes[index] = !showWastes[index])}>
				<dt class="flex items-center gap-x-4 my-4">
					<img class="object-contain w-12 h-12" src={waste.img} alt="waste-type" />
					<div class="group flex items-center gap-x-4">
						<h4 style="margin:0;" class="mb-0 group-hover:text-v-lilac">{waste.name}</h4>
						<ChevronRight
							class="h-6 w-6 group-hover:text-v-lilac {showWastes[index]
								? 'rotate-90'
								: 'rotate-0'} "
						/>
					</div>
				</dt>
			</button>
			{#if showWastes[index]}
				<dd in:fly={{ duration: 250 }}>
					<p class="pl-16 mb-0">
						{waste.description}
					</p>
				</dd>
			{/if}
		</div>
	{/each}
</dl>
