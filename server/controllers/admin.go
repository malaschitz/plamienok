package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/malaschitz/plamienok/server"
	"github.com/malaschitz/plamienok/server/model"
)

func Status(c echo.Context) error {
	fmt.Println("Status is ok")
	data := map[string]interface{}{"status": "OK"}
	//return errorApiResponse(c, errors.New("Problem"))
	return okApiResponse(c, data)
}

func Info(c echo.Context) error {
	v := c.(*PlContext)
	info := struct {
		App     string
		Version string
		User    *model.User
	}{
		App:     server.AppName,
		Version: server.Version,
		User:    v.User,
	}

	return okApiResponse(c, info)
}

func Login(c echo.Context) error {
	var data struct {
		Email    string
		Password string
	}
	err := c.Bind(&data)
	if err != nil {
		return errorApiResponse(c, err)
	}
	return okApiResponse(c, "OK")
}

func errorApiResponse(c echo.Context, err error) error {
	log.Println("ERROR", err)
	return c.JSON(http.StatusBadRequest, struct{ Error string }{Error: err.Error()})
}

func okApiResponse(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, data)
}

type PlContext struct {
	echo.Context
	User *model.User
}
