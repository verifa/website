import { getBlogs } from "$lib/posts/posts";

export async function get({ }) {

    return {
        body: {
            ...getBlogs()
        }
    };
}

