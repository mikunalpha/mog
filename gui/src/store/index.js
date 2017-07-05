import Vue from 'vue'
import Vuex from 'vuex'
import MockAuth from './modules/mock/auth'
import MockPosts from './modules/mock/posts'

let auth = {}
let posts = {}

if (process.env.STORE_ENV === 'mock') {
  auth = MockAuth
  posts = MockPosts
}

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    auth,
    posts
  }
})
