all: api
	go build cmd/main.go

api:
	oapi-codegen -package=api \
	-generate "types" \
	api/openapi.yaml > internal/api/types.gen.go;
	oapi-codegen -package=api \
	-generate "std-http" \
	api/openapi.yaml > internal/api/server.gen.go;
