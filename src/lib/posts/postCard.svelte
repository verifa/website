<script lang="ts">
	import { getPostUrl, PostType, type Post } from './posts';
	import PostBadges from './postBadges.svelte';

	export let post: Post;
	export let showBadges: boolean = true;

	const dateOptions = {
		weekday: 'long',
		year: 'numeric',
		month: 'long',
		day: 'numeric'
	};

	// Make url reactive so that it updates when the post changes
	$: url = getPostUrl(post);
</script>

<div>
	{#if post.previewImage}
		<a href={url}>
			<img src={post.previewImage} alt={post.title} class="mb-8 w-2/3" />
		</a>
	{/if}
	{#if showBadges}
		<div class="mb-4">
			<PostBadges type={post.type} tags={post.tags} />
		</div>
	{/if}
	<a href={url}>
		<p class="mb-0">{new Date(post.date).toLocaleDateString()}</p>
		<h4>{post.title}</h4>
		<p>{post.subheading}</p>
	</a>
</div>
