package db

import (
	"time"

	"github.com/aldisaputra17/go-micro/constant"
)

type connectionOption struct {
	maxIdleConns      int
	maxOpenConns      int
	maxLifetime       time.Duration
	keepAliveInterval time.Duration
}

func defaultConnectionOption() *connectionOption {
	return &connectionOption{
		maxIdleConns:      constant.DefaultMaxIdle,
		maxOpenConns:      constant.DefaultMaxOpen,
		maxLifetime:       constant.DefaultMaxLifetime,
		keepAliveInterval: constant.DefaultKeepAliveInterval,
	}
}
