import service from '@/utils/request'

export const getApiList = (data) => {
  return service({
    url: '/permission/list',
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


export const getAllApis = (data) => {
  return service({
    url: '/permission/listAll',
    method: 'post',
    data
  })
}

