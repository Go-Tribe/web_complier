<template>
  <div class="nav-bar">
    <a :href="baseConfig.baseUrl" class="logo">
      <img src="https://gotribe.cn/_nuxt/img/logo.1694234.png" alt="在线运行" class="logo-img" />
      <span>在线运行</span>
    </a>
    <ul class="nav-link">
      <li v-for="(nav, index) in navArr" :key="'top-nav-'+index">
        <a :href="nav.link" target="_blank">{{nav.title}}</a>
      </li>
    </ul>
  </div>
</template>
<script>
  import baseConfig from '~/web_complier.js';
  export default {
    name: 'TopBar',
    data() {
      return {
        baseConfig,
      }
    },
    computed: {
      navArr() {
        return this.baseConfig.topNavArr;
      }
    },
    methods: {
    }
  }
</script>
<style lang="scss" scoped>
  .nav-bar {
    position: relative;
    height: 100%;
    margin: 0 auto;
    padding: 0 16px;
    display: flex;
    align-items: center;
    justify-content: space-between;
    .logo {
      // position: absolute;
      left: 16px;
      display: flex;
      align-items: center;
      color: #222226;
      font-weight: bold;
      .logo-img {
        width: 30px;
        height: 30px;
        margin-right: 12px;
      }
    }
    .nav-link {
      height: 100%;
      display: flex;
      align-items: center;
      li {
        margin: 0 4px;
        padding: 8px;
      }
    }
  }
</style>