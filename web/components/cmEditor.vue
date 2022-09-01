<template>
  <section class="container">
    <div class="editor-container">
      <div class="header">
        <div>
          <select id="languageSelect" name="language_type" class="language">
            <option v-for="(lang, index) in languages" :key="'lang-'+index" :value="lang.value" :selected="lang.default">{{lang.name}}</option>
          </select>
          <button class="share-code" @click="shareCode">
            分享
            <svg-icon iconClass="share" />
          </button>
          <input v-show="shareLink" v-model="shareLink" style="height: 26px;width:200px" />
        </div>
        <div class="action-btn">
          <button class="run-code" @click="submit" :disabled="isRunning">
            <span v-show="isRunning" class="loading" :class="{'rotate': isRunning}"></span>
            <span>执行代码</span>
            <svg-icon iconClass="run" />
          </button>
          <!-- <button class="submit" @click="getContent">提交</button> -->
        </div>
      </div>
      <codemirror
        ref="myCm"
        v-model="code"
        :options="cmOptions"
        class="cm-editor"
        @change="onChange"
        @blur="onBlur"
        @ready="loadFromLocal"
        @inputRead="onCursorActivity"
      >
      </codemirror>
    </div>
    <div class="resize-line"></div>
    <div class="result-container">
      <pre>{{result}}</pre>
    </div>
  </section>
</template>
<script>
  import 'codemirror/theme/monokai.css';
  import 'codemirror/theme/midnight.css';
  import 'codemirror/addon/hint/show-hint.css';
  import goTemplate from '~/code_template/template_go.js';
  import baseConfig from '~/web_complier.js';
  import Vue from 'vue';

  const shareParamName = 's';

  export default {
    name: 'CMEditor',
    data() {
      const code = goTemplate;
      return {
        code: code,
        cmOptions: {
          tabSize: 2,
          foldGutter: true,
          theme: 'monokai',
          lineNumbers: true,
          line: true,
          spellcheck: true,
          highlightDifferences: true,
          viewportMargin: 100,
          // 代码提示配置
          hintOptions: {
            completeSingle: false
          },
          // 代码折叠配置
          gutters: [
            "CodeMirror-lint-markers",
            "CodeMirror-linenumbers",
            "CodeMirror-foldgutter",
          ],
          foldGutter: true,
        },
        timer: null,
        languages: [
          {lang: 'golang', value: 'go', name: 'Go 1.8', default: true},
          // {lang: 'c', value: 'c', name: 'c', default: false},
          // {lang: 'c++', value: 'c++', name: 'c++', default: false},
          // {lang: 'java', value: 'java', name: 'Java', default: false},
          // {lang: 'python', value: 'python', name: 'Python', default: false},
          // {lang: 'javascript', value: 'javascript', name: 'Javascript', default: true},
        ],
        langMap: {
          'go': 'golang',
          'c': 'c',
          'c++': 'c++',
          'javascript': 'javascript',
          'java': 'java',
          'python': 'python'
        },
        curLang: '',
        isRunning: false,
        result: '',
        shareLink: '',
        shareContent: {
          code: '',
        },
        baseUrl: baseConfig.baseUrl,
      }
    },
    computed: {
      codemirror() {
        return this.$refs.myCm.codemirror;
      }
    },
    created() {
      if(process.browser) {
        // 语言
        require('codemirror/mode/javascript/javascript.js');
        require('codemirror/mode/go/go.js');
        // 折叠
        require('codemirror/addon/fold/foldgutter.css');
        require('codemirror/addon/fold/foldcode');
        require('codemirror/addon/fold/foldgutter');
        require('codemirror/addon/fold/brace-fold');
        require('codemirror/addon/fold/comment-fold');
        require('codemirror/addon/fold/indent-fold');
        // 代码提示
        require('codemirror/addon/hint/show-hint.js');
        require('codemirror/addon/hint/javascript-hint.js');
        require('codemirror/addon/hint/show-hint.css');
        const VueCodemirror = require('vue-codemirror');
        Vue.use(VueCodemirror);
      }
    },
    mounted() {
      this.codemirror.setSize('auto', 'calc(100vh - 80px)');
      const languageSelect = document.querySelector('#languageSelect');
      this.curLang = this.langMap[languageSelect.value] || languageSelect.value;
      // 是分享链接
      if(this.$route.query[shareParamName]) {
        this.getCodeByLink(this.$route.query[shareParamName]).then(res => {
          this.codemirror.setOption('mode', res.lang);
          this.codemirror.setValue(res.code);
        });
      } else {
        // 正常进入
        this.codemirror.setOption('mode', languageSelect.value);
      }
      languageSelect.onchange = () => {
        this.codemirror.setOption('mode', languageSelect.value);
        this.curLang = this.langMap[languageSelect.value] || languageSelect.value;
        console.log(languageSelect.value);
      }
    },
    methods: {
      getContent() {
        // const content = this.codemirror.getValue();
        const content = this.codemirror.getValue();
        console.log(content);
      },
      onChange(editor, changeObj) {
        console.log('change了', editor, changeObj);
        if(this.timer) clearTimeout(this.timer);
        this.timer = setTimeout(this.saveContent2Local, 5000);
      },
      // 从浏览器localStorage中获取历史缓存内容
      loadFromLocal() {
        const localCode = localStorage.getItem(`${this.curLang}_code`);
        if(localCode) {
          this.codemirror.setValue(localCode);
        }
      },
      // 将内容存到localStorage
      saveContent2Local() {
        console.log('将内容缓存到localStorage');
        localStorage.setItem(`${this.curLang}_code`, this.codemirror.getValue())
      },
      onBlur(editor, obj) {
        console.log('blur', editor, obj);
      },
      submit() {
        this.isRunning = true;
        const paramsObj = {
          lang: this.curLang,
          code: this.code,
        }
        console.log(paramsObj);
        this.$http.post(
          `http://www.run.com/api/v1/run`,
          paramsObj,
          {headers: {'Content-Type': 'application/json'}}
        ).then(res => {
          console.log(res);
          this.result = res.stdout;
        }).catch((error) => {
          console.log(error);
        }).finally(() => {
          this.isRunning = false;
        })
      },
      onCursorActivity(cm, obj) {
        if(obj.text && obj.text.length) {
          let c = obj.text[0][obj.text[0].length - 1];
          if(/^[a-zA-z]*$/.test(c)) {
            this.codemirror.showHint();
          }
        }
      },
      // 分享代码
      shareCode() {
        const paramsObj = {
          lang: this.curLang,
          code: this.code,
        }
        this.$http.post(
          `/api/v1/share`,
          paramsObj,
          {headers: {'Content-Type': 'application/json'}}
        ).then(res => {
          console.log('保存', res);
          this.shareLink = baseConfig.baseUrl + `?${shareParamName}=${res.url}`;
        })
      },
      // 根据分享链接获取分享代码内容
      getCodeByLink(shareLink='') {
        if(!shareLink) return;
        return new Promise((resolve, reject) => {
          this.$http.get(
            `/api/v1/code`,
            {params: {gid: shareLink}}
          ).then(res => {
            console.log('获得分享内容', res);
            this.shareContent = res;
            resolve(this.shareContent);
          }).catch((error) => {
            console.error(error);
          })
        })
      }
    },
    beforeDestroy() {
      if(this.timer) clearTimeout(this.timer);
    }
  }
</script>
<style lang="scss" scoped>
  .container {
    display: flex;
    width: 100%;
    margin: 0 auto;
    top: 0;
    background-color: #313741;
    .editor-container {
      flex: 1 1 0;
      height: 100%;
      .header {
        height: 40px;
        padding: 0 10px;
        display: flex;
        align-items: center;
        justify-content: space-between;
        .language {
          // padding: 4px;
        }
        .share-code {
          min-width: 70px;
          color: #2db55d;
          border: 1px solid #2db55d;
          border-radius: 4px;
          background-color: #fff;
          padding: 4px 8px;
        }
        .action-btn {
          // padding: 0px 12px;
          display: flex;
          column-gap: 12px;
          align-items: center;
          justify-content: flex-end;
          .run-code {
            min-width: 70px;
            color: #2db55d;
            border: 1px solid #2db55d;
            border-radius: 4px;
            background-color: #fff;
            padding: 4px 8px;
            >* {
              vertical-align: middle;
            }
            .loading {
              display: inline-block;
              width: 14px;
              height: 14px;
              border: 1px solid;
              border-top-color: transparent;
              border-radius: 100%;
            }
            .rotate {
              animation: circle infinite 0.75s linear;
            }

            // 转转转动画
            @keyframes circle {
              0% {
                transform: rotate(0);
              }
              100% {
                transform: rotate(360deg);
              }
            }
          }
          .submit {
            padding: 4px 8px;
            color: #fff;
            border: 1px solid #2db55d;
            border-radius: 4px;
            background-color: #2db55d;
          }
        }
      }
      .cm-editor {
        font-size: 16px;
      }
      :deep().CodeMirror {
        height: auto;
      }
    }
    .resize-line {
      flex: 0 0 8px;
      // height: 100vh;
      background-color: aquamarine;
      cursor: col-resize;
    }
    .result-container {
      flex: 1 0 0;
      white-space: pre-line;
      color: #fff;
    }
  }
</style>
