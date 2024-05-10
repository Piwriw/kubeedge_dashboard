import {
  fetchRuleList,
  fetchRuleEndPointList,
  deleteRule,
  deleteRuleEndPoint,
  fetchRuleItem,
  fetchRuleEndPointItem,
  createRule,
  createRuleEndPoint

} from '@/api/edge-router'

const state = null
const mutations = {
  SET_RULE_NAME: (state, name) => {
    state.rule = name
  },
  SET_RULEENDPOINT_NAME: (state, name) => {
    state.ruleendpoint = name
  }
}
const actions = {

  createRule({}, data) {
    return new Promise((resolve, reject) => {
      createRule(data).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },

  createRuleEndPoint({}, data) {
    return new Promise((resolve, reject) => {
      createRuleEndPoint(data).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },

  //路由规则列表
  getRuleList({}) {
    return new Promise((resolve, reject) => {
      fetchRuleList().then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },

  //路由节点列表
  getRuleEndPointList({}) {
    return new Promise((resolve, reject) => {
      fetchRuleEndPointList().then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },

  //删除路由规则
  deleteRule({ commit, state }) {
    return new Promise((resolve, reject) => {
      deleteRule({
        'name': state.rule
      }).then(response => {
        commit('SET_RULE_NAME', '')
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  //删除路由节点
  deleteRuleEndPoint({ commit, state }) {
    return new Promise((resolve, reject) => {
      deleteRuleEndPoint({ 'name': state.ruleendpoint }).then(response => {
        commit('SET_RULEENDPOINT_NAME', '')
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  getRuleItem({ commit, state }) {
    return new Promise((resolve, reject) => {
      fetchRuleItem(state.rule).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  getRuleEndPointItem({ commit, state }) {
    return new Promise((resolve, reject) => {
      fetchRuleEndPointItem(state.ruleendpoint).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  //缓存当前Rule名称
  setRuleName({ commit }, name) {
    return new Promise(resolve => {
      commit('SET_RULE_NAME', name)
      resolve()
    })
  },

  //缓存当前RuleEndPoint名称
  setRuleEndPointName({ commit }, name) {
    return new Promise(resolve => {
      commit('SET_RULEENDPOINT_NAME', name)
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
