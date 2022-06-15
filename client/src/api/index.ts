import { Apis, Configuration } from '@/api/generated'

const api = new Apis(new Configuration({ basePath: '/api' }))

export default api
export * from '@/api/generated'
