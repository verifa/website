<script lang="ts">
	import Column from '$lib/column.svelte';
	import Columns from '$lib/columns.svelte';
	import CompanyValues from '$lib/company/companyValues.svelte';

	import PostGrid from '$lib/posts/postGrid.svelte';
	import HeaderLine from '$lib/headerLine.svelte';
	import { seo } from '$lib/seo/store';
	import CareersForm from '$lib/careers/careersForm.svelte';
	import type { PageData } from './$types';
	import type { Post } from '$lib/posts/posts';

	export let data: PageData;

	seo.reset();
	$seo.title = 'Join us: we care about our work and the impact we have';
	$seo.image.url = '/round-table.svg';

	const jobs: Post[] = data.posts.posts;
</script>

<section>
	<HeaderLine />
	<h4>Join us</h4>
	<h1>We care about our work and the impact we have.</h1>
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
	<Columns>
		<Column class="self-start">
			<h2>What we look for.</h2>
			<p>
				We look for people passionate about continuous practices, cloud architecture and all the
				wonderful things that will help our customers deliver greatness. If you like to be
				challenged and want to keep developing, then you should get in touch.
			</p>
		</Column>
		<Column class="self-start">
			<h2>What we offer.</h2>
			<p>
				Our goal is to provide a great place to work where you can continue to develop and learn.
				You'll get to work with cool tech and discuss with others who share a passion to deliver. To
				ensure fairness, we have an open salary policy.
			</p>
		</Column>
	</Columns>
</section>
<section id="open-positions">
	<HeaderLine />
	<h4>Open positions</h4>
	<h1>Check out our open positions.</h1>
</section>
<section>
	<PostGrid posts={jobs} showBadges={false} />
</section>
<section>
	<HeaderLine />
	<h4>Apply now</h4>
	<h1>Join our crew.</h1>
</section>
<section>
	<CareersForm />
</section>
