import {
  fetchList,
  deleteAiModel,
  distributeAiModel,
  getDistributeProgress,
  fetchAiModelItem
} from '@/api/ai-model'

const state = {
  version: ''
}
const mutations = {
  SET_MODEL_VERSION: (state, version) => {
    state.version = version
  }
}
const actions = {

  //模型下发
  distributeModel({}, data) {
    return new Promise((resolve, reject) => {
      distributeAiModel(data).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  //获取模型下发的进度
  getDistributeProgress({}, serializeId) {
    return new Promise((resolve, reject) => {
      getDistributeProgress(serializeId)
        .then(response => {
          resolve(response)
        })
        .catch(error => {
          reject(error)
        })
    })
  },
  //应用列表
  getAiModelList({}) {
    return new Promise((resolve, reject) => {
      fetchList().then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  //删除AI模型
  deleteAiModel({ commit, state }) {
    return new Promise((resolve, reject) => {
      let params = {
        'version': state.version
      }
      deleteAiModel(params).then(response => {
        commit('SET_MODEL_VERSION', '')
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  getAiModelItem({}, version) {
    return new Promise((resolve, reject) => {
      fetchAiModelItem(version).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  setAiModelVersion({ commit, state }, version) {
    return new Promise(resolve => {
      commit('SET_MODEL_VERSION', version)
      resolve()
    })
  },
  getAiModelVersion({ commit, state }) {
    return new Promise(resolve => {
      resolve(state.version)
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
