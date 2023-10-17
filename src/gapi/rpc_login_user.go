package gapi

import (
	"context"
	"database/sql"

	"github.com/aldisaputra17/go-micro/constant"
	"github.com/aldisaputra17/go-micro/helper"
	"github.com/aldisaputra17/go-micro/src/domain/payload"
	"github.com/aldisaputra17/go-micro/src/domain/pb"
	"github.com/aldisaputra17/go-micro/src/domain/repository/query"
	"github.com/aldisaputra17/go-micro/src/domain/repository/query/dto"
	"github.com/aldisaputra17/go-micro/src/token"
	"github.com/aldisaputra17/go-micro/toolkit/log"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *GRPCServer) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (res *pb.LoginUserResponse, err error) {
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

	q := query.NewUserQuery(tx)

	user, err := q.GetUserByEmail(ctx, req.GetEmail())
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed get user")
		return
	}

	if err = helper.CompareHashPassword(user.Password, req.Password); err != nil {
		return nil, status.Errorf(codes.Internal, "password not match: %s", err)
	}

	tknPayload := token.AccessTokenPayload{
		GUID:     uuid.NewString(),
		UserGUID: user.GUID,
	}

	accessToken, err := token.CreateAccessToken(tknPayload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed create access token:")
	}

	refreshToken, err := token.CreateRefreshToken(token.RefreshTokenPayload(tknPayload))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed create refresh token:")
	}

	mtdt := server.extractMetadata(ctx)
	arg := dto.CreateSessionParam{
		GUID:         uuid.NewString(),
		UserGUID:     user.GUID,
		ClientIP:     mtdt.ClientIP,
		UserAgent:    mtdt.UserAgent,
		RefreshToken: refreshToken.Token,
		ExpiredAt:    refreshToken.ExpiredAt,
	}

	session, err := q.CreateSession(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed create session:")
	}

	userRsp := payload.ToReadUserResponse(user)

	res = &pb.LoginUserResponse{
		User: &pb.User{
			Guid:        userRsp.GUID,
			Email:       userRsp.Email,
			PhoneNumber: userRsp.PhoneNumber,
			CreatedAt:   timestamppb.New(userRsp.CreatedAt),
			CreatedBy:   userRsp.CreatedBy,
		},
		SessionGuid:           session.GUID,
		AccessToken:           accessToken.Token,
		AccessTokenExpiresAt:  timestamppb.New(accessToken.ExpiredAt),
		RefreshToken:          refreshToken.Token,
		RefreshTokenExpiresAt: timestamppb.New(refreshToken.ExpiredAt),
	}

	return
}
