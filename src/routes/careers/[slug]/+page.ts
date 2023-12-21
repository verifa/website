import { error } from '@sveltejs/kit';
export const load = async ({ params }) => {
	try {
		const post = await import(`../../../posts/${params.slug}.md`);
		return {
			post: post.metadata,
			component: post.default
		};
	} catch (e) {
		error(404, "Cannot find job opening: " + params.slug);
	}
};
