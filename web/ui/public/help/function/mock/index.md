## Mockjs
  Mock.js 是一款模拟数据生成器，旨在帮助前端攻城师独立于后端进行开发，帮助编写单元测试。提供了以下模拟功能：

* 根据数据模板生成模拟数据
* 模拟 Ajax 请求，生成并返回模拟数据
* 基于 HTML 模板生成模拟数据

具体用法请参见 [http://mockjs.com/](http://mockjs.com/)

### 怎么新增模拟数据

如我们模拟一个随机的表格数据：

首先在`mock`文件夹下面新建一个`table.js`,内容如下

```javascript
import Mock from 'mockjs'

const data = Mock.mock({
  'items|30': [{
    id: '@id',
    title: '@sentence(10, 20)',
    'status|1': ['published', 'draft', 'deleted'],
    author: 'name',
    display_time: '@datetime',
    pageviews: '@integer(300, 5000)'
  }]
})

export default [
  {
    url: '/vue-admin-template/table/list',
    type: 'get',
    response: config => {
      const items = data.items
      return {
        code: 20000,
        data: {
          total: items.length,
          items: items
        }
      }
    }
  }
]
```

第二步在`@/api/`目录下新建`table.js`,内容如下

```javascript
import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/vue-admin-template/table/list',
    method: 'get',
    params
  })
}
```

调用方式，在vue或者js文件内引入并调用

```javascript
// 引入
import { getList } from '@/api/table'
// 调用 并返回promise对象
getList().then(response => {
  // to do something
})
```

### 怎么移除mock server

首先移除vue.config.js内的devServer

```json
devServer: {
  port: port,
  open: true,
  overlay: {
    warnings: false,
    errors: true
  },
  before: require('./mock/mock-server.js')
}
```

移除`main.js`内的如下代码

```javascript
if (process.env.NODE_ENV === 'production') {
  const { mockXHR } = require('../mock')
  mockXHR()
}
```

>**注** 移除后需要重启

### 根路径

`Mockserver`的根路径和环境变量对应

```bash
// dev
http://localhost:9528/dev-api/
// prod
http://localhost:9528/prod-api/
// stage
http://localhost:9528/stage-api/
```

>所以调用`table`数据访问的完整路径为(`dev`模式下)[`http://localhost:9528/dev-api/vue-admin-template/table/list`](http://localhost:9528/dev-api/vue-admin-template/table/list)
>>**注** 可根据需要自行改造多`Server`模式
