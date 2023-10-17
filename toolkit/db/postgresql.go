package db

import (
	"fmt"
	"log"
	"net/url"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func NewPostgresql(opt *databaseOption) (db *sqlx.DB, err error) {
	var (
		host       string
		dataSource string
	)

	if opt.port != 0 {
		// Using port
		connURL := &url.URL{
			Scheme: opt.driver,
			User:   url.UserPassword(opt.username, opt.password),
			Host:   fmt.Sprintf("%s:%d", opt.host, opt.port),
			Path:   opt.schema,
		}
		q := connURL.Query()
		q.Add("sslmode", "disable")
		connURL.RawQuery = q.Encode()

		host = connURL.Host
		dataSource = connURL.String()
	} else {
		// Not using port (for cloudsql handling conn)
		host = opt.host
		dataSource = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
			opt.host, opt.username, opt.password, opt.schema)
	}

	db, err = sqlx.Open(opt.driver, dataSource)
	if err != nil {
		err = errors.Wrap(err, "postgres: failed to open connection")
		return
	}

	db.SetMaxIdleConns(opt.connectionOption.maxIdleConns)
	db.SetMaxOpenConns(opt.connectionOption.maxOpenConns)
	db.SetConnMaxLifetime(opt.connectionOption.maxLifetime)

	log.Println("successfully connected to postgres", host)

	// db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)

	go keepAliveConnection(db, opt.driver, opt.schema, opt.keepAliveInterval)

	return
}
