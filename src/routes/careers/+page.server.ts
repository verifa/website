import { getPostsGlob, PostType } from '$lib/posts/posts';

/** @type {import('./$types').PageServerLoad} */
export async function load() {
	return {
		posts: getPostsGlob({
			types: [PostType.Job],
			jobActive: true
		})
	};
}

