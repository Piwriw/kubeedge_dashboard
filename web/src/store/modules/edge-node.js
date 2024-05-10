import {
  fetchList,
  fetchItem,
  deleteNodeLabel,
  createNodeLabel,
  getJoinToken

} from '@/api/edge-node'

const state = {
  nodeName: ''
}
const mutations = {
  SET_NODE_NAME: (state, nodeName) => {
    state.nodeName = nodeName
  }
}
const actions = {
  //节点列表
  getNodeList({}) {
    return new Promise((resolve, reject) => {
      fetchList().then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  //节点详情
  getNodeItem({ commit }, nodename) {
    return new Promise((resolve, reject) => {
      fetchItem(nodename).then(response => {
        commit('SET_NODE_NAME', nodename)
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  //删除节点标签
  deleteNodeLabel({ commit, state }, data) {
    data['node_name'] = state.nodeName
    return new Promise((resolve, reject) => {
      deleteNodeLabel(data).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  //添加/更新节点标签
  createNodeLabel({ commit, state }, data) {
    data['node_name'] = state.nodeName
    return new Promise((resolve, reject) => {
      createNodeLabel(data).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  getJoinToken({}) {
    return new Promise((resolve, reject) => {
      getJoinToken().then(response => {
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
