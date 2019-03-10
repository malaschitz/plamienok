package controllers

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/malaschitz/plamienok/server/db"
)

// return all visists for user
func Visits(c echo.Context) error {
	id := c.Param("id")
	fmt.Println("PERSON ID", id)
	person, err := db.PersonByID(id)
	if err == nil {
		visits := db.DtoVisits(person)
		return okApiResponse(c, visits)
	}
	return errorApiResponse(c, err)
}
