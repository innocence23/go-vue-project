import service from '@/utils/request'

export const asyncMenu = () => {
  return service({
    url: '/menu/treeList',
    method: 'post'
  })
}

export const getMenuList = (data) => {
  return service({
    url: '/menu/list',
    method: 'post',
    data
  })
}

export const addBaseMenu = (data) => {
  return service({
    url: '/menu/create',
    method: 'post',
    data
  })
}

export const getBaseMenuById = (data) => {
  return service({
    url: '/menu/show',
    method: 'post',
    data
  })
}

export const deleteBaseMenu = (data) => {
  return service({
    url: '/menu/delete',
    method: 'post',
    data
  })
}

export const updateBaseMenu = (data) => {
  return service({
    url: '/menu/update',
    method: 'post',
    data
  })
}

export const getBaseMenuTree = () => {
  return service({
    url: '/menu/getUidMenu',
    method: 'post'
  })
}

export const getMenuAuthority = (data) => {
  return service({
    url: '/menu/getRoleMenu',
    method: 'post',
    data
  })
}

export const addMenuAuthority = (data) => {
  return service({
    url: '/menu/addRoleMenu',
    method: 'post',
    data
  })
}