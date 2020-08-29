import request_api from '@/utils/request_api'

export function getAppList(data) {
  return request_api({
    url: '/app/list',
    method: 'post',
    data
  })
}

export function saveApp(data) {
  return request_api({
    url: '/app/save',
    method: 'post',
    data
  })
}

