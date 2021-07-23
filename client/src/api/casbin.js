import service from '@/utils/request'


export const UpdateCasbin = (data) => {
  return service({
    url: '/casbin/update',
    method: 'post',
    data
  })
}

export const getPolicyPathByAuthorityId = (data) => {
  return service({
    url: '/casbin/getPermByRoleId',
    method: 'post',
    data
  })
}
