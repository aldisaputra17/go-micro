package echokit

import (
	"context"

	"github.com/aldisaputra17/go-micro/src/api"
	"github.com/aldisaputra17/go-micro/src/module"
	"github.com/aldisaputra17/go-micro/toolkit/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RunEchoHTTP(ctx context.Context, mdl *module.Module) {
	runtimeCfg := NewRuntimeConfig()

	e := echo.New()

	e.HideBanner = true

	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: runtimeCfg.RequestTimeoutConfig.Timeout,
		Skipper: runtimeCfg.RequestTimeoutConfig.Skipper,
	}))

	e.Validator = config.NewValidator()

	api.Routes(e, mdl)

	// run actual server
	RunServerWithContext(ctx, e, runtimeCfg)
}
