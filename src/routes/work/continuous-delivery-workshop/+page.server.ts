import { blogTypes, getPostsGlob, PostType } from '$lib/posts/posts';

/** @type {import('./$types').PageServerLoad} */
export async function load() {
	return {
		posts: getPostsGlob({
			types: [...blogTypes, PostType.Case],
			limit: 3,
			keywords: ['Value Stream Mapping']
		})
	};
}
