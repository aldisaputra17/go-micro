package service

import (
	"context"
	"database/sql"

	"github.com/aldisaputra17/go-micro/constant"
	"github.com/aldisaputra17/go-micro/src/domain/payload"
	"github.com/aldisaputra17/go-micro/src/domain/repository/model"
	"github.com/aldisaputra17/go-micro/src/domain/repository/query"
	"github.com/aldisaputra17/go-micro/toolkit/log"
	"github.com/pkg/errors"
)

func (s *UserService) CreateUser(
	ctx context.Context,
	req payload.CreateUserRequest,
) (data model.User, err error) {
	tx, err := s.db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed begin tx")
		err = errors.WithStack(constant.ErrUnknownSource)

		return
	}

	defer func() {
		if err != nil {
			if errRollback := tx.Rollback(); errRollback != nil {
				log.FromCtx(ctx).Error(err, "error rollback", errRollback)
				err = errors.WithStack(constant.ErrUnknownSource)

				return
			}
		}
	}()

	q := query.NewUserQuery(tx)

	data, err = q.CreateUser(ctx, req.ToParam())
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed create user")
		return
	}

	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "error commit")
		err = errors.WithStack(constant.ErrUnknownSource)
	}

	return
}
