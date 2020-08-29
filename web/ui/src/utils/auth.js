import Cookies from 'js-cookie'

const TokenKey = 'ucen_template_token'

/**
 * @description 获取单点登录token
 */
export function getToken() {
  return Cookies.get(TokenKey)
}

/**
 * @description 设置单点登录token
 * @param {*} token
 */
export function setToken(token) {
  return Cookies.set(TokenKey, token)
}

/**
 * 移除单点登录token
 */
export function removeToken() {
  return Cookies.remove(TokenKey)
}
