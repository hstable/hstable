function getFromCookie(cookie, key) {
  if (!cookie) {
    return ''
  }
  const target = cookie.split('; ').find((row) => row.startsWith(key + '='))
  if (!target) {
    return ''
  }
  return target.split('=')[1]
}

export { getFromCookie }
