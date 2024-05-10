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
    name: 'CreateDeviceDialog',
    data() {
      return {
        form: {
          properties: [
            {
              name: 'name',
              type: 'text',
              label: {
                title: '设备名称',
                desc: '设备名称'
              },
              value: '',
              rules: [
                { required: true, message: '请输入设备名称', trigger: 'blur' }
              ],
              placeholder: '请输入设备名称'
            },
            {
              name: 'dm_name',
              type: 'options',
              label: {
                title: '设备模板',
                desc: '设备模板'
              },
              value: '',
              options: [],
              rules: [
                { required: true, message: '请选择设备模板', trigger: 'blur' }
              ],
              placeholder: '请选择设备模板'
            },
            {
              name: 'node_name',
              type: 'options',
              label: {
                title: '绑定节点',
                desc: '绑定节点'
              },
              value: '',
              options: [],
              rules: [
                { required: true, message: '请选择绑定节点', trigger: 'blur' }
              ],
              placeholder: '请选择绑定节点'
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
        this.getNodeList()
        this.getDmList()
      },
      //获取所有的节点数据
      getNodeList() {
        this.$store.dispatch('edge-node/getNodeList').then(response => {
          response=response.data
          this.form.properties.forEach((item, index0) => {
            if (item.name === 'node_name') {
              response.forEach((node, index1) => {
                item.options.push(node.metadata.name)
              })
            }
          })
        })
      },
      //获取所有的设备模板数据
      getDmList() {
        this.$store.dispatch('edge-device/getDeviceModelList').then(response => {
          response=response.data
          this.form.properties.forEach((item, index0) => {
            if (item.name === 'dm_name') {
              response.items.forEach((dm, index1) => {
                item.options.push(dm.metadata.name)
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
            this.$store.dispatch('edge-device/createDevice', params)
              .then(_ => {
                Message({
                  message: '设备创建成功',
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
