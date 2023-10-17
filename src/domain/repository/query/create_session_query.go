package query

import (
	"context"

	"github.com/aldisaputra17/go-micro/src/domain/repository/model"
	"github.com/aldisaputra17/go-micro/src/domain/repository/query/dto"
)

func (q *UserQuery) CreateSession(
	ctx context.Context,
	arg dto.CreateSessionParam,
) (data model.Session, err error) {
	const stmt = `
		INSERT INTO sessions
			(guid, user_guid, client_ip, refresh_token, user_agent, expired_at, created_at)
		VALUES
			($1, $2, $3, $4, $5, $6, (now() at time zone 'UTC')::TIMESTAMP)
		RETURNING
			guid, user_guid, client_ip, refresh_token, user_agent, expired_at, created_at
	`

	row := q.db.QueryRowxContext(ctx, stmt,
		arg.GUID,
		arg.UserGUID,
		arg.ClientIP,
		arg.RefreshToken,
		arg.UserAgent,
		arg.ExpiredAt,
	)

	if err = row.Scan(
		&data.GUID,
		&data.UserGUID,
		&data.ClientIP,
		&data.RefreshToken,
		&data.UserAgent,
		&data.ExpiredAt,
		&data.CreatedAt,
	); err != nil {
		return
	}

	return
}
