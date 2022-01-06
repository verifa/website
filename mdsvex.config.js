const config = {
	extensions: ['.svelte.md', '.md', '.svx'],

	layout: {
		_: "./src/lib/posts/layout.svelte"
	},

	smartypants: {
		dashes: 'oldschool'
	},

	remarkPlugins: [],
	rehypePlugins: []
};

export default config;
