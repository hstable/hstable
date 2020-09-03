import Cookies from 'js-cookie'

export default ({ app }, inject) => {
  inject('cookies', Cookies)
}
