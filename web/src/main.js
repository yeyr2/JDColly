import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './css/common.css'
import './plugins/elements.js'
import axios from 'axios'
import * as echarts from "echarts"

Vue.config.productionTip = false
// 全局挂载axios
axios.defaults.baseURL = 'http://www.iyeyr2.top'
// 
// axios请求拦截
axios.interceptors.request.use(config => {
  // 为请求头对象，添加Token验证的Authorization字段
  config.headers.Authorization = window.sessionStorage.getItem('token')
  return config
})

Vue.prototype.$http = axios //将axios放到全局原型上，利用继承机制在实例里通过$http使用axios
Vue.prototype.$echarts = echarts;
new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
