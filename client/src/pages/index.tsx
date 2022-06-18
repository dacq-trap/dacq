import { Box } from '@mui/material'
import { AxiosError, AxiosRequestConfig } from 'axios'
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
    const option: AxiosRequestConfig = {
      headers: ctx.req.headers.cookie
        ? { cookie: ctx.req.headers.cookie }
        : undefined,
    }
    const { data } = await api.getMe(option)
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
