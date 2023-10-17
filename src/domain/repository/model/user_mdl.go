package model

import (
	"database/sql"
	"time"
)

type User struct {
	GUID        string         `db:"guid"`
	Email       string         `db:"email"`
	Password    string         `db:"password"`
	PhoneNumber string         `db:"phone_number"`
	CreatedAt   time.Time      `db:"created_at"`
	CreatedBy   string         `db:"created_by"`
	UpdatedAt   sql.NullTime   `db:"updated_at"`
	UpdatedBy   sql.NullString `db:"updated_by"`
	DeletedAt   sql.NullTime   `db:"deleted_at"`
	DeletedBy   sql.NullString `db:"deleted_by"`
}
