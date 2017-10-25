import axios from 'axios'
import cookies from 'browser-cookies'

// createAccount creates a new account to server.
let createAccount = ({newAccount, successCallback, errorCallback}) => {
  // Deal with newPost here ...
  // ...

  // Send request
  axios.post('/mogapis/v1/account', {
    data: newAccount
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

export default {
  createAccount
}
