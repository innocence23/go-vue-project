import service from '@/utils/request'


export const getApiList = (data) => {
  return service({
    url: '/permission/list',
    method: 'post',
    data
  })
}


export const createApi = (data) => {
  return service({
    url: '/permission/create',
    method: 'post',
    data
  })
}


export const getApiById = (data) => {
  return service({
    url: '/permission/show',
    method: 'post',
    data
  })
}

export const updateApi = (data) => {
  return service({
    url: '/permission/update',
    method: 'post',
    data
  })
}


export const getAllApis = (data) => {
  return service({
    url: '/permission/listAll',
    method: 'post',
    data
  })
}


export const deleteApi = (data) => {
  return service({
    url: '/permission/delete',
    method: 'post',
    data
  })
}


export const deleteApisByIds = (data) => {
  return service({
    url: '/permission/deleteByIds',
    method: 'delete',
    data
  })
}
