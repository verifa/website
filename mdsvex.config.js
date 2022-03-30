import rehypeExternalLinks from "rehype-external-links";
import rehypeSlug from "rehype-slug";
import rehypeAutolinkHeadings from "rehype-autolink-headings";

const config = {
	extensions: ['.svelte.md', '.md', '.svx'],

	layout: {
		_: "./src/lib/posts/layout.svelte"
	},

	smartypants: {
		dashes: 'oldschool'
	},

	remarkPlugins: [],
	rehypePlugins: [
		rehypeExternalLinks, rehypeSlug, rehypeAutolinkHeadings
	]
};

export default config;
