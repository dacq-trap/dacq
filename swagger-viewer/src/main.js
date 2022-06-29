import SwaggerUI from 'swagger-ui'
import SwaggerUIStandalonePreset from 'swagger-ui-dist/swagger-ui-standalone-preset'
import { EnumNamesPlugin } from 'swagger-ui-plugin-enum-names'
import urls from './urls'

import './index.css'
import 'swagger-ui/dist/swagger-ui.css'
import 'swagger-ui-plugin-enum-names/dist/index.css'

document.addEventListener('DOMContentLoaded', () => {
  SwaggerUI({
    urls,
    dom_id: '#app',
    deepLinking: true,
    presets: [
      SwaggerUI.presets.apis,
      SwaggerUIStandalonePreset
    ],
    plugins: [
      SwaggerUI.plugins.DownloadUrl,
      EnumNamesPlugin
    ],
    layout: 'StandaloneLayout',
    queryConfigEnabled: true
  })
})
