## 结构化规范

``` bash
|-- babel.config.js       babel代码转换工具配置
|-- jest.config.js        jest单元测试配置
|-- jsconfig.json         指定根文件和JavaScript语言服务提供的功能选项
|-- package.json          模块基本信息(依赖包、版本、项目名称、描述等)
|-- README.md             项目说明文档
|-- vue.config.js         webpack配置
|-- build
|   |-- index.js
|-- config                插件配置
|   |-- plugin.config.js    插件配置
|   |-- serverMap.js        根据编译类型配置rest服务地址
|   |-- settings.js         项目初始化默认配置
|-- mock                  前端数据模拟
|   |-- index.js
|   |-- mock-server.js
|   |-- table.js
|   |-- user.js
|-- public                静态文件目录（不参与编译）
|   |-- favicon.ico
|   |-- index.html          入口页面
|   |-- help                帮助文档
|       |-- readMe.md
|-- src                   项目源码目录
|   |-- App.vue             根组件
|   |-- main.js             入口js文件
|   |-- permission.js       路由拦截js
|   |-- api                 restful api接口调用（所有调用后台或者mock接口的书写目录）
|   |   |-- table.js
|   |   |-- user.js
|   |-- assets              资源目录，这里的资源会被wabpack构建
|   |   |-- 404_images
|   |       |-- 404.png
|   |       |-- 404_cloud.png
|   |-- components          组件目录
|   |   |-- index.js
|   |   |-- index.scss        组件公共样式/变量（组件切换主题样式动态变换）
|   |   |-- common            项目内组件
|   |   |   |-- .gitkeep
|   |   |-- ucen-ui           可以被公司其它项目组复用的公共组件
|   |       |-- themeSetting
|   |           |-- index.js
|   |           |-- index.scss
|   |           |-- index.vue
|   |-- core                框架初核心库
|   |   |-- bootstrap.js      框架初始化
|   |-- icons               Svgicon
|   |   |-- index.js
|   |   |-- svgo.yml
|   |   |-- svg
|   |       |-- dashboard.svg
|   |-- layout              框架基本布局
|   |-- router              路由
|   |   |-- index.js
|   |-- store               Vuex
|   |   |-- getters.js
|   |   |-- index.js
|   |   |-- mutation-types.js
|   |   |-- modules
|   |       |-- app.js
|   |       |-- settings.js
|   |       |-- user.js
|   |-- styles              样式
|   |   |-- defines.scss
|   |   |-- element-ui.scss
|   |   |-- index.js
|   |   |-- index.scss
|   |   |-- mixin.scss
|   |   |-- sidebar.scss
|   |   |-- transition.scss
|   |   |-- variables.scss
|   |   |-- element-theme
|   |       |-- theme-changed.scss
|   |       |-- theme-element-variables.scss
|   |-- utils               工具
|   |   |-- auth.js
|   |   |-- get-page-title.js
|   |   |-- index.js
|   |   |-- moduleEvent.js
|   |   |-- request.js
|   |   |-- themeColorHelper.js
|   |   |-- validate.js
|   |-- views               页面
```
