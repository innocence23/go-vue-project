import service from '@/utils/request'

export const jsonInBlacklist = () => {
  return service({
    url: '/jwt/inBlacklist',
    method: 'post'
  })
}
