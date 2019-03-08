package controllers

import (
	"github.com/labstack/echo"
	"github.com/malaschitz/plamienok/server/db"
)

func Sessions(c echo.Context) error {
	cars, err := db.Cars(true)
	if err != nil {
		return errorApiResponse(c, err)
	} else {
		return okApiResponse(c, cars)
	}
}
