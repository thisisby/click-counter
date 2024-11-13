package app

import (
	"click-counter/internal/config"
	"click-counter/internal/routes"
	"click-counter/internal/utils"
	"click-counter/pkg/httpserver"
	"click-counter/pkg/logger"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
	"os/signal"
	"syscall"
)

func MustRun() {

	logger.ZeroLogger.Info().Msg("Setting up default postgre connection...")
	conn, err := utils.SetupDefaultPostgreConnection()
	if err != nil {
		logger.ZeroLogger.Fatal().Msg(fmt.Errorf("app - MustRun - utils.SetupDefaultPostgreConnection: %w", err).Error())
	}
	defer conn.Close()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogMethod:  true,
		LogLatency: true,
		LogURI:     true,
		LogStatus:  true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.ZeroLogger.Info().
				Str("method", v.Method).
				Str("URI", v.URI).
				Int("status", v.Status).
				Str("latency", fmt.Sprintf("%dms", v.Latency.Milliseconds())).
				Msg("Request -> ")

			return nil
		},
	}))

	v1 := e.Group("/api/v1")

	// setup routers
	setupRouters(conn, v1)

	// running server
	logger.ZeroLogger.Info().Msg("Starting http server...")
	httpServer := httpserver.New(e, httpserver.Port(config.Config.Port))

	// waiting signal
	logger.ZeroLogger.Info().Msg("Configuring graceful shutdown...")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logger.ZeroLogger.Info().Msg(fmt.Sprintf("app - Run - signal: " + s.String()))
	case err = <-httpServer.Notify():
		logger.ZeroLogger.Info().Msg(fmt.Errorf("app - Run - httpServer.Notify: %w", err).Error())
	}

	// Graceful shutdown
	logger.ZeroLogger.Info().Msg("Shutting down...")
	err = httpServer.Shutdown()
	if err != nil {
		logger.ZeroLogger.Fatal().Msg(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err).Error())
	}

}

func setupRouters(conn *sqlx.DB, e *echo.Group) {
	routes.NewClickRouter(conn, e).Register()
}
