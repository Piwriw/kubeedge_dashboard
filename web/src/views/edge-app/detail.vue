<template>
  <div class="app-container">
    <usage-card />
    <labels-card :labels="labels" />
  </div>
</template>

<script>
import LabelsCard from '@/components/Card/LabelsCard'
import UsageCard from '@/components/Card/NodeUsageCard'
export default {
  name: 'Detail',
  components: { LabelsCard, UsageCard },
  data() {
    return {
      // url传递
      id: this.$route.params.id,
      labels: ['beta.kubernetes.io/arch=arm64', 'beta.kubernetes.io/os=linux']
    }
  },
  created() {
    const title = '应用详情'
    this.tempRoute = Object.assign({}, this.$route)
    this.setPageTitle(title)
    this.setTagsViewTitle(title)
  },
  methods: {
    setPageTitle(title) {
      document.title = `${title} - ${this.id}`
    },
    setTagsViewTitle(title) {
      const route = Object.assign({}, this.tempRoute, { title: `${title}-${this.id}` })
      this.$store.dispatch('tagsView/updateVisitedView', route)
    }
  }
}
</script>

<style scoped>
  .app-container > * {
    margin: 20px 0;
  }
</style>
