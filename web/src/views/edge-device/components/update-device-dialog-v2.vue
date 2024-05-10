<template>
  <div>
    <el-row style="margin-bottom: 10px">
      <el-col :span="7">
        <el-row>
          <el-col :span="8" style="line-height: 30px;text-align: center">
            <strong>模型版本：</strong>
          </el-col>
          <el-col :span="16">
            <el-select placeholder="请选择AI模型" @change="changeValue" v-model="modeInfo.value">
              <el-option
                v-for="op in modeInfo.options"
                :key="op"
                :label="op"
                :value="op"
              />
            </el-select>
          </el-col>
        </el-row>
      </el-col>
    </el-row>
    <el-row v-for="(item,key) in twins" :key="key">
      <el-col :span="7">
        <el-row style="text-align: center">
          <el-col :span="8" style="line-height: 30px">
            字段名称：
          </el-col>
          <el-col :span="16">
            <el-input :disabled="true" v-model="item.propertyName" />
          </el-col>
        </el-row>
      </el-col>
      <el-col :span="7">
        <el-row style="text-align: center">
          <el-col :span="6" style="line-height: 30px">
            上报值：
          </el-col>
          <el-col :span="18">
            <el-input :disabled="true" v-model="item.reported.value" />
          </el-col>
        </el-row>
      </el-col>
      <el-col :span="7" style="padding-left: 10px;text-align: center">
        <el-row>
          <el-col :span="6" style="line-height: 30px">
            期望值：
          </el-col>
          <el-col :span="18">
            <el-input v-model="item.desired.value" />
          </el-col>
        </el-row>
      </el-col>
      <el-col :span="3" style="text-align:center;line-height: 40px">
        <i class="el-icon-refresh" @click="syncDeviceStatus(key)" />
      </el-col>
    </el-row>
    <el-row style="text-align: center;margin-top: 10px">
      <el-button type="primary" @click="syncDeviceStatusAll">一键同步所有字段</el-button>
    </el-row>
  </div>
</template>
<script>
  import { Message } from 'element-ui'

  export default {
    name: 'UpdateDeviceDialogV2',
    data() {
      return {
        modeInfo: {
          options: [],
          value: ''
        },
        twins: null
      }
    },
    created() {
      this.init()
    },
    methods: {
      changeValue(value) {
        this.$store.dispatch('edge-aimodel/getAiModelItem', value)
          .then(response => {
            //依次填充twins期望值
            this.twins.forEach((item, index) => {
              item.desired.value = response[item.propertyName]
            })
          })
      },
      init() {
        this.$store.dispatch('edge-device/getDeviceItem')
          .then(response => {
            this.twins = response.data.status.twins
          })
        this.$store.dispatch('edge-aimodel/getAiModelList').then(response => {
          console.log(response)
          response.data.forEach((item, index) => {
            this.modeInfo.options.push(item.version)
          })
        })
      },
      syncDeviceStatusAll() {
        let _temp = {}
        this.twins.forEach((item, index) => {
          _temp[item['propertyName']] = item['desired']['value']
        })
        let params = {
          'desired': _temp
        }

        this.$store.dispatch('edge-device/updateDevice', params)
          .then(_ => {
            Message({
              message: '设备状态更新成功',
              type: 'success',
              duration: 5 * 1000
            })
          })
      },
      //同步设备状态
      syncDeviceStatus(index) {
        let item = this.twins[index]
        let _temp = {}
        _temp[item['propertyName']] = item['desired']['value']
        let params = {
          'desired': _temp
        }
        this.$store.dispatch('edge-device/updateDevice', params)
          .then(_ => {
            Message({
              message: '设备属性' + item['propertyName'] + '状态更新成功',
              type: 'success',
              duration: 5 * 1000
            })
          })
      }
    }
  }
</script>
<style scoped>
  .el-icon-refresh {
    font-size: 20px;
    cursor: pointer;
  }

  .el-icon-refresh:hover {
    color: #409EFF;
  }
</style>
