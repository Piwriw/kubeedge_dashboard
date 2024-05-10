import request from '@/utils/request'
import { getToken } from '@/utils/auth'
import { Message } from 'element-ui'

export function fetchList() {
  return request({
    url: '/node/node/',
    method: 'get'
  })
}

export function getJoinToken() {
  return request({
    url: '/node/join/',
    method: 'get'
  })
}

export function fetchItem(nodename) {
  return request({
    url: '/node/node/' + nodename + '/',
    method: 'get'
  })
}

export function deleteNodeLabel(data) {

  return request({
    url: '/node/label/',
    method: 'delete',
    data
  })

}

export function createNodeLabel(data) {

  return request({
    url: '/node/label/',
    method: 'put',
    data
  })

}
