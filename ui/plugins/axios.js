// import { getFromCookie } from '../assets/js/cookieTool'

export default function ({ $axios, redirect, req, app }) {
  $axios.interceptors.request.use(
    (config) => {
      // console.log(config.baseURL, config.url, process.browser)
      if (process.browser) {
        let token = ''
        token = app.$cookies.get('token')
        // 备用存储地点
        if (!token && localStorage.token) {
          token = localStorage.token
        }
        // console.log('token', token)
        if (!token && !location.pathname.startsWith('/login')) {
          redirect('/login')
        }
        token && (config.headers.Authorization = token)
      }
      return config
    },
    (err) => {
      return Promise.reject(err)
    }
  )

  $axios.interceptors.response.use(
    (response) => {
      return response
    },
    (error) => {
      if (error.response && error.response.status === 401) {
        redirect('/login')
      }
      return Promise.reject(error)
    }
  )
}
