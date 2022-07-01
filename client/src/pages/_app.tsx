import { CacheProvider, type EmotionCache } from '@emotion/react'
import { CssBaseline, ThemeProvider } from '@mui/material'
import type { AppProps } from 'next/app'
import Head from 'next/head'
import { RecoilRoot } from 'recoil'
import createEmotionCache from '@/common/createEmotionCache'
import theme from '@/styles/theme'

const clientSiideEmotionCache = createEmotionCache()

interface MyAppProps extends AppProps {
  emotionCache?: EmotionCache
}

function MyApp(props: MyAppProps) {
  const { Component, emotionCache = clientSiideEmotionCache, pageProps } = props
  return (
    <RecoilRoot>
      <CacheProvider value={emotionCache}>
        <Head>
          <meta name='viewport' content='initial-scale=1, width=device-width' />

          <title>DacQ</title>
          <link rel='icon' href='/favicon.ico' />
        </Head>
        <ThemeProvider theme={theme}>
          <CssBaseline />
          <Component {...pageProps} />
        </ThemeProvider>
      </CacheProvider>
    </RecoilRoot>
  )
}

export default MyApp
