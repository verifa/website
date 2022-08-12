<script lang="ts">
	import { children } from 'svelte/internal';

	interface Link {
		text: string;
		url: string;
		children?: Link[];
		showChildren?: boolean;
	}

	let menuVisible = false;

	let links: Link[] = [
		{
			text: 'What we do',
			url: '/work/',
			children: [
				{
					text: 'Continuous Delivery',
					url: '/work#continuous-delivery'
				},
				{
					text: 'Cloud Architecture',
					url: '/work#cloud-architecture'
				},
				{
					text: 'Workshops',
					url: '/work#workshops'
				},
				{
					text: 'Coaching',
					url: '/work#coaching'
				},
				{
					text: 'Implementation',
					url: '/work#implementation'
				}
			],
			showChildren: false
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

	function hideMenu() {
		menuVisible = false;
		links.forEach((link) => {
			collapse(link);
		});
	}

	function collapse(link) {
		if (link.children) {
			link.showChildren = false;
			link.children.forEach((child) => {
				collapse(child);
			});
		}
	}

	function handleClick(index: number) {
		let link = links[index];
		if (link.children) {
			links[index].showChildren = !link.showChildren;
		} else {
			hideMenu();
		}
	}
</script>

<div class="flex items-center md:hidden">
	<button
		type="button"
		class="p-2 text-v-black hover:bg-v-white active:bg-v-white focus:bg-v-white focus:outline-none"
		aria-expanded="false"
		on:click={() => (menuVisible = !menuVisible)}
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

{#if menuVisible}
	<div class="absolute -z-1 top-0 inset-x-0 p-2 md:hidden">
		<div class="bg-v-white ring-2 ring-v-black ring-opacity-5 overflow-hidden">
			<div class="px-5 pt-4 flex items-center justify-between">
				<div>
					<a href="/" on:click={() => hideMenu()}>
						<img class="h-8 w-28" src="/verifa-logo.svg" alt="" />
					</a>
				</div>
				<button
					type="button"
					class="bg-v-white rounded-md p-2 inline-flex items-center justify-center text-v-black hover:bg-v-white active:bg-v-white focus:bg-v-white focus:outline-none"
					on:click={() => hideMenu()}
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
			<div class="px-5 z-1 flex flex-col gap-y-3 py-3">
				{#each links as link, index}
					<a
						href={link.url}
						class="z-1 block rounded-md text-xl text-v-black hover:text-v-lilac flex gap-x-2 items-center"
						on:click={() => handleClick(index)}
						>{link.text}
						{#if link.children}
							{#if link.showChildren}
								<!-- chevron pointing up -->
								<svg
									xmlns="http://www.w3.org/2000/svg"
									class="h-5 w-5"
									fill="none"
									viewBox="0 0 24 24"
									stroke="currentColor"
									stroke-width="2"
								>
									<path stroke-linecap="round" stroke-linejoin="round" d="M5 15l7-7 7 7" />
								</svg>
							{:else}
								<!-- chevron pointing down -->
								<svg
									xmlns="http://www.w3.org/2000/svg"
									class="h-5 w-5"
									fill="none"
									viewBox="0 0 24 24"
									stroke="currentColor"
									stroke-width="2"
								>
									<path stroke-linecap="round" stroke-linejoin="round" d="M19 9l-7 7-7-7" />
								</svg>
							{/if}
						{/if}
					</a>
					{#if link.children && link.showChildren}
						<div class="px-4 z-1 flex flex-col gap-y-3">
							{#each link.children as child}
								<a
									href={child.url}
									class="z-1 block rounded-md text-base text-v-black hover:text-v-lilac"
									on:click={() => hideMenu()}>{child.text}</a
								>
							{/each}
						</div>
					{/if}
				{/each}
			</div>
		</div>
	</div>
{/if}
