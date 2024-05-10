<template>
  <div class="app-container">
    <div class="filter-container">
      <!--<router-link :to="'/edge-device/device-model/create'">-->
      <el-button class="filter-item" type="primary" @click="dialogVisible2create = true">
        创建设备模板
        <i class="el-icon-circle-plus"/>
      </el-button>
      <!--</router-link>-->
    </div>
    <el-table :data="list" fit style="width: 100%;">
      <el-table-column prop="name" label="名称" width="200">
        <template slot-scope="{row}">
          <svg-icon icon-class="status" style="color: green;margin-right: 10px"/>
          {{ row.name }}
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
      <el-table-column label="字段属性" width="400">
        <template slot-scope="{row}">
          <el-tag class="my-tag" v-for="(item,key) in row.properties" :key="key">{{ item }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作">
        <template slot-scope="{row}">
          <el-dropdown style="cursor: pointer" placement="bottom-start">
            <span class="el-dropdown-link">
              <svg-icon icon-class="3points-vertical"/>
            </span>
            <el-dropdown-menu slot="dropdown">
              <el-dropdown-item @click.native="showDeviceModelDetail(row.name)">查看</el-dropdown-item>
              <el-dropdown-item @click.native="deleteDeviceModelView(row.name)">删除</el-dropdown-item>
              <el-dropdown-item @click.native="updateDm(row.name)">更新</el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>
        </template>
      </el-table-column>
    </el-table>
    <el-dialog
      title="提示"
      :visible.sync="dialogVisible"
      width="30%"
      :before-close="handleClose"
    >
      <span>删除设备模板后无法恢复，确定删除？</span>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="deleteDeviceModel">确 定</el-button>
      </span>
    </el-dialog>
    <el-dialog
      title="设备模板详情"
      :visible.sync="dialogVisible2detail"
      width="30%"
    >
            <pre>
              {{detailContent}}
            </pre>
<!--      <pre>-->
<!--      <json-viewer-->
<!--        :value=JSON.parse(detailContent)-->
<!--        :expand-depth=5-->
<!--        copyable="true"-->
<!--        boxed-->
<!--        sort></json-viewer>-->
<!--</pre>-->
    </el-dialog>
    <el-dialog
      title="创建设备模板"
      :visible.sync="dialogVisible2create"
      width="30%"
    >
      <create-dm-dialog @reload="reload"/>
    </el-dialog>
    <el-dialog
      title="更新设备模板"
      :visible.sync="dialogVisible2update"
      width="30%"
    >
      <update-dm-dialog @reload="reload"/>
    </el-dialog>
  </div>
</template>
<script>
import {fetchDeviceModelList} from '@/api/edge-device'
import {dateTimeFormat} from '@/utils'
import {Message} from 'element-ui'
import CreateDmDialog from '../components/create-dm-dialog'
import UpdateDmDialog from '../components/update-dm-dialog'

export default {
  name: 'DevicesModel',
  components: {CreateDmDialog, UpdateDmDialog},
  data() {
    return {
      dialogVisible2detail: false,
      dialogVisible2create: false,
      dialogVisible2update: false,
      dialogVisible: false,
      list: null,
      detailContent: null
    }
  },
  created() {
    this.getList()
  },
  methods: {
    dateTimeFormat,
    reload() {
      this.dialogVisible2update = false
      this.dialogVisible2create = false
      this.getList()
    },
    updateDm(name) {
      this.$store.dispatch('edge-device/setDeviceModelName', name)
      this.dialogVisible2update = true
    },
    //设备模板详情展示
    showDeviceModelDetail(name) {
      this.dialogVisible2detail = true
      this.$store.dispatch('edge-device/setDeviceModelName', name)
      this.$store.dispatch('edge-device/getDeviceModelItem')
        .then(response => {
          this.detailContent = JSON.stringify(response.data, null, 4)
        })
    },
    //更新设备模板信息
    updateDeviceModelView(name) {
      this.$store.dispatch('edge-device/setDeviceModelName', name)
      this.$router.push({path: '/edge-device/device-model/update'})
    },
    deleteDeviceModelView(name) {
      this.$store.dispatch('edge-device/setDeviceModelName', name)
      this.dialogVisible = true
    },
    deleteDeviceModel() {
      this.$store.dispatch('edge-device/deleteDeviceModel').then(
        () => {
          Message({
            message: '删除模板创建成功',
            type: 'success',
            duration: 5 * 1000
          })
          this.dialogVisible = false
          //刷新数据
          this.getList()
        }
      )
    },
    handleClose(done) {
      this.$confirm('确认删除？')
        .then(_ => {
          done()
        })
        .catch(_ => {
        })
    },
    getList() {
      this.$store.dispatch('edge-device/getDeviceModelList')
        .then(response => {
          let result = []
          response.data.items.forEach((item, index) => {
            let properties = []
            //遍历所有的字段属性
            item.spec.properties.forEach((item_prop, index_prop) => {
              properties.push(item_prop.name)
            })
            result.push(
              {
                name: item.metadata.name,
                createTime: item.metadata.creationTimestamp,
                properties: properties
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
