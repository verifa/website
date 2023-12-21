import rehypeExternalLinks from "rehype-external-links";
import rehypeSlug from "rehype-slug";
import rehypeAutolinkHeadings from "rehype-autolink-headings";

import { h, s } from 'hastscript'

const config = {
	extensions: ['.svelte.md', '.md', '.svx'],

	smartypants: {
		dashes: 'oldschool'
	},

	remarkPlugins: [],
	rehypePlugins: [
		rehypeExternalLinks, rehypeSlug,
		[
			rehypeAutolinkHeadings,
			{
				// Append link to header
				behavior: 'append',
				properties: {
					// Add class to <a> element
					class: 'group header-anchor-link'
				},
				content(node) {
					// Add class to each header
					node.properties.class = 'group header-anchor'
					return [
						h('span.hidden', ' permalink'),
						// Add link svg from heroicons
						s('svg', {
							xmlns: 'http://www.w3.org/2000/svg',
							// Add custom classes, including header-anchor-icon-hx depending on size of header
							class: `header-anchor-icon header-anchor-icon-${node.tagName}`,
							fill: "none",
							viewBox: "0 0 24 24",
							stroke: "currentColor",
							'stroke-width': "2"
						}, [
							s('path', {
								'stroke-linecap': "round",
								'stroke-linejoin': "round",
								d: "M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1"
							})
						])
					]

				}
			}
		]
	]
};

export default config;
