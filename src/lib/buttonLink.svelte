<script>
	export let href;
</script>

<div class="my-5">
	<a
		{href}
		class="inline-flex px-5 py-3 border border-transparent text-v-white bg-v-black hover:bg-v-lilac"
	>
		<slot />
	</a>
</div>
