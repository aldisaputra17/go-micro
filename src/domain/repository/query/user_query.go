package query

import "github.com/aldisaputra17/go-micro/src/module"

type UserQuery struct {
	db module.DB
}

func NewUserQuery(db module.DB) *UserQuery {
	return &UserQuery{
		db: db,
	}
}
