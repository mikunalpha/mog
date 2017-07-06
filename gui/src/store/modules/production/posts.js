import * as types from '../../mutations'

// initial state
const state = {
  posts: [],
  post: {}
}

// getters
const getters = {
  posts: state => state.posts,
  post: state => state.post
}

// actions
const actions = {
  getPosts ({ commit }, {options, success, error}) {
    let posts = []
    for (let i = 1; i <= 10; i++) {
      posts.push({
        id: 'p' + i,
        title: 'Post ' + i + ' Title',
        content: fake,
        created_at: 1479855084,
        published: false
      })
    }
    commit(types.RECEIVE_POSTS, {posts})
  },

  getPost ({ commit }, {options, success, error}) {
    let post = {
      id: 'p1',
      title: 'Post 1 Title',
      content: fake,
      created_at: 1479855084,
      published: true
    }
    commit(types.RECEIVE_POST, {post})
  },

  clearPost ({ commit }) {
    console.log('123')
    commit(types.CLEAR_POST)
  }
}

// mutations
const mutations = {
  [types.RECEIVE_POSTS] (state, {posts}) {
    state.posts = posts
  },

  [types.RECEIVE_POST] (state, {post}) {
    state.post = post
  },

  [types.CLEAR_POST] (state) {
    state.post = {}
  }
}

let fake = '<p>In a professional context it often happens that private or corporate clients corder a publication to be made and presented with the actual content still not being ready. Think of a news blog that\'s filled with content hourly on the day of going live. However, reviewers tend to be distracted by comprehensible content, say, a random text copied from a newspaper or the internet. The are likely to focus on the text, disregarding the layout and its elements. Besides, random text risks to be unintendedly humorous or offensive, an unacceptable risk in corporate environments. Lorem ipsum and its many variants have been employed since the early 1960ies, and quite likely since the sixteenth century.</p>'

export default {
  state,
  getters,
  actions,
  mutations
}
