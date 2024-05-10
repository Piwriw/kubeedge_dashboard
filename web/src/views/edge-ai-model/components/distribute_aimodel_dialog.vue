<template>
  <div>
    <el-form ref="form" :model="form">
      <el-form-item v-for="(item,key) in form.properties" :key="key"
                    :rules="item.rules"
                    :prop="'properties['+key+'].value'"
      >
        <span slot="label">
          {{ item.label.title }}
        </span>
        <el-input v-if="item.type=='text'" :disabled="item.disabled" v-model="item.value"
                  :placeholder="item.placeholder" />
        <el-select v-else-if="item.type=='options'" v-model="item.value" style="width: 100%"
                   :placeholder="item.placeholder">
          <el-option
            v-for="op in item.options"
            :key="op"
            :label="op"
            :value="op"
          />
        </el-select>
        <el-input v-else type="textarea" v-model="item.value"></el-input>
      </el-form-item>
      <el-form-item v-if="dialogVisible2progress">
        <p>{{distributeStatus}}</p>
        <el-progress :percentage="percentage" :status="progressStatus"></el-progress>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm('form')">立即下发</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
  import { Message } from 'element-ui'

  const DISTRIBUTE_STATUS_ORIGIN = '准备开始'
  const PERCENTAGE = 0

  export default {
    name: 'DistributeAiModelDialog',
    data() {
      return {
        timer: null,
        dialogVisible2progress: false,
        distributeStatus: DISTRIBUTE_STATUS_ORIGIN,
        percentage: PERCENTAGE,
        progressStatus: '',
        form: {
          properties: [
            {
              name: 'version',
              type: 'text',
              label: {
                title: '模型版本',
                desc: '模型版本'
              },
              disabled: true,
              value: '',
              rules: [],
              placeholder: '请输入模型版本'
            },
            {
              name: 'node_name',
              type: 'options',
              label: {
                title: '边缘节点',
                desc: '边缘节点'
              },
              value: '',
              options: [],
              rules: [
                { required: true, message: '请选择边缘节点', trigger: 'blur' }
              ],
              disabled: false,
              placeholder: '请选择边缘节点'
            },
            {
              name: 'rule_name',
              type: 'options',
              label: {
                title: '路由规则',
                desc: '路由规则'
              },
              disabled: false,
              value: '',
              options: [],
              rules: [
                { required: true, message: '请选择路由规则', trigger: 'blur' }
              ],
              placeholder: '请选择路由规则'
            }
          ]
        }
      }
    },
    created() {
      this.init()
    },
    methods: {
      init() {
        //进度条初始化
        this.dialogVisible2progress = false
        this.percentage = PERCENTAGE
        this.distributeStatus = DISTRIBUTE_STATUS_ORIGIN
        //初始模型版本
        this.$store.dispatch('edge-aimodel/getAiModelVersion').then(
          version => {
            this.form.properties.forEach((item, index0) => {
              if (item.name === 'version') {
                item.value = version
              }
            })
          }
        )
        //初始化路由规则选项
        this.$store.dispatch('edge-router/getRuleList')
          .then(response => {
            this.form.properties.forEach((item, index0) => {
              if (item.name === 'rule_name') {
                response.items.forEach((rule, index1) => {
                  item.options.push(rule.metadata.name)
                })
              }
            })
          })
        //初始化节点选型
        this.$store.dispatch('edge-node/getNodeList')
          .then(response => {
            this.form.properties.forEach((item, index0) => {
              if (item.name === 'node_name') {
                response.forEach((node, index1) => {
                  item.options.push(node.metadata.name)
                })
              }
            })
          })
      },
      submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            //构造数据
            let params = {}
            this.form.properties.forEach((item, index) => {
              try {
                params[item['name']] = JSON.parse(item['value'])
              } catch (e) {
                params[item['name']] = item['value']
              }
            })
            this.$store.dispatch('edge-aimodel/distributeModel', params)
              .then(response => {
                let serializeId = response.detail
                //显示进度条
                this.dialogVisible2progress = true
                this.timer = window.setInterval(() => {
                  this.$store.dispatch('edge-aimodel/getDistributeProgress', serializeId)
                    .then(progress => {
                      let detail = progress.detail
                      if (detail === 'done') {
                        this.percentage = 100
                        this.distributeStatus = '模型下发完毕'
                        this.progressStatus = 'success'
                        window.clearInterval(this.timer)
                        return
                      }
                      //判断是否为数字
                      if (!isNaN(detail)) {
                        this.percentage = Math.round(parseFloat(detail) * 100)
                        this.distributeStatus = '模型下发中...'
                      } else {
                        this.distributeStatus = detail
                      }
                    })
                    .catch(() => {
                      //异常中断接口轮询
                      window.clearInterval(this.timer)
                    })
                }, 1000)
              })
          } else {
            return false
          }
        })
      }
    }
  }
</script>
<style scoped>
</style>
