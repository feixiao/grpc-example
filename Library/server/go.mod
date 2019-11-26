module main

go 1.12

require (
	api v0.0.0
	github.com/grpc-ecosystem/grpc-gateway v1.12.1
	google.golang.org/grpc v1.25.1
)

replace api v0.0.0 => ../api
