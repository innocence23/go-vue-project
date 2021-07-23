import service from '@/utils/request'

export const getSystemState = () => {
  return service({
    url: '/machine/info',
    method: 'get',
    donNotShowLoading: true
  })
}
