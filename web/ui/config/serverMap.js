var configs = {
  // 本地开发环境的接口地址(npm run dev)
  dev: {
    base: '', // 本地开发调试用的服务器地址，修改不会影响发布
    node_api: '' // 本地开发调试用的服务器地址，修改不会影响发布
  },
  // 待发布的开发环境的接口地址(npm run build-dev)
  dev_build: {
    base: '//xxx.dev61.xxx.com',
    node_api: '//xxx-node.dev61.xxx.com'
  },
  // 测试环境的接口地址
  test: {
    base: '//aaa-service-test.xxx.com',
    node_api: '//bbb-node-service-test.xxx.com'
  },
  // 演示环境的接口地址
  demo: {
    base: '//aaa-demo.xxx.com',
    node_api: '//bbb-node-service-demo.xxx.com'
  },
  // 生产环境的接口地址
  prod: {
    base: '//aaa.xxx.com',
    node_api: '//bbb-node-service.xxx.com'
  },
  stage: {
    base: '//aaa.xxx.com',
    node_api: '//bbb-node-service.xxx.com'
  }
}
module.exports = configs[process.env.ENV_CONFIG]
