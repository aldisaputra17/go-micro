package db

import (
	"strconv"

	"github.com/pkg/errors"
)

type databaseOption struct {
	driver   string
	host     string
	port     int
	username string
	password string
	schema   string
	*connectionOption
}

func newDatabaseOption(driver, host, portStr, username,
	password, schema string, conn *connectionOption,
) (*databaseOption, error) {
	if portStr == "" {
		portStr = "0"
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, errors.Wrapf(err, "error parse int on port db env : %s", portStr)
	}

	if host == "" {
		return nil, errors.Wrapf(errors.New("invalid datasource host or port"), "db: host=%s port=%d", host, port)
	}

	if conn == nil || conn.maxOpenConns == 0 || conn.maxOpenConns < conn.maxIdleConns {
		conn = defaultConnectionOption()
	}

	return &databaseOption{
		driver:           driver,
		host:             host,
		port:             port,
		username:         username,
		password:         password,
		schema:           schema,
		connectionOption: conn,
	}, nil
}
