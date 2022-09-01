import path from 'path';
import env from './env/index';
export default {
  env: {
    ...env,
  },
  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    title: 'web_complier_fe',
    htmlAttrs: {
      lang: 'en'
    },
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' },
      { name: 'format-detection', content: 'telephone=no' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
      { rel: 'stylesheet' ,href: "https://cdn.jsdelivr.net/npm/@mdi/font@4.x/css/materialdesignicons.min.css"},
    ]
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: [
    // lib css
    'codemirror/lib/codemirror.css',
    // merge css
    'codemirror/addon/merge/merge.css',
    {src: '~/assets/css/reset.scss', lang: 'scss'}
  ],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: [
    '~/plugins/axios.js',
    '~/plugins/icon.js',
  ],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
  ],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
    '@nuxtjs/axios'
  ],

  axios: {
    // baseUrl: '',
    credential: true,
    debug: false,
    proxyHeaders: true,
    retry: false,
    proxy: {},
    progress: true,
  },

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {
    extend(config, ctx) {
      // 在默认处理svg文件中exclude掉
      const svgRule = config.module.rules.find(rule => rule.test.test('.svg'));
      svgRule.exclude = [path.resolve(__dirname, 'assets/images/svg')];
      // 使用svg-sprite-loader处理指定文件夹下的svg
      config.module.rules.push({
        test: /\.svg$/,
        loader: 'svg-sprite-loader',
        include: [path.resolve(__dirname, 'assets/images/svg')],
        options: {
          limit: 1000,
          symbolId: 'icon-[name]'
        },
      });
    },
  }
}
