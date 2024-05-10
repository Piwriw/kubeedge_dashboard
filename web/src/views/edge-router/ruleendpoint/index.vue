<template>
  <div class="app-container">
    <div class="filter-container">
      <el-button class="filter-item" type="primary" @click="dialogVisible2create=true">
        创建路由节点 <i class="el-icon-circle-plus" />
      </el-button>
    </div>
    <el-table :data="list" fit style="width: 100%;">
      <el-table-column prop="name" label="节点名称" width="200">
        <template slot-scope="{row}">
          <svg-icon icon-class="status" style="color: green;margin-right: 10px" />
          <a style="color: blue" @click="showDetail(row.name)">
            {{ row.name }}
          </a>
        </template>
      </el-table-column>
      <el-table-column prop="age" label="创建时间" width="150">
        <template slot-scope="{row}" style="cursor: pointer">
          <el-tooltip class="item" effect="dark" :content="row.createTime" placement="right-start">
            <el-button style="border: none;outline: none">{{ dateTimeFormat(row.createTime) }} <i
              class="el-icon-info" /></el-button>
          </el-tooltip>
        </template>
      </el-table-column>
      <el-table-column label="操作">
        <template slot-scope="{row}">
          <el-dropdown style="cursor: pointer" placement="bottom-start">
            <span class="el-dropdown-link">
              <svg-icon icon-class="3points-vertical" />
            </span>
            <el-dropdown-menu slot="dropdown">
              <el-dropdown-item @click.native="showDetail(row.name)">
                查看
              </el-dropdown-item>
              <el-dropdown-item @click.native="deleteRuleEndPointView(row.name)">删除</el-dropdown-item>
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
    >
      <span>删除路由节点后无法恢复，确定删除？</span>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="deleteRuleEndPoint">确 定</el-button>
      </span>
    </el-dialog>
    <el-dialog
      title="路由规则详情"
      :visible.sync="dialogVisible2detail"
      width="30%"
    >
      <pre>
        {{detailContent}}
      </pre>

    </el-dialog>
    <el-dialog
      title="创建路由节点"
      :visible.sync="dialogVisible2create"
      width="40%"
    >
      <create-dialog @reload="reload" />
    </el-dialog>
  </div>
</template>
<script>
  import { dateTimeFormat } from '@/utils'
  import CreateDialog from '../components/create-ruleendpoint-dailog'
  import { Message } from 'element-ui'

  export default {
    name: 'RuleEndPoint',
    components: { CreateDialog },
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
      reload(){
        this.dialogVisible2create = false
        this.getList()
      },
      deleteRuleEndPoint() {
        this.$store.dispatch('edge-router/deleteRuleEndPoint').then(
          () => {
            Message({
              message: '删除路由节点成功',
              type: 'success',
              duration: 5 * 1000
            })
            this.dialogVisible = false
            //刷新数据
            this.getList()
          }
        )
      },
      deleteRuleEndPointView(name) {
        this.$store.dispatch('edge-router/setRuleEndPointName', name)
        this.dialogVisible = true
      },
      //详情展示
      showDetail(name) {
        this.dialogVisible2detail = true
        this.$store.dispatch('edge-router/setRuleEndPointName', name)
        this.$store.dispatch('edge-router/getRuleEndPointItem')
          .then(response => {
            this.detailContent = JSON.stringify(response.data, null, 4)
          })
      },
      getList() {
        this.$store.dispatch('edge-router/getRuleEndPointList')
          .then(response => {
            let result = []
            response.data.items.forEach((item, index) => {
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
  .el-dropdown {
    vertical-align: top;
  }

  .el-dropdown + .el-dropdown {
    margin-left: 15px;
  }

  .el-icon-arrow-down {
    font-size: 12px;
  }
</style>
