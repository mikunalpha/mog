import axios from 'axios'
import cookies from 'browser-cookies'

// fetchStatus gets the status of server.
let fetchStatus = ({successCallback, errorCallback}) => {
  axios.get('/mogapis/v1/status', {
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
    if (typeof error === 'function') {
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
  fetchStatus
}
