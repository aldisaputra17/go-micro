package echokit

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/aldisaputra17/go-micro/toolkit/log"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
)

// RunServerWithContext run graceful restapi server with existing background context
// provides default '/actuator/health' as healthcheck endpoint
// provides '/metrics' as prometheus metrics endpoint.
// set echo.Validator using `web.Validator` from `web` package.
func RunServerWithContext(ctx context.Context, e *echo.Echo, cfg *RuntimeConfig) {
	logger := log.FromCtx(ctx)

	// prometheus
	e.Use(echoprometheus.NewMiddleware(cfg.Name))  // register middleware to gather metrics from requests
	e.GET("/metrics", echoprometheus.NewHandler()) // register route to serve gathered metrics in Prometheus format

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeoutDuration)
		defer cancel()

		<-time.After(cfg.ShutdownWaitDuration)

		if err := e.Shutdown(ctx); err != nil {
			logger.Error(err, "ERROR shutdown server")
		}
	}()

	// error fallback handler
	e.HTTPErrorHandler = loggerHTTPErrorHandler(e.HTTPErrorHandler)

	// PrintRoutes(e)

	// start server
	logger.Info("serving REST HTTP server", "config", cfg)

	if err := e.Start(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Error(err, "starting http server")
	}
}

// PrintRoutes prints *echo.Echo routes.
func PrintRoutes(e *echo.Echo) {
	log.Println("== initializing http routes")

	for _, r := range e.Routes() {
		handlerNames := strings.Split(r.Name, "/")
		log.Printf("=====> %s %s %s", r.Method, r.Path, handlerNames[len(handlerNames)-1:][0])
	}
}
