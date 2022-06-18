import { Button } from '@mui/material'
import { setCookie } from 'nookies'

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
  setCookie(null, 'dq_session', 'session-code-for-dev-login', {
    maxAge: 30 * 24 * 60 * 60,
    path: '/',
  })
}

export default DevLogin
