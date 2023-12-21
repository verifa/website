<script lang="ts">
	import ExclamationTriangle from '$lib/icons/exclamationTriangle.svelte';
	import InfoCircle from '$lib/icons/infoCircle.svelte';
	import LightBulb from '$lib/icons/lightBulb.svelte';
	import { error } from '@sveltejs/kit';
	import type { ComponentType } from 'svelte';

	export let type = '';
	export let title = '';

	let defaultTitle = '';
	let bgColor = '';
	let icon: ComponentType;

	switch (type) {
		case 'idea':
			defaultTitle = 'Idea';
			bgColor = 'bg-v-lilac';
			icon = LightBulb;
			break;
		case 'info':
			defaultTitle = 'Info';
			bgColor = 'bg-v-green';
			icon = InfoCircle;
			break;
		case 'warning':
			defaultTitle = 'Warning';
			bgColor = 'bg-v-pink';
			icon = ExclamationTriangle;
			break;
		default:
			error(400, 'unknown admonition type ' + type);
	}
</script>

<div class="my-4 shadow-sm">
	<div class="{bgColor} py-2 px-4 flex gap-4 items-center">
		<svelte:component this={icon} class="text-v-black w-6 h-6" />
		<h4 class="!m-0">{title === '' ? defaultTitle : title}</h4>
	</div>
	<div class="{bgColor} bg-opacity-30 py-2 px-4">
		<slot />
	</div>
</div>
