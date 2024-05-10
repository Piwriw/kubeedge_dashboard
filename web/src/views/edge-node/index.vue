<template>
  <div class="app-container">
    <div class="filter-container">
      <el-button class="filter-item" type="primary">
        <router-link :to="'/edge-node/edge-node/create'">
          节点纳管 <i class="el-icon-circle-plus"/>
        </router-link>
      </el-button>
    </div>
    <el-table :data="tableData" style="width: 100%;">
      <el-table-column prop="name" label="节点名称" width="180">
        <template slot-scope="{row}">
          <svg-icon icon-class="status" style="margin-right: 10px"
                    :style="{color:(row.status===0)?'green':'red'}"/>
          <a style="color: blue" @click="getMore(row.name,row.status)">{{ row.name }}</a>
        </template>
      </el-table-column>
      <el-table-column prop="type" label="节点架构类型" width="120">
        <template slot-scope="{row}">
          {{ row.architecture }}
        </template>
      </el-table-column>
      <el-table-column label="节点标签" width="800">
        <template slot-scope="{row}">
          <el-tag class="my-tag" v-for="(value,key,index) in row.labels">{{key}}:{{value}}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="age" label="创建时间" width="150">
        <template slot-scope="{row}" style="cursor: pointer">
          <el-tooltip class="item" effect="dark" :content="row.creationTimestamp" placement="right-start">
            <el-button style="border: none;outline: none">{{ dateTimeFormat(row.creationTimestamp) }} <i
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
              <el-dropdown-item>
                <router-link :to="'/edge-node/detail/'+row.name">
                  查看
                </router-link>
              </el-dropdown-item>
              <!--<el-dropdown-item @click.native="dialogVisible = true">删除</el-dropdown-item>-->
              <!--<el-dropdown-item>更新</el-dropdown-item>-->
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
      <span>删除以后节点无法恢复，确定删除？</span>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="dialogVisible = false">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>
<script>
  import {dateTimeFormat} from '@/utils'
  import {Message} from 'element-ui'

  export default {
    name: 'EdgeNode',
    data() {
      return {
        dialogVisible: false,
        tableData: null
      }
    },
    created() {
      this.getList()
    },
    methods: {
      dateTimeFormat,
      //节点列表
      getList() {
        this.$store.dispatch("edge-node/getNodeList")
          .then(response => {
            let result = []
            response.data.forEach((item, index) => {
              result.push({
                "creationTimestamp":item.metadata.creationTimestamp,
                "name":item.metadata.name,
                "status":item.status.conditions.pop().status ==='True'?0:1,
                "architecture":item.status.nodeInfo.architecture,
                "labels":item.metadata.labels
              })
            })
            this.tableData = result
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
      //页面跳转，查看节点详情
      getMore(node_name, node_status) {
        if (node_status) {
          Message({
            message: '节点异常',
            type: 'error',
            duration: 5 * 1000
          })
        } else {
          this.$router.push(
            {path: '/edge-node/detail/' + node_name}
          )
        }
      }
    }
  }
</script>
<style scoped>
</style>
