package service

import (
	"context"

	"github.com/aldisaputra17/go-micro/src/domain/repository/model"
	"github.com/aldisaputra17/go-micro/src/domain/repository/query"
	"github.com/aldisaputra17/go-micro/toolkit/log"
)

func (s *UserService) GetUser(
	ctx context.Context,
	email string,
) (data model.User, err error) {
	q := query.NewUserQuery(s.db)

	data, err = q.GetUserByEmail(ctx, email)
	if err != nil {
		log.FromCtx(ctx).Error(err, "verify user failed")
		return

	}

	return
}
