package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"simplebattle/gachapon"
	"simplebattle/webapp/services"

	"github.com/labstack/echo/v4"
)

type EchoRoute struct {
	Service services.IGachaService
	Port    string

	app *echo.Echo
}

func (route *EchoRoute) Start() {
	app := echo.New()
	route.app = app

	app.POST("/roll", generateRollFeature(route))
	app.POST("/recharger", generateRechargeFeature(route))

	listen := fmt.Sprintf(":%s", route.Port)
	app.Logger.Fatal(app.Start(listen))
}

func generateRollFeature(route *EchoRoute) echo.HandlerFunc {
	return func(context echo.Context) error {
		var request = struct {
			User  string `json:"user"`
			Times int    `json:"times"`
		}{}

		body, _ := ioutil.ReadAll(context.Request().Body)
		err := json.Unmarshal(body, &request)

		if err != nil {
			return echo.ErrBadRequest
		}

		var items []gachapon.Item
		items, err = route.Service.Roll(request.User, request.Times)

		if err != nil {
			return context.JSON(400, fmt.Sprintf("%v", err))
		}

		return context.JSON(200, items)
	}
}

func generateRechargeFeature(route *EchoRoute) echo.HandlerFunc {
	return func(context echo.Context) error {
		var request = struct {
			User  string `json:"user"`
			Point int    `json:"point"`
		}{}

		body, err := ioutil.ReadAll(context.Request().Body)
		err = json.Unmarshal(body, &request)

		if err == nil {
			points := route.Service.Recharge(request.User, request.Point)
			return context.JSON(200, fmt.Sprintf("success : current point = %d", points))
		} else {
			return context.JSON(500, err)
		}
	}
}
