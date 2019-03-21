package controllers

import (
	"strconv"

	"github.com/malaschitz/plamienok/server/model"

	"github.com/labstack/echo"
	"github.com/malaschitz/plamienok/server/db"
)

func Meds(c echo.Context) error {
	filterBy := c.QueryParam("filter")
	meds := db.GetLieky(filterBy, 200)
	return okApiResponse(c, meds)
}

func Diagnoses(c echo.Context) error {
	filterBy := c.QueryParam("filter")
	maxS := c.QueryParam("max")
	max, err := strconv.Atoi(maxS)
	if err != nil || max == 0 {
		max = 200
	}
	dgs := db.GetDiagnozy(filterBy, max)
	return okApiResponse(c, dgs)
}

func DiagnosesAll(c echo.Context) error {
	dgs := db.GetDiagnozyAll()

	return okApiResponse(c, dgs)
}

func Relationships(c echo.Context) error {
	return okApiResponse(c, model.RelationshipsNames)
}
