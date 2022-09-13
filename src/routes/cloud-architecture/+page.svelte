<script lang="ts">
	import CtaButton from '$lib/ctaButton.svelte';
	import Grid from '$lib/grid.svelte';
	import { seo } from '$lib/seo/store';
	seo.reset();
	$seo.title = 'Cloud Architecture';
	// TODO: description
	$seo.description = 'TODO: something about cloud architecture';
	$seo.image.url = '/cloud-architecture.svg';
	$seo.image.alt = 'cloud-architecture';
</script>

<section id="cloud-architecture">
	<h1>Cloud Architecture.</h1>
</section>
<section>
	<h3>
		Architecting for the cloud requires a lot of knowledge about different things. We want to make
		our knowledge and experience that we have gathered useful in our projects by sharing it with the
		teams that we work with through design sessions and workshops. Getting the right architecture
		and plan for the cloud give you a great leap in the right direction, so let us help you build an
		awesome cloud architecture.
	</h3>
</section>
<section>
	<Grid>
		<div class="self-start">
			<h3>Cloud Adoption.</h3>
			<p class="mb-0">
				We help teams adopt cloud (public or private) by helping in the architecting process and
				coaching teams with best practices (such as Infrastructure as Code). Our goal is to reduce
				the amount of time it takes for you to be benefiting from the many wonderful things that the
				cloud delivers.
			</p>
		</div>
		<div class="self-start">
			<h3>Cloud Operating Model.</h3>
			<p class="mb-0">
				A Cloud Operating Model defines how we operate the cloud, and typically for us, it defines
				how we provision, secure, connect, deliver and run things in the cloud. We can help you
				unlock your cloud operating model by applying our knowledge and experience for each of these
				activities.
			</p>
		</div>
		<div class="self-start">
			<h3>Zero Trust Security.</h3>
			<p class="mb-0">
				Security is a multi-layer discipline and we apply the zero trust security model to
				everything we do. Through building awareness, approrpiate tooling and the necessary
				knowledge transfer want to deliver scalable and performant systems without compromising on
				security.
			</p>
		</div>
		<div class="self-start">
			<h3>Cloud Native.</h3>
			<p class="mb-0">
				Developing cloud-native software requires knowledge about the cloud ecosystem. We want to
				help by sharing our knowledge, best practices and lessons learn to help teams adopt a modern
				approach to cloud-native development.
			</p>
		</div>
	</Grid>
</section>
<section>
	<CtaButton />
</section>
