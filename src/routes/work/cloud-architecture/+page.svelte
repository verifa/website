<script lang="ts">
	import Column from '$lib/column.svelte';
	import Columns from '$lib/columns.svelte';
	import { seo } from '$lib/seo/store';
	import ButtonLink from '$lib/buttonLink.svelte';
	import HeaderLine from '$lib/headerLine.svelte';
	import CtaButton from '$lib/ctaButton.svelte';

	seo.reset();
	$seo.title = 'What we do: Cloud Architecture';
</script>

<section>
	<Columns reverse={false}>
		<Column>
			<h1>Cloud Architecture.</h1>
			<h3>Designing and implementing cloud-native architectures that are robust yet evolvable.</h3>
		</Column>
		<Column>
			<img
				class="object-contain h-full w-full"
				src="/cloud-architecture.svg"
				alt="An illustration of a group of people designing a complex 3-dimensional object"
			/>
		</Column>
	</Columns>
</section>

<section>
	<HeaderLine />
	<h4>Scalable, fast and reliable cloud</h4>
	<Columns reverse={false}>
		<Column>
			Migrating to the cloud is essential for ambitious businesses to achieve better performance and
			efficiencies - scalability, flexibility, self-service, cost savings, security, innovation,
			faster time to delivery. Architecting for the cloud requires careful analysis, planning and
			execution to make sure it serves your own unique business requirements.
		</Column>
		<Column>
			<ul>
				<li>Private or Public Cloud</li>
				<li>Orchestration</li>
				<li>Infrastructure as Code</li>
				<li>Monitoring & Observability</li>
				<li>Cost Optimisation</li>
				<li>Zero-Trust Security</li>
			</ul>
		</Column>
	</Columns>
	<h3 class="mb-0">
		Our expert cloud architects can help you navigate the complexities of the cloud, and help you
		build and adopt an awesome cloud architecture!
	</h3>
</section>

<section id="how-we-help">
	<HeaderLine />
	<h4>How we can help</h4>
	<p class="mb-0">
		We help teams adopt cloud (public or private) by helping in the architecting process and
		coaching teams with best practices. such as Infrastructure as Code. Our goal is to reduce the
		amount of time it takes for you to start benefiting from the many wonderful things that the
		cloud delivers. Through design sessions and workshops, we share our knowledge and experience to
		help you get the right architecture and plan for the cloud.
	</p>

	<Columns reverse={false}>
		<Column>
			<CtaButton
				text={'Need help migrating to the cloud?'}
				contact={'Let’s chat'}
				fullScreen={false}
			/>
		</Column>
		<Column>
			<h2>
				‘We were looking for help with professionalising our Continuous Integration and software
				delivery platform whilst moving the setup to the cloud. Verifa was a very reliable partner
				in implementing a modern setup using their years of experience’.
			</h2>
			<h4>— Laurent Muller, Vyaire Medical.</h4>
		</Column>
	</Columns>
</section>
