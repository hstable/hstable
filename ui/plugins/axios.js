import axios from 'axios'
import router from 'vue-router'

axios.interceptors.request.use(
  (config) => {
    if (localStorage.token) {
      config.headers.Authorization = `Bearer ${localStorage.token}`
    }
    return config
  },
  (err) => {
    return Promise.reject(err)
  }
)

axios.interceptors.response.use(
  (response) => {
    return response
  },
  (error) => {
    if (error.response && error.response.status === 401) {
      router.replace({
        path: 'login',
      })
    }
    return Promise.reject(error)
  }
)
