/**
 * @Author Piwriw.
 * @Email Piwriw@163.com
 * @motto 你不能做我的诗，正如我不能做你的梦.
 */
import request from '@/utils/request'

export function fetchList(pageInfo) {
  return request({
    url: '/ie_data/ie_data',
    method: 'get',
    params:{
      'page':pageInfo.page,
      'size':pageInfo.size,
    }
  })
}
