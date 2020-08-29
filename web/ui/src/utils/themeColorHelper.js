import client from 'webpack-theme-color-replacer/client'
import forElementUI from 'webpack-theme-color-replacer/forElementUI'
import appConfig from '../../config/settings.js'

export let curColor = appConfig.primaryColor

// 动态切换主题色
export function changeThemeColor(newColor) {
  var options = {
    newColors: [...forElementUI.getElementUISeries(newColor), '#ff0000', '#ffff00']
  }
  return client.changer.changeColor(options, Promise)
    .then(t => {
      curColor = newColor
      // localStorage.setItem('theme_color', curColor)
    })
}

// 动态切换所有颜色 包括自定义颜色
export function changeThemeColors(newPrimaryColor, newCustomColorArr, theme) {
  theme = theme || appConfig.theme
  newCustomColorArr = newCustomColorArr || getColorsByTheme(theme)
  customCoverStyles(theme)
  const themeColors = newCustomColorArr ? forElementUI.getElementUISeries(newPrimaryColor).concat(newCustomColorArr) : forElementUI.getElementUISeries(newPrimaryColor)
  var options = {
    newColors: themeColors
  }
  return client.changer.changeColor(options, Promise)
    .then(t => {
      curColor = newPrimaryColor
      // localStorage.setItem('theme_color', curColor)
    })
}

/**
 * 添加样式覆盖头
 * @param {string} theme
 */
function customCoverStyles(theme) {
  switch (theme) {
    case 'light':
      document.body.classList.remove('dark')
      document.body.classList.add('light')
      break
    case 'dark':
      document.body.classList.remove('light')
      document.body.classList.add('dark')
      break
    default:
      break
  }
}

const getColorsByTheme = (theme) => {
  let colors = []
  switch (theme) {
    case 'light':
      colors = appConfig.lightThemeColors
      break
    case 'dark':
      colors = appConfig.darkThemeColors
      break
    default:
      colors = []
      break
  }
  return colors
}

export function initThemeColor() {
  const savedColor = localStorage.getItem('theme_color')
  if (savedColor) {
    curColor = savedColor
    changeThemeColor(savedColor)
  }
}
