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
		rehypeExternalLinks, rehypeSlug,
		[
			rehypeAutolinkHeadings,
			{
				behavior: 'prepend',
				// TODO: would be nice to get a proper icon ref
				// content: {
				// 	type: 'element',
				// 	tagName: 'span',
				// 	properties: {
				// 		className: ['hidden']
				// 	},
				// 	children: [
				// 		{
				// 			type: 'text',
				// 			value: '#'
				// 		}
				// 	]
				// }
			}
		]
	]
};

export default config;
