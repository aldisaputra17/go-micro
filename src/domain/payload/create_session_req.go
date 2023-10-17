package payload

import "time"

type CreateSessionRequest struct {
	GUID         string    `json:"guid"`
	UserGUID     string    `json:"user_guid"`
	ClientIP     string    `json:"client_ip"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	CreatedAt    time.Time `json:"created_at"`
}
