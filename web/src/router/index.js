import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../components/Login'
import Main from '../components/Main'
import CommodityCenter from '../components/CommodityCenter'
import PersonalCenter from '../components/PersonalCenter'
import CommodityDetails from '../components/CommodityDetails'
import Echarts from '../components/Echarts'
import Reshow from '../components/Reshow'
import SearchRecords from '../components/SearchRecords'
import Home from '../components/Home'
import Redirect from '../components/Redirect'
Vue.use(VueRouter)

//解决重复点击导航出现的bug
const originalPush = VueRouter.prototype.push
VueRouter.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
}

const routes = [
  { path: '/', redirect: '/login' },
  { path: '/login', component: Login },
  {
    path: '/main',
    component: Main,
    children: [
      {
        path: '/Home', component: Home
      },
      {
        path: '/CommodityCenter', component: CommodityCenter
      },
      {
        path: '/Redirect', component: Redirect
      },
      {
        path: '/CommodityDetails', component: CommodityDetails
      },
      {
        path: '/PersonalCenter',
        component: PersonalCenter,
      },
      {
        path: '/Echarts', component: Echarts
      },
      {
        path: '/Reshow', component: Reshow
      },
      {
        path: '/SearchRecords', component: SearchRecords
      }
    ]
  },
]

const router = new VueRouter({
  routes
})


// // 判断token字段
// router.beforeEach((to, from, next) => {
//   if (to.path === '/login') return next();
//   const tokenStr = window.sessionStorage.getItem('token');
//   if (tokenStr === 'undefined' || tokenStr === null) return next('/login');
//   next();
// })

export default router



