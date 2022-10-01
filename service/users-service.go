package service

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	proto "grpc/gen"
	"grpc/internal/repository"
)

func (s *MicroserviceServer) GetUserList(ctx context.Context, empty *emptypb.Empty) (*proto.UsersListResponse, error) {
	fmt.Println("TEST")
	usersRepo := repository.NewUserRepository()
	userService := NewUserService(usersRepo)
	result := userService.GetUsers()

	usersResponse := make([]*proto.UsersListResponse_Users, 0, len(result))
	for _, user := range result {
		usersResponse = append(usersResponse, &proto.UsersListResponse_Users{
			Id:      user.ID,
			UserNam: user.Name,
			Email:   user.Email,
		})
	}
	return &proto.UsersListResponse{
		Code: &proto.UsersListResponse_Code{
			Code:    "200",
			Message: "Success",
		},
		Users: usersResponse,
	}, nil
}
