
import { blogTypes, getPostsGlob, type PostsData } from '$lib/posts/posts';
import { error } from '@sveltejs/kit';

/** @type {import('./$types').PageServerLoad} */
export async function load({ params }) {
	try {
		const post = await import(`../../../posts/${params.slug}.md`);
		const { title, tags } = post.metadata;
		const postsData = getPostsGlob({
			types: blogTypes,
			limit: 3,
			skipTitle: title,
			keywords: tags
		})
		return {
			relatedPosts: postsData.posts
		};
	} catch (e) {
		throw error(404, "post cannot be found: " + params.slug)
	}
}
