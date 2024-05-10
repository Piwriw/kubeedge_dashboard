import request from '@/utils/request'

export function fetchList() {
  return request({
    url: '/ai_model/ai_model/',
    method: 'get'
  })
}

export function fetchAiModelItem(version) {
  return request({
    url: '/ai_model/ai_model/' + version + '/',
    method: 'get'
  })
}

export function create(data) {
  return request({
    url: '/ai_model/ai_model/',
    method: 'post',
    data
  })
}

export function deleteAiModel(data) {
  return request({
    url: '/ai_model/ai_model/',
    method: 'delete',
    data
  })
}

//模型下发
export function distributeAiModel(data) {
  return request({
    url: '/ai_model/distribute/',
    method: 'post',
    data
  })
}

//获取模型下发的进度
export function getDistributeProgress(serializeId) {
  return request({
    url: '/ai_model/distribute/' + serializeId + '/',
    method: 'get'
  })
}

