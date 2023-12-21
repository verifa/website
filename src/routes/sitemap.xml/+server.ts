import { getPostsGlob, PostType } from "$lib/posts/posts";

/** @type {import('./$types').RequestHandler} */
export async function GET({ }) {
    function sitemapDate(date: Date): string {
        return date.toISOString().split('T')[0]
    }
    interface Page {
        loc: string;
        priority: number;
        lastmod?: Date;
    }
    const defaultLastMod = new Date()

    let pages: Page[] = [
        {
            loc: "https://verifa.io",
            priority: 1.0
        },
        {
            loc: "https://verifa.io/work/value-stream-assessment",
            priority: 0.8
        },
        {
            loc: "https://verifa.io/work/software-delivery-platforms",
            priority: 0.8
        },
        {
            loc: "https://verifa.io/work/team-topologies",
            priority: 0.8
        },
        {
            loc: "https://verifa.io/work/cloud-architecture",
            priority: 0.8
        },
        {
            loc: "https://verifa.io/work/implementation",
            priority: 0.8
        },
        {
            loc: "https://verifa.io/company",
            priority: 0.8
        },
        {
            loc: "https://verifa.io/clients",
            priority: 0.8
        },
        {
            loc: "https://verifa.io/careers",
            priority: 0.8
        },
    ]

    const posts = getPostsGlob({})

    posts.posts.forEach((post) => {
        pages.push({
            loc: `https://verifa.io/blog/${post.slug}`,
            priority: post.featured ? 0.6 : 0.5,
            lastmod: new Date(post.date)
        })
    })

    const body = `<?xml version="1.0" encoding="UTF-8" ?>
<urlset
    xmlns="https://www.sitemaps.org/schemas/sitemap/0.9"
    xmlns:news="https://www.google.com/schemas/sitemap-news/0.9"
    xmlns:xhtml="https://www.w3.org/1999/xhtml"
    xmlns:mobile="https://www.google.com/schemas/sitemap-mobile/1.0"
    xmlns:image="https://www.google.com/schemas/sitemap-image/1.1"
    xmlns:video="https://www.google.com/schemas/sitemap-video/1.1"
>
${pages.map((page) =>
        `
<url>
    <loc>${page.loc}</loc>
    <lastmod>${page.lastmod ? sitemapDate(page.lastmod) : sitemapDate(defaultLastMod)}</lastmod>
    <priority>${page.priority}</priority>
    <changefreq>daily</changefreq>
</url>`).join('')
        }
</urlset>
    `

    return new Response(body);
}
