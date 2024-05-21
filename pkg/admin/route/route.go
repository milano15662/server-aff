package route

import (
	routemiddleware "git.selly.red/Cashbag-B2B/server-aff/internal/middleware"
	"github.com/labstack/echo/v4"
)

// Init ...
func Init(e *echo.Echo) {
	// middlewares
	e.Use(routemiddleware.CORS())
	e.Use(routemiddleware.RateLimiter())
	e.Use(routemiddleware.SetContext)
	e.Use(routemiddleware.Auth)

	// components
	common(e)
}
