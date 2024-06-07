package user_usecase

import (
	"context"
	"fullcycle-auction_go/internal/entity/user_entity"
	"fullcycle-auction_go/internal/internal_error"
)

type UserInputDTO struct {
	Name string `json:"name"`
}

type UserOutputDTO struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UserUseCase struct {
	userRepositoryInterface user_entity.UserRepositoryInterface
}

type UserUseCaseInterface interface {
	CreateUser(
		ctx context.Context,
		userInput UserInputDTO) *internal_error.InternalError

	FindUserById(
		ctx context.Context, id string) (*UserOutputDTO, *internal_error.InternalError)
}

func NewUserUseCase(
	userRepositoryInterface user_entity.UserRepositoryInterface) UserUseCaseInterface {
	return &UserUseCase{
		userRepositoryInterface: userRepositoryInterface,
	}
}

func (au *UserUseCase) CreateUser(
	ctx context.Context,
	userInput UserInputDTO) *internal_error.InternalError {
	user, err := user_entity.CreateUser(userInput.Name)
	if err != nil {
		return err
	}

	if err := au.userRepositoryInterface.CreateUser(
		ctx, user); err != nil {
		return err
	}

	return nil
}
