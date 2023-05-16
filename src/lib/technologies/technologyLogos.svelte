<script lang="ts">
	import { technologyLogosWhite } from '$lib/technologies/technologies';
	import { ScreenSize } from '$lib/screenSizes';

	$: outerWidth = 0;

	let logos = technologyLogosWhite();

	// showIndex determines whether a logo should be shown based on its index
	// and the width of the screen
	$: showIndex = (index: number): boolean => {
		if (index <= 5) {
			return true;
		}
		return false;
	};
</script>

<!-- Bind the window width to outerWidth -->
<svelte:window bind:outerWidth />

<div class="grid grid-cols-6 gap-0.5 md:grid-cols-6 lg:grid-cols-6">
	{#each logos as logo, index}
		{#if showIndex(index)}
			<div class="col-span-3 flex justify-center md:col-span-3 lg:col-span-3 py-8 px-8 bg-v-lilac">
				<img src={logo.image} alt={logo.name} class="h-12 w-full object-contain" />
			</div>
		{/if}
	{/each}
</div>
