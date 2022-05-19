<script lang="ts">
	import type { Member } from '$lib/crew/crew';

	export let member: Member;
</script>

<div class="space-y-4">
	<a href={member.id} class="group">
		<div class="flex flex-col space-y-4">
			<div>
				<img
					class="h-64 w-64 object-contain transition-all duration-200 group-hover:scale-110"
					src={member.image}
					alt={member.id}
				/>
			</div>

			<div>
				<h3 class="mb-0 group-hover:text-v-lilac">{member.name}</h3>
				<div class="flex items-center gap-x-4">
					<p class="my-0 group-hover:text-v-lilac">{member.position}</p>
				</div>
			</div>
		</div>
	</a>
</div>
