package controllers

import (
	"github.com/labstack/echo"
	"github.com/malaschitz/plamienok/server/db"
	"github.com/malaschitz/plamienok/server/model/dto"
)

func Persons(c echo.Context) error {
	var filter dto.PersonFilter
	err := c.Bind(&filter)
	if err != nil {
		return errorApiResponse(c, err)
	}
	persons, err := db.PersonsFiltered(filter)
	if err != nil {
		return errorApiResponse(c, err)
	} else {
		return okApiResponse(c, persons)
	}
}

func Person(c echo.Context) error {
	id := c.Param("id")
	person, err := db.PersonByID(id)
	if err != nil {
		return errorApiResponse(c, err)
	} else {
		return okApiResponse(c, person)
	}
}
