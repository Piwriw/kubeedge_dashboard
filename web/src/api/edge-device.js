import request from '@/utils/request'

//DeviceModel列表
export function fetchDeviceModelList() {
  return request({
    url: '/device/devicemodel/',
    method: 'get'
  })
}
//创建设备模板
export function createDeviceModel(data) {
  return request({
    url: '/device/devicemodel/',
    method: 'post',
    data
  })
}

//更新设备模板
export function updateDeviceModel(data) {
  return request({
    url: '/device/devicemodel/',
    method: 'put',
    data
  })
}

//更新设备模板
export function updateDevice(data) {
  console.log("data",data)
  return request({
    url: '/device/device/',
    method: 'put',
    data
  })
}

//删除设备模板
export function deleteDeviceModel(data) {
  return request({
    url: '/device/devicemodel/',
    method: 'delete',
    data
  })
}

//删除设备
export function deleteDevice(data) {
  return request({
    url: '/device/device/',
    method: 'delete',
    data
  })
}

//创建设备
export function createDevice(data) {
  console.log("data: " ,data)
  return request({
    url: '/device/device/',
    method: 'post',
    data
  })
}
//设备模板详情
export function fetchDeviceModelItem(name) {
  return request({
    url: '/device/devicemodel/default/'+name+"/",
    method: 'get'
  })
}

//设备详情
export function fetchDeviceItem(name) {
  return request({
    url: '/device/device/default/'+name+"/",
    method: 'get'
  })
}


//设备列表
export function fetchDeviceInstanceList() {
  return request({
    url: '/device/device/',
    method: 'get'
  })
}
