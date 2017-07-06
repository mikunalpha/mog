import axios from 'axios'
import * as types from '../../mutations'

// initial state
const state = {
  status: {
    loaded: false,
    administered: true,
    blog: {
      posts: {
        amount: 0
      }
    }
  }
}

// getters
const getters = {
  status: state => state.status
}

// actions
const actions = {
  getStatus ({ commit }, {success, error}) {
    axios.get('/mogapis/v1/status').then(function (response) {
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
  }
}

// mutations
const mutations = {
  [types.RECEIVE_STATUS] (state, { status }) {
    state.status = status
    state.status.loaded = true
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}
