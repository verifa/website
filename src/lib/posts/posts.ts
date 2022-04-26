
export interface Blogs {
    blogs: Post[]
    keywords: string[]
}

export interface Jobs {
    jobs: Post[]
    keywords: string[]
}

export interface Post {
    slug: string;
    type: PostType;
    title: string;
    subheading: string;
    authors: string[];
    tags: string[];
    date: Date;
    image: string;
    featured: boolean;

    jobActive: boolean;
}

export enum PostType {
    Blog = "Blog",
    Event = "Event",
    Job = "Job"
}

export const blogTypes = [PostType.Blog, PostType.Event]
export const jobTypes = [PostType.Job]

const getSlug = (key: string): string =>
    key.substring('../../posts/'.length, key.lastIndexOf('.'));

// orderKeywords takes a list of keywords (with possible duplicates)
// and returns an ordered list by the number of occurences of each keyword
const orderKeywords = (keywords: string[]): string[] => {
    var countMap = keywords.reduce(function (p, c) {
        p[c] = (p[c] || 0) + 1;
        return p;
    }, {});

    return Object.keys(countMap).sort(function (a, b) {
        return countMap[b] - countMap[a];
    });
}
// We have some old blogs using keywords that are not so relevant...
// Let's ignore those rather than updating old blogs.
const ignoredKeywords: string[] = [
    "JFrog",
    "Machine Learning",
    "Data Science"
]

export async function getRelatedBlogs(fetch: any, title: string, keywords: string[]): Promise<Blogs> {
    const res = await fetch("/posts/blogs.json");
    if (res.ok) {
        return {
            blogs: (await res.json()).blogs.filter((blog) => {
                // Make sure we don't return the current blog post as a similar one
                if (blog.title === title) {
                    return false
                }
                for (let index = 0; index < keywords.length; index++) {
                    const keyword = keywords[index];
                    // If there is even one match, then return it
                    if (blog.tags.includes(keyword)) {
                        return true
                    }
                }
                return false
            }).slice(0, 3),
            // Not needed for similar posts
            keywords: []
        }
    } else {
        return Promise.reject(res.text())
    }
}

export const getBlogs = (limit: number = -1, featured: boolean = false): Blogs => {
    let posts: Post[] = [];
    let keywords: string[] = [];

    const rawPosts = import.meta.globEager("../../posts/*.md")
    for (const key in rawPosts) {
        const rawPost = rawPosts[key]
        const post: Post = {
            slug: getSlug(key),
            ...rawPost.metadata,
        }
        // Check filters
        if (!blogTypes.includes(post.type)) {
            continue
        }
        posts.push(post)
        post.tags.forEach((tag) => {
            if (!ignoredKeywords.includes(tag)) {
                keywords.push(tag)
            }
        })
    }

    // Sort the posts by date
    posts.sort((a, b) => new Date(b.date).valueOf() - new Date(a.date).valueOf())
    // If searching for featured posts, first remove those that are not featured
    if (featured) {
        posts = posts.filter((post) => post.featured)
    }
    // Apply any limit on them
    if (limit > 0) {
        posts = posts.slice(0, limit)
    }

    return {
        blogs: posts,
        keywords: orderKeywords(keywords),
    }
}

export const getJobs = (limit: number = -1): Jobs => {
    let posts: Post[] = [];
    let keywords: string[] = []

    const rawPosts = import.meta.globEager("../../posts/*.md")
    for (const key in rawPosts) {
        const rawPost = rawPosts[key]
        const post: Post = {
            slug: getSlug(key),
            ...rawPost.metadata,
        }
        // Check filters
        if (!jobTypes.includes(post.type)) {
            continue
        }

        // Skip inactive jobs
        if (!post.jobActive) {
            continue
        }
        posts.push(post)
        post.tags.forEach((tag) => {
            if (!ignoredKeywords.includes(tag)) {
                keywords.push(tag)
            }
        })
    }

    // Sort the posts by date and apply any limit on them
    posts.sort((a, b) => new Date(b.date).valueOf() - new Date(a.date).valueOf())
    if (limit > 0) {
        posts = posts.slice(0, limit)
    }

    return {
        jobs: posts,
        keywords: orderKeywords(keywords),
    }
}

export const getPostUrl = (post: Post): string => {
    switch (post.type.toLowerCase()) {
        case "blog":
        case "event":
            return `/blog/${post.slug}`
        case "job":
            return `/careers/${post.slug}`
        default:
            return "/404"
    }
}
