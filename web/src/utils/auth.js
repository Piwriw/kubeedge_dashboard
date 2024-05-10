import Cookies from 'js-cookie'

const TokenKey = 'Admin-Token'
const UserName = "USERNAME"

export function getToken() {
  return Cookies.get(TokenKey)
}

export function setToken(token) {
  return Cookies.set(TokenKey, token)
}

export function setUsername(username) {
  return Cookies.set(UserName, username)
}

export function getUsername() {
  return Cookies.get(UserName)
}

export function removeToken() {
  return Cookies.remove(TokenKey)
}
