rm -rf generated
mkdir generated
oapi-codegen --config=config.yaml api/openapi.yaml
