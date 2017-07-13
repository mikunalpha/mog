import axios from 'axios'
import * as types from '../../mutations'
import cookies from 'browser-cookies'

// initial state
const state = {
  status: {
    gotAt: null,
    administered: null,
    blog: {
      posts: {
        amount: 0
      }
    }
  },
  scroll: {
    x: 0,
    y: 0
  }
}

// getters
const getters = {
  status: state => state.status,
  scroll: state => state.scroll
}

// actions
const actions = {
  getStatus ({ commit }, {success, error}) {
    axios.get('/mogapis/v1/status', {
      headers: {
        'Authorization': 'Bearer ' + cookies.get('at')
      }
    })
    .then(function (response) {
      commit(types.RECEIVE_STATUS, {status: response.data.data})

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

  saveScroll ({ commit }, {x, y}) {
    commit(types.RECEIVE_SCROLL, {x, y})
  }
}

// mutations
const mutations = {
  [types.RECEIVE_STATUS] (state, { status }) {
    state.status = status
    state.status.gotAt = Date.now()
  },

  [types.RECEIVE_SCROLL] (state, { x, y }) {
    state.scroll.x = x
    state.scroll.y = y
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}
