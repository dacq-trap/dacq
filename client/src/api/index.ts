import axios from 'axios'
import { wrapper } from 'axios-cookiejar-support'
import { CookieJar } from 'tough-cookie'
import { Apis } from '@/api/generated'

const jar = new CookieJar()
const axiosInstance = wrapper(axios.create({ withCredentials: true, jar }))

// TODO: 開発環境と本番環境とでbasePathを分けられるようにする
const api = new Apis(undefined, 'http://localhost:4010', axiosInstance)

export default api
export * from '@/api/generated'
