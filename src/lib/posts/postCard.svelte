<script lang="ts">
	import { PostType, type Post } from './posts';
	import PostBadges from './postBadges.svelte';

	export let post: Post;
	export let showBadges: boolean = true;

	const dateOptions = {
		weekday: 'long',
		year: 'numeric',
		month: 'long',
		day: 'numeric'
	};

	const urlPrefix = post.type === PostType.Job ? 'careers' : 'blog';
</script>

<div>
	{#if showBadges}
		<PostBadges type={post.type} tags={post.tags} />
	{/if}
	<a href="/{urlPrefix}/{post.slug}" class="block mt-4">
		<p class="mb-0">{new Date(post.date).toLocaleDateString()}</p>
		<h4>{post.title}</h4>
		<p>{post.subheading}</p>
	</a>
	<!-- <div class="mt-6 flex items-center">
		<div class="flex-shrink-0">
			<a href="#">
				<span class="sr-only">{post.author}</span>
				<img class="h-16 w-16 rounded-full" src="/crew/jlarfors.svg" alt="" />
			</a>
		</div>
		<div class="ml-3">
			<p class="font-semibold m-0">{post.author}</p>
			<p class="m-0">{new Date(post.date).toDateString()}</p>
		</div>
	</div> -->
</div>
