import axios from 'axios'
import cookies from 'browser-cookies'
import querystring from 'querystring'

// fetchPosts gets multiple posts from server.
let fetchPosts = ({options, successCallback, errorCallback}) => {
  // Deal with options here ...
  // ...

  // Send request
  axios.get('/mogapis/v1/posts?' + querystring.stringify(options), {
    headers: {
      'Authorization': 'Bearer ' + cookies.get('at')
    }
  })
  .then((response) => {
    if (typeof successCallback === 'function') {
      successCallback(response.data.data)
    }
  })
  .catch((e) => {
    if (typeof errorCallback === 'function') {
      if (typeof e.response === 'undefined' || e.response.status === 504) {
        return errorCallback({
          status: 500,
          error: {
            code: 'ServerError',
            message: 'Can not connect to server.'
          }
        })
      }
      errorCallback({
        status: e.response.status,
        error: e.response.data.error
      })
    }
  })
}

// fetchPost gets an identified post from server.
let fetchPost = ({id, options, successCallback, errorCallback}) => {
  // Deal with options here ...
  // ...

  // Send request
  axios.get('/mogapis/v1/post/' + id + '?' + querystring.stringify(options), {
    headers: {
      'Authorization': 'Bearer ' + cookies.get('at')
    }
  })
  .then((response) => {
    if (typeof successCallback === 'function') {
      successCallback(response.data.data)
    }
  })
  .catch((e) => {
    if (typeof errorCallback === 'function') {
      if (typeof e.response === 'undefined' || e.response.status === 504) {
        return errorCallback({
          status: 500,
          error: {
            code: 'ServerError',
            message: 'Can not connect to server.'
          }
        })
      }
      errorCallback({
        status: e.response.status,
        error: e.response.data.error
      })
    }
  })
}

// createPost creates a new post to server.
let createPost = ({newPost, successCallback, errorCallback}) => {
  // Deal with newPost here ...
  // ...

  // Send request
  axios.post('/mogapis/v1/post', {
    data: newPost
  }, {
    headers: {
      'Authorization': 'Bearer ' + cookies.get('at')
    }
  })
  .then((response) => {
    if (typeof successCallback === 'function') {
      successCallback(response.data.data)
    }
  })
  .catch((e) => {
    if (typeof errorCallback === 'function') {
      if (typeof e.response === 'undefined' || e.response.status === 504) {
        return errorCallback({
          status: 500,
          error: {
            code: 'ServerError',
            message: 'Can not connect to server.'
          }
        })
      }
      errorCallback({
        status: e.response.status,
        error: e.response.data.error
      })
    }
  })
}

// updatePost updates a exist post.
let updatePost = ({post, successCallback, errorCallback}) => {
  // Deal with post here ...
  // ...

  // Send request
  axios.patch('/mogapis/v1/post/' + post.id, {
    data: post
  }, {
    headers: {
      'Authorization': 'Bearer ' + cookies.get('at')
    }
  })
  .then((response) => {
    if (typeof successCallback === 'function') {
      successCallback(response.data.data)
    }
  })
  .catch((e) => {
    if (typeof errorCallback === 'function') {
      if (typeof e.response === 'undefined' || e.response.status === 504) {
        return errorCallback({
          status: 500,
          error: {
            code: 'ServerError',
            message: 'Can not connect to server.'
          }
        })
      }
      errorCallback({
        status: e.response.status,
        error: e.response.data.error
      })
    }
  })
}

// deletePost deletes a exist post.
let deletePost = ({id, successCallback, errorCallback}) => {
  // Deal with id here ...
  // ...

  // Send request
  axios.delete('/mogapis/v1/post/' + id, {
    headers: {
      'Authorization': 'Bearer ' + cookies.get('at')
    }
  })
  .then((response) => {
    if (typeof successCallback === 'function') {
      successCallback(response.data.data)
    }
  })
  .catch((e) => {
    if (typeof errorCallback === 'function') {
      if (typeof e.response === 'undefined' || e.response.status === 504) {
        return errorCallback({
          status: 500,
          error: {
            code: 'ServerError',
            message: 'Can not connect to server.'
          }
        })
      }
      errorCallback({
        status: e.response.status,
        error: e.response.data.error
      })
    }
  })
}

export default {
  fetchPosts,
  fetchPost,
  createPost,
  updatePost,
  deletePost
}
