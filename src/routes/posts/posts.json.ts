import { getPostsGlob, type PostsQuery } from "$lib/posts/posts";

export async function get({ url }) {
    const query: PostsQuery = {
        author: url.searchParams.get("author"),
        limit: url.searchParams.get("limit"),
        featured: url.searchParams.get("featured"),
        keywords: url.searchParams.get("keywords") ? url.searchParams.get("keywords").split(",") : null,
        types: url.searchParams.get("types") ? url.searchParams.get("types").split(",") : null,
        skipTitle: url.searchParams.get("skip_title")
    }
    return {
        body: {
            ...getPostsGlob(query)
        }
    };
}

