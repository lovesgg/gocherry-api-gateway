import request_api from '@/utils/request_api'

export function login(data) {
  return request_api({
    url: '/login',
    method: 'post',
    data
  })
}

export function getUserList(data) {
  return request_api({
    url: '/user/list',
    method: 'post',
    data
  })
}


export function saveUser(data) {
  return request_api({
    url: '/user/save',
    method: 'post',
    data
  })
}
