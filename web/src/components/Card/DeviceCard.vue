<template>
  <div class="card-container">
    <div class="card-header">
      {{ data.title }}
    </div>
    <div class="card-body">
      <el-table :data="devices" fit style="width: 100%;">
        <el-table-column prop="name" label="设备名称" width="200">
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
      </el-table>
    </div>
  </div>
</template>
<script>
  import {dateTimeFormat} from '@/utils'

  export default {
    name: 'DeviceCard',
    props: ['devices'],
    data() {
      return {
        data: {
          title: '设备列表'
        },
      }
    },
    methods: {
      dateTimeFormat
    }
  }
</script>
<style scoped>
</style>
