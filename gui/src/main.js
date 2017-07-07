// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import Moment from 'moment'
import KeenUI from 'keen-ui'
import App from './app'
import Store from './store'
import Router from './router'

Vue.config.productionTip = false

Vue.use(KeenUI)

Vue.filter('fromNow', (t) => {
  let d = new Moment(t * 1000)
  return d.fromNow()
})

Vue.filter('toDatetime', (t) => {
  let d = new Moment(t * 1000)
  return d.format('YYYY.MM.DD HH:mm:ss')
})

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router: Router,
  store: Store,
  template: '<App/>',
  components: { App }
})
