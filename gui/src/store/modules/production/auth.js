import axios from 'axios'
import * as types from '../../mutations'
import cookies from 'browser-cookies'

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

  loginAccount ({ commit }, {credentials, success, error}) {
    axios.post('/mogapis/v1/auth/login', {
      data: credentials
    }, {
      headers: {
        'Authorization': 'Bearer ' + cookies.get('at')
      }
    }).then(function (response) {
      if (typeof success === 'function') {
        success(response.data.data)
      }
    }).catch(function (e) {
      // console.log(e)
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
    axios.get('/mogapis/v1/auth/info', {
      headers: {
        'Authorization': 'Bearer ' + cookies.get('at')
      }
    })
    .then(function (response) {
      commit(types.RECEIVE_AUTH_INFO, {info: response.data.data})

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
