import { Button } from '@mui/material'
import { AxiosError } from 'axios'
import api from '@/api'

const DevLogin = () => {
  return (
    <Button
      variant='contained'
      onClick={() => {
        onClickDevLogin()
      }}
    >
      Dev Login
    </Button>
  )
}

const onClickDevLogin = async () => {
  try {
    const { data } = await api.postOauthCode({ code: 'code-for-dev-login' })
  } catch (error) {
    const err = error as AxiosError
    console.log(err)
  }
}

export default DevLogin
