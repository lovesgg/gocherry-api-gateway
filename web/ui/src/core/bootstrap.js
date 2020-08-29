// import Vue from 'vue'
import store from '@/store/'
import {
  DEFAULT_COLOR,
  DEFAULT_THEME
} from '@/store/mutation-types'
import config from '../../config/settings'

export default function Initializer() {
  console.log(`API_URL: ${process.env.VUE_APP_API_BASE_URL}`)
  // store.dispatch('settings/toggleTheme', localStorage.getItem(DEFAULT_THEME) || config.theme)
  // store.dispatch('settings/togglePrimaryColor', localStorage.getItem(DEFAULT_COLOR) || config.primaryColor)
  const color = localStorage.getItem(DEFAULT_COLOR) === 'undefined' ? undefined : localStorage.getItem(DEFAULT_COLOR)
  const theme = localStorage.getItem(DEFAULT_THEME) === 'undefined' ? undefined : localStorage.getItem(DEFAULT_THEME)
  const param = {
    primaryColor: color || config.primaryColor,
    theme: theme || config.theme
  }
  store.dispatch('settings/initTheme', param)
  // 这里也可以添加一些动态布局初始化
}
