<script>
	import '../app.css';
	import '../fonts.css';
	import Header from './_header.svelte';
	import Footer from './_footer.svelte';
	import Subscribe from '$lib/subscribe.svelte';
	import SEO from '$lib/seo/seo.svelte';
</script>

<div class="bg-v-white">
	<Header />
	<main class="bg-v-white mx-auto max-w-7xl px-8 sm:px-16 pt-12 sm:pt-20 md:pt-28">
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
