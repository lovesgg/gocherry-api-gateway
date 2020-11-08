import Vue from 'vue'
import Router from 'vue-router'
/* Layout */
import Layout from '@/layout'
import {LoginToken} from '../enum/enum'

Vue.use(Router)

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'             the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true,
    name: "Login"
  },

  {
    path: '/404',
    component: () => import('@/views/404'),
    hidden: true
  },

  {
    path: '/',
    component: Layout,
    redirect: '/index',
    children: [{
      path: 'index',
      name: 'Index',
      component: () => import('@/views/index'),
      meta: { title: '主页', icon: 'dashboard' }
    }]
  },
  {
    path: '/app',
    component: Layout,
    redirect: '/app/list',
    name: 'App',
    meta: {title: 'App', icon: 'example'},
    children: [
      {
        path: 'list',
        name: 'AppList',
        component: () => import('@/views/app/list'),
        meta: {title: '应用列表', icon: 'table'}
      },
      {
        path: 'edit',
        name: 'AppEdit',
        component: () => import('@/views/app/edit'),
        meta: {title: '其他', icon: 'table' }
      }
    ]
  },
  {
    path: '/cluster',
    component: Layout,
    redirect: '/cluster/list',
    name: 'Cluster',
    meta: {title: 'Cluster', icon: 'example'},
    children: [
      {
        path: 'list',
        name: 'ClusterList',
        component: () => import('@/views/cluster/list'),
        meta: {title: '服务列表', icon: 'table'}
      },
      {
        path: 'list1',
        name: 'ClusterList',
        component: () => import('@/views/cluster/list'),
        meta: {title: '其他', icon: 'table'}
      },
    ]
  },
  {
    path: '/server',
    component: Layout,
    redirect: '/server/list',
    name: 'Server',
    meta: { title: 'Server', icon: 'example' },
    children: [
      {
        path: 'list',
        name: 'ServerList',
        component: () => import('@/views/server/list'),
        meta: { title: '服务节点', icon: 'table' }
      },
      {
        path: 'edit',
        name: 'ServerEdit',
        component: () => import('@/views/server/edit'),
        meta: { title: '其他', icon: 'table' }
      }
    ]
  },
  {
    path: '/api',
    component: Layout,
    redirect: '/api/list',
    name: 'Api',
    meta: { title: 'Api', icon: 'example' },
    children: [
      {
        path: 'list',
        name: 'ApiList',
        component: () => import('@/views/api/list'),
        meta: { title: '接口列表', icon: 'table' }
      },
      {
        path: 'edit',
        name: 'ApiEdit',
        component: () => import('@/views/api/edit'),
        meta: { title: '添加接口', icon: 'table' }
      }
    ]
  },
  {
    path: '/user',
    component: Layout,
    redirect: '/user/list',
    name: 'User',
    meta: { title: 'User', icon: 'example' },
    children: [
      {
        path: 'list',
        name: 'UserList',
        component: () => import('@/views/user/list'),
        meta: { title: '用户列表', icon: 'table' }
      },
      {
        path: 'edit',
        name: 'UserEdit',
        component: () => import('@/views/user/edit'),
        meta: { title: '其他', icon: 'table' }
      }
    ]
  },
  {
    path: '/other',
    component: Layout,
    redirect: '/other/list',
    name: 'Other',
    meta: { title: 'Other', icon: 'example' },
    children: [
      {
        path: 'list',
        name: 'OtherList',
        component: () => import('@/views/other/list'),
        meta: { title: '其他', icon: 'table' }
      },
      {
        path: 'edit',
        name: 'OtherEdit',
        component: () => import('@/views/other/edit'),
        meta: { title: '其他', icon: 'table' }
      }
    ]
  },

  { path: '*', redirect: '/404', hidden: true }
]

const createRouter = new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})
//登录路由校验
createRouter.beforeEach((to, from, next) => {
  let token = localStorage.getItem(LoginToken);
  if (token !== undefined && token !== "") {
    // 如果已经登录的话
    next();
  } else {
    if (to.path === '/login') {
      // 如果是登录页面的话，直接next()
      next();
    } else {
      // 否则 跳转到登录页面
      next({
        path: '/login'
      });
    }
  }
})

const router = createRouter

export function resetRouter() {
  const newRouter = createRouter
  router.matcher = newRouter.matcher
}

export default router
