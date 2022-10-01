package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	proto "grpc/gen"
	"grpc/internal/database"
	"grpc/service"
	"log"
	"net"
	"net/http"
)

func main() {
	mux := runtime.NewServeMux()
	_ = godotenv.Load()
	database.ConnectDatabase()

	err := proto.RegisterUserServiceHandlerServer(context.Background(), mux, service.NewMicroservice())
	if err != nil {
		log.Println("cannot register this service")
	}
	log.Fatalln(http.ListenAndServe(":8081", mux))
}
func grpcServer() {
	listener, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Server starting")
	grpcServer := grpc.NewServer()
	//proto.RegisterAddServiceServer(grpcServer, &server{})
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalln(err)
	}

}
