const { execSync } = require('child_process')

const SWAGGER_PATH = '../schema/openapi/openapi.yaml'
const GENERATED_DIR = 'src/api/generated'

;(async () => {
  try {
    execSync(
      `npm exec -- @openapitools/openapi-generator-cli generate -g typescript-axios -i ${SWAGGER_PATH} -o ${GENERATED_DIR}`
    )
  } catch (e) {
    console.error(e)
  }
})()
