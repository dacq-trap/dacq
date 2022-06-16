import { Apis, Configuration } from '@/api/generated'

// TODO: 開発環境と本番環境とでbasePathを分けられるようにする
const api = new Apis(new Configuration({ basePath: 'http://localhost:4010' }))

export default api
export * from '@/api/generated'
