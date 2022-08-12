<script lang="ts">
	import ChevronRight from './icons/chevronRight.svelte';

	export let question: string;

	let show: boolean = false;
</script>

<div class="pt-6">
	<dt>
		<!-- Expand/collapse question button -->
		<button
			type="button"
			class="text-left w-full flex justify-between items-start text-v-black"
			aria-controls="faq-0"
			aria-expanded="false"
			on:click={() => (show = !show)}
		>
			<h4 class="mb-0">{question}</h4>
			<span class="ml-6 h-7 flex items-center">
				<ChevronRight class="{show ? 'rotate-90' : 'rotate-0'} h-6 w-6 transform" />
			</span>
		</button>
	</dt>
	{#if show}
		<dd class="mt-2 pr-12" id="faq-0">
			<slot />
		</dd>
	{/if}
</div>
