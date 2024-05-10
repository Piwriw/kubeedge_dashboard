/**
 * @Author Piwriw.
 * @Email Piwriw@163.com
 * @motto 你不能做我的诗，正如我不能做你的梦.
 */
import {
  fetchList
} from '@/api/ie-data'

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
