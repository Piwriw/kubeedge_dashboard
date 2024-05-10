<template>
  <div class="card-panel" @click="handleSetLineChartData('newVisitis')">
    <div class="card-panel-icon-wrapper icon-people">
      <!--<svg-icon icon-class="peoples" class-name="card-panel-icon" />-->
      <svg-icon icon-class="nodejs" class-name="card-panel-icon"/>
    </div>
    <div class="card-panel-description">
      <div class="card-panel-text">
        边缘节点
      </div>
      <count-to :start-val="0" :end-val="nodeNum" :duration="2600" class="card-panel-num"/>
    </div>
  </div>
</template>

<script>
    export default {
        name: "Panel"
    }
</script>

<style scoped>

</style>
