package controllers

import (
	"github.com/labstack/echo"
	"github.com/malaschitz/plamienok/server/db"
	"github.com/malaschitz/plamienok/server/model"
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

// new person
func PersonPost(c echo.Context) error {
	p := c.(*PlContext)
	var data struct {
		FirstName, Surname, BirthDate string
	}
	err := c.Bind(&data)
	var person model.Person
	if err == nil {
		person = model.Person{FirstName: data.FirstName, Surname: data.Surname, BirthDate: model.String2Date(data.BirthDate), IsPatient: true}
		err = db.SavePerson(&person, p.User.ID)
	}
	if err != nil {
		return errorApiResponse(c, err)
	} else {
		return okApiResponse(c, person)
	}
}
