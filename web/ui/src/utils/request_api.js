import axios from 'axios'
import {LoginToken} from '../enum/enum'

var serverMap = require('../../config/serverMap.js')

// create an axios instance
// console.log(process.env.VUE_APP_BASE_API)
const service = axios.create({
  baseURL: serverMap.base, // process.env.VUE_APP_BASE_API, // url = base url + request url
  // withCredentials: true, // send cookies when cross-domain requests
  timeout: 5000 // request timeout
})

// request interceptor
service.interceptors.request.use(
  config => {
    // do something before request is sent

    config.headers['AppName'] = 'mall_shop'
    config.headers['LoginToken'] = localStorage.getItem(LoginToken) || ""
    return config
  },
  error => {
    // do something with request error
    return Promise.reject(error)
  }
)

// response interceptor
service.interceptors.response.use(
  /**
   * If you want to get http information such as headers or status
   * Please return  response => response
   */

  /**
   * Determine the request status by custom code
   * Here is just an example
   * You can also judge the status by HTTP Status Code
   */
  response => {
    const res = response.data
    return res
  },
  error => {
    return Promise.reject(error)
  }
)

export default service
