import Vue from 'vue'
import Vuex from 'vuex'

import Status from './modules/production/status'
import Auth from './modules/production/auth'
import Posts from './modules/production/posts'

import MockStatus from './modules/mock/status'
import MockAuth from './modules/mock/auth'
import MockPosts from './modules/mock/posts'

// console.log(process.env.STORE_ENV)

let status = Status
let auth = Auth
let posts = Posts

if (process.env.STORE_ENV === 'mock') {
  status = MockStatus
  auth = MockAuth
  posts = MockPosts
}

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    status,
    auth,
    posts
  }
})
