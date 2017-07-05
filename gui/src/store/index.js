import Vue from 'vue'
import Vuex from 'vuex'
import MockStatus from './modules/mock/status'
import MockAuth from './modules/mock/auth'
import MockPosts from './modules/mock/posts'

let status = {}
let auth = {}
let posts = {}

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
