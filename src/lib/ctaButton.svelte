<script lang="ts">
	import Column from './column.svelte';
	import Columns from './columns.svelte';

	export let text: string = 'Want to find out how we can help you?';
	export let contact: string = 'Get in touch.';
	export let fullScreen: boolean = true;
</script>

{#if fullScreen}
	<Columns>
		<Column />
		<Column>
			<a href="/contact" class="bg-v-black px-8 py-8 flex flex-col gap-y-10">
				<h3 class="text-v-white">{text}</h3>
				<h3 class="text-v-white text-right mb-0">{contact}</h3>
			</a>
		</Column>
	</Columns>
{:else}
	<a href="/contact" class="bg-v-black px-8 py-8 flex flex-col gap-y-10">
		<h3 class="text-v-white">{text}</h3>
		<h3 class="text-v-white text-right mb-0">{contact}</h3>
	</a>
{/if}
