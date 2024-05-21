package handler

import (
	"git.selly.red/Cashbag-B2B/server-aff/internal/response"
	"github.com/labstack/echo/v4"
)

// Common ...
type Common struct{}

// Ping godoc
// @tags Common
// @summary Ping
// @id common-ping
// @accept json
// @produce json
// @router /ping [get]
func (Common) Ping(c echo.Context) error {
	return response.R200(c, "", echo.Map{})
}
