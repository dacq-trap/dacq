#!/bin/sh
docker run --init --rm \
-v $(pwd)/schema/openapi:/tmp \
-p 4010:4010 \
stoplight/prism:4 mock \
-h 0.0.0.0 \
"/tmp/openapi.yaml"