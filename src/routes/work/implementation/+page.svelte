<script lang="ts">
	import Column from '$lib/column.svelte';
	import Columns from '$lib/columns.svelte';
	import CtaButton from '$lib/ctaButton.svelte';
	import { seo } from '$lib/seo/store';
	import HeaderLine from '$lib/headerLine.svelte';
	import CDWorkshop from '../_cdWorkshop.svelte';

	seo.reset();
	$seo.title = 'How we do it';
</script>

<section>
	<Columns reverse={false}>
		<Column>
			<h1>How we do it.</h1>
			<h3>Working together to create sustainable solutions that you own.</h3>
		</Column>
		<Column>
			<img
				class="object-contain h-full w-full"
				src="/continuous-delivery-partner.svg"
				alt="An illustration of several people constructing a building-sized dashboard. One person is operating a crane."
			/>
		</Column>
	</Columns>
</section>

<section id="workshops">
	<HeaderLine />
	<h2>Assess.</h2>

	<h4>Kickstarting improvement processes and framing core issues.</h4>
	<Columns>
		<Column class="self-start">
			<h3>Workshops.</h3>
			<p>
				Our workshops focus on knowledge sharing and self-discovery. By discussing and analysing
				value streams, cloud architectures, development workflows, deployment strategies, and many
				more topics, we help teams understand more about the ecosystem and help them arrive at their
				own conclusions. This is fundamental transforming the way teams work because everyone knows
				<i>why</i> there is change.
			</p>
		</Column>
		<Column class="self-start">
			<h3>Who are they for?</h3>
			<p>
				Our workshops are for any teams looking to change the way they work. They can span from a
				few hours to multiple days, meaning there is no big upfront investment. Whether you're
				starting a new project or looking to improve your current one, we are happy to hear from you
				and try to help. The very least we can offer is a trusted opinion of your ways of working.
			</p>
		</Column>
	</Columns>
	<CDWorkshop />
</section>

<section id="coaching">
	<HeaderLine />
	<h2>Coach.</h2>
	<h4>Building competence within teams.</h4>
	<Columns>
		<Column class="self-start">
			<h3>You own your work.</h3>
			<p>
				To make changes long-lasting it is important that the knowledge and practices introduced
				stay with the teams who own the work. Hiring consultants to come in and implement everything
				for you can lead to a lack of knowledge and practices to own the work. We believe in
				coaching teams to empower them learn new practices and technology, therefore ensuring the
				ownership is clear.
			</p>
		</Column>
		<Column class="self-start">
			<h3>High performing teams.</h3>
			<p>
				Building high performing teams means bringing everyone with you. Working alongside engineers
				with experience will boost productivity whilst building competence. We help teams adopt
				technologies like the Cloud, HashiStack and Kubernetes by doing the day-to-day work
				together, thereby training and solving real problems at the same time.
			</p>
		</Column>
	</Columns>
	<h2>
		‘A key and crucial part of our Continuous Integration journey was having Verifa's involvement in
		holding-our-hands and giving us confidence while we moved our active projects to a CI
		methodology.’.
	</h2>
	<h4>— David Hoslett, Siemens Mobility.</h4>
</section>

<section id="implementation">
	<HeaderLine />
	<h2>Build and Evolve.</h2>
	<h4>Building sustainable solutions.</h4>
	<p>
		We have long-lasting relationships with customers who trust us to work alongside their teams and
		help solve all kinds of problems. We can deliver both short-term and long-term work. Whether you
		need help with CI/CD workflows, Platform Engineering, Cloud, the HashiStack, Kubernetes, or any
		related topic, we're here to help you build, implement and evolve sustainable solutions tailored
		for your unique business needs.
	</p>
</section>
<section>
	<CtaButton />
</section>
