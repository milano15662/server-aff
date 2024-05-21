package server

import (
	"git.selly.red/Cashbag-B2B/server-aff/internal/config"
	"git.selly.red/Cashbag-B2B/server-aff/internal/logger"
	"git.selly.red/Cashbag-B2B/server-aff/internal/mongodb"
	"git.selly.red/Cashbag-B2B/server-aff/internal/sentryio"
	"git.selly.red/Cashbag-B2B/server-aff/pkg/admin/route"
	"github.com/labstack/echo/v4"
)

// Bootstrap ...
func Bootstrap(e *echo.Echo) {
	// config
	config.InitWithLoadENV()
	cfg := config.GetENV()

	// logger
	logger.Init(cfg.Environment)

	// database
	mongodb.Connect(cfg.MongoDB)

	// sentry
	if config.IsRelease() {
		sentryio.Init(e, cfg.SentryDSN, cfg.SentryMachine, cfg.Environment)
	}
	// routes
	route.Init(e)
}
