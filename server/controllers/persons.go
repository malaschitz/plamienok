package controllers

import (
	"github.com/labstack/echo"
	"github.com/malaschitz/plamienok/server/db"
	"github.com/malaschitz/plamienok/server/model"
	"github.com/malaschitz/plamienok/server/model/dto"
	"github.com/pkg/errors"
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
	var car model.Car
	err := c.Bind(&car)
	if err == nil {
		if car.Name == "" {
			err = errors.New("Nesprávny názov")
		}
		err = db.SaveCar(&car, p.User.ID)
	}
	if err != nil {
		return errorApiResponse(c, err)
	} else {
		return okApiResponse(c, car)
	}
}
