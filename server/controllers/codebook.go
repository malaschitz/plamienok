/*
 * Developed by Richard Malaschitz on 3/4/19 1:19 PM
 * Last modified 3/4/19 9:49 AM
 * Copyright (c) 2019. All right reserved.
 *
 */

package controllers

import (
	"github.com/labstack/echo"
	"github.com/malaschitz/plamienok/server/db"
)

func Meds(c echo.Context) error {
	filterBy := c.QueryParam("filter")
	meds := db.GetLieky(filterBy)
	return okApiResponse(c, meds)
}

func Diagnoses(c echo.Context) error {
	filterBy := c.QueryParam("filter")
	dgs := db.GetDiagnozy(filterBy)
	return okApiResponse(c, dgs)
}
