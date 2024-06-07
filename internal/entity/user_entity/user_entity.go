package user_entity

import (
	"context"
	"fullcycle-auction_go/internal/internal_error"

	"github.com/google/uuid"
)

type User struct {
	Id   string
	Name string
}

type UserRepositoryInterface interface {
	CreateUser(
		ctx context.Context,
		userEntity *User) *internal_error.InternalError

	FindUserById(
		ctx context.Context, userId string) (*User, *internal_error.InternalError)
}

func (au *User) Validate() *internal_error.InternalError {
	if len(au.Name) <= 1 {
		return internal_error.NewBadRequestError("invalid user object")
	}
	return nil
}

func CreateUser(name string) (*User, *internal_error.InternalError) {
	user := &User{
		Id:   uuid.New().String(),
		Name: name,
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}
