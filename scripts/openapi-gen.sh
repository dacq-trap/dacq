#!/bin/sh
docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
    -i /local/schema/openapi/openapi.yaml \
    -g typescript-axios \
    -o /local/client/src/api/generated