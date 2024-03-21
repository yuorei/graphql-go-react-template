package port

import (
	"context"

	"github.com/yuorei/member-list/app/domain"
)

// adaputerがusecase層を呼び出されるメソッドのインターフェースを定義
type UserInputPort interface {
	GetUser(context.Context, string) (*domain.User, error)
}

type UserRepository interface {
	GetUserFromDB(context.Context, string) (*domain.User, error)
}
