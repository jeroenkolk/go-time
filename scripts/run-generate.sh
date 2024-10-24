docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate -i /local/openapi.yaml -g go-server -o /local/generated-models
#--skip-operations --skip-validators
