package serverFactory

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GrpcFactory struct {
}

func (sf *GrpcFactory) Build() {
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
