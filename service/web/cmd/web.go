package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/selcux/terraform-azure-sample/service/web/api"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	e.Validator = &CustomValidator{validator: validator.New()}
	registerRoutes(e)
	run(e, "", 9000)
}

func run(e *echo.Echo, host string, port int) {
	// Start server
	go func() {
		if err := e.Start(fmt.Sprintf("%s:%d", host, port)); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func registerRoutes(e *echo.Echo) {
	controller := api.NewController()
	e.GET("/:name", controller.Read)
}
