import Cookies from 'js-cookie'

const TokenKey = 'afkser_admin_live_house_sk'
const UserIdKey = 'afkser_admin_live_house_uk';

export function getToken() {
  return Cookies.get(TokenKey)
}

export function setToken(token) {
  return Cookies.set(TokenKey, token)
}

export function removeToken() {
  Cookies.remove(UserIdKey)
  return Cookies.remove(TokenKey)
}

export function getUserId(){
  return Cookies.get(UserIdKey)
}

export function setUserId(userId){
  return Cookies.set(UserIdKey, userId)
}

export function removeUserId(){
  Cookies.remove(TokenKey)
  return Cookies.remove(UserIdKey)
}