## 环境变量

以不同的方式运行或者打包框架会写入不同的环境变量，代码中我们可以通过`process.env.[NODE_ENV|ENV_CONFIG]`获取

```bash
"dev": "cross-env NODE_ENV=development ENV_CONFIG=dev vue-cli-service serve",
"build:prod": "cross-env NODE_ENV=production ENV_CONFIG=prod vue-cli-service build",
"build:stage": "cross-env NODE_ENV=production ENV_CONFIG=stage vue-cli-service build --mode staging",
```

上面是框架的运行或打包配置`dev`对应`development`,`build:prod`和`build:stage`对应`production`。

### 怎么利用环境变量

1. 项目中我们可以通过环境变量加载不同的插件 比如`development`环境中需要热加载和快速编译构建 `production`环境中需要编译出的代码尽可能的压缩
2. 项目中利用不同的环境变量加载不同的配置

  如：在`config/serverMap.js`中有如下配置

  ```bash
  // 本地开发环境的接口地址(npm run dev)
  dev: {
    base: '/dev-api', // 本地开发调试用的服务器地址，修改不会影响发布
    node_api: '//127.0.0.1:7001' // 本地开发调试用的服务器地址，修改不会影响发布
  },
  // 生产环境的接口地址(npm run build:prod)
  prod: {
    base: 'xxx.xxx.com',
    node_api: 'xxx.xxx.com'
  },
  // 生产环境的接口地址(npm run build:stage)
  stage: {
    base: 'xxx.xxx.com',
    node_api: 'xxx.xxx.com'
  }
  ```

  不同的环境读取不同的配置，不用每次打包都修改配置地址

  同理也可以配置扩展，如下

  ```bash
  "build:henan": "cross-env NODE_ENV=production ENV_CONFIG=henan vue-cli-service build"

  // 生产环境的接口地址(npm run build:henan)
  henan: {
    base: 'xxx.xxx.com',
    node_api: 'xxx.xxx.com'
  }
  ```

  以上添加了一个针对河南的部署配置，打包的时候只需要运行`npm run build:henan`
