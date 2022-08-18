
export interface PostsData {
    posts: Post[]
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
    previewImage?: string;
    image: string;
    featured: boolean;

    jobActive: boolean;
}

export enum PostType {
    Blog = "Blog",
    Podcast = "Podcast",
    Case = "Case",
    Event = "Event",
    Job = "Job"
}

export const blogTypes: PostType[] = [
    PostType.Blog, PostType.Podcast, PostType.Event
]

// getSlug takes a path to a post and returns the name of the file removing all
// paths and the extension.
// E.g.
// input path: posts/cases/this-is-some-case.md
// returns: this-is-some-case
const getSlug = (path: string): string =>
    path.split("/").at(-1).replace(".md", "");

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
export interface PostsQuery {
    types?: PostType[];
    keywords?: string[];
    limit?: number;
    featured?: boolean;
    author?: string;
    relatedPostTitle?: string;
    skipTitle?: string;
}


export const getPostsGlob = (query: PostsQuery = {}): PostsData => {
    let posts: Post[] = [];

    const rawPosts = import.meta.globEager("../../posts/**/*.md")
    for (const key in rawPosts) {
        const rawPost = rawPosts[key]
        const post: Post = {
            slug: getSlug(key),
            ...rawPost.metadata,
        }
        posts.push(post)
    }

    // Sort the posts by date
    posts.sort((a, b) => new Date(b.date).valueOf() - new Date(a.date).valueOf())

    return preparePostsData(posts, query)
}

const preparePostsData = (posts: Post[], query: PostsQuery = {}): PostsData => {
    posts = filterPosts(posts, query)
    return {
        posts: posts,
        keywords: keywordsFromPosts(posts),
    }
}

export const filterPosts = (posts: Post[], query: PostsQuery): Post[] => {
    posts = posts.filter((post) => {
        // Check types
        if (query.types && !query.types.includes(post.type)) {
            return false
        }
        // Check keywords
        if (query.keywords && !hasMatchingKeyword(query.keywords, post.tags)) {
            return false
        }
        // If searching for featured posts, first remove those that are not featured
        if (query.featured && !post.featured) {
            return false
        }
        // Check if filter by author
        if (query.author && !post.authors.includes(query.author)) {
            return false
        }
        // Check skip title
        if (query.skipTitle && post.title === query.skipTitle) {
            return false
        }
        return true
    })
    // Apply any limit on them
    if (query.limit > 0) {
        posts = posts.slice(0, query.limit)
    }
    return posts
}

const keywordsFromPosts = (posts: Post[]): string[] => {
    let keywords: string[] = [];
    // Populate keywords from the posts
    posts.forEach((post) => post.tags.forEach((tag) => {
        if (!ignoredKeywords.includes(tag)) {
            keywords.push(tag)
        }
    }))
    return orderKeywords(keywords)
}


export const getPostUrl = (post: Post): string => {
    switch (post.type) {
        case PostType.Blog:
        case PostType.Podcast:
        case PostType.Event:
            return `/blog/${post.slug}`
        case PostType.Job:
            return `/careers/${post.slug}`
        case PostType.Case:
            return `/clients/${post.slug}`
        default:
            return "/404"
    }
}

const hasMatchingKeyword = (keywords: string[], tags: string[]): boolean => {
    for (let index = 0; index < keywords.length; index++) {
        const keyword = keywords[index];
        // If there is even one match, then return it
        if (tags.includes(keyword)) {
            return true
        }
    }
    return false
}
