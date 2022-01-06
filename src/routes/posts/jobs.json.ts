import { getJobs } from "$lib/posts/posts";

export async function get({ }) {

    return {
        body: {
            ...getJobs()
        }
    };
}

