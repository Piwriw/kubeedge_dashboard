<template>
  <div class="card-container">
    <div class="card-header">
      {{ data.title }}
    </div>
    <div class="card-body">
      <el-table :data="apps" fit style="width: 100%;">
        <el-table-column prop="name" label="应用名称" width="400">
          <template slot-scope="{row}">
            <svg-icon icon-class="status" style="margin-right: 10px;"
                      :style="{color:(row.status===0)?'green':'red'}" />
            {{ row.name }}
          </template>
        </el-table-column>
        <el-table-column label="镜像" width="320">
          <template slot-scope="{row}">
            <p v-for="(item,key) in row.images">{{item}}</p>
          </template>
        </el-table-column>
        <el-table-column label="标签">
          <template slot-scope="{row}">
            <el-tag class="my-tag" v-for="(value,key,index) in row.labels">{{key}}:{{value}}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="age" label="创建时间">
          <template slot-scope="{row}" style="cursor: pointer">
            <el-tooltip class="item" effect="dark" :content="row.createTime" placement="right-start">
              <el-button style="border: none;outline: none">{{ dateTimeFormat(row.createTime) }} <i
                class="el-icon-info" /></el-button>
            </el-tooltip>
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
  </div>
</template>
<script>
  import {dateTimeFormat} from '@/utils'
export default {
  name: 'ApplicationCard',
  props:['apps'],
  data() {
    return {
      dialogVisible: false,
      data: {
        title: '应用列表'
      },
      tableData: null
    }
  },
  methods: {
    dateTimeFormat,
    handleClose(done) {
      this.$confirm('确认删除？')
        .then(_ => {
          done()
        })
        .catch(_ => {
        })
    }
  }
}
</script>
<style scoped>
</style>
