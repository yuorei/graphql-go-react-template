package application

import "github.com/yuorei/member-list/app/application/port"

type UseCase struct {
	port.UserInputPort
}

func NewUseCase(application *Application) *UseCase {
	return &UseCase{
		UserInputPort: application,
	}
}
