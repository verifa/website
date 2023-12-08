<script lang="ts">
	import type { Member } from '$lib/crew/crew';

	export let member: Member;

	let image: string = member.avatar;

	function handleMouseEnter() {
		image = member.sillyProfile;
	}

	function handleMouseLeave() {
		image = member.avatar;
	}
</script>

<div class="space-y-4">
	<a href={member.id} class="group">
		<div class="flex flex-col space-y-4">
			<div>
				<img
					class="h-64 w-64 object-contain transition-all duration-200 group-hover:scale-110 rounded-lg"
					src={image}
					alt={member.id}
					on:mouseenter={handleMouseEnter}
					on:mouseleave={handleMouseLeave}
				/>
			</div>

			<div>
				<h3 class="mb-0 group-hover:text-v-lilac">{member.name}</h3>
				<div class="flex items-center gap-x-4">
					<span class="my-0 group-hover:text-v-lilac">{member.position}</span>
				</div>
			</div>
		</div>
	</a>
</div>
