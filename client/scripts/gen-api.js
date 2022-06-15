const { execSync } = require('child_process')
const addApis = require('./add-apis')

const SWAGGER_PATH = '../schema/openapi/openapi.yaml'
const GENERATED_DIR = 'src/api/generated'

;(async () => {
  // generate api
  try {
    execSync(
      `npm exec -- @openapitools/openapi-generator-cli generate -g typescript-axios -i ${SWAGGER_PATH} -o ${GENERATED_DIR}`
    )
  } catch (e) {
    console.error(e)
  }

  // generate Apis class
  await addApis(GENERATED_DIR)
})()
