<template>
  <div :id="id" :class="className" :style="{height:height,width:width}" />
</template>
<script>
  import echarts from 'echarts'
  import resize from './mixins/resize'

  export default {
    name: 'UsageChart',
    mixins: [resize],
    props: {
      scale:{
        type: Number,
        default: 1
      },
      unit:{
        type: String,
        default: null
      },
      name:{
        type: String,
        default: null
      },
      metrics: {
        type: String,
        default: null
      },
      className: {
        type: String,
        default: 'chart'
      },
      id: {
        type: String,
        default: 'chart'
      },
      width: {
        type: String,
        default: '200px'
      },
      height: {
        type: String,
        default: '200px'
      }
    },
    data() {
      return {
        chart: null,
        usage: 0
      }
    },
    mounted() {
      // console.log(this.metrics)
      // console.log(JSON.stringify(this.metrics))
      // this.initChart(this.metrics)
    },
    beforeDestroy() {
      if (!this.chart) {
        return
      }
      this.chart.dispose()
      this.chart = null
    },
    watch: {
      'metrics': {
        handler(newVale) {
          this.initChart(newVale)
        }
      }
    },
    methods: {
      initChart(metrics) {
        metrics = JSON.parse(metrics)
        const usage =Math.round((metrics.usage / metrics.allocatable) * 10000)/100 + "%"
        this.chart = echarts.init(document.getElementById(this.id))
        console.log(this.scale)
        this.chart.setOption({
          title: {
            text: this.name +': '+
              Math.round(metrics.usage / this.scale)+
              this.unit + ' / '+Math.round(metrics.allocatable / this.scale) + this.unit,
            x: 'center',
            y: 'bottom',
            textStyle: {
              color: '#000',
              fontSize: 20,
              fontWeight: 'normal',
              fontFamily: 'Microsoft YaHei'
            }
          },
          grid: {
            left: '5%',
            right: '5%'
          },
          series: [
            {
              name: '资源使用',
              type: 'pie',
              radius: ['40%', '70%'],
              avoidLabelOverlap: false,
              labelLine: {
                show: true
              },
              label: {
                normal: {
                  show: true,
                  position: 'center',
                  formatter: function(data) { // 设置圆饼图中间文字排版
                    return usage
                  },
                  textStyle: {
                    fontSize: '30',
                    color: '#333'
                  }
                }
              },
              data: [
                { value: metrics.usage, name: '已使用' },
                { value: metrics.allocatable - metrics.usage, name: '未使用' }
              ],
              itemStyle: {
                normal: {
                  color: function(params) {
                    if (params.name === '已使用') {
                      return '#409EFF'
                    } else {
                      return '#DCDCDC'
                    }
                  }
                }
              }
            }
          ]
        })
      }
    }
  }
</script>
<style scoped>
</style>
