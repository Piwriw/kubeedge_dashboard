<template>
  <div class="card-container">
    <div class="card-header">
      节点资源使用情况
    </div>
    <div class="card-body">
      <el-row>
        <el-col :span="12">
          <div class="grid-content bg-purple">
            <usage-chart :name="'CPU'" :unit="'n'" :scale="1000" :metrics="JSON.stringify(metrics.cpu)" :id="chartName1" :class-name="chartName1" height="400px" width="100%" />
          </div>
        </el-col>
        <el-col :span="12">
          <div class="grid-content bg-purple-light">
            <usage-chart :name="'内存'" :unit="'Mi'" :scale="1024" :metrics="JSON.stringify(metrics.memory)" :id="chartName2" :class-name="chartName2" height="400px" width="100%" />
          </div>
        </el-col>
      </el-row>
    </div>
  </div>
</template>
<script>
  import UsageChart from '@/components/Charts/UsageChart'

  export default {
    name: 'NodeUsageCard',
    components: { UsageChart },
    props: ['metrics'],
    data() {
      return {
        chartName1: 'chart1',
        chartName2: 'chart2',
      }
    },
  }
</script>
<style scoped>
</style>
