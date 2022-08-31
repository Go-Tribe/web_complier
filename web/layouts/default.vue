<template>
  <div style="position: relative;">
    <header class="top-bar"><TopBar /></header>
    <main class="main-container"><nuxt /></main>
  </div>
</template>
<script>
import TopBar from "../components/topBar.vue";
export default {
  name: 'DefaultLayout',
  components: {
    TopBar,
  }
}
</script>
<style lang="scss">
  $topHeight: 40px;
  $footerHeight: 120px;
  .top-bar {
    display: block;
    height: $topHeight;
    width: 100%;
    position: fixed;
    top: 0;
    background-color: #fff;
    box-shadow: 0px 16px 32px rgba(0, 0, 0, 0.04);
    z-index: 10;
    // border-bottom: 1px solid #000;
  }
  .main-container {
    padding-top: $topHeight;
  }
</style>