import service from '@/utils/request'

export const createDict = (data) => {
  return service({
    url: '/dict/create',
    method: 'post',
    data
  })
}

export const deleteDict = (data) => {
  return service({
    url: '/dict/delete',
    method: 'delete',
    data
  })
}

export const updateDict = (data) => {
  return service({
    url: '/dict/update',
    method: 'put',
    data
  })
}

export const findDict = (params) => {
  return service({
    url: '/dict/show',
    method: 'get',
    params
  })
}

export const getDictList = (params) => {
  return service({
    url: '/dict/list',
    method: 'get',
    params
  })
}
