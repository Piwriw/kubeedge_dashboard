<template>
  <div class="app-container">
    <div v-for="item in cmds" :key="item.title">
      <h3>
        {{ item.title }}
      </h3>
      <div class="code-block">
        <div class="code">
          {{ item.cmd }}
        </div>
        <div v-clipboard:copy="item.cmd" v-clipboard:success="clipboardSuccess" class="copy">
          复制
        </div>
      </div>
    </div>
  </div>
</template>
<script>
  import clip from '@/utils/clipboard' // use clipboard directly
  import clipboard from '@/directive/clipboard/index.js' // use clipboard by v-directive
  export default {
    name: 'EdgeNodeCreate',
    directives: {
      clipboard
    },
    data() {
      return {
        cmds: [
          {
            id: 1,
            title: '检查Docker环境',
            cmd: 'docker --version'
          },
          {
            id: 2,
            title: '检查edgecore状态',
            cmd: 'systemctl status edgecore.service'
          },
          {
            id: 3,
            title: '修改edgecore.yaml的token,token如下',
            cmd: null
          }
        ]
      }
    },
    created() {
      this.getJoinToken()
    },
    methods: {
      getJoinToken() {
        this.$store.dispatch('edge-node/getJoinToken')
          .then(response => {
            let token = response.data
            this.cmds.forEach((item, index) => {
              //防止乱序
              if (item.id === 3) {
                this.cmds[index].cmd = token
              }
            })
          })
      },
      handleCopy(text, event) {
        clip(text, event)
      },
      clipboardSuccess() {
        this.$message({
          message: 'Copy successfully',
          type: 'success',
          duration: 1500
        })
      }
    }
  }
</script>
<style scoped>
  .code-block {
    background: #eeeeee;
    padding: 15px;
    width: 80%;
    border-radius: 10px;
  }

  .code-block > * {
    display: inline-block;
  }

  .code {
    width: 95%;
    word-wrap: break-word
  }

  .copy {
    float: right;
    cursor: pointer;
    color: #5a5e66;
  }

  .output {
    padding: 10px;
    margin: 5px 0;
    border-left: 5px solid dodgerblue;
  }
</style>
