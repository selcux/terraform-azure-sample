package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct{}

func (c *Controller) Read(ctx echo.Context) error {
	name := ctx.Param("name")

	return ctx.JSON(http.StatusOK, name)
}

func NewController() *Controller {
	return &Controller{}
}
