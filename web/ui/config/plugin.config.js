const webpack = require('webpack')
const JoinFileContentPlugin = require('join-file-content-plugin')
const ThemeColorReplacer = require('webpack-theme-color-replacer')
const forElementUI = require('webpack-theme-color-replacer/forElementUI')
var appConfig = require('./settings.js')
// var isProd = process.env.ENV_CONFIG === 'prod'

// const themePluginOption = {
//   fileName: 'css/theme-colors.[contenthash:8].css',
//   matchColors: forElementUI.getElementUISeries(appConfig.primaryColor).concat(appConfig.lightThemeColors),// element-ui主色系
//   changeSelector: forElementUI.changeSelector
//   // isJsUgly: config.isBuild,
// }

// const createThemeColorReplacerPlugin = () => new ThemeColorReplacer(themePluginOption)
var matchColors = forElementUI.getElementUISeries(appConfig.primaryColor).concat(appConfig.lightThemeColors)

module.exports = function(isBuild) {
  const plugins = [
    new webpack.DefinePlugin({
      'process.env': {
        NODE_ENV: JSON.stringify(process.env.NODE_ENV),
        ENV_CONFIG: JSON.stringify(process.env.ENV_CONFIG)
      }
    }),
    // 将theme-changed.scss应用到element-ui，供babel-plugin-component按需加载
    new JoinFileContentPlugin({
      file: 'node_modules/element-theme-chalk/src/common/var.scss',
      prependFile: 'src/styles/element-theme/theme-changed.scss'
    }),
    // 生成仅包含颜色的替换样式（主题色等）
    new ThemeColorReplacer({
      fileName: 'css/theme-colors.[contenthash:8].css',
      matchColors: matchColors, // forElementUI.getElementUISeries(appConfig.primaryColor).concat(appConfig.lightThemeColors),// element-ui主色系
      changeSelector: forElementUI.changeSelector,
      isJsUgly: isBuild
      // injectCss: false,
      // resolveCss(resultCss) { // optional. Resolve result css code as you wish.
      //     return resultCss + youCssCode
      // }
    })
  ]
  return plugins
}
