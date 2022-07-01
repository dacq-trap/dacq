import { Button } from '@mui/material'
import { setCookie } from 'nookies'
import { useSetRecoilState } from 'recoil'
import { fetchMe, meState } from '@/store/domain/me'

const DevLogin = () => {
  const setMe = useSetRecoilState(meState)
  return (
    <Button
      variant='contained'
      onClick={async () => {
        onClickDevLogin()
        await fetchMe(setMe)
      }}
    >
      Dev Login
    </Button>
  )
}

const onClickDevLogin = () => {
  setCookie(null, 'dq_session', 'session-code-for-dev-login', {
    maxAge: 30 * 24 * 60 * 60,
    path: '/',
  })
}

export default DevLogin
