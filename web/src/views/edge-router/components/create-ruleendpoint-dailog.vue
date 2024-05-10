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
            v-for="op in item.options"
            :key="op"
            :label="op"
            :value="op"
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
    name: 'CreateRuleendpointDailog',
    data() {
      return {
        form: {
          properties: [
            {
              name: 'name',
              type: 'text',
              label: {
                title: '路由节点名称',
                desc: '路由节点名称'
              },
              value: '',
              rules: [
                { required: true, message: '请输入路由节点名称', trigger: 'blur' }
              ],
              placeholder: '请输入路由节点名称'
            },
            {
              name: 'rule_endpoint_type',
              type: 'options',
              label: {
                title: '路由节点类型',
                desc: '路由节点类型'
              },
              options: ['rest', 'servicebus', 'eventbus'],
              value: '',
              rules: [
                { required: true, message: '请输入路由节点类型', trigger: 'blur' }
              ],
              placeholder: '请输入路由节点类型'
            },
            {
              name: 'properties',
              type: 'textarea',
              label: {
                title: '路由节点配置信息',
                desc: '路由节点配置信息'
              },
              value: '',
              rules: [
              ],
              placeholder: '请输入路由节点配置信息'
            }
          ]
        }
      }
    },
    methods: {
      submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            //构造数据
            let params = {}
            this.form.properties.forEach((item, _) => {
              try {
                params[item['name']] = JSON.parse(item['value'])
              } catch (e) {
                params[item['name']] = item['value']
              }
            })
            let that = this
            this.$store.dispatch('edge-router/createRuleEndPoint', params)
              .then(_ => {
                Message({
                  message: '路由节点创建成功',
                  type: 'success',
                  duration: 5 * 1000
                })
                that.$emit("reload")
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
