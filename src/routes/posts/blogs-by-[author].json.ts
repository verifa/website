import { getBlogs } from "$lib/posts/posts";

export async function get({ params }) {
    return {
        body: {
            ...getBlogs({ author: params.author })
        }
    };
}

