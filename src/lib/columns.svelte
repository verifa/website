<script lang="ts">
	export { className as class };

	export let reverse: boolean = false;
	export let fullWidth: boolean = false;
	let className: string = '';
	// export let xPad: boolean = true;
</script>

<div
	class="max-w-xl mx-auto flex flex-col gap-y-8 sm:gap-y-16 lg:flex-row lg:max-w-none lg:items-center {fullWidth
		? 'px-0'
		: 'lg:px-8'} lg:gap-x-24 
	{className} {reverse ? 'lg:flex-row-reverse' : ''}"
>
	<slot />
</div>
