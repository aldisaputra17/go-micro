package query

import (
	"context"

	"github.com/aldisaputra17/go-micro/src/domain/repository/model"
)

func (q *UserQuery) GetUserByEmail(
	ctx context.Context,
	email string,
) (data model.User, err error) {
	const stmt = `
		SELECT
			guid, email, password, phone_number, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by
		FROM
			users
		WHERE
			email = $1
			AND deleted_at IS NULL
			AND deleted_by IS NULL 
	`

	if err = q.db.GetContext(ctx, &data, stmt, email); err != nil {
		return
	}

	return
}
