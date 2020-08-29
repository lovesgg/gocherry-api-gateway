import request_api from '@/utils/request_api'

export function getServerList(data) {
  return request_api({
    url: '/server/list',
    method: 'post',
    data
  })
}

export function saveServer(data) {
  return request_api({
    url: '/server/save',
    method: 'post',
    data
  })
}

