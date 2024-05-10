<template>
  <div class="dashboard-container">
    <!--component是vue自己定义的标签 它的用途就是可以动态绑定我们的组件，根据数据的不同更换不同的组件-->
    <admin-dashboard/>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import adminDashboard from './admin'

export default {
  name: 'Dashboard',
  components: { adminDashboard },
  data() {
    return {
      currentRole: 'adminDashboard'
    }
  },
  computed: {
    // 等价于roles():{this.$state.roles}//其他地方改变该值得时候会跟着变化
    ...mapGetters(['roles'])
  },
  created() {
    if (!this.roles.includes('admin')) {
      this.currentRole = 'editorDashboard'
    }
  }
}
</script>
