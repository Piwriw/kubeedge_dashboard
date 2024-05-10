import request from '@/utils/request'

export function fetchList() {
  return request({
    url: '/vue-element-admin/secret/list',
    method: 'get'
  })
}
