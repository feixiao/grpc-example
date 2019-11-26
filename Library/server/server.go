package main

import (
	"context"
	"log"
	"net/http"
	"sync"
	"api"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func startHTTP() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Connect to the GRPC server
	conn, err := grpc.Dial("localhost:5566", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	// Register grpc-gateway
	rmux := runtime.NewServeMux()

	err = api.RegisterLibraryHandlerServer(ctx, rmux, &LibServer{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("REST server ready...")
	err = http.ListenAndServe("localhost:8080", rmux)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	go startHTTP()

	// Block forever
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()

}


type LibServer struct {
}

func (s *LibServer)GetBook(ctx context.Context, in *api.GetBookRequest) (*api.Book, error) {
	return nil, nil
}

func (s *LibServer)  CreateBook(ctx context.Context, in *api.CreateBookRequest) (*api.Book, error) {
	return nil, nil
}