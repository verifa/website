<script>
	export let href;
	export { className as class };
	let className = '';
</script>

<div class="my-5">
	<a
		{href}
		class="inline-flex px-5 py-3 border border-transparent text-v-white bg-v-black hover:bg-v-lilac {className}"
	>
		<slot />
	</a>
</div>
