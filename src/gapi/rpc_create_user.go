package gapi

import (
	"context"
	"database/sql"

	"github.com/aldisaputra17/go-micro/constant"
	"github.com/aldisaputra17/go-micro/helper"
	"github.com/aldisaputra17/go-micro/src/domain/pb"
	"github.com/aldisaputra17/go-micro/src/domain/repository/query"
	"github.com/aldisaputra17/go-micro/src/domain/repository/query/dto"
	"github.com/aldisaputra17/go-micro/toolkit/log"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *GRPCServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (res *pb.CreateUserResponse, err error) {
	tx, err := server.mdl.GetDB().BeginTxx(ctx, &sql.TxOptions{})
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

	query := query.NewUserQuery(tx)

	hashedPassword, err := helper.GeneratePassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed hash password: %s", err)
	}

	arg := dto.CreateUserParam{
		GUID:        uuid.NewString(),
		Email:       req.GetEmail(),
		Password:    hashedPassword,
		PhoneNumber: req.GetPhoneNumber(),
		CreatedBy:   req.GetCreatedBy(),
	}

	usr, err := query.CreateUser(ctx, arg)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed create user")
		return
	}

	res = &pb.CreateUserResponse{
		User: &pb.User{
			Guid:        usr.GUID,
			Email:       usr.Email,
			PhoneNumber: usr.PhoneNumber,
			CreatedAt:   timestamppb.New(usr.CreatedAt),
			CreatedBy:   usr.CreatedBy,
		},
	}

	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "error commit")
		err = errors.WithStack(constant.ErrUnknownSource)
	}

	return
}
