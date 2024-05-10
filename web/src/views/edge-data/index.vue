<template>
  <div class="app-container">
    <div class="filter-container">
      <el-input placeholder="搜索内容" style="width: 200px;" class="filter-item"
      />
      <el-button class="filter-item" type="primary" icon="el-icon-search">
        搜索
      </el-button>
    </div>
    <el-table
      :data="list"
      fit
      highlight-current-row
      style="width: 100%;"
    >
      <el-table-column label="图像数据" width="200" align="center">
        <template slot-scope="{row}">
          <el-image
            style="width: 100px; height: 100px"
            :src="'data:image/png;base64,'+row.image_base64"
            :fit="'contain'"
          />
        </template>
      </el-table-column>
      <el-table-column prop="class_name" label="识别结果" width="200" align="center"/>
      <el-table-column prop="age" label="创建时间" width="150">
        <template slot-scope="{row}" style="cursor: pointer">
          <el-tooltip class="item" effect="dark" :content="row.upload_time" placement="right-start">
            <el-button style="border: none;outline: none">{{ dateTimeFormat(row.upload_time) }} <i
              class="el-icon-info"/></el-button>
          </el-tooltip>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>
<script>
import {dateTimeFormat} from '@/utils'
import {Base64ToImage} from '@/utils/base64-to-img'

export default {
  name: 'ComplexTable',
  data() {
    return {
      list: null,
    }
  },
  created() {
    this.getList()
  },
  methods: {
    dateTimeFormat,
    getList() {
      this.$store.dispatch("edge-data/getDataList").then(res => {
        this.list = res.data
        console.log(this.list)
      })
    }
  }
}
</script>
