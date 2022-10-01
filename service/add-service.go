package service

import (
	"context"
	proto "grpc/gen"
)

func (s *MicroserviceServer) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a := request.GetA()
	b := request.GetB()
	res := a + b

	return &proto.Response{
		Result: res,
	}, nil
}
