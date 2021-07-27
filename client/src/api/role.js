import service from '@/utils/request'


export const setRoleUser = (data) => {
  return service({
    url: '/role/setRoleUser',
    method: 'post',
    data
  })
}

export const setDataRole = (data) => {
  return service({
    url: '/authority/setDataRole',
    method: 'post',
    data
  })
}


export const getRoleList = (data) => {
  return service({
    url: '/role/list',
    method: 'post',
    data
  })
}

export const createRole = (data) => {
  return service({
    url: '/authority/createRole',
    method: 'post',
    data
  })
}





export const updateRole = (data) => {
  return service({
    url: '/authority/updateRole',
    method: 'put',
    data
  })
}

export const deleteRole = (data) => {
  return service({
    url: '/authority/deleteRole',
    method: 'post',
    data
  })
}

