package infrastructure

import "github.com/yuorei/member-list/app/driver/db"

type Infrastructure struct {
	db *db.DB
}

func NewInfrastructure() *Infrastructure {
	return &Infrastructure{
		db: db.NewDB(),
	}
}
