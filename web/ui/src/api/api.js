import request_api from '@/utils/request_api'

export function getApiList(data) {
  return request_api({
    url: '/api/list',
    method: 'post',
    data
  })
}

export function saveApi(data) {
  return request_api({
    url: '/api/save',
    method: 'post',
    data
  })
}

export function getOneApi(data) {
  return request_api({
    url: '/api/get_one',
    method: 'post',
    data
  })
}

export function delApi(data) {
  return request_api({
    url: '/api/del',
    method: 'post',
    data
  })
}

