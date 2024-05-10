import {
  fetchList,
  fetchItem,
  deleteApp

} from '@/api/edge-app'

const state = {
  name: '',
  namespace: '',
}
const mutations = {
  SET_APPNAME: (state, name) => {
    state.name = name
  },
  SET_APPNAMESPACE:(state,namespace) => {
    state.namespace = namespace
  }
}
const actions = {
  //应用列表
  getAppList({}) {
    return new Promise((resolve, reject) => {
      fetchList().then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  //应用详情
  getAppItem({},data) {
    return new Promise((resolve, reject) => {
      fetchItem(data).then(response => {
        resolve(response)
        commit('SET_APPNAME', '')
      }).catch(error => {
        reject(error)
      })
    })
  },
  //删除应用
  deleteApp({ commit, state }) {
    return new Promise((resolve, reject) => {
      let params = {
        'name': state.name,
        'namespace': state.namespace,
      }
      deleteApp(params).then(response => {
        commit('SET_APPNAME', '')
        commit('SET_APPNAMESPACE','')
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  setAppName({ commit, state }, data) {
    return new Promise(resolve => {
      commit('SET_APPNAME', data.name)
      commit('SET_APPNAMESPACE',data.namespace)
      resolve()
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
