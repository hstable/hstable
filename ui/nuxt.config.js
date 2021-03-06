export default {
  server: {
    port: 3000, // default: 3000
    host: '0.0.0.0', // default: localhost
  },
  /*
   ** Nuxt rendering mode
   ** See https://nuxtjs.org/api/configuration-mode
   */
  mode: 'universal',
  /*
   ** Nuxt target
   ** See https://nuxtjs.org/api/configuration-target
   */
  target: 'server',
  /*
   ** Headers of the page
   ** See https://nuxtjs.org/api/configuration-head
   */
  head: {
    title: 'HSTable | 哈深格子',
    meta: [
      { charset: 'utf-8' },
      {
        name: 'viewport',
        content:
          'width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no',
      },
      {
        hid: 'description',
        name: 'description',
        content: process.env.npm_package_description || '',
      },
    ],
    link: [{ rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }],
  },
  /*
   ** Global CSS
   */
  css: [],
  /*
   ** Plugins to load before mounting the App
   ** https://nuxtjs.org/guide/plugins
   */
  plugins: [
    { src: '~/plugins/easytable.js' },
    { src: '~/plugins/iview.js' },
    { src: '~/plugins/vant.js' },
    { src: '~/plugins/axios.js' },
    { src: '~/plugins/cookies.client.js' },
  ],
  /*
   ** Auto import components
   ** See https://nuxtjs.org/api/configuration-components
   */
  components: true,
  /*
   ** Nuxt.js dev-modules
   */
  buildModules: [
    // Doc: https://github.com/nuxt-community/eslint-module
    '@nuxtjs/eslint-module',
  ],
  /*
   ** Nuxt.js modules
   */
  modules: [
    // Doc: https://axios.nuxtjs.org/usage
    '@nuxtjs/axios',
    '@nuxtjs/pwa',
    'nuxt-babel',
  ],
  /*
   ** Axios module configuration
   ** See https://axios.nuxtjs.org/options
   */
  axios: {
    baseURL: process.env.API_URL_BROWSER || '//localhost:8000',
  },
  publicRuntimeConfig: {
    axios: {
      baseURL: process.env.BASE_URL,
      browserBaseURL: process.env.API_URL_BROWSER || '/',
    },
  },
  privateRuntimeConfig: {
    axios: {
      baseURL: process.env.BASE_URL,
    },
  },
  /*
   ** Build configuration
   ** See https://nuxtjs.org/api/configuration-build/
   */
  build: {},
  transpile: [/vant.*?less/],

  babel: {
    plugins: [
      [
        'import',
        {
          libraryName: 'vant',
          libraryDirectory: 'es',
          style: (name) => {
            return `${name}/style/index.js`
          },
        },
        'vant',
      ],
      [
        'import',
        {
          libraryName: 'iview',
          libraryDirectory: 'src/components',
        },
        'iview',
      ],
    ],
  },
  pwa: {
    workbox: {
      skipWaiting: true,
    },
  },
}
