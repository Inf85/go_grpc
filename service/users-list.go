package service

import (
	"fmt"
	"grpc/internal/datastruct"
	"grpc/internal/repository"
)

type userService struct {
	repo repository.UserQuery
}

func NewUserService(repo repository.UserQuery) *userService {
	return &userService{
		repo: repo,
	}
}

func (u *userService) GetUsers() []datastruct.User {
	users, err := u.repo.GetAllUsers()
	if err != nil {
		panic(err)
	}
	fmt.Println(users)
	return users
}
