import { AxiosError } from 'axios'
import { atom, useRecoilCallback } from 'recoil'
import api from '@/api'

interface Me {
  name: string
  iconUrl: string
  isLogin: boolean
}

const defaultData: Me = {
  name: '',
  iconUrl: '',
  isLogin: false,
}

export const meState = atom<Me>({
  key: 'meState',
  default: defaultData,
})

export const useMeDispatch = () => {
  const fetchMe = useRecoilCallback(({ set }) => async () => {
    try {
      const { data } = await api.getMe()
      const me: Me = {
        ...data,
        isLogin: true,
      }
      set(meState, me)
    } catch (error) {
      const err = error as AxiosError
      console.log(err)
    }
  })

  return {
    fetchMe,
  }
}
