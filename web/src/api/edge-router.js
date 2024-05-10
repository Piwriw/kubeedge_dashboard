import request from '@/utils/request'
//路由节点 RuleEndpoint 列表
export function fetchRuleEndPointList() {
  return request({
    url: '/router/ruleendpoint/',
    method: 'get'
  })
}
//路由规则 Rule 列表
export function fetchRuleList() {
  return request({
    url: '/router/rule/',
    method: 'get'
  })
}

//路由规则详情
export function fetchRuleItem(name) {
  return request({
    url: '/router/rule/default/'+name+"/",
    method: 'get'
  })
}

export function fetchRuleEndPointItem(name) {
  return request({
    url: '/router/ruleendpoint/default/'+name+"/",
    method: 'get'
  })
}


//删除路由规则 Rule
export function deleteRule(data) {
  return request({
    url: '/router/rule/',
    method: 'delete',
    data
  })
}

//删除路由节点 Rule
export function deleteRuleEndPoint(data) {
  return request({
    url: '/router/ruleendpoint/',
    method: 'delete',
    data
  })
}

//创建路由规则 Rule
export function createRule(data) {
  console.log("data",data)
  return request({
    url: '/router/rule/',
    method: 'post',
    data
  })
}



//创建路由节点 RuleEndPoint
export function createRuleEndPoint(data) {
  return request({
    url: '/router/ruleendpoint/',
    method: 'post',
    data
  })
}

