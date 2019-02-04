package controllers

import (
	"github.com/labstack/echo"
	"github.com/malaschitz/plamienok/server/db"
)

func Stats(c echo.Context) error {
	s := []stats{
		stats{"Počet užívateľov", db.UserCount()},
		stats{"Počet návštev", db.VisitCount()},
	}
	return okApiResponse(c, s)
}

type stats struct {
	Text  string
	Value int
}
