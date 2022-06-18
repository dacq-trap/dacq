import { AppBar, Container, Toolbar, Typography } from '@mui/material'
import Link from 'next/link'
import DevLogin from './devLoginButton'

const Header = () => {
  return (
    <AppBar position='static'>
      <Container maxWidth='xl'>
        <Toolbar disableGutters>
          <Link href='/' passHref>
            <Typography
              variant='h6'
              noWrap
              component='a'
              sx={{
                display: { xs: 'none', md: 'flex' },
                fontWeight: 700,
                color: 'inherit',
                textDecoration: 'none',
              }}
            >
              DacQ
            </Typography>
          </Link>
          <DevLogin />
        </Toolbar>
      </Container>
    </AppBar>
  )
}

export default Header
