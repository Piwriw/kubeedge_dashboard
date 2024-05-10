<template>
  <div class="app-container">
    <div class="filter-container">
      <el-button class="filter-item" type="primary" @click="createDeviceView">
        创建实例 <i class="el-icon-circle-plus"/>
      </el-button>
    </div>
    <el-table :data="list" fit style="width: 100%;">
      <el-table-column prop="name" label="设备名称" width="200">
        <template slot-scope="{row}">
          <svg-icon icon-class="status" style="color: green;margin-right: 10px"/>
          <a @click="showDeviceDetail(row.name)" style="color: blue">
            {{ row.name }}
          </a>
        </template>
      </el-table-column>
      <el-table-column prop="age" label="创建时间" width="150">
        <template slot-scope="{row}" style="cursor: pointer">
          <el-tooltip class="item" effect="dark" :content="row.createTime" placement="right-start">
            <el-button style="border: none;outline: none">{{ dateTimeFormat(row.createTime) }} <i
              class="el-icon-info"/></el-button>
          </el-tooltip>
        </template>
      </el-table-column>
      <el-table-column label="操作">
        <template slot-scope="{row}">
          <el-dropdown style="cursor: pointer" placement="bottom-start">
            <span class="el-dropdown-link">
              <svg-icon icon-class="3points-vertical"/>
            </span>
            <el-dropdown-menu slot="dropdown">
              <el-dropdown-item @click.native="showDeviceDetail(row.name)">查看</el-dropdown-item>
              <el-dropdown-item @click.native="deleteDeviceView(row.name)">删除</el-dropdown-item>
              <el-dropdown-item @click.native="updateDeviceView(row.name)">状态同步</el-dropdown-item>
              <el-dropdown-item @click.native="updateDeviceViewV2(row.name)">状态同步v2</el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>
        </template>
      </el-table-column>
    </el-table>
    <el-dialog
      title="提示"
      :visible.sync="dialogVisible"
      width="30%"
    >
      <span>删除设备后无法恢复，确定删除？</span>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="deleteDevice">确 定</el-button>
      </span>
    </el-dialog>
    <el-dialog
      title="设备详情"
      :visible.sync="dialogVisible2detail"
      width="30%"
    >
            <pre>
              {{detailContent}}
            </pre>


<!--      <json-viewer-->
<!--        :value=detailContent-->
<!--        :expand-depth=5-->
<!--        copyable="true"-->
<!--        boxed-->
<!--        sort></json-viewer>-->
    </el-dialog>
    <el-dialog
      title="设备状态同步"
      :visible.sync="dialogVisible2status"
      width="40%"
    >
      <update-device-dialog/>
    </el-dialog>
    <el-dialog
      title="设备状态同步V2"
      :visible.sync="dialogVisible2statusV2"
      width="65%"
    >
      <update-device-dialog-v2/>
    </el-dialog>
    <el-dialog
      title="创建设备实例"
      :visible.sync="dialogVisible2create"
      width="30%"
    >
      <create-device-dialog @reload="reload"/>
    </el-dialog>
  </div>
</template>
<script>
import {fetchDeviceInstanceList} from '@/api/edge-device'
import {dateTimeFormat} from '@/utils'
import {Message} from 'element-ui'
import UpdateDeviceDialog from '../components/update-device-dialog'
import UpdateDeviceDialogV2 from '../components/update-device-dialog-v2'
import CreateDeviceDialog from '../components/create-device-dialog'

export default {
  name: 'DeviceInstance',
  components: {
    UpdateDeviceDialog,
    UpdateDeviceDialogV2,
    CreateDeviceDialog
  },
  data() {
    return {
      dialogVisible2detail: false,
      dialogVisible: false,
      dialogVisible2statusV2: false,
      dialogVisible2create: false,
      dialogVisible2status: false,
      list: null,
      detailContent: '',
      twins: null
    }
  },
  created() {
    this.getList()
  },
  methods: {
    dateTimeFormat,
    reload() {
      this.dialogVisible2create = false
      this.getList()
    },
    createDeviceView() {
      this.dialogVisible2create = true
    },
    //更新设备信息
    updateDeviceView(name) {
      this.$store.dispatch('edge-device/setDeviceName', name)
      this.dialogVisible2status = true
    },
    updateDeviceViewV2(name) {
      this.$store.dispatch('edge-device/setDeviceName', name)
      this.dialogVisible2statusV2 = true
    },
    //设备详情展示
    showDeviceDetail(name) {
      this.dialogVisible2detail = true
      this.$store.dispatch('edge-device/setDeviceName', name)
      this.$store.dispatch('edge-device/getDeviceItem')
        .then(response => {
          this.detailContent = JSON.stringify(response.data, null, 4)
        })
    },
    //删除设备展示
    deleteDeviceView(name) {
      this.$store.dispatch('edge-device/setDeviceName', name)
      this.dialogVisible = true
    },
    //删除设备
    deleteDevice() {
      this.$store.dispatch('edge-device/deleteDevice').then(
        () => {
          Message({
            message: '删除设备成功',
            type: 'success',
            duration: 5 * 1000
          })
          this.dialogVisible = false
          //刷新数据
          this.getList()
        }
      )
    },
    getList() {
      fetchDeviceInstanceList().then(response => {
        let result = []
        response.data.items.forEach((item, _) => {
          result.push(
            {
              name: item.metadata.name,
              createTime: item.metadata.creationTimestamp
            }
          )
        })
        this.list = result
      })
    }
  }
}
</script>
<style scoped>
</style>
