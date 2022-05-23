import { AppBar, Container, Toolbar, Typography } from '@mui/material'
import { GetServerSideProps } from 'next'
import Link from 'next/link'

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
        </Toolbar>
      </Container>
    </AppBar>
  )
}

export const getServerSideProps: GetServerSideProps = async (ctx) => {
  return {
    props: {
      data: null,
    },
  }
}

export default Header
