package log

import (
	"strings"

	"github.com/aldisaputra17/go-micro/constant"
)

// Level is log output level.
type Level int

var loggerCtxKey contextKey = contextKey{name: "internal-ctx-log"}

func levelString(l Level) string {
	switch l {
	case constant.LogLevelError:
		return "ERROR"
	case constant.LogLevelWarn:
		return "WARN"
	case constant.LogLevelInfo:
		return "INFO"
	default:
		return "DEBUG"
	}
}

// GetLevelFromString return error level based on config string.
func GetLevelFromString(level string) Level {
	switch strings.ToLower(level) {
	case "warn":
		return constant.LogLevelWarn
	case "debug":
		return constant.LogLevelDebug
	case "error":
		return constant.LogLevelError
	default:
		return constant.LogLevelInfo
	}
}
