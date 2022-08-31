import Vue from 'vue';
import IconSvg from '~/components/iconComponent.vue';

Vue.component('SvgIcon', IconSvg);
const req = require.context('~/assets/images/svg', false, /\.svg$/);
req.keys().map(req);
