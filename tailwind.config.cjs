
const config = {
	mode: 'jit',
	content: ["./**/*.templ"],

	safelist: [
		'ml-2',
		'no-underline',
		'hover:text-v-lilac',
	],

	theme: {
		colors: {
			'v-black': "#0d0e12",
			'v-pink': "#fc9cac",
			'v-green': "#ccecef",
			'v-gray': "#c4d0dd",
			'v-lilac': "#ad9ce3",
			'v-beige': "#8f8379",
			'v-white': "#f9fafb",

			// 'v-caution': '#f83d5a',
			'v-note': '#316DCA',
			'v-tip': '#347D39',
			'v-important': '#8256D0',
			'v-warning': '#966600',
			'v-caution': '#22272E',
		},
		container: {
			padding: {
				DEFAULT: '1rem',
				sm: '2rem',
				md: '3rem',
				lg: '4rem',
				xl: '5rem',
				'2xl': '6rem',
			},
		},
		fontFamily: {
			sans: ['Outfit', 'sans-serif'],
		},
		extend: {
			lineHeight: {
				'pizzo': '1.1',
			}
		}
	},

	plugins: [
		require('@tailwindcss/typography'),
		require('@tailwindcss/forms'),
	]
};

module.exports = config;
