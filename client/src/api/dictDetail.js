import service from '@/utils/request'

export const createDictDetail = (data) => {
  return service({
    url: '/dict-detail/create',
    method: 'post',
    data
  })
}

export const deleteDictDetail = (data) => {
  return service({
    url: '/dict-detail/delete',
    method: 'delete',
    data
  })
}

export const updateDictDetail = (data) => {
  return service({
    url: '/dict-detail/update',
    method: 'put',
    data
  })
}

export const findDictDetail = (params) => {
  return service({
    url: '/dict-detail/show',
    method: 'get',
    params
  })
}

export const getDictDetailList = (params) => {
  return service({
    url: '/dict-detail/list',
    method: 'get',
    params
  })
}
