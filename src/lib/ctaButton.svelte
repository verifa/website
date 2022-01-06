<script>
	import Column from './column.svelte';
	import Columns from './columns.svelte';
</script>

<Columns>
	<Column />
	<Column>
		<a href="/contact" class="bg-v-black px-8 py-8 flex flex-col gap-y-10">
			<h3 class="text-v-white">Want to find out how we can help you?</h3>
			<h3 class="text-v-white text-right mb-0">Get in touch.</h3>
		</a>
	</Column>
</Columns>
