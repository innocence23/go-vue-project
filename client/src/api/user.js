import service from '@/utils/request'

export const login = (data) => {
  return service({
    url: '/basic/login',
    method: 'post',
    data: data
  })
}

export const captcha = (data) => {
  return service({
    url: '/basic/captcha',
    method: 'post',
    data: data
  })
}

export const register = (data) => {
  return service({
    url: '/user/register',
    method: 'post',
    data: data
  })
}

export const changePassword = (data) => {
  return service({
    url: '/user/changePassword',
    method: 'post',
    data: data
  })
}

export const getUserList = (data) => {
  return service({
    url: '/user/list',
    method: 'post',
    data: data
  })
}

export const setUserAuthority = (data) => {
  return service({
    url: '/user/setRole',
    method: 'post',
    data: data
  })
}

export const deleteUser = (data) => {
  return service({
    url: '/user/delete',
    method: 'delete',
    data: data
  })
}

export const setUserInfo = (data) => {
  return service({
    url: '/user/update',
    method: 'put',
    data: data
  })
}
