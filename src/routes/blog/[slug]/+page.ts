
import { error } from '@sveltejs/kit';

/** @type {import('./$types').PageLoad} */
export async function load({ data }) {
	try {
		const post = await import(`../../../posts/${data.slug}.md`);
		return {
			post: post.metadata,
			component: post.default,
			relatedPosts: data.relatedPosts
		};
	} catch (e) {
		error(404, "post cannot be found: " + data.slug);
	}
}
