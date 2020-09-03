import { getFromCookie } from '../assets/js/cookieTool'

export default function ({ $axios, redirect, req, app }) {
  $axios.interceptors.request.use(
    (config) => {
      let token = ''
      if (process.browser) {
        token = app.$cookies.token
        if (!token && !location.pathname.startsWith('/login')) {
          redirect('/login')
        }
      } else {
        token = getFromCookie(req.headers.cookie, 'token')
      }
      config.headers.Authorization = `Bearer ${token}`
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
