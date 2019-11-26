module client

go 1.12

require (
	github.com/grpc-ecosystem/grpc-gateway v1.12.1
	github.com/scottyw/grpc-example/D.adding-endpoint/factory v0.0.0
	google.golang.org/grpc v1.25.1
)

replace github.com/scottyw/grpc-example/D.adding-endpoint/factory v0.0.0 => ../factory
