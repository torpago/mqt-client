set -ex


# wget "https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/7.7.0/openapi-generator-cli-7.7.0.jar"

java -jar openapi-generator-cli-7.7.0.jar generate -i https://raw.githubusercontent.com/marqeta/marqeta-openapi/main/yaml/CoreAPI.yaml  -g go -p packageName=openapi -p structPrefix=true -p enumClassPrefix=true
