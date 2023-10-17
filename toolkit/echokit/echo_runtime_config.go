package echokit

import (
	"os"
	"strconv"
	"time"

	"github.com/aldisaputra17/go-micro/constant"
	"github.com/iancoleman/strcase"
	"github.com/labstack/echo/v4/middleware"
)

// RuntimeConfig defines echo REST API runtime config with healthcheck.
type RuntimeConfig struct {
	Name                    string         `json:"name,omitempty"`
	Host                    string         `json:"host,omitempty"`
	Port                    int            `json:"port,omitempty"`
	RequestTimeoutConfig    *TimeoutConfig `json:"request_timeout_config,omitempty"`
	ShutdownWaitDuration    time.Duration  `json:"shutdown_wait_duration,omitempty"`
	ShutdownTimeoutDuration time.Duration  `json:"shutdown_timeout_duration,omitempty"`
}

type TimeoutConfig struct {
	Timeout time.Duration      `json:"timeout,omitempty"`
	Skipper middleware.Skipper `json:"-"`
}

func (cfg *RuntimeConfig) validate() {
	// port
	if cfg.Port == 0 {
		cfg.Port = constant.DefaultPort
	}

	// check for timeout setting
	if cfg.RequestTimeoutConfig == nil {
		cfg.RequestTimeoutConfig = &TimeoutConfig{}
	}

	if cfg.RequestTimeoutConfig.Timeout == 0 {
		cfg.RequestTimeoutConfig.Timeout = constant.DefaultTimeout
	}

	if cfg.RequestTimeoutConfig.Skipper == nil {
		cfg.RequestTimeoutConfig.Skipper = middleware.DefaultSkipper
	}
}

func NewRuntimeConfig() *RuntimeConfig {
	r := RuntimeConfig{}

	r.Name = os.Getenv("APP_NAME")
	r.Host = os.Getenv("APP_HOST")

	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if port > 0 && err == nil {
		r.Port = port
	} else {
		r.Port = 8000
	}

	r.RequestTimeoutConfig = &TimeoutConfig{
		Skipper: middleware.DefaultSkipper,
	}

	if requestTimeout, err := time.ParseDuration(os.Getenv("APP_REQUEST_TIMEOUT")); requestTimeout > 0 && err != nil {
		r.RequestTimeoutConfig.Timeout = requestTimeout
	}

	if shutdownWait, err := time.ParseDuration(os.Getenv("APP_SHUTDOWN_WAIT")); shutdownWait > 0 && err != nil {
		r.ShutdownTimeoutDuration = shutdownWait
	}

	if shutdownTimeout, err := time.ParseDuration(os.Getenv("APP_SHUTDOWN_TIMEOUT")); shutdownTimeout > 0 && err != nil {
		r.ShutdownWaitDuration = shutdownTimeout
	}

	r.Name = strcase.ToSnake(r.Name)
	r.validate()

	return &r
}
