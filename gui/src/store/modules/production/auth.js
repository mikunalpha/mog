import api from '@/api/auth'
import * as types from '@/store/mutations'

// initial state
const state = {
  info: {
    gotAt: null,
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
  login ({commit}, {credentials, successCallback, errorCallback}) {
    api.login({
      credentials,
      successCallback,
      errorCallback
    })
  },

  getAuthInfo ({commit}, {successCallback, errorCallback}) {
    api.fetchAuthInfo({
      successCallback: (info) => {
        commit(types.RECEIVE_AUTH_INFO, {info})

        if (typeof successCallback === 'function') {
          successCallback(info)
        }
      },
      errorCallback
    })
  }
}

// mutations
const mutations = {
  [types.RECEIVE_AUTH_INFO] (state, {info}) {
    state.info = info
    state.info.gotAt = Date.now()
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}
