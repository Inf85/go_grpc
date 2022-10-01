package service

import (
	proto "grpc/gen"
)

type MicroserviceServer struct {
	proto.UnimplementedAddServiceServer
	proto.UnimplementedUserServiceServer
}

func NewMicroservice() *MicroserviceServer {
	return &MicroserviceServer{}
}
