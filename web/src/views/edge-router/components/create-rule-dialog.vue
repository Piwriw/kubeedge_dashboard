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
        <el-input v-if="item.type=='text'" v-model="item.value" :placeholder="item.placeholder" />
        <el-select v-else-if="item.type=='options'" v-model="item.value" style="width: 100%"
                   :placeholder="item.placeholder">
          <el-option
            v-for="ruleEndpoint in ruleEndpointList"
            :key="ruleEndpoint.metadata.name"
            :label="ruleEndpoint.metadata.name+'('+ruleEndpoint.spec.ruleEndpointType+')'"
            :value="ruleEndpoint.metadata.name"
          />
        </el-select>
        <el-input v-else type="textarea" v-model="item.value"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm('form')">立即创建</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
  import { Message } from 'element-ui'

  export default {
    name: 'CreateDialog',
    data() {
      return {
        ruleEndpointList: [],
        form: {
          properties: [
            {
              name: 'name',
              type: 'text',
              label: {
                title: '路由规则名称',
                desc: '路由规则名称'
              },
              value: '',
              rules: [
                { required: true, message: '请输入路由规则名称', trigger: 'blur' }
              ],
              placeholder: '请输入路由规则名称'
            },
            {
              name: 'source',
              type: 'options',
              label: {
                title: '路由起点',
                desc: '路由起点'
              },
              value: '',
              rules: [
                { required: true, message: '请输入路由起点', trigger: 'blur' }
              ],
              placeholder: '请输入路由起点'
            },
            {
              name: 'sourceResource',
              type: 'textarea',
              label: {
                title: '路由起点配置信息',
                desc: '路由起点配置信息'
              },
              value: '',
              rules: [
                { required: true, message: '请输入路由起点配置信息', trigger: 'blur' }
              ],
              placeholder: '请输入路由起点配置信息'
            },
            {
              name: 'target',
              type: 'options',
              label: {
                title: '路由终点',
                desc: '路由终点'
              },
              value: '',
              rules: [
                { required: true, message: '请输入路由终点', trigger: 'blur' }
              ],
              placeholder: '请输入路由终点'
            },
            {
              name: 'targetResource',
              type: 'textarea',
              label: {
                title: '路由终点配置信息',
                desc: '路由终点配置信息'
              },
              value: '',
              rules: [
                { required: true, message: '请输入路由终点配置信息', trigger: 'blur' }
              ],
              placeholder: '请输入路由终点配置信息'
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
        this.$store.dispatch('edge-router/getRuleEndPointList')
          .then(response => {
            this.ruleEndpointList = response.data.items
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
            this.$store.dispatch('edge-router/createRule', params)
              .then(_ => {
                Message({
                  message: '路由创建成功',
                  type: 'success',
                  duration: 5 * 1000
                })
                this.$emit('reload')
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
