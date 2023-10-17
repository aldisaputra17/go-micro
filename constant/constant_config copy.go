package constant

import "time"

// config echo runtime.
const (
	DefaultAppPort           = 8000
	DefaultAppRequestTimeout = 10 * time.Second
)

// config logger.
const (
	LogLevelDisabled = -1
	LogLevelDebug    = 0
	LogLevelInfo     = 1
	LogLevelWarn     = 2
	LogLevelError    = 3

	LogSkipCallerCount              = 4
	LogDefaultStdLogSkipCallerCount = 2
)

// config db connection.
const (
	DefaultDBMaxIdle           = 10
	DefaultDBMaxOpen           = 100
	DefaultDBMaxLifetime       = 10 * time.Minute
	DefaultDBKeepAliveInterval = 1 * time.Minute
)

// config middleware.
const (
	DefaultMdwHeaderToken  = "Authorization"
	DefaultMdwHeaderBearer = "Bearer"
	DefaultMdwRateLimiter  = 20
)
