package repository

import (
	"database/sql"
	"fmt"
	"grpc/internal/database"
	"grpc/internal/datastruct"
)

type UserQuery interface {
	GetAllUsers() ([]datastruct.User, error)
}

type userQuery struct {
}

func NewUserRepository() *userQuery {

	return &userQuery{}
}

func (u *userQuery) GetAllUsers() ([]datastruct.User, error) {
	var records []datastruct.User
	fmt.Println(database.Database())
	var err error

	if err = database.Database().Select(&records, "SELECT id, name, email FROM users"); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return records, nil
}
