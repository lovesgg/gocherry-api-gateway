import defaultSettings from '../../../config/settings'
import {
  DEFAULT_THEME,
  DEFAULT_COLOR
} from '@/store/mutation-types'
import { changeThemeColors } from '@/utils/themeColorHelper'

const { showSettings, fixedHeader, sidebarLogo } = defaultSettings

const state = {
  showSettings: showSettings,
  fixedHeader: fixedHeader,
  sidebarLogo: sidebarLogo,
  primaryColor: '',
  theme: ''
}

const mutations = {
  CHANGE_SETTING: (state, { key, value }) => {
    if (state.hasOwnProperty(key)) {
      state[key] = value
    }
  },
  TOGGER_PRIMARY_COLOR: (state, value) => {
    localStorage.setItem(DEFAULT_COLOR, value)
    state.primaryColor = value
    changeThemeColors(value, null, state.theme)
  },
  TOGGER_THEME: (state, value) => {
    localStorage.setItem(DEFAULT_THEME, value)
    state.theme = value
    changeThemeColors(state.primaryColor || defaultSettings.primaryColor, null, value)
  },
  INIT_THEME: (state, data) => {
    state.primaryColor = data.primaryColor
    state.theme = data.theme
    localStorage.setItem(DEFAULT_COLOR, state.primaryColor)
    localStorage.setItem(DEFAULT_THEME, state.theme)
    changeThemeColors(state.primaryColor, null, state.theme)
  }
}

const actions = {
  changeSetting({ commit }, data) {
    commit('CHANGE_SETTING', data)
  },
  togglePrimaryColor({ commit }, data) {
    commit('TOGGER_PRIMARY_COLOR', data)
  },
  toggleTheme({ commit }, data) {
    commit('TOGGER_THEME', data)
  },
  initTheme({ commit }, data) {
    commit('INIT_THEME', data)
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}

