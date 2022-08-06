<script lang="ts">
	import { page } from '$app/stores';
	import { onMount } from 'svelte';

	interface Link {
		text: string;
		url: string;
	}
	const navLinks: Link[] = [
		{
			text: 'What we do',
			url: '/work/'
		},
		{
			text: 'About us',
			url: '/company/'
		},
		{
			text: 'Crew',
			url: '/crew/'
		},
		{
			text: 'Clients',
			url: '/clients/'
		},
		{
			text: 'Careers',
			url: '/careers/'
		},
		{
			text: 'Blog',
			url: '/blog/'
		},
		{
			text: 'Contact',
			url: '/contact/'
		}
	];

	let showMenu: boolean = false;
	let mobileMenu: HTMLElement = null;

	let previousY: number;
	let currentY: number;
	let clientHeight: number;

	const isScrollingUp = (y: number) => {
		const scrollUp = !previousY || previousY < y ? false : true;
		previousY = y;

		return scrollUp;
	};

	$: scrollUp = isScrollingUp(currentY);
	$: offscreen = !scrollUp && currentY > clientHeight * 4;

	onMount(() => {
		const handleOutsideClick = (event) => {
			if (showMenu && !mobileMenu.contains(event.target)) {
				showMenu = false;
			}
		};

		const handleEscape = (event) => {
			if (showMenu && event.key === 'Escape') {
				showMenu = false;
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

<svelte:window bind:scrollY={currentY} />

<header
	class="sticky top-0 py-6 px-8 sm:px-16 bg-v-white/50 backdrop-blur-sm transition-transform ease-in"
	class:motion-safe:-translate-y-full={offscreen}
	bind:clientHeight
>
	<div bind:this={mobileMenu}>
		<nav class="mx-auto flex items-center justify-between gap-x-8" aria-label="Global">
			<a class="flex-none" href="/">
				<span class="sr-only">verifa</span>
				<img class="h-8 w-full md:h-12 " src="/verifa-logo.svg" alt="verifa-logo" />
			</a>

			<div class="hidden md:flex md:items-center md:space-x-10 md:flex-wrap">
				{#each navLinks as link, index}
					<a
						href={link.url}
						class="text-xl py-2 text-v-black hover:text-v-lilac font-medium border-b-2 border-v-black transition-all ease-in-out duration-150  {link.url ===
						$page.url.pathname
							? 'border-solid'
							: 'border-transparent'}">{link.text}</a
					>
				{/each}
			</div>
			<!-- Mobile menu button -->
			<div class="flex items-center md:hidden">
				<button
					type="button"
					class="p-2 text-v-black"
					aria-expanded="false"
					on:click={() => (showMenu = !showMenu)}
				>
					<span class="sr-only">Open main menu</span>
					<!-- Heroicon name: outline/menu -->
					<svg
						class="h-6 w-6"
						xmlns="http://www.w3.org/2000/svg"
						fill="none"
						viewBox="0 0 24 24"
						stroke="currentColor"
						aria-hidden="true"
					>
						<path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M4 6h16M4 12h16M4 18h16"
						/>
					</svg>
				</button>
			</div>
		</nav>

		{#if showMenu}
			<div class="absolute -z-1 top-0 inset-x-0 p-2 md:hidden">
				<div class="bg-v-white ring-2 ring-v-black ring-opacity-5 overflow-hidden">
					<div class="px-5 pt-4 flex items-center justify-between">
						<div>
							<a href="/" on:click={() => (showMenu = false)}>
								<img class="h-8 w-full" src="/verifa-logo.svg" alt="" />
							</a>
						</div>
						<button
							type="button"
							class="bg-v-white rounded-md p-2 inline-flex items-center justify-center text-v-black hover:text-gray-700 hover:bg-gray-100 focus:outline-none"
							on:click={() => (showMenu = false)}
						>
							<span class="sr-only">Close menu</span>
							<!-- Heroicon name: outline/x -->
							<svg
								class="h-6 w-6"
								xmlns="http://www.w3.org/2000/svg"
								fill="none"
								viewBox="0 0 24 24"
								stroke="currentColor"
								aria-hidden="true"
							>
								<path
									stroke-linecap="round"
									stroke-linejoin="round"
									stroke-width="2"
									d="M6 18L18 6M6 6l12 12"
								/>
							</svg>
						</button>
					</div>
					<div class="px-5 pt-2 pb-3 z-1">
						{#each navLinks as link}
							<a
								href={link.url}
								class="z-1 block py-2 rounded-md text-xl text-v-black hover:text-v-lilac"
								on:click={() => (showMenu = false)}>{link.text}</a
							>
						{/each}
					</div>
				</div>
			</div>
		{/if}
	</div>
</header>
