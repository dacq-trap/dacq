import { Box } from '@mui/material'
import { AxiosError } from 'axios'
import type { GetServerSideProps, NextPage } from 'next'
import Link from 'next/link'
import api from '@/api'
import Header from '@/components/header'

type Props = {
  userName: string
}

const Home = (props: Props) => {
  return (
    <>
      <Header />
      <Box padding={4}>
        ユーザー名 {props.userName}
        <Link href='/competitions/sample-competition'>
          サンプルコンペページ
        </Link>
      </Box>
    </>
  )
}

export const getServerSideProps: GetServerSideProps = async (ctx) => {
  let userName = ''
  try {
    const { data } = await api.getUsersMe()
    userName = data.name
  } catch (error) {
    const err = error as AxiosError
    console.log(err)
  }
  return {
    props: {
      userName: userName,
    },
  }
}

export default Home
