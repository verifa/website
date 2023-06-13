<script lang="ts">
	export { className as class };

	let className: string = '';
</script>

<div
	class="mx-auto max-w-xl lg:max-w-none grid grid-cols-1 lg:grid-cols-2 gap-y-8 md:gap-y-16 lg:gap-y-24 lg:px-8 lg:gap-x-24 {className}"
>
	<slot />
</div>
