package echocontext

import (
	"git.selly.red/Cashbag-B2B/server-aff/internal/appcontext"
	"github.com/labstack/echo/v4"
)

const KeyContext = "ctx"

// GetContext ...
func GetContext(c echo.Context) *appcontext.AppContext {
	return c.Get(KeyContext).(*appcontext.AppContext)
}

// SetContext ...
func SetContext(c echo.Context, value interface{}) {
	c.Set(KeyContext, value)
}
