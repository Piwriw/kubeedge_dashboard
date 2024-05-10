import {
  fetchList
} from '@/api/edge-data'

const state = null
const mutations = null
const actions = {
  //节点列表
  getDataList({}, ) {
    return new Promise((resolve, reject) => {
      fetchList().then(response => {
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
