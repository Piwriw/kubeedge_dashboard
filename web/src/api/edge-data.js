import request from '@/utils/request'

export function fetchList() {
  return request({
    url: '/svc_data/svc_data/',
    method: 'get',
  })
}
