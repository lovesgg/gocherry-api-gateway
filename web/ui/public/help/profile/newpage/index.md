## 如何新增页面

首先在 `@/router/index.js` 中增加你需要添加的路由。

**以新增`404`页面为例**

***路由配置***

```json
{
  path: '/404',
  component: Layout,
  redirect: '/404/index',
  name: '404',
  meta: {
    title: '404',
    icon: '404'
  },
  children: [{
    path: 'index',
    name: 'Index',
    component: () => import('@/views/404'),
    meta: { title: '404', icon: 'notfondicon' }
  }]
}
```

***如下图新建页面***

![404.png](/help/pic/404.png)

> **注：**
> 访问 [http://localhost:9528/404](http://localhost:9528/#/404) 会跳转到基于`Layout`布局的`test`页面
---

## 如何新增组件

第一步在 `@/components/ucen-ui` 中增加如下目录结构。

![addcomp.png](/help/pic/addcomp.png)

> **注：**
> 这里新建一个名为`themeSetting`的组件 里面分别包含了三个基础文件`index.js`组件统一注册逻辑 `index.scss`组件样式独立封装方便样式动态切换 `index.vue`组件页面

第二步在 @/components/index.js 中注册添加的组件

![addcomp2.png](/help/pic/addcomp2.png)

> **注：**
> 新建一般控件同上
