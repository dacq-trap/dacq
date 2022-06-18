import { AxiosRequestConfig } from 'axios'
import { GetServerSidePropsContext } from 'next'

export const requestOption = (
  ctx: GetServerSidePropsContext
): AxiosRequestConfig => {
  return {
    headers: ctx.req.headers.cookie
      ? { cookie: ctx.req.headers.cookie }
      : undefined,
  }
}
