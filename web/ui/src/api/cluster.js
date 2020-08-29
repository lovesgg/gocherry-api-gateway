import request_api from '@/utils/request_api'

export function getClusterList(data) {
  return request_api({
    url: '/cluster/list',
    method: 'post',
    data
  })
}

export function saveCluster(data) {
  return request_api({
    url: '/cluster/save',
    method: 'post',
    data
  })
}

