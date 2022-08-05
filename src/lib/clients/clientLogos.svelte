<script lang="ts">
	import { clientLogosWhite } from '$lib/clients/clients';
	import { ScreenSize } from '$lib/screenSizes';

	$: outerWidth = 0;

	let logos = clientLogosWhite();

	// showIndex determines whether a logo should be shown based on its index
	// and the width of the screen
	$: showIndex = (index: number): boolean => {
		if (index <= 4) {
			return true;
		}
		// If outerWidth has not been initialised yet, do not show the extra element
		// as it causes a screen stutter
		if (outerWidth == 0) {
			return false;
		}
		if (index == 5 && outerWidth < ScreenSize.Large) {
			return true;
		}
		return false;
	};
</script>

<!-- Bind the window width to outerWidth -->
<svelte:window bind:outerWidth />

<div class="grid grid-cols-2 gap-0.5 md:grid-cols-6 lg:grid-cols-5">
	{#each logos as logo, index}
		{#if showIndex(index)}
			<div class="col-span-1 flex justify-center md:col-span-2 lg:col-span-1 py-8 px-8  bg-v-lilac">
				<img src={logo.image} alt={logo.name} class="h-8 w-full object-contain" />
			</div>
		{/if}
	{/each}
</div>
