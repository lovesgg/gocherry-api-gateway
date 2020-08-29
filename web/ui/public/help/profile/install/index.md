## 项目安装

``` bash
# install cnpm 国内npm镜像库 只需要安装一次
npm install -g cnpm --registry=https://registry.npm.taobao.org

# install dependencies 第一次拉取项目或者添加了新依赖后需要执行该命令
cnpm install

# serve with hot reload at localhost:8080 开发环境运行
npm run dev

# build for production with minification 生产版本打包
npm run build:prod

# build for staging with minification 阶段版本打包
npm run build:stage

# run unit tests 单元测试
npm run test:unit
```
