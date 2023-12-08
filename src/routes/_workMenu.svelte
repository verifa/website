<script lang="ts">
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import { fly } from 'svelte/transition';
	import { headerVisible } from './store';

	const links: {
		title: string;
		url: string;
	}[] = [
		{
			title: 'Value Stream Assessment',
			url: '/work/value-stream-assessment'
		},
		{
			title: 'Software Delivery Platforms',
			url: '/work/software-delivery-platforms'
		},
		{
			title: 'Team Topologies',
			url: '/work/team-topologies'
		},
		{
			title: 'Cloud Architecture',
			url: '/work/cloud-architecture'
			//description: 'Designing cloud solutions and facilitating adoption',
		},
		{
			title: 'How We Do It',
			url: '/work/implementation'
		}
	];

	let workMenuDiv;
	let shown = false;

	// Subscribe to headerVisible store so that if the header disappears and we
	// have the menu open, then the menu is also closed
	headerVisible.subscribe((visible) => {
		if (!visible && shown) {
			shown = false;
		}
	});

	function hide() {
		shown = false;
	}

	function toggleShow() {
		shown = !shown;
	}

	onMount(() => {
		const handleOutsideClick = (event) => {
			if (shown && !workMenuDiv.contains(event.target)) {
				hide();
			}
		};

		const handleEscape = (event) => {
			if (shown && event.key === 'Escape') {
				hide();
			}
		};

		// add events when element is added to the DOM
		document.addEventListener('click', handleOutsideClick, false);
		document.addEventListener('keyup', handleEscape, false);

		// remove events when element is removed from the DOM
		return () => {
			document.removeEventListener('click', handleOutsideClick, false);
			document.removeEventListener('keyup', handleEscape, false);
		};
	});
</script>

<div bind:this={workMenuDiv}>
	<button
		on:click={toggleShow}
		class="text-xl py-2 text-v-black hover:text-v-lilac focus:outline-none font-medium border-b-2 border-v-black transition-all ease-in-out duration-150 {'/work/' ===
		$page.url.pathname
			? 'border-solid'
			: 'border-transparent'}"
	>
		<div class="flex gap-x-2 items-center ">
			What we do
			{#if shown}
				<svg
					xmlns="http://www.w3.org/2000/svg"
					class="h-5 w-5 pointer-events-none"
					fill="none"
					viewBox="0 0 24 24"
					stroke="currentColor"
					stroke-width="2"
				>
					<path stroke-linecap="round" stroke-linejoin="round" d="M5 15l7-7 7 7" />
				</svg>
			{:else}
				<svg
					xmlns="http://www.w3.org/2000/svg"
					class="h-5 w-5 pointer-events-none"
					fill="none"
					viewBox="0 0 24 24"
					stroke="currentColor"
					stroke-width="2"
				>
					<path stroke-linecap="round" stroke-linejoin="round" d="M19 9l-7 7-7-7" />
				</svg>
			{/if}
		</div>
	</button>
	{#if shown}
		<div
			in:fly={{ y: -35, duration: 250 }}
			out:fly={{ y: -35, duration: 250 }}
			class="absolute z-10 transform -translate-x-1/3 mt-3 px-2 max-w-md"
		>
			<div class="shadow-lg ring-1 ring-v-black ring-opacity-5 overflow-hidden">
				<div class="relative grid gap-6 bg-v-white px-5 py-6 sm:gap-8 sm:p-8">
					{#each links as link}
						<a
							on:click={hide}
							href={link.url}
							class="-m-3 p-3 group block hover:bg-v-gray hover:bg-opacity-10 transition ease-in-out duration-150"
						>
							<span class="-p-3 mb-0 group-hover:text-v-lilac">{link.title}</span>
						</a>
					{/each}
				</div>
			</div>
		</div>
	{/if}
</div>
