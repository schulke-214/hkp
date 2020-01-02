import pkg from './package';

const API_URL = (() => {
	if (process.env.API_URL) return process.env.API_URL;

	console.warn('using http://localhost:8080 as api url');
	return 'http://localhost:8080/';
})();

export default {
	mode: 'universal',

	/*
	 ** Headers of the page
	 */
	head: {
		title: pkg.name,
		meta: [
			{ charset: 'utf-8' },
			{ name: 'viewport', content: 'width=device-width, initial-scale=1' },
			{ hid: 'description', name: 'description', content: pkg.description }
		],
		link: [{ rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }]
	},

	/*
	 ** Customize the progress-bar color
	 */
	loading: '~/components/ui/loading.vue',

	env: {
		API_URL
	},

	/*
	 ** Global CSS
	 */
	css: [],

	/*
	 ** Plugins to load before mounting the App
	 */
	plugins: ['~/plugins/api.js', '~/plugins/vee-validate.js'],

	/*
	 ** Nuxt.js modules
	 */
	modules: [
		// Doc: https://axios.nuxtjs.org/usage
		'@nuxtjs/axios',
		// Doc: https://buefy.github.io/#/documentation
		'nuxt-buefy',
		'@nuxtjs/pwa'
	],
	/*
	 ** Axios module configuration
	 */
	axios: {
		baseURL: API_URL
	},

	/*
	 ** Build configuration
	 */
	build: {
		/*
		 ** You can extend webpack config here
		 */
		extend(config, ctx) {},
		transpile: ['vee-validate']
	}
};
