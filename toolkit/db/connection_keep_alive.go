package db

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

func keepAliveConnection(
	db *sqlx.DB, driver, schema string,
	duration time.Duration,
) {
	for {
		err := db.Ping()
		if err != nil {
			log.Printf("ERROR db.keepAlive driver=%s schema=%s \n%s \n \n", driver, schema, err)
		}
		time.Sleep(duration)
	}
}
