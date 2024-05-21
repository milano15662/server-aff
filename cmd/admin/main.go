package main

import (
	"context"
	"errors"
	"git.selly.red/Cashbag-B2B/server-aff/internal/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "git.selly.red/Cashbag-B2B/server-aff/docs/admin"
	"git.selly.red/Cashbag-B2B/server-aff/pkg/admin/server"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Server AFF - API ADMIN
// @version 1.0
// @description All APIs ADMIN for Server AFF.
// @description
// @termsOfService https://cashbag-b2b.vn
// @contact.name Cashbag-B2B
// @contact.url https://cashbag-b2b.vn
// @contact.email
// @basePath

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// echo instance
	e := echo.New()

	// middleware
	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "${time_rfc3339} | ${remote_ip} | ${method} ${uri} - ${status} - ${latency_human}\n",
	// }))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 5}))

	// bootstrap things
	server.Bootstrap(e)

	if os.Getenv("ENVIRONMENT") == config.EnvRelease {
		e.Use(middleware.Recover())
	} else {
		// swagger
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	// start server
	go func() {
		if err := e.Start(":3000"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
