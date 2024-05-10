<template>
  <div class="dashboard-editor-container">
    <iframe
      :src="url"
      width="100%"
      height="1000px"
      frameborder="0"></iframe>
    <panel-group/>
    <!--<el-row style="background:#fff;padding:16px 16px 0;margin-bottom:32px;">-->
    <!--<el-col :span="12">-->
    <!--<resource-range-chart :chart-data="lineChartData" :title="'CPU使用情况'" :color="'#409EFF'" />-->
    <!--</el-col>-->
    <!--<el-col :span="12">-->
    <!--<resource-range-chart :chart-data="lineChartData" :title="'内存使用情况'" :color="'#30B08F'" />-->
    <!--</el-col>-->
    <!--</el-row>-->

    <el-row :gutter="32">
      <el-col :xs="24" :sm="24" :lg="8">
        <div class="chart-wrapper">
          <ratio-pie-chart :title="'节点架构类型' " :data="nodeArchPie"/>
        </div>
      </el-col>
      <el-col :xs="24" :sm="24" :lg="8">
        <div class="chart-wrapper">
          <ratio-pie-chart :title="'节点健康状态'" :data="nodeStatusPie"/>
        </div>
      </el-col>
      <el-col :xs="24" :sm="24" :lg="8">
        <div class="chart-wrapper">
          <ratio-pie-chart :title="'应用健康状态'" :data="appStatusPie"/>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>
  import PanelGroup from './components/PanelGroup'
  import RatioPieChart from '@/components/Charts/RatioPieChart'

  export default {
    name: 'DashboardAdmin',
    components: {
      PanelGroup,
      RatioPieChart
    },
    data() {
      return {
        url: 'http://192.168.1.110:32029/d/efa86fd1d0c121a26444b636a3f509a8/kubernetes-compute-resources-cluster?orgId=1&refresh=10s&kiosk',
        nodeArchPie: [],
        nodeStatusPie: [],
        appStatusPie: [],
      }
    },
    created() {
      this.initNodePie()
      this.initAppPie()
    },
    methods: {
      obj2list(obj) {
        let result = []
        for (let key in obj) {
          result.push({
            "name": key,
            "value": obj[key]
          })
        }
        return result
      },
      initAppPie() {
        this.$store.dispatch('edge-app/getAppList').then(response => {
          let appStatusObj = {}
          response.data.items.forEach((item, _) => {
            let status = item.status.phase === 'Running' ? "运行中" : "故障"
            if (appStatusObj.hasOwnProperty(status)) {
              appStatusObj[status] = appStatusObj[status] + 1
            } else {
              appStatusObj[status] = 1
            }
          })
          this.appStatusPie = this.obj2list(appStatusObj)
        })
      },
      initNodePie() {
        this.$store.dispatch("edge-node/getNodeList").then(res => {
          let nodeArchObj = {}
          let nodeStatusObj = {}
          res.data.forEach((item, index) => {
            let arch = item.status.nodeInfo.architecture
            let status = item.status.conditions.pop().status === 'True' ? "运行中" : "故障"
            //节点架构
            if (nodeArchObj.hasOwnProperty(arch)) {
              nodeArchObj[arch] = nodeArchObj[arch] + 1
            } else {
              nodeArchObj[arch] = 1
            }
            //节点状态
            if (nodeStatusObj.hasOwnProperty(status)) {
              nodeStatusObj[status] = nodeStatusObj[status] + 1
            } else {
              nodeStatusObj[status] = 1
            }
          })
          this.nodeArchPie = this.obj2list(nodeArchObj)
          this.nodeStatusPie = this.obj2list(nodeStatusObj)
        })
      }
    }
  }
</script>

<style lang="scss" scoped>
  .dashboard-editor-container {
    padding: 32px;
    background-color: rgb(240, 242, 245);
    position: relative;

    .github-corner {
      position: absolute;
      top: 0px;
      border: 0;
      right: 0;
    }

    .chart-wrapper {
      background: #fff;
      padding: 16px 16px 0;
      margin-bottom: 32px;
    }
  }

  @media (max-width: 1024px) {
    .chart-wrapper {
      padding: 8px;
    }
  }
</style>
