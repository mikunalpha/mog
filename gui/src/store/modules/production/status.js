import api from '@/api/status'
import * as types from '@/store/mutations'

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
  keepAlive: true
}

// getters
const getters = {
  status: state => state.status,
  keepAlive: state => state.keepAlive
}

// actions
const actions = {
  getStatus ({commit}, {successCallback, errorCallback}) {
    api.fetchStatus({
      successCallback: (status) => {
        commit(types.RECEIVE_STATUS, {status})

        if (typeof success === 'function') {
          successCallback(status)
        }
      },
      errorCallback
    })
  },

  enableKeepAlive ({commit}) {
    commit(types.ENABLE_KEEP_ALIVE)
  },

  disableKeepAlive ({commit}) {
    commit(types.DISABLE_KEEP_ALIVE)
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
  },

  [types.ENABLE_KEEP_ALIVE] (state) {
    state.keepAlive = true
  },

  [types.DISABLE_KEEP_ALIVE] (state) {
    state.keepAlive = false
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}
