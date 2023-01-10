import { AxiosError } from 'axios'
import { atom, type SetterOrUpdater, useRecoilCallback } from 'recoil'
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

export const fetchMe = async (setMe: SetterOrUpdater<Me>) => {
  try {
    const { data } = await api.getMe()
    const me: Me = {
      ...data,
      isLogin: true,
    }
    setMe(me)
  } catch (error) {
    const err = error as AxiosError
    console.log(err)
  }
}
