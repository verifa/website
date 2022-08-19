import { error } from '@sveltejs/kit';
export const load = async ({ params }) => {
	try {
		const post = await import(`../../../posts/${params.slug}.md`);
		return {
			post: post.default
		};
	} catch (e) {
		throw error(404, "Cannot find job opening: " + params.slug)
	}
};
