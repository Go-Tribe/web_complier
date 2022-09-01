<template>
  <no-ssr placeholder="codemirror loading...">
    <cm-editor />
  </no-ssr>
</template>

<script>
  export default {
    name: 'IndexPage',
    head() {
      return {
        title: '在线代码运行',
      }
    },
    data() {
      return {

      }
    },
    computed: {
      isClient() {
        return process.browser;
      }
    },
    mounted() {
      // if(process.browser) {
      //   require('codemirror/mode/javascript/javascript.js');
      // //   // require('codemirror/addon/hint/show-hint.js');
      // //   // require('codemirror/addon/hint/javascript-hint.js');
      // //   // require('codemirror/addon/hint/show-hint.css');
      //   const VueCodemirror = require('vue-codemirror');
      //   Vue.use(VueCodemirror);
      // }
    }
  }
</script>
<style scoped>
</style>
