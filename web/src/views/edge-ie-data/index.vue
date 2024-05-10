<template>
  <div class="app-container">


    <el-table
      :data="list"
      fit
      highlight-current-row
      style="width: 100%;"
    >

      <el-table-column prop="edge_device_id" label="边缘设备号" width="200" align="center"/>
      <el-table-column prop="voltage" label="电压" width="200" align="center"/>
      <el-table-column prop="electric_current" label="电流" width="200" align="center"/>
      <el-table-column prop="power" label="功率" width="200" align="center"/>
      <el-table-column prop="power_factor" label="功率因素" width="200" align="center"/>
      <el-table-column prop="frequency" label="频率" width="200" align="center"/>
      <el-table-column prop="electricity_consumption" label="累积耗电量" width="200" align="center"/>
      <el-table-column prop="label" label="电器类别" width="200" align="center"/>
      <el-table-column prop="age" label="创建时间" width="150">
        <template slot-scope="{row}" style="cursor: pointer">
          <el-tooltip class="item" effect="dark" :content="row.upload_time" placement="right-start">
            <el-button style="border: none;outline: none">{{ dateTimeFormat(row.upload_time) }} <i
              class="el-icon-info"/></el-button>
          </el-tooltip>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      background
      layout="prev, pager, next"
      @current-change="changePage"
      @size-change="changeSize"
      :current-page="pageInfo.page"
      :page-size="pageInfo.size"
      :page-sizes="[1,5,10,20]"
      :total="total">
    </el-pagination>
    <!--    <div class="filter-container">-->
    <!--      <el-input placeholder="搜索内容" style="width: 200px;" class="filter-item"-->
    <!--      />-->
    <!--      <el-button class="filter-item" type="primary" icon="el-icon-search">-->
    <!--        搜索-->
    <!--      </el-button>-->
    <!--    </div>-->

  </div>
</template>

<script>
import {dateTimeFormat} from "@/utils";
import {fetchList} from "@/api/ie-data";

export default {
  name: 'ComplexTable',
  data() {
    return {
      list: null,
      total: 0,
      pageInfo: {
        page: 1,
        size: 10
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    dateTimeFormat,
    getList() {
      fetchList(this.pageInfo).then(res => {
        this.list = res.data.data
        this.total = res.data.total
      })
    },
    changePage(page) {
      this.pageInfo.page=page
      console.log(this.pageInfo)
      this.getList()
    },
    changeSize(size){
      this.pageInfo.size=size
      this.getList()
    }
  }
}
</script>

<style scoped>

</style>
