<template>
  <div :class="className" :style="{height:height,width:width}"/>
</template>

<script>
  import echarts from 'echarts'

  require('echarts/theme/macarons') // echarts theme
  import resize from './mixins/resize'

  export default {
    mixins: [resize],
    props: {
      data: {
        type: Array,
        default: []
      },
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
        default: '300px'
      },
      title: {
        type: String,
        request: true
      }
    },
    data() {
      return {
        chart: null
      }
    },
    mounted() {
      this.$nextTick(() => {
        this.initChart()
      })
    },
    watch: {
      'data': {
        handler(newVale) {
          this.initChart(newVale)
        }
      }
    },
    beforeDestroy() {
      if (!this.chart) {
        return
      }
      this.chart.dispose()
      this.chart = null
    },
    methods: {
      initChart(newVale) {
        this.chart = echarts.init(this.$el, 'macarons')
        const option = {
          title: {
            text: this.title + "占比"
          },
          tooltip: {
            trigger: 'item',
            formatter: '{a} <br/>{b} : {c} ({d}%)'
          },
          legend: {
            left: 'center',
            bottom: '10'
          },
          series: [
            {
              name: this.title,
              type: 'pie',
              radius: '50%',
              data: newVale,
              emphasis: {
                itemStyle: {
                  shadowBlur: 10,
                  shadowOffsetX: 0,
                  shadowColor: 'rgba(0, 0, 0, 0.5)'
                }
              },
              // itemStyle: {
              //   normal: {
              //     color: function(params) {
              //       if (params.name === 'running') {
              //         return '#30B08F'
              //       } else {
              //         return '#DCDCDC'
              //       }
              //     }
              //   }
              // }
            }
          ]
        }
        this.chart.setOption(option)
      }
    }
  }
</script>
