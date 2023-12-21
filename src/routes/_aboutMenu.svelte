<script lang="ts">
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import { fly } from 'svelte/transition';
	import { headerVisible } from './store';

	const links: {
		title: string;
		description: string;
		url: string;
	}[] = [
		{
			title: 'Company',
			description: 'About us, our values and our strategy',
			url: '/company/'
		},
		{
			title: 'Crew',
			description: 'The wonderful crew that make up Verifa',
			url: '/crew/'
		},
		{
			title: 'Clients',
			description: 'Who we continuously deliver for',
			url: '/clients/'
		}
	];

	let workMenuDiv: any;
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

	function isActive(pathname: string): boolean {
		return links.findIndex((link) => link.url === pathname) >= 0;
	}

	onMount(() => {
		const handleOutsideClick = (event: any) => {
			if (shown && !workMenuDiv.contains(event.target)) {
				hide();
			}
		};

		const handleEscape = (event: any) => {
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
		class="text-xl py-2 text-v-black hover:text-v-lilac focus:outline-none font-medium border-b-2 transition-all ease-in-out duration-150 {isActive(
			$page.url.pathname
		)
			? 'border-v-black'
			: 'border-v-black/0'}"
	>
		<div class="flex gap-x-2 items-center">
			About us
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
							<span class="pb-3 group-hover:text-v-lilac">{link.title}</span>
							<br />
							<span class="mb-0 font-normal text-base group-hover:text-v-lilac">
								{link.description}
							</span>
						</a>
					{/each}
				</div>
			</div>
		</div>
	{/if}
</div>
