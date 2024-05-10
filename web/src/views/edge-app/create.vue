<template>
  <div class="app-container">
    <MonacoEditor
      height="600"
      width="100%"
      language="yaml"
      :code="code"
      :editorOptions="options"
      @mounted="onMounted"
      @codeChange="onCodeChange"
    >
    </MonacoEditor>
    <el-row style="margin-top: 10px">
      <el-button type="primary" @click="createApp">提交</el-button>
    </el-row>
  </div>
</template>
<script>
  import MonacoEditor from 'vue-monaco-editor'
  import { create } from '@/api/edge-app'
  import { Message } from 'element-ui'

  export default {
    components: {
      MonacoEditor
    },
    data() {
      return {
        editor: null,
        code: '',
        options: {
          selectOnLineNumbers: false
        },
        randomkey: 1231231
      }
    },
    methods: {
      onMounted(editor) {
        this.editor = editor
      },
      onCodeChange(editor) {
        this.code = editor.getValue()
      },
// createRamdomKey随机生成值，其值类似于id。该方法最为重要，在给code赋值之后，调用这个方法
      createRandomKey() {
        return Math.floor(Math.random() * (10, 1000000012313))
      },
      getDetail() {
        // res为请求返回的值
        this.$nextTick((res) => {
          this.code = res.sourceFormula
          this.randomkey = this.createRandomKey()
        })
      },
      createApp() {
        //创建Device
        return new Promise((resolve, reject) => {
          create(this.code
            // {
            //   declare: this.code,
            // }
          ).then(response => {
            Message({
              message: '创建成功',
              type: 'success',
              duration: 5 * 1000
            })
            console.log(response)
            //路由跳转
            resolve()
            this.$router.push({ path: '/edge-app' })
          }).catch(error => {
            reject(error)
          })
        })
      }
    }
  }
</script>
<style scoped>
</style>
