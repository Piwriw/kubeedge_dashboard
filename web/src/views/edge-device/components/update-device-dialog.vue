<template>
  <div>
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
          <el-col :span="8" style="line-height: 30px">
            上报值：
          </el-col>
          <el-col :span="16">
            <el-input :disabled="true" v-model="item.reported.value" />
          </el-col>
        </el-row>
      </el-col>
      <el-col :span="7" style="padding-left: 10px;text-align: center">
        <el-row>
          <el-col :span="8" style="line-height: 30px">
            期望值：
          </el-col>
          <el-col :span="16">
            <el-input v-model="item.desired.value" />
          </el-col>
        </el-row>
      </el-col>
      <el-col :span="3" style="text-align:center;line-height: 40px">
        <i class="el-icon-refresh" @click="syncDeviceStatus(key)" />
      </el-col>
    </el-row>
  </div>
</template>
<script>
  import { Message } from 'element-ui'

  export default {
    name: 'UpdateDeviceDialog',
    data() {
      return {
        twins: null
      }
    },
    created() {
      this.init()
    },
    methods: {
      init() {
        this.$store.dispatch('edge-device/getDeviceItem')
          .then(response => {
            this.twins = response.data.status.twins
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
