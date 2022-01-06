<script lang="ts">
	export let quote;
</script>

<h3>‘{quote.text}’</h3>
<h4>— {quote.person}</h4>
<div class="mt-10">
	<img class="object-contain object-left max-h-20 w-2/3" src={quote.logo} alt={quote.logo} />
</div>
