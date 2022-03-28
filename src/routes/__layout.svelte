<script>
	import '../app.css';
	import Header from './_header.svelte';
	import Footer from './_footer.svelte';
	import Subscribe from '$lib/subscribe.svelte';
	import SEO from '$lib/seo/seo.svelte';
</script>

<div class="bg-gray-50 overflow-auto">
	<Header />
	<main>
		<slot />
	</main>
	<section>
		<Subscribe />
	</section>
	<Footer />
</div>

<!-- 
	This has to come last to make sure our seo store is updated correctly by
	all the possible pages above
 -->
<SEO />
