package service

import (
	"context"
	"database/sql"

	"github.com/aldisaputra17/go-micro/constant"
	"github.com/aldisaputra17/go-micro/src/domain/repository/model"
	"github.com/aldisaputra17/go-micro/src/domain/repository/query"
	"github.com/aldisaputra17/go-micro/src/domain/repository/query/dto"
	"github.com/aldisaputra17/go-micro/toolkit/log"
	"github.com/pkg/errors"
)

func (s *UserService) CreateSession(
	ctx context.Context,
	arg dto.CreateSessionParam,
) (data model.Session, err error) {
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

	data, err = q.CreateSession(ctx, arg)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed create session")
		return

	}

	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "error commit")
		err = errors.WithStack(constant.ErrUnknownSource)
	}

	return

}
