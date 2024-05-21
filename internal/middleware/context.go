package routemiddleware

import (
	"git.selly.red/Cashbag-B2B/server-aff/internal/appcontext"
	"git.selly.red/Cashbag-B2B/server-aff/internal/echocontext"
	"github.com/labstack/echo/v4"
)

// SetContext ...
func SetContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		appCtx := appcontext.New(c.Request().Context())
		echocontext.SetContext(c, appCtx)

		return next(c)
	}
}
