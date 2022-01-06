<script lang="ts">
	import { page } from '$app/stores';
	import { onMount } from 'svelte';

	interface Link {
		text: string;
		url: string;
	}
	const links: Link[] = [
		{
			text: 'What we do',
			url: '/work'
		},
		{
			text: 'About us',
			url: '/company'
		},
		{
			text: 'Clients',
			url: '/clients'
		},
		{
			text: 'Careers',
			url: '/careers'
		},
		{
			text: 'Blog',
			url: '/blog'
		},
		{
			text: 'Contact',
			url: '/contact'
		}
	];

	let showMenu: boolean = false;
	let mobileMenu: HTMLElement = null;

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

<div class="relative mb-0 sm:mb-8 md:mb-16 lg:mb-16">
	<div class="hidden sm:block sm:absolute sm:inset-y-0 sm:h-full sm:w-full" aria-hidden="true">
		<div class="relative h-full max-w-7xl mx-auto" />
	</div>

	<div class="relative pt-6 px-8 sm:px-16 pb-12">
		<div bind:this={mobileMenu}>
			<div class="mx-auto">
				<nav
					class="relative flex items-center justify-between sm:h-10 md:justify-end"
					aria-label="Global"
				>
					<div class="flex items-center flex-1 md:absolute md:inset-y-0 md:left-0">
						<div class="flex items-center justify-between w-full md:w-auto">
							<a href="/">
								<span class="sr-only">verifa</span>
								<img class="mt-2 h-8 w-auto sm:h-12" src="/verifa-logo.svg" alt="" />
							</a>
							<div class="-mr-2 flex items-center md:hidden">
								<button
									type="button"
									class="bg-gray-50 rounded-md p-2 inline-flex items-center justify-center text-v-black hover:text-gray-700 hover:bg-gray-100 focus:outline-none "
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
						</div>
					</div>

					<div class="hidden md:flex md:space-x-10">
						{#each links as link}
							<a
								href={link.url}
								class="text-xl text-v-black hover:text-gray-900 font-normal {link.url ===
								$page.url.pathname
									? 'border-b-2 border-v-black'
									: ''}">{link.text}</a
							>
						{/each}
					</div>
				</nav>
			</div>

			{#if showMenu}
				<div class="absolute -z-1 top-0 inset-x-0 p-2 md:hidden">
					<div class="bg-v-white ring-2 ring-v-black ring-opacity-5 overflow-hidden">
						<div class="px-5 pt-4 flex items-center justify-between">
							<div>
								<a href="/" on:click={() => (showMenu = false)}>
									<img class="h-8 w-auto" src="/verifa-logo.svg" alt="" />
								</a>
							</div>
							<div class="-mr-2">
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
						</div>
						<div class="px-2 pt-2 pb-3 z-1">
							{#each links as link}
								<a
									href={link.url}
									class="z-1 block px-3 py-2 rounded-md text-xl text-v-black hover:text-gray-700 hover:bg-gray-50"
									on:click={() => (showMenu = false)}>{link.text}</a
								>
							{/each}
						</div>
					</div>
				</div>
			{/if}
		</div>
	</div>
</div>
