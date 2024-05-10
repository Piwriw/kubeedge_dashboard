<template>
  <div class="app-container">
    <div class="filter-container">
      <router-link :to="'/edge-app/create'">
        <el-button class="filter-item" type="primary">
          创建应用 <i class="el-icon-circle-plus"/>
        </el-button>
      </router-link>
    </div>
    <el-table :data="list" fit style="width: 100%;">
      <el-table-column prop="name" label="NAME" width="320">
        <template slot-scope="{row}">
          <svg-icon icon-class="status" style="margin-right: 10px;"
                    :style="{color:(row.status===0)?'green':'red'}"/>
          <a style="color: blue" @click="showAppDetail(row.name)">
            {{ row.name }}
          </a>
        </template>
      </el-table-column>
      <el-table-column label="镜像" width="320">
        <template slot-scope="{row}">
          <p v-for="(item,key) in row.images">{{ item }}</p>
        </template>
      </el-table-column>
      <el-table-column label="标签">
        <template slot-scope="{row}">
          <el-tag class="my-tag" v-for="(value,key,index) in row.labels">{{ key }}:{{ value }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="node" label="节点" width="150"/>
      <el-table-column prop="age" label="创建时间">
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
              <el-dropdown-item @click.native="showAppDetail(row.name,row.namespace)">查看</el-dropdown-item>
              <el-dropdown-item @click.native="showDeleteApp(row.name,row.namespace)">删除</el-dropdown-item>
              <!--<el-dropdown-item>升级</el-dropdown-item>-->
              <!--<el-dropdown-item>日志</el-dropdown-item>-->
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
      <span>删除以后应用无法恢复，确定删除？</span>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="deleteApp">确 定</el-button>
      </span>
    </el-dialog>
    <el-dialog
      title="应用详情"
      :visible.sync="dialogVisible2detail"
      width="30%"
    >
      <pre>
        {{ detailContent }}
      </pre>
<!--      <json-viewer-->
<!--        :value=detailContent-->
<!--        :expand-depth=5-->
<!--        copyable="true"-->
<!--        boxed-->
<!--        sort></json-viewer>-->
    </el-dialog>
  </div>
</template>
<script>
import {dateTimeFormat} from '@/utils'
import {Message} from 'element-ui'

export default {
  name: 'Applications',
  data() {
    return {
      dialogVisible: false,
      dialogVisible2detail: false,
      detailContent: '',
      list: null
    }
  },
  created() {
    this.getList()
  },
  methods: {
    dateTimeFormat,
    showAppDetail(name, namespace) {
      this.dialogVisible2detail = true
      let data = {
        name,
        namespace
      }
      this.$store.dispatch('edge-app/setAppName', name)
      this.$store.dispatch('edge-app/getAppItem', data)
        .then(response => {
          this.detailContent = JSON.stringify(response.data, null, 4)
        })
    },
    showDeleteApp(name, namespace) {
      this.dialogVisible = true
      let data = {
        name, namespace
      }
      this.$store.dispatch('edge-app/setAppName', data)
    },
    deleteApp() {
      this.$store.dispatch('edge-app/deleteApp')
        .then(_ => {
          Message({
            message: '删除应用成功',
            type: 'success',
            duration: 5 * 1000
          })
          this.dialogVisible = false
          this.getList()
        })
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
      this.$store.dispatch('edge-app/getAppList').then(response => {
        let result = []
        response.data.items.forEach((item, _) => {
          let images = []
          //计算镜像
          item.spec.containers.forEach((item_c, index_c) => {
            images.push(
              item_c.image
            )
          })
          result.push({
            name: item.metadata.name,
            createTime: item.metadata.creationTimestamp,
            labels: item.metadata.labels,
            images: images,
            node: item.spec.nodeName,
            status: item.status.phase === 'Running' ? 0 : 1,
            namespace: item.metadata.namespace
          })
        })
        this.list = result
      })
    }
  }
}
</script>
<style scoped>
</style>
