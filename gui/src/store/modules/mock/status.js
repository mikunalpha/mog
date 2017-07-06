import * as types from '../../mutations'

// initial state
const state = {
  status: {
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
  stauts: state => state.status
}

// actions
const actions = {
  getStatus ({ commit }, {success, error}) {
    let status = {
      administered: true,
      blog: {
        posts: {
          amount: 0
        }
      }
    }
    commit(types.RECEIVE_STATUS, {status})
    if (typeof success === 'function') {
      success(status)
    }
  }
}

// mutations
const mutations = {
  [types.RECEIVE_STATUS] (state, { status }) {
    state.status = status
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}
