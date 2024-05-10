<template>
  <div class="card-container">
    <div class="card-header">
      {{ data.title }}
    </div>
    <div class="card-body">
      <el-row>
        <!--:key必须写-->
        <el-col v-for="(value, key) in data.content" :key="key" :span="12" style="padding: 15px 0">
          <strong>{{ key }}:</strong> {{ value }}
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
export default {
  name: 'NodeCard',
  props: ['data'],
  data() {
    return {
      name: 'node-01'
    }
  }
}
</script>

<style scoped>

</style>
