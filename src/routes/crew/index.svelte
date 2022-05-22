<script lang="ts">
	import { allActiveCrewShuffle } from '$lib/crew/crew';
	import type { Member } from '$lib/crew/crew';
	import { seo } from '$lib/seo/store';
	import MemberCard from './_memberCard.svelte';

	const crew: Member[] = allActiveCrewShuffle();

	seo.reset();
</script>

<section>
	<h4>Our Crew</h4>
	<h1>Experts in all things Continuous and Cloud.</h1>
</section>
<section>
	<ul
		class="space-y-12 sm:grid sm:grid-cols-2 sm:gap-x-6 sm:gap-y-12 sm:space-y-0 lg:grid-cols-3 lg:gap-x-8 list-none"
	>
		{#each crew as member}
			<li>
				<MemberCard {member} />
			</li>
		{/each}
	</ul>
</section>
