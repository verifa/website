<script lang="ts">
	import type { Post } from '$lib/posts/posts';
	import PostCard from '$lib/posts/postCard.svelte';

	export let showBadges: boolean = true;
	export let showPreview: boolean = true;
	export let posts: Post[] = [];
</script>

<div class="grid gap-8 md:grid-cols-2 lg:grid-cols-3 lg:gap-x-5 lg:gap-y-12">
	{#each posts as post}
		<PostCard {showBadges} {showPreview} {post} />
	{/each}
</div>
