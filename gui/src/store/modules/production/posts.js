import axios from 'axios'
import * as types from '../../mutations'
import cookies from 'browser-cookies'
import querystring from 'querystring'

// initial state
const state = {
  posts: {
    gotAt: null,
    data: []
  },
  post: {
    gotAt: null,
    data: {}
  }
}

// getters
const getters = {
  posts: state => state.posts,
  post: state => state.post
}

// actions
const actions = {
  getPosts ({ commit }, {options, success, error}) {
    axios.get('/mogapis/v1/posts?' + querystring.stringify(options), {
      headers: {
        'Authorization': 'Bearer ' + cookies.get('at')
      }
    })
    .then(function (response) {
      commit(types.RECEIVE_POSTS, {posts: response.data.data})

      if (typeof success === 'function') {
        success(response.data.data)
      }
    }).catch(function (e) {
      console.log(e)
      if (typeof error === 'function') {
        if (e.response.status === 504) {
          return error(
            504,
            {
              code: 'ServerError',
              message: 'Can not connect to server.'
            }
          )
        }
        error(e.response.status, e.response.data.error)
      }
    })
  },

  getPost ({ commit }, {id, options, success, error}) {
    axios.get('/mogapis/v1/post/' + id + '?', {
      headers: {
        'Authorization': 'Bearer ' + cookies.get('at')
      }
    })
    .then(function (response) {
      commit(types.RECEIVE_POST, {post: response.data.data})

      if (typeof success === 'function') {
        success(response.data.data)
      }
    }).catch(function (e) {
      console.log(e)
      if (typeof error === 'function') {
        if (e.response.status === 504) {
          return error(
            504,
            {
              code: 'ServerError',
              message: 'Can not connect to server.'
            }
          )
        }
        error(e.response.status, e.response.data.error)
      }
    })
  },

  createPost ({ commit }, {post, success, error}) {
    axios.post('/mogapis/v1/post', {
      data: post
    }, {
      headers: {
        'Authorization': 'Bearer ' + cookies.get('at')
      }
    }).then(function (response) {
      if (typeof success === 'function') {
        success(response.data.data)
      }
    }).catch(function (e) {
      console.log(e)
      if (typeof error === 'function') {
        if (e.response.status === 504) {
          return error(
            504,
            {
              code: 'ServerError',
              message: 'Can not connect to server.'
            }
          )
        }
        error(e.response.status, e.response.data.error)
      }
    })
  },

  updatePost ({ commit }, {id, post, success, error}) {
    axios.patch('/mogapis/v1/post/' + id, {
      data: post
    }, {
      headers: {
        'Authorization': 'Bearer ' + cookies.get('at')
      }
    }).then(function (response) {
      if (typeof success === 'function') {
        success(response.data.data)
      }
    }).catch(function (e) {
      console.log(e)
      if (typeof error === 'function') {
        if (e.response.status === 504) {
          return error(
            504,
            {
              code: 'ServerError',
              message: 'Can not connect to server.'
            }
          )
        }
        error(e.response.status, e.response.data.error)
      }
    })
  },

  deletePost ({ commit }, {id, success, error}) {
    axios.delete('/mogapis/v1/post/' + id + '?', {
      headers: {
        'Authorization': 'Bearer ' + cookies.get('at')
      }
    })
    .then(function (response) {
      commit(types.CLEAR_POST, {post: response.data.data})

      if (typeof success === 'function') {
        success(response.data.data)
      }
    }).catch(function (e) {
      console.log(e)
      if (typeof error === 'function') {
        if (e.response.status === 504) {
          return error(
            504,
            {
              code: 'ServerError',
              message: 'Can not connect to server.'
            }
          )
        }
        error(e.response.status, e.response.data.error)
      }
    })
  },

  refreshPost ({ commit }, {post}) {
    commit(types.RECEIVE_POST, {post})
  },

  clearPosts ({ commit }) {
    commit(types.CLEAR_POSTS)
  },

  clearPost ({ commit }) {
    commit(types.CLEAR_POST)
  }
}

// mutations
const mutations = {
  [types.RECEIVE_POSTS] (state, {posts}) {
    state.posts.data = posts
    state.posts.gotAt = Date.now()
  },

  [types.RECEIVE_POST] (state, {post}) {
    state.post.data = post
    state.post.gotAt = Date.now()
  },

  [types.CLEAR_POSTS] (state) {
    state.posts.data = []
  },

  [types.CLEAR_POST] (state) {
    state.post.data = {}
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}
