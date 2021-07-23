import service from '@/utils/request'


export const deleteSysOperationRecord = (data) => {
  return service({
    url: '/opLog/delete',
    method: 'delete',
    data
  })
}

export const deleteSysOperationRecordByIds = (data) => {
  return service({
    url: '/opLog/deleteByIds',
    method: 'delete',
    data
  })
}


export const getSysOperationRecordList = (params) => {
  return service({
    url: '/opLog/list',
    method: 'get',
    params
  })
}
