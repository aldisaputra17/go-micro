package db

import (
	"os"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func NewDatabaseConnection() (db *sqlx.DB, err error) {
	var (
		driver   string = os.Getenv("DB_DRIVER")
		host     string = os.Getenv("DB_HOST")
		port     string = os.Getenv("DB_PORT")
		username string = os.Getenv("DB_USERNAME")
		password string = os.Getenv("DB_PASSWORD")
		schema   string = os.Getenv("DB_SCHEMA")
	)
	connOpt := defaultConnectionOption()

	if maxIdle, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONN")); maxIdle > 0 && err == nil {
		connOpt.maxIdleConns = maxIdle
	}
	if maxOpen, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONN")); maxOpen > 0 && err == nil {
		connOpt.maxOpenConns = maxOpen
	}
	if maxLifetime, err := time.ParseDuration(os.Getenv("DB_MAX_LIFETIME_CONN")); maxLifetime > 0 && err == nil {
		connOpt.maxLifetime = maxLifetime
	}

	if keepAliveInterval, err := time.ParseDuration(os.Getenv("DB_KEEP_ALIVE_INTERVAL_CONN")); keepAliveInterval > 0 && err == nil {
		connOpt.keepAliveInterval = keepAliveInterval
	}
	opt, err := newDatabaseOption(driver, host, port, username, password, schema, connOpt)
	if err != nil {
		return
	}
	switch opt.driver {
	case "pgx":
		db, err = NewPostgresql(opt)
	default:
		err = errors.Wrapf(errors.New("invalid datasource driver"), "db: driver=%s", opt.driver)
	}
	return

}
