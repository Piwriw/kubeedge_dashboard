/**
 * @Author Piwriw.
 * @Email Piwriw@163.com
 * @motto 你不能做我的诗，正如我不能做你的梦.
 */
import request from '@/utils/request'

export function getSecretList(namespace) {
  return request({
    url: `/secret/${namespace}`,
    method: 'get',
  })
}
export function getSecretDetails(name,namespace) {
  return request({
    url: `/secret/${namespace}`,
    method: 'get',
    params:{
      name: name
    }
  })
}

export function getConfigMapList(namespace) {
  return request({
    url: `/configmap/${namespace}`,
    method: 'get',
  })
}
export function getConfigMapDetails(name,namespace) {
  return request({
    url: `/configmap/${namespace}`,
    method: 'get',
    params:{
      name: name
    }
  })
}
