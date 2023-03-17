<script>
	import '../app.css';
	import '../fonts.css';
	import Header from './_header.svelte';
	import Footer from './_footer.svelte';
	import Subscribe from '$lib/subscribe.svelte';
	import SEO from '$lib/seo/seo.svelte';
</script>

<div
	class="bg-v-white prose lg:prose-lg xl:prose-xl 2xl:prose-2xl max-w-none prose-xl-img:m-0 prose-2xl-img:m-0 prose-h2:leading-tight"
>
	<Header />
	<main class="bg-v-white mx-auto max-w-7xl px-8 sm:px-16 pt-8 sm:pt-12">
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
