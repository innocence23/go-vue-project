import service from '@/utils/request'




export const getSystemState = () => {
  return service({
    url: '/system/getServerInfo',
    method: 'post',
    donNotShowLoading: true
  })
}
