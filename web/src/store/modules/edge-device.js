import {
  fetchDeviceInstanceList,
  fetchDeviceModelList,
  createDeviceModel,
  createDevice,
  updateDeviceModel,
  updateDevice,
  deleteDeviceModel,
  deleteDevice,
  fetchDeviceModelItem,
  fetchDeviceItem
} from '@/api/edge-device'

const state = {
  deviceModelName: '',
  deviceName: ''
}
const mutations = {
  SET_DEVICEMODEL_NAME: (state, name) => {
    state.deviceModelName = name
  },
  SET_DEVICE_NAME: (state, name) => {
    state.deviceName = name
  }
}
const actions = {
  //设备列表
  getDeviceList({}) {
    return new Promise((resolve, reject) => {
      fetchDeviceInstanceList().then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  //设备模板列表
  getDeviceModelList({}) {
    return new Promise((resolve, reject) => {
      fetchDeviceModelList().then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  //更新设备模板
  updateDeviceModel({}, data) {
    return new Promise((resolve, reject) => {
      updateDeviceModel(data).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  //同步设备字段信息
  updateDevice({ commit, state }, data) {
    data['name'] = state.deviceName
    return new Promise((resolve, reject) => {
      updateDevice(data).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  //创建设备模板
  createDeviceModel({}, data) {
    return new Promise((resolve, reject) => {
      createDeviceModel(data).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  //创建设备模板
  createDevice({}, data) {
    return new Promise((resolve, reject) => {
      createDevice(data).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  //设备模板详情
  getDeviceModelItem({ commit, state }) {
    return new Promise((resolve, reject) => {
      fetchDeviceModelItem(state.deviceModelName).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  getDeviceItem({ commit, state }) {
    return new Promise((resolve, reject) => {
      fetchDeviceItem(state.deviceName).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  //删除设备模板
  deleteDeviceModel({ commit, state }) {
    return new Promise((resolve, reject) => {
      //缓存获取设备模板名称
      let params = {
        'name': state.deviceModelName
      }
      deleteDeviceModel(params).then(response => {
        commit('SET_DEVICEMODEL_NAME', '')
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  //删除设备
  deleteDevice({ commit, state }) {
    return new Promise((resolve, reject) => {
      //缓存获取设备模板名称
      let params = {
        'name': state.deviceName
      }
      deleteDevice(params).then(response => {
        commit('SET_DEVICE_NAME', '')
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  //缓存当前设备模板名称
  setDeviceModelName({ commit }, deviceModelName) {
    return new Promise(resolve => {
      commit('SET_DEVICEMODEL_NAME', deviceModelName)
      resolve()
    })
  },
  //缓存当前设备模板名称
  setDeviceName({ commit }, deviceName) {
    return new Promise(resolve => {
      commit('SET_DEVICE_NAME', deviceName)
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
