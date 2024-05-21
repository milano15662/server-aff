package echocontext

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

const keyQuery = "query"
const keyBody = "body"
const keyFile = "file"

// GetQuery ...
func GetQuery(c echo.Context) interface{} {
	return c.Get(keyQuery)
}

// SetQuery ...
func SetQuery(c echo.Context, value interface{}) {
	c.Set(keyQuery, value)
}

// GetBody ...
func GetBody(c echo.Context) interface{} {
	return c.Get(keyBody)
}

// SetBody ...
func SetBody(c echo.Context, value interface{}) {
	c.Set(keyBody, value)
}

// GetFile ...
func GetFile(c echo.Context) *multipart.FileHeader {
	return c.Get(keyFile).(*multipart.FileHeader)
}

// SetFile ...
func SetFile(c echo.Context, value *multipart.FileHeader) {
	c.Set(keyFile, value)
}
