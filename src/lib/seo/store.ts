import { writable } from "svelte/store";

export interface SEOMetadata {
    title: string
    description: string
    canonical?: string

    robots: {
        noindex: boolean
    }

    article?: {
        published_time: Date
        modified_time: Date
        authors: string[]
        tags: string[]
    }

    image: SEOImage
    moreImages?: SEOImage[]
}

interface SEOImage {
    url: string
    alt: string
    // Width and height are pixels
    width?: number
    height?: number
}

const socialConfig = {
    facebookPage: "https://www.facebook.com/verifaio",
    githubPage: "https://github.com/verifa",
    linkedinProfile: "https://www.linkedin.com/company/verifa",
    twitterUsername: "verifaio",
    twitterProfile: "https://twitter.com/verifaio",
}

export const config = {
    entity: 'Verifa',
    siteLanguage: 'en',
    siteTitle: 'Verifa: Your trusted crew for all things Continuous and Cloud',
    siteShortTitle: 'Verifa Website',
    siteUrl: "https://verifa.io",
    contactEmail: "info@verifa.io",
    ...socialConfig,
    sameAs: [
        socialConfig.githubPage, socialConfig.linkedinProfile, socialConfig.twitterProfile, socialConfig.facebookPage,
    ]
};

// Create a custom store which has a reset function
function createSEOStore() {
    // For some reason, we cannot put the default store value in a variable and
    // reference it in the set method. So for now we just duplicate the default
    // values
    const store = writable<SEOMetadata>({
        title: "Verifa: Your trusted crew for all things Continuous and Cloud",
        description: "We are an experienced crew of DevOps and Cloud professionals dedicated to helping our customers with Continuous practices and Cloud adoption.",
        robots: {
            noindex: false
        },
        image: {
            url: "/verifa-logo.svg",
            alt: config.siteTitle,
        }
    });

    return {
        ...store,
        reset: () => store.set({
            title: "Verifa: Your trusted crew for all things Continuous and Cloud",
            description: "We are an experienced crew of DevOps and Cloud professionals dedicated to helping our customers with Continuous practices and Cloud adoption.",
            robots: {
                noindex: false
            },
            image: {
                url: "/verifa-logo.svg",
                alt: config.siteTitle,
            }
        })
    };
}

// Set some default values for the SEO
export const seo = createSEOStore();
