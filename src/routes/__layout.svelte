<script>
	import '../app.css';
	import Header from './_header.svelte';
	import Footer from './_footer.svelte';
	import Subscribe from '$lib/subscribe.svelte';
	import SEO from '$lib/seo/seo.svelte';
</script>

<SEO />
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
