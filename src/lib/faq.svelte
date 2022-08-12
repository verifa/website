<script lang="ts">
	export let title: string;
</script>

<div class="mx-auto divide-y-2 divide-v-gray">
	{#if title}
		<h3>{title}</h3>
	{/if}
	<dl class="mt-6 space-y-6 divide-y divide-v-gray">
		<slot />
	</dl>
</div>
