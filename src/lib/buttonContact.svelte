<script>
	export let href;
</script>

<div>
	<a
		{href}
		class="inline-flex px-7 py-5 border-v-black border-8 text-v-black bg-v-white hover:bg-v-gray"
	>
		<slot />
	</a>
</div>
