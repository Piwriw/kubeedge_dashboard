import request from '@/utils/request'

// export function fetchList() {
//   return request({
//     url: '/node/node/',
//     method: 'get'
//   })
// }

export function fetchNodeMetrics(nodename) {
  console.log(nodename)
  return request({
    url: '/node/metrics/' + nodename + '/',
    method: 'get'
  })
}

