package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/malaschitz/plamienok/server/constants"
	"github.com/malaschitz/plamienok/server/controllers"
	"github.com/malaschitz/plamienok/server/db"
	"github.com/malaschitz/plamienok/server/routers"
	"github.com/malaschitz/plamienok/server/service"
)

func main() {
	constants.InitConst()
	service.InitEmail()

	log.Info("start server " + constants.AppName)
	db.InitDB()

	e := echo.New()

	e.Debug = true
	e.Use(handleUser)
	routers.InitAccesible(e)

	r := e.Group("/r")
	r.Use(restrictUser)
	routers.InitRestricted(r)

	e.Logger.Fatal(e.Start(":" + constants.ServerPort))
}

func handleUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		vc := &controllers.PlContext{Context: c}
		var token string
		z := c.Request().Header["Token"]
		if len(z) > 0 {
			token = z[0]
		}
		if token != "" {
			user, err := db.ValidToken(token)
			if err == nil {
				vc.User = &user
			}
		}
		fmt.Println("handle user", c.Request().URL, token)
		return next(vc)
	}
}

func restrictUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("restrict user", c.Request().URL)
		p := c.(*controllers.PlContext)
		if p.User == nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Neprihlásený užívateľ")
		}
		return next(c)
	}
}
