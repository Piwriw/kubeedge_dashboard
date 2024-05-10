<template>
  <div :class="className" :style="{height:height,width:width}" />
</template>

<script>
import echarts from 'echarts'
require('echarts/theme/macarons') // echarts theme
import resize from './mixins/resize'
export default {
  mixins: [resize],
  props: {
    className: {
      type: String,
      default: 'chart'
    },
    width: {
      type: String,
      default: '100%'
    },
    height: {
      type: String,
      default: '230px'
    },
    autoResize: {
      type: Boolean,
      default: true
    },
    chartData: {
      type: Object,
      required: true
    },
    title: {
      type: String,
      required: true
    },
    color: {
      type: String,
      required: true
    }
  },
  data() {
    return {
      chart: null
    }
  },
  watch: {
    chartData: {
      deep: true,
      handler(val) {
        this.setOptions(val)
      }
    }
  },
  mounted() {
    this.$nextTick(() => {
      this.initChart()
    })
  },
  beforeDestroy() {
    if (!this.chart) {
      return
    }
    this.chart.dispose()
    this.chart = null
  },
  methods: {
    initChart() {
      this.chart = echarts.init(this.$el, 'macarons')
      this.setOptions(this.chartData)
    },
    setOptions({ expectedData, actualData } = {}) {
      const xAxisData = []
      var currentDate = new Date()
      for (let i = 0; i < 10; i++) {
        const temp = new Date(currentDate.getTime() - i * 1000 * 60)
        // console.log(temp.getHours())
        // console.log(temp.getMinutes())
        xAxisData.push(temp.getHours() + ':' + temp.getMinutes())
      }
      var data = [820, 932, 901, 934, 1290, 1330, 1320, 820, 932, 901, 934, 1290, 1330, 1320]
      console.log(xAxisData)
      const option = {
        title: {
          text: this.title,
          x: 'left'
        },
        xAxis: {
          type: 'category',
          boundaryGap: false,
          splitLine: {
            show: true
          },
          data: xAxisData
        },
        yAxis: {
          type: 'value',
          splitLine: {
            show: true
          },
          axisTick: {
            show: false
          }
        },
        grid: {
          left: 10,
          right: 20,
          bottom: 20,
          top: 50,
          containLabel: true
        },
        // legend: {
        //   data: ['memory']
        // },
        tooltip: {
          trigger: 'axis',
          axisPointer: {
            type: 'cross'
          },
          padding: [5, 10]
        },
        series: [{
          name: 'memory',
          data: data,
          type: 'line',
          smooth: true,
          areaStyle: {},
          itemStyle: {
            normal: {
              color: this.color
            },
            lineStyle: {
              color: this.color,
              width: 2
            },
            areaStyle: {
              // '#409EFF'
              color: this.color
            }
          },
          animationDuration: 2800,
          animationEasing: 'cubicInOut'
        }]
      }
      this.chart.setOption(option)
    }
  }
}
</script>
