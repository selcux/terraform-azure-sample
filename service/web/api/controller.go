package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/selcux/terraform-azure-sample/pkg/client"
)

type Controller struct{}

func (c *Controller) Read(ctx echo.Context) error {
	name := ctx.Param("name")

	greetClient := client.NewGreetClient()
	err := greetClient.Connect()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	defer greetClient.Close()

	resp, err := greetClient.SayHello(name)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, resp)
}

func NewController() *Controller {
	return &Controller{}
}
