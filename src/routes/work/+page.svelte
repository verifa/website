<script lang="ts">
	import Column from '$lib/column.svelte';
	import Columns from '$lib/columns.svelte';
	import CtaButton from '$lib/ctaButton.svelte';
	import Grid from '$lib/grid.svelte';
	import { seo } from '$lib/seo/store';
	import HeaderLine from '$lib/headerLine.svelte';
	import Faq from '../../lib/faq.svelte';
	import CDWorkshop from './_cdWorkshop.svelte';
	import FaqItem from '$lib/faqItem.svelte';

	seo.reset();
	$seo.title = 'What we do: Continuous Delivery and Cloud Architecture';
</script>

<section id="continuous-delivery">
	<HeaderLine />
	<h4>Our Specialities <br /> 1 / 2</h4>
	<h1>Continuous Delivery.</h1>
</section>
<section>
	<Columns>
		<Column class="self-start">
			<h3>What is Continuous Delivery?</h3>
			<p class="mb-0">
				Continuous Delivery is about building a process where software is always releasable whilst
				ensuring short release cycles. Often the pains and struggles to help achieve this are not
				related to technology but stem from a fragmented release process with multiple handovers
				causing friction and delays.
			</p>
		</Column>
		<Column class="self-start">
			<h3>What we believe</h3>
			<p class="mb-0">
				Every software team is slightly different and best practices are not a one-size-fits all.
				Our workshops and coaching help you unlock new potential by guiding you through your own
				Continuous Delivery journey. We will save you time by guard-railing you from common
				pitfalls, whilst learning from each other as we go.
			</p>
		</Column>
	</Columns>
</section>
<section>
	<CDWorkshop />
</section>
<section>
	<Columns>
		<Column>
			<div class="flex flex-col gap-y-6">
				<div class="flex space-x-6 items-center">
					<div>
						<img class="h-4 w-4 object-contain" src="/logo-element.png" alt="logo-ement" />
					</div>
					<div>
						<h4 class="mb-0">Visualize your release process</h4>
					</div>
				</div>
				<div class="flex space-x-6 items-center">
					<div>
						<img class="h-4 w-4 object-contain" src="/logo-element.png" alt="logo-ement" />
					</div>
					<div>
						<h4 class="mb-0">Pinpoint opportunities and waste</h4>
					</div>
				</div>
				<div class="flex space-x-6 items-center">
					<div>
						<img class="h-4 w-4 object-contain" src="/logo-element.png" alt="logo-ement" />
					</div>
					<div>
						<h4 class="mb-0">Get your whole team onboard</h4>
					</div>
				</div>
				<div class="flex space-x-6 items-center">
					<div>
						<img class="h-4 w-4 object-contain" src="/logo-element.png" alt="logo-ement" />
					</div>
					<div>
						<h4 class="mb-0">Justify investment</h4>
					</div>
				</div>
			</div>
		</Column>
		<Column>
			<CtaButton
				fullScreen={false}
				text={'Ready to unlock your CD potential?'}
				contact={'Get started.'}
			/>
		</Column>
	</Columns>
</section>
<section>
	<h2>‘Best practices are a good solution to someone else's problem’</h2>
	<h4>— Bosse Nyström, Chief of Wisdom, Verifa</h4>
</section>

<section id="cloud-architecture">
	<HeaderLine />
	<h4>Our Specialities <br /> 2 / 2</h4>
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

<section id="workshops">
	<HeaderLine />
	<h4>How we do it <br /> 1 / 3</h4>
	<h1>Workshops.</h1>
</section>
<section>
	<Columns>
		<Column class="self-start">
			<h3>What are they?</h3>
			<p>
				Our workshops focus on knowledge sharing and self-discovery. By discussing and analyzing
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
				few hours to multiple days, meaning there is no big upfront investment. Whether you are
				starting a new project or looking to improve your current one, we are happy to hear from you
				and try to help. The very least we can offer a trusted opinion of your ways of working.
			</p>
		</Column>
	</Columns>
</section>
<section>
	<CDWorkshop />
</section>
<section class="hidden">
	<h2>‘Do you want change, or do you want to change?’</h2>
	<h4>— Jacob Lärfors, Consultant, Verifa</h4>
</section>

<section id="coaching">
	<HeaderLine />
	<h4>How we do it <br /> 2 / 3</h4>
	<h1>Coaching.</h1>
</section>
<section>
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
</section>

<section id="implementation">
	<HeaderLine />
	<h4>How we do it <br /> 3 / 3</h4>
	<h1>Implementation.</h1>
</section>
<section>
	<h3>
		We have long-lasting relationships with customers who trust us to work alongside their teams and
		help solve all kinds of problems. We can deliver both short-term and long-term work. Whether you
		need help with Cloud, the HashiStack, Kubernetes, CI/CD workflows, or any related topic, we are
		here to be a trusted extension of your team.
	</h3>
</section>
<section>
	<CtaButton />
</section>
