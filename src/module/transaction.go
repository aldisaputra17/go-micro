package module

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type DB interface {
	SelectContext(context.Context, interface{}, string, ...interface{}) error
	GetContext(context.Context, interface{}, string, ...interface{}) error
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	QueryxContext(context.Context, string, ...interface{}) (*sqlx.Rows, error)
	QueryRowxContext(context.Context, string, ...interface{}) *sqlx.Row
}
