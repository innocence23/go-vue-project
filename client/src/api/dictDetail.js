import service from '@/utils/request'

export const createDictDetail = (data) => {
  return service({
    url: '/dictDetail/createdictDetail',
    method: 'post',
    data
  })
}

export const deleteDictDetail = (data) => {
  return service({
    url: '/dictDetail/deleteDictDetail',
    method: 'delete',
    data
  })
}

export const updateDictDetail = (data) => {
  return service({
    url: '/dictDetail/updateDictDetail',
    method: 'put',
    data
  })
}

export const findDictDetail = (params) => {
  return service({
    url: '/dictDetail/findDictDetail',
    method: 'get',
    params
  })
}

export const getDictDetailList = (params) => {
  return service({
    url: '/dictDetail/getDictDetailList',
    method: 'get',
    params
  })
}
