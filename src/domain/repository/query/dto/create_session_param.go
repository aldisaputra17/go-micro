package dto

import "time"

type CreateSessionParam struct {
	GUID         string    `db:"guid"`
	UserGUID     string    `db:"user_guid"`
	ClientIP     string    `db:"client_ip"`
	RefreshToken string    `db:"refresh_token"`
	UserAgent    string    `db:"user_agent"`
	ExpiredAt    time.Time `db:"expired_at"`
}
