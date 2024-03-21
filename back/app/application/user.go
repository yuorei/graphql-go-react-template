package application

import (
	"context"

	"github.com/yuorei/member-list/app/application/port"
	"github.com/yuorei/member-list/app/domain"
)

type UserUseCase struct {
	userRepository port.UserRepository
}

func NewUserUseCase(userRepository port.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepository: userRepository,
	}
}

func (a *Application) GetUser(ctx context.Context, id string) (*domain.User, error) {
	return a.User.userRepository.GetUserFromDB(ctx, id)
}
