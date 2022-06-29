import { defineConfig } from 'vite'

export default defineConfig({
  build: {
    rollupOptions: {
      output: {
        manualChunks: id => {
          if (id.includes('/swagger-ui/')) {
            return 'core'
          }
          if (id.includes('/swagger-client/')) {
            return 'swagger-client'
          }
          if (id.includes('/autolinker/') || id.includes('/remarkable') || id.includes('/dompurify/') || id.includes('xml-but-prettier') || id.includes('js-yaml')) {
            return 'library-misc'
          }
          if (id.includes('/react') || id.includes('/redux') || id.includes('/immutable') || id.includes('/prop-types/')) {
            return 'react'
          }
          if (id.includes('/lodash') || id.includes('core-js') || id.includes('babel')) {
            return 'lodash-and-babel'
          }
          return null
        }
      }
    }
  }
})
