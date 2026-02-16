package api

//go:generate go tool oapi-codegen -package=api -generate "types" -o types.gen.go openapi.yaml
//go:generate go tool oapi-codegen -package=api -generate "std-http" -o server.gen.go openapi.yaml
