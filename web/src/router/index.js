import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'
// import tableRouter from './modules/table'

/* Router Modules */
// import componentsRouter from './modules/components'
// import chartsRouter from './modules/charts'
// import tableRouter from './modules/table'
// import nestedRouter from './modules/nested'

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
    icon: 'svg-name'/'el-icon-x' the icon show in the sidebar
    noCache: true                if set true, the page will no be cached(default is false)
    affix: true                  if set true, the tag will affix in the tags-view
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
    component: () => import('@/views/login/index_new'),
    hidden: true
  },
  {
    path: '/auth-redirect',
    component: () => import('@/views/login/auth-redirect'),
    hidden: true
  },
  {
    path: '/404',
    component: () => import('@/views/error-page/404'),
    hidden: true
  },
  {
    path: '/401',
    component: () => import('@/views/error-page/401'),
    hidden: true
  },
  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    children: [
      {
        path: 'dashboard',
        component: () => import('@/views/dashboard/index'),
        name: 'Dashboard',
        meta: {title: 'Dashboard', icon: 'dashboard', affix: true}
      }
    ]
  },
  {
    path: '/edge-node',
    component: Layout,
    redirect: '/edge-node/index',
    hidden: false,
    children: [
      {
        path: 'index',
        component: () => import('@/views/edge-node/index'),
        name: 'EdgeNode',
        meta: {title: '边缘节点', icon: 'nodejs', noCache: true}
      },
      {
        path: 'detail/:id(.+)',
        component: () => import('@/views/edge-node/detail'),
        name: 'EditArticle',
        meta: {title: '边缘节点详情', noCache: true, activeMenu: '/edge-node/index'},
        hidden: true
      },
      {
        path: 'edge-node/create',
        component: () => import('@/views/edge-node/create'),
        name: 'EditCreate',
        meta: {title: '添加边缘节点', noCache: true, activeMenu: '/edge-node/index'},
        hidden: true
      }
    ]
  },
  {
    path: '/edge-device',
    component: Layout,
    // tag点击跳转用
    redirect: '/edge-device/device-model',
    hidden: false,
    meta: {title: '设备管理', icon: 'device-fill', noCache: true},
    children: [
      {
        path: 'device-model',
        component: () => import('@/views/edge-device/device-model/index'),
        name: 'DeviceModel',
        meta: {title: '设备模板', noCache: true}
      },
      {
        path: 'device-instance',
        component: () => import('@/views/edge-device/device-instance/index'),
        name: 'DevicesInstance',
        meta: {title: '设备实例', noCache: true}
      },
      {
        path: 'device-model/create',
        component: () => import('@/views/edge-device/device-model/create'),
        name: 'DevicesModelCreate',
        hidden: true,
        meta: {title: '创建设备模板', noCache: true, activeMenu: '/edge-device/device-model'}
      },
      {
        path: 'device-model/update',
        component: () => import('@/views/edge-device/device-model/update'),
        name: 'DevicesModelUpdate',
        hidden: true,
        meta: {title: '更新设备模板', noCache: true, activeMenu: '/edge-device/device-model'}
      }
    ]
  },
  {
    path: '/edge-router',
    component: Layout,
    // tag点击跳转用
    redirect: '/edge-router/ruleendpoint',
    hidden: false,
    meta: {title: '路由管理', icon: 'device-fill', noCache: true},
    children: [
      {
        path: 'ruleendpoint',
        component: () => import('@/views/edge-router/ruleendpoint/index'),
        name: 'DevicesInstance',
        meta: {title: '路由节点', noCache: true}
      },
      {
        path: 'rule',
        component: () => import('@/views/edge-router/rule/index'),
        name: 'DeviceModel',
        meta: {title: '路由规则', noCache: true}
      }
    ]
  },
  {
    path: '/resource',
    component: Layout,
    // tag点击跳转用
    redirect: '/resource/configmap',
    hidden: false,
    meta: {title: 'Resource', icon: 'data', noCache: true},
    children: [
      {
        path: 'configmap',
        component: () => import('@/views/resource/configmap/index'),
        name: 'ConfigMap',
        meta: {title: 'ConfigMap', icon: 'cm', noCache: true}
      },
      {
        path: 'secret',
        component: () => import('@/views/resource/secret/index'),
        name: 'Secret',
        meta: {title: 'Secret', icon: 'secret', noCache: true}
      }
    ]
  },
  {
    path: '/edge-app',
    component: Layout,
    redirect: '/edge-app/index',
    hidden: false,
    children: [
      {
        path: 'index',
        component: () => import('@/views/edge-app/index'),
        name: 'Applications',
        meta: {title: '边缘应用', icon: 'application', noCache: true}
      },
      {
        path: 'detail/:id(.+)',
        component: () => import('@/views/edge-app/detail'),
        name: 'Applications',
        hidden: true,
        meta: {title: '应用详情', noCache: true, activeMenu: '/edge-app/index'}
      },
      {
        path: 'create',
        component: () => import('@/views/edge-app/create'),
        name: 'ApplicationCreate',
        hidden: true,
        meta: {title: '创建应用', noCache: true, activeMenu: '/edge-app/index'}
      }
    ]
  },
  // {
  //   path: '/edge-ai-model',
  //   component: Layout,
  //   redirect: '/edge-ai-model/index',
  //   hidden: false,
  //   children: [
  //     {
  //       path: 'index',
  //       component: () => import('@/views/edge-ai-model/index'),
  //       name: 'AiModel',
  //       meta: { title: '边缘AI模型', icon: 'model2', noCache: true }
  //     }
  //   ]
  // },
  {
    path: '/edge-data',
    component: Layout,
    // tag点击跳转用
    redirect: '/edge-data/index',
    hidden: false,
    meta: {title: '边缘数据', icon: 'data', noCache: true},
    children: [
      {
        path: 'edge_data',
        component: () => import('@/views/edge-data/index'),
        name: 'EdgeData',
        meta: {title: '果蔬识别边缘数据', noCache: true}
      },
      {
        path: 'edge_ie_data',
        component: () => import('@/views/edge-ie-data/index'),
        name: 'EdgeIEData',
        meta: {title: 'IE边缘数据', noCache: true}
      }
    ]
  },
  // {
  // path: '/edge-data',
  // component: Layout,
  // redirect: '/edge-data/index',
  // hidden: false,
  // children: [
  //   {
  //     path: 'index',
  //     component: () => import('@/views/edge-data/index'),
  //     name: 'EdgeData',
  //     meta: { title: '果蔬边缘数据', icon: 'data', noCache: true }
  //   },
  // ]
  // }
]

/**
 * asyncRoutes
 * the routes that need to be dynamically loaded based on user roles
 */
export const asyncRoutes = [
  // {
  //   path: '/icon',
  //   component: Layout,
  //   children: [
  //     {
  //       path: 'index',
  //       component: () => import('@/views/icons/index'),
  //       name: '图标库',
  //       meta: { title: '图标库', icon: 'icon', noCache: true }
  //     }
  //   ]
  // },
  // 404 page must be placed at the end !!!
  {path: '*', redirect: '/404', hidden: true}
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({y: 0}),
  routes: constantRoutes.concat(asyncRoutes)
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
