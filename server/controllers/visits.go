package controllers

import (
	"github.com/labstack/echo"
	"github.com/malaschitz/plamienok/server/db"
)

// return all visists for user
func Visits(c echo.Context) error {
	id := c.Param("id")
	person, err := db.PersonByID(id)
	if err == nil {
		visits, err := db.DtoVisists(person)
		if err == nil {
			return okApiResponse(c, visits)
		}
	}
	return errorApiResponse(c, err)
}
