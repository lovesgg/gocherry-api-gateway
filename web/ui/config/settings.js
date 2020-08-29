module.exports = {

  title: 'gocherry API网关',

  /**
   * @type {boolean} true | false
   * @description Whether fix the header
   */
  fixedHeader: false,

  /**
   * @type {boolean} true | false
   * @description Whether show the logo in sidebar
   */
  sidebarLogo: false,
  primaryColor: '#42b983',
  theme: 'light',
  /**
   * @type {Array}
   * @description 亮色主题
   */
  lightThemeColors: [
    // '#000000d9',
    // '#000000a6',
    // '#00000073',
    '#050505',
    '#fbf7f0',
    // 'rgba(255, 255, 255, 0.89)',
    '#f6f6f6'
  ],
  /**
   * @type {Array}
   * @description 暗色主题
   */
  darkThemeColors: [
    // '#ffffffd9',
    // '#ffffffa6',
    // '#ffffff73',
    '#f5f5f5',
    '#363d43',
    // 'rgba(0, 0, 0, 0.89)',
    '#363d42'
  ],
  /**
   * @description vue-ls options
   */
  storageOptions: {
    namespace: 'pro__', // key prefix
    name: 'ls', // name variable Vue.[ls] or this.[$ls],
    storage: 'local' // storage name session, local, memory
  },
  /**
   * 是否打开单点登录
   */
  ssoOpen: false,
  /**
   * 部署目录
   */
  publicPath: '/'
}
