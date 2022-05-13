import { getCases } from "$lib/posts/posts";

export async function get({ }) {

    return {
        body: {
            ...getCases()
        }
    };
}
