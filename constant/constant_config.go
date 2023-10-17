package constant

import "time"

// config db connection.
const (
	DefaultMaxIdle           = 10
	DefaultMaxOpen           = 100
	DefaultMaxLifetime       = 10 * time.Minute
	DefaultKeepAliveInterval = 1 * time.Minute
)

// config echo runtime.
const (
	DefaultPort    = 8000
	DefaultTimeout = 10 * time.Second
)
