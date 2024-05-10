/**
 * @Author Piwriw.
 * @Email Piwriw@163.com
 * @motto 你不能做我的诗，正如我不能做你的梦.
 */
import request from '@/utils/request'

export function getAllNamespace() {
  return request({
    url: '/namespace/',
    method: 'get',
  })
}
