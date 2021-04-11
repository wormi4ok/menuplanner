run:
	@go run main.go

test:
	@go test ./...

api-docs:
	@docker run --rm \
       -v ${PWD}:/local openapitools/openapi-generator-cli generate \
       -i /local/api.yaml \
       -g html2 \
       -o /local/docs
