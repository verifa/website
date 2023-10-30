<script lang="ts">
	import ButtonLink from '$lib/buttonLink.svelte';
	import Column from '$lib/column.svelte';
	import Columns from '$lib/columns.svelte';
	import CompanyValues from '$lib/company/companyValues.svelte';
	import Grid from '$lib/grid.svelte';
	import { config, seo } from '$lib/seo/store';
	import HeaderLine from '$lib/headerLine.svelte';

	seo.reset();
	$seo.title = 'About us: Consulting is our game';

	const cloudLogos: { name: string; image: string }[] = [
		{
			image: '/clouds/google-cloud-white.svg',
			name: 'google-cloud'
		},
		{
			image: '/clouds/aws-white.svg',
			name: 'aws'
		},
		{
			image: '/clouds/azure-white.svg',
			name: 'azure'
		},
		{
			image: '/clouds/upcloud-white.svg',
			name: 'upcloud'
		}
	];

	const certLogos: { name: string; image: string }[] = [
		{
			image: '/certifications/hashicorp-chip-terraform-lilac.png',
			name: 'hashicorp-terraform-chip'
		},
		{
			image: '/certifications/hashicorp-chip-vault-lilac.png',
			name: 'hashicorp-vault-chip'
		},
		{
			image: '/certifications/hashicorp-terraform-associate-lilac.png',
			name: 'hashicorp-terraform-associate'
		},
		{
			image: '/certifications/hashicorp-vault-associate-lilac.png',
			name: 'hashicorp-vault-associate-lilac'
		},
		{
			image: '/certifications/hashicorp-consul-associate-badge-lilac.png',
			name: 'hashicorp-consul-associate'
		},
		{
			image: '/certifications/kubernetes-certified-administrator-lilac.png',
			name: 'kubernetes-certified-administrator'
		},
		{
			image: '/certifications/gcp-professiona-cloud-architect-lilac.png',
			name: 'gcp-professiona-cloud-architect-lilac'
		},
		{
			image: '/certifications/google-cloud-devops-engineer-lilac.png',
			name: 'google-cloud-devops-engineer-lilac'
		},
		{
			image: '/certifications/azure-administrator-associate-lilac.png',
			name: 'azure-administrator-associate-lilac'
		},
		{
			image: '/certifications/aws-cloud-practitioner-lilac.png',
			name: 'aws-cloud-practitioner-lilac'
		},
		{
			image: '/certifications/aws-certified-developer-associate-badge-lilac.png',
			name: 'aws-certified-developer-associate'
		},
		{
			image: '/certifications/aws-certified-solutions-architect-associate-badge-lilac.png',
			name: 'aws-certified-solutions-architect-associate'
		}
	];
</script>

<section>
	<HeaderLine />
	<h4>About Us</h4>
	<h1>Consulting is our game.</h1>
</section>
<section>
	<h3>
		We are an experienced crew of DevOps and Cloud professionals dedicated to helping our customers
		with Continuous practices and Cloud adoption. We each bring our own specialities, experiences
		and know-how when collaborating on customer projects, tailoring services to our customers'
		individual needs.
	</h3>
</section>
<section>
	<HeaderLine />
	<h4>Our Values</h4>
	<CompanyValues />
</section>

<section>
	<HeaderLine />
	<h4>Our Strategy</h4>
	<h2>Open Source and Cloud.</h2>
</section>
<section>
	<Grid>
		<img
			class="object-contains w-full h-full"
			src="/everything-is-connected.svg"
			alt="everything is connected"
		/>
		<h4>
			Open source has and will continue to change the software industry. We want to do our part and
			support the open source community and help others build trust with open source.
		</h4>
		<div class="lg:order-last grid grid-cols-2 gap-0.5 md:grid-cols-2">
			{#each cloudLogos as logo}
				<div class="col-span-2 sm:col-span-1 flex justify-center py-8 px-8  bg-v-lilac">
					<img src={logo.image} alt={logo.name} class="h-16 w-full object-contain" />
				</div>
			{/each}
		</div>
		<div class="">
			<h4>
				We made friends with the cloud. It enables us to move faster and more flexibly. We can help
				you with your cloud journey so you can also benefit from it.
			</h4>
		</div>
	</Grid>
</section>

<section>
	<HeaderLine />
	<h4>Our Crew</h4>
	<h2>Do you want to get to know us better? Find out who we are as people and connect with us.</h2>
	<ButtonLink href="/crew">Our Crew</ButtonLink>
</section>
<section>
	<h4>Qualified and certified to prove it.</h4>
	<div class="grid grid-cols-2 gap-0.5 md:grid-cols-6 lg:grid-cols-5">
		{#each certLogos as logo}
			<div class="col-span-1 flex justify-center md:col-span-2 lg:col-span-1 py-8 px-8  bg-v-lilac">
				<img src={logo.image} alt={logo.name} class="h-38 w-full object-contain" />
			</div>
		{/each}
	</div>
</section>

<section>
	<Columns>
		<Column>
			<img class="object-contain h-full w-full" src="/round-table.svg" alt="round table" />
		</Column>
		<Column>
			<div class="flex flex-col gap-y-0">
				<h2>Join us.</h2>
				<h4>
					We are always on the lookout for our next colleague to join us on our Veri-nice journey.
					Like what you see and want to know more?
				</h4>
				<div>
					<ButtonLink href="/careers">Careers</ButtonLink>
				</div>
			</div>
		</Column>
	</Columns>
</section>
