package module

import "github.com/jmoiron/sqlx"

type Module struct {
	db *sqlx.DB
}

func NewModule(db *sqlx.DB) *Module {
	return &Module{
		db: db,
	}
}

func (s *Module) GetDB() *sqlx.DB {
	return s.db
}
