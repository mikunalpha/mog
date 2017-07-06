import axios from 'axios'
import * as types from '../../mutations'

// initial state
const state = {
  info: {
    role: 1
  },
  roles: {
    Admin: 9999,
    Anoymous: 1
  }
}

// getters
const getters = {
  authInfo: state => state.info,
  roles: state => state.roles
}

// actions
const actions = {
  createAdmin ({ commit }, {account, success, error}) {
    axios.post('/mogapis/v1/account', {
      data: account
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

  getAuthInfo ({ commit }, {success, error}) {
    let info = {
      id: 'iii',
      username: 'uuu',
      role: state.roles.Admin
    }
    commit(types.RECEIVE_AUTH_INFO, {info})
    if (typeof success === 'function') {
      success(info)
    }
  }

  // login ({ commit }, {credentials, success, error}) {
  //   AuthAPI.loginAccount(credentials, success, error)
  // }
}

// mutations
const mutations = {
  [types.RECEIVE_AUTH_INFO] (state, { info }) {
    state.info = info
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}
