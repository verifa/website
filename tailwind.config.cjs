const typography = require('@tailwindcss/typography');
const forms = require('@tailwindcss/forms');

const config = {
	mode: 'jit',
	purge: ['./src/**/*.{html,js,svelte,ts}'],

	theme: {
		extend: {}
	},

	plugins: [forms, typography]
};

module.exports = config;
