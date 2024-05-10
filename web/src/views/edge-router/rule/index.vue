<template>
  <div class="app-container">
    <div class="filter-container">
      <el-button class="filter-item" type="primary" @click="dialogVisible2create = true">
        创建路由规则 <i class="el-icon-circle-plus"/>
      </el-button>
    </div>
    <el-table :data="list" fit style="width: 100%;">
      <el-table-column prop="name" label="名称" width="200">
        <template slot-scope="{row}">
          <svg-icon icon-class="status" style="color: green;margin-right: 10px"/>
          <a style="color: blue" @click="showRuleDetail(row.name)">
            {{ row.name }}
          </a>
        </template>
      </el-table-column>
      <el-table-column prop="source" label="路由起点" width="200"/>
      <el-table-column prop="target" label="路由终点" width="200"/>
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
              <el-dropdown-item @click.native="showRuleDetail(row.name)">查看</el-dropdown-item>
              <el-dropdown-item @click.native="deleteRuleView(row.name)">删除</el-dropdown-item>
              <el-dropdown-item>更新</el-dropdown-item>
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
      <span>删除路由规则后无法恢复，确定删除？</span>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="deleteRule">确 定</el-button>
      </span>
    </el-dialog>
    <el-dialog
      title="路由规则详情"
      :visible.sync="dialogVisible2detail"
      width="30%"
    >
      <pre>
        {{ detailContent }}
    </pre>

    </el-dialog>
    <el-dialog
      title="创建路由规则"
      :visible.sync="dialogVisible2create"
      width="40%"
    >
      <create-dialog @reload="reload"/>
    </el-dialog>
  </div>
</template>
<script>
import {fetchRuleList} from '@/api/edge-router'
import {dateTimeFormat} from '@/utils'
import {Message} from 'element-ui'
import CreateDialog from "../components/create-rule-dialog"

export default {
  name: 'Rule',
  components: {CreateDialog},
  data() {
    return {
      dialogVisible2detail: false,
      dialogVisible2create: false,
      dialogVisible: false,
      detailContent: null,
      list: null
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
    //路由规则详情展示
    showRuleDetail(name) {
      this.dialogVisible2detail = true
      this.$store.dispatch('edge-router/setRuleName', name)
      this.$store.dispatch('edge-router/getRuleItem')
        .then(response => {
          this.detailContent = JSON.stringify(response.data, null, 4)
        })
    },
    deleteRuleView(name) {
      this.$store.dispatch('edge-router/setRuleName', name)
      this.dialogVisible = true
    },
    deleteRule() {
      this.$store.dispatch('edge-router/deleteRule').then(
        () => {
          Message({
            message: '删除路由规则成功',
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
      this.$store.dispatch('edge-router/getRuleList')
        .then(response => {
          let result = []
          response.data.items.forEach((item, _) => {
            result.push(
              {
                name: item.metadata.name,
                createTime: item.metadata.creationTimestamp,
                source: item.spec.source,
                sourceResource: item.spec.sourceResource,
                targetResource: item.spec.targetResource,
                target: item.spec.target
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
