package route

import (
	"git.selly.red/Cashbag-B2B/server-aff/pkg/admin/handler"
	"github.com/labstack/echo/v4"
)

// common ...
func common(e *echo.Echo) {
	var (
		g = e.Group("")
		h = handler.Common{}
	)

	// Ping ...
	g.GET("/ping", h.Ping)
}
