import { error } from "@sveltejs/kit";

export const load = async ({ params }) => {
	try {
		const post = await import(`../../../posts/cases/${params.slug}.md`);
		return {
			post: post.default,
		};
	} catch (e) {
		error(404, "Cannot find case: " + params.slug);
	}
};
