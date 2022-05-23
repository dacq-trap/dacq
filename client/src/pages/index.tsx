import { Box } from '@mui/material'
import type { NextPage } from 'next'
import Link from 'next/link'
import Header from '@/components/header'

const Home: NextPage = () => {
  return (
    <>
      <Header />
      <Box padding={4}>
        <Link href='/problems/sample-problem'>サンプル問題ページ</Link>
      </Box>
    </>
  )
}

export default Home
