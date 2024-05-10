<template>
  <div>
    <el-form ref="form" :rules="rules" :model="form" label-width="130px">
      <el-form-item
        prop="name.value"
      >
        <span slot="label">
          {{ form.name.label.title }}
          <el-tooltip placement="right">
            <div slot="content">
              {{ form.name.label.desc }}
            </div>
            <i class="el-icon-question" />
          </el-tooltip>
        </span>
        <el-input style="width: 80%" v-model="form.name.value" :placeholder="form.name.placeholder" />
      </el-form-item>
      <el-form-item
        v-for="(item,key) in form.properties" :key="key"
        :prop="'properties['+key+'].value'"
      >
        <span slot="label">
          {{ item.label.title + (key+1) }}
          <el-tooltip placement="right">
            <div slot="content">
              {{item.label.desc }}
            </div>
            <i class="el-icon-question" />
          </el-tooltip>
        </span>
        <el-input style="width: 80%" v-model="item.value" :placeholder="item.placeholder" />
        <div class="property-container">
          <i v-if="key==form.properties.length-1" @click="addPropView" class="el-icon-circle-plus" />
          <i v-else class="el-icon-remove" @click="removePropView(key)" />
        </div>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit('form')">立即创建</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
  import { Message } from 'element-ui'
  //字段属性
  const PROPERTY_ITEM = {
    label: {
      title: '字段属性',
      desc: '该模板定义的字段'
    },
    value: '',
    placeholder: '请输入设备模板字段属性'
  }
  //校验规则
  const CHECK_REQUIRED = { required: true, message: '请输入设备模板字段属性', trigger: 'blur' }
  export default {
    name: 'CreateDmDialog',
    data() {
      return {
        form: {
          name: {
            label: {
              title: '名称',
              desc: '设备模板名称'
            },
            value: '',
            placeholder: '请输入设备模板名称'
          },
          properties: [
            JSON.parse(JSON.stringify(PROPERTY_ITEM))
          ]
        },
        rules: {
          'name.value': [
            { required: true, message: '请输入模板名称', trigger: 'blur' }
          ],
          'properties[0].value': [
            CHECK_REQUIRED
          ]
        }
      }
    },
    methods: {
      //移除属性
      removePropView(index) {
        this.form.properties.splice(index, 1)
      },
      //添加属性
      addPropView() {
        //深拷贝
        this.form.properties.push(
          JSON.parse(JSON.stringify(PROPERTY_ITEM))
        )
        let key = 'properties[' + (this.form.properties.length - 1) + '].value'
        this.rules[key] = [CHECK_REQUIRED]
      },
      onSubmit(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            let properties = []
            this.form.properties.forEach((item, index) => {
              properties.push(item.value)
            })
            let params = {
              name: this.form.name.value,
              properties: properties
            }
            this.$store.dispatch(
              'edge-device/createDeviceModel', params
            ).then(() => {
              Message({
                message: '设备模板创建成功',
                type: 'success',
                duration: 5 * 1000
              })
              //页面跳转
              this.$router.push({ path: '/edge-device/device-model' })
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
  .property-container {
    /*background: red;*/
    width: 15%;
    text-align: center;
    display: inline-block
  }
  i {
    font-size: 20px;
    cursor: pointer;
  }

  i:hover {
    color: #409EFF;
  }
</style>
