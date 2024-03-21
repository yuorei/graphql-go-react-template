package infrastructure

import (
	"context"

	"github.com/yuorei/member-list/app/domain"
)

func (i *Infrastructure) GetUserFromDB(ctx context.Context, id string) (*domain.User, error) {
	// var user *domain.User
	user := &domain.User{}
	err := i.db.Database.Where("id = ?", id).First(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
