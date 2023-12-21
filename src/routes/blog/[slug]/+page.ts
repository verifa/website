
import { error } from '@sveltejs/kit';

/** @type {import('./$types').PageLoad} */
export async function load({ params, data }) {
	try {
		const post = await import(`../../../posts/${params.slug}.md`);
		return {
			post: post.default,
			relatedPosts: data.relatedPosts
		};
	} catch (e) {
		error(404, "post cannot be found: " + params.slug);
	}
}
