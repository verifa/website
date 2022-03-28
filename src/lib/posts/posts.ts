
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
    type: string;
    title: string;
    subheading: string;
    authors: string[];
    tags: string[];
    date: Date;
    image: string;
    featured: boolean;

    jobActive: boolean;
}

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

export const getPost = (slug: string): any => {
    const rawPosts = import.meta.globEager("../../posts/*.md")
    for (const key in rawPosts) {
        if (slug === getSlug(key)) {
            return rawPosts[key].default
        }
    }
    return null
}

export const getSimilarBlogs = (title: string, keywords: string[]): Blogs => {
    const blogs = getBlogs()
    return {
        blogs: blogs.blogs.filter((blog) => {
            // Make sure we don't return the current blog post as a similar one
            if (blog.title === title) {
                return false
            }
            for (let index = 0; index < keywords.length; index++) {
                const keyword = keywords[index];
                if (!blog.tags.includes(keyword)) {
                    return false
                }
            }
            return true
        }).slice(0, 3),
        // Not needed for similar posts
        keywords: []
    }
}

export const getBlogs = (limit: number = -1, featured: boolean = false): Blogs => {
    let posts: Post[] = [];
    let keywords: string[] = [];

    const postTypes: string[] = ["blog", "event"]
    const rawPosts = import.meta.globEager("../../posts/*.md")
    for (const key in rawPosts) {
        const rawPost = rawPosts[key]
        const post: Post = {
            slug: getSlug(key),
            ...rawPost.metadata,
        }
        // Check filters
        if (postTypes.length > 0 && !postTypes.includes(post.type.toLowerCase())) {
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

    const postTypes: string[] = ["job"]
    const rawPosts = import.meta.globEager("../../posts/*.md")
    for (const key in rawPosts) {
        const rawPost = rawPosts[key]
        const post: Post = {
            slug: getSlug(key),
            ...rawPost.metadata,
        }
        // Check filters
        if (postTypes.length > 0 && !postTypes.includes(post.type.toLowerCase())) {
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
