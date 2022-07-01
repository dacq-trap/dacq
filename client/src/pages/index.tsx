import { Box } from '@mui/material'
import Link from 'next/link'
import { useRecoilValue } from 'recoil'
import Header from '@/components/header'
import { meState } from '@/store/domain/me'

const Home = () => {
  const me = useRecoilValue(meState)
  return (
    <>
      <Header />
      <Box padding={4}>
        ユーザー名 {me.name}
        <Link href='/competitions/9e2b5722-e1b8-c7de-44db-1af41a7f6c12'>
          サンプルコンペページ
        </Link>
      </Box>
    </>
  )
}

export default Home
