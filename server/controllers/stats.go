package controllers

import (
	"github.com/labstack/echo"
	"github.com/malaschitz/plamienok/server/db"
)

func Stats(c echo.Context) error {
	s := []stats{
		stats{"Počet užívateľov", db.UserCount()},
		stats{"Počet návštev", db.VisitCount()},
		stats{"Počet áut", db.CarCount()},
		stats{"Počet osôb (pacienti, deti, príbuzní)", db.PersonsCount()},
	}
	return okApiResponse(c, s)
}

type stats struct {
	Text  string
	Value int
}
