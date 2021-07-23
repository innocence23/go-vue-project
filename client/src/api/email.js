import service from '@/utils/request'

export const emailTest = (data) => {
  return service({
    url: '/email/test',
    method: 'post',
    data
  })
}
