package query

import (
	"context"

	"github.com/aldisaputra17/go-micro/src/domain/repository/model"
	"github.com/aldisaputra17/go-micro/src/domain/repository/query/dto"
)

func (q *UserQuery) CreateUser(
	ctx context.Context,
	arg dto.CreateUserParam,
) (data model.User, err error) {
	const stmt = `
		INSERT INTO users
			(guid, email, password, phone_number, created_at, created_by)
		VALUES
			($1, $2, $3, $4, (now() at time zone 'UTC')::TIMESTAMP, $5)
		RETURNING
			guid, email, password, phone_number, created_at, created_by
	`

	row := q.db.QueryRowxContext(ctx, stmt,
		arg.GUID,
		arg.Email,
		arg.Password,
		arg.PhoneNumber,
		arg.CreatedBy,
	)

	if err = row.Scan(
		&data.GUID,
		&data.Email,
		&data.Password,
		&data.PhoneNumber,
		&data.CreatedAt,
		&data.CreatedBy,
	); err != nil {
		return
	}

	return
}
