package controllers

import (
	"github.com/labstack/echo"
	"github.com/malaschitz/plamienok/server/db"
)

func Stats(c echo.Context) error {
	stats := struct {
		Users  int
		Visits int
	}{
		Users:  db.UserCount(),
		Visits: db.VisitCount(),
	}
	return okApiResponse(c, stats)
}
