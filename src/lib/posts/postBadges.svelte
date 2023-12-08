<script lang="ts">
	export let type: string;
	export let tags: string[];
</script>

<div class="flex flex-col space-y-4">
	{#if type}
		<div>
			<span class="inline-flex items-center px-3 py-0.5 bg-v-lilac">
				<span class="m-0 capitalize text-v-white">{type}</span>
			</span>
		</div>
	{/if}

	{#if tags}
		<!-- 
			Margin allows for space between rows when elemnts wrap.
			Negative margin is to negate the margin for the first row
		 -->
		<div class="-my-2 flex flex-wrap gap-x-4">
			{#each tags as tag}
				<span class="inline-flex items-center my-2 px-3 py-0.5 bg-v-gray">
					<span class="m-0 capitalize text-v-white">{tag}</span>
				</span>
			{/each}
		</div>
	{/if}
</div>
