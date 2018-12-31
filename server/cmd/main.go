package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/gommon/log"

	"github.com/labstack/echo"
	"github.com/malaschitz/plamienok/server"
	"github.com/malaschitz/plamienok/server/controllers"
	"github.com/malaschitz/plamienok/server/db"
	"github.com/malaschitz/plamienok/server/routers"
)

func main() {
	server.InitConst()
	log.Info("start server " + server.AppName)
	db.InitDB()

	e := echo.New()
	e.Debug = true
	e.Use(handleUser)
	routers.InitAccesible(e)

	r := e.Group("/r")
	r.Use(restrictUser)
	routers.InitRestricted(r)

	e.Logger.Fatal(e.Start(":" + server.ServerPort))
}

func handleUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("handle user", c.Request().URL)
		vc := &controllers.PlContext{Context: c}
		var token string
		z := c.Request().Header["Token"]
		if len(z) > 0 {
			token = z[0]
		}
		if token != "" {
			user, err := db.ValidToken(token)
			if err != nil {
				return err
			}
			vc.User = &user
		}
		return next(vc)
	}
}

func restrictUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		p := c.(*controllers.PlContext)
		if p.User == nil {
			return echo.NewHTTPError(http.StatusForbidden, "Neprihlásený užívateľ")
		}
		return next(c)
	}
}
