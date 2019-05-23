import Vue from 'vue'
import axios from 'axios'
import VueAxios from 'vue-axios'
import Antd from 'ant-design-vue'
import App from './App'
import 'ant-design-vue/dist/antd.css'
import router from './router'
import store from './store'
import moment from "moment";
import VueMomentJS from "vue-momentjs";
import lineClamp from 'vue-line-clamp'
import uuidv4 from 'uuid/v4'
import blake from 'blakejs'
import querystring from 'querystring'

Vue.config.productionTip = false


let axiosConfigs = {
  timeout: 30000,
  type: 'json'
}
let axiosInstance
if (process.env.NODE_ENV === "production") {
  axiosInstance = axios.create({
    baseURL: '',
    ...axiosConfigs
  });
} else {
  axiosInstance = axios.create({
    baseURL: 'http://localhost:3050',
    ...axiosConfigs
  });
}

Vue.use(Antd)
Vue.use(VueAxios, axiosInstance)
Vue.use(VueMomentJS, moment)
Vue.use(lineClamp)

Vue.prototype.$prepareCredentials = (password) => {
  const SEPARATOR= "|"
  let key = uuidv4()
  let encryptedPassword = blake.blake2bHex(`${key}${SEPARATOR}${password}`)
  return {key, encryptedPassword}
}
Vue.prototype.$queryString = querystring

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
