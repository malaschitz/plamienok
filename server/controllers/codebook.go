package controllers

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/malaschitz/plamienok/server/db"
)

func Meds(c echo.Context) error {
	filterBy := c.QueryParam("filter")
	meds := db.GetLieky(filterBy, 20)
	return okApiResponse(c, meds)
}

func Diagnoses(c echo.Context) error {
	filterBy := c.QueryParam("filter")
	maxS := c.QueryParam("max")
	max, err := strconv.Atoi(maxS)
	if err != nil || max == 0 {
		max = 20
	}
	dgs := db.GetDiagnozy(filterBy, max)
	return okApiResponse(c, dgs)
}
