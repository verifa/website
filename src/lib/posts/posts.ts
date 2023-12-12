
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
    hidden: boolean;
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
//   input path: posts/cases/this-is-some-case.md
//   returns: this-is-some-case
const getSlug = (path: string): string =>
    path.split("/").at(-1).replace(".md", "");

// We have some old blogs using keywords that are not so relevant...
// Let's ignore those rather than updating old blogs.
const ignoredKeywords: string[] = [
    "Conference",
    "Data Science",
    "JFrog",
    "Machine Learning",
]

export interface PostsQuery {
    // types, if provided, filters by type (e.g. Blog, Case, Job)
    types?: PostType[];
    // keywords, if provided, filters by keywords (i.e. "tags" in posts)
    keywords?: string[];
    // limit, if provided, limits the results returned. It is done after all
    // other filters and ordering has been performed
    limit?: number;
    // featured, if provided, filters by the featured posts
    featured?: boolean;
    // author, if provided, filters by the author
    author?: string;
    // skipTitle, if provided, filters the post with the same title.
    // This is useful if getting related posts, so that it excludes
    // the same post from related posts
    skipTitle?: string;
    // allKeywords, if provided, returns all the keywords before the filter
    // is applied. If not provided, the keywords returned are those from
    // the filtered posts that are returned
    allKeywords?: boolean;
    // jobActive filters job types if they should be active or not
    jobActive?: boolean;
}


export const getPostsGlob = (query: PostsQuery = {}): PostsData => {
    let posts: Post[] = [];

    const rawPosts = import.meta.glob("../../posts/**/*.md", { eager: true })
    for (const key in rawPosts) {
        const rawPost = rawPosts[key]
        const post: Post = {
            slug: getSlug(key),
            ...rawPost["metadata"],
        }
        // Filter by type here, as some logic depends on only having the requested
        // types in the list of all posts.
        // We filter the rest of the query later.
        if (query.types && !query.types.includes(post.type)) {
            continue
        }
        posts.push(post)
    }

    // Sort the posts by date
    posts.sort((a, b) => new Date(b.date).valueOf() - new Date(a.date).valueOf())

    const filteredPosts = filterPosts(posts, query)

    let keywords: string[] = []
    if (query.allKeywords) {
        keywords = keywordsFromPosts(posts)
    } else {
        keywords = keywordsFromPosts(filteredPosts)
    }

    return {
        posts: filteredPosts,
        keywords: keywords,
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
        // Check job active filter
        if (query.jobActive && post.type === PostType.Job && !post.jobActive) {
            return false
        }

        // Do not filter hidden posts here.
        // For the static build to render a blog, there needs to be an href
        // to that blog somewhere on the site.
        // Therefore, return hidden posts and hide them in the HTML instead.
        // That way we have an href to the blog, meaning it gets built, but
        // it is not visible on the site.
        return true
    })
    // Apply any limit on them
    if (query.limit > 0) {
        posts = posts.slice(0, query.limit)
    }
    return posts
}

const keywordsFromPosts = (posts: Post[]): string[] => {
    let keywords: Set<string> = new Set();
    // Populate keywords from the posts
    posts.forEach((post) => post.tags.forEach((tag) => {
        if (!ignoredKeywords.includes(tag)) {
            keywords.add(tag)
        }
    }))
    return Array.from(keywords).sort()
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
    const lowerTags = tags.map((tag) => tag.toLowerCase())
    for (let index = 0; index < keywords.length; index++) {
        const keyword = keywords[index];
        // If there is even one match, then return it
        if (lowerTags.includes(keyword.toLowerCase())) {
            return true
        }
    }
    return false
}
