import {
  fetchNodeMetrics

} from '@/api/metrics'

const state = null
const mutations = null
const actions = {
  //节点列表
  getNodeMetrics({}, nodename) {
    return new Promise((resolve, reject) => {
      fetchNodeMetrics(nodename).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
