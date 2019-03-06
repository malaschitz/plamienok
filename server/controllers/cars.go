/*
 * Developed by Richard Malaschitz on 3/4/19 1:19 PM
 * Last modified 3/4/19 9:49 AM
 * Copyright (c) 2019. All right reserved.
 *
 */

package controllers

import (
	"time"

	"github.com/labstack/echo"
	"github.com/malaschitz/plamienok/server/db"
	"github.com/malaschitz/plamienok/server/model"
	"github.com/pkg/errors"
)

func Cars(c echo.Context) error {
	cars, err := db.Cars(true)
	if err != nil {
		return errorApiResponse(c, err)
	} else {
		return okApiResponse(c, cars)
	}
}

func CarPost(c echo.Context) error {
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

func CarDelete(c echo.Context) error {
	p := c.(*PlContext)
	var id string
	id = c.Param("id")
	car, err := db.CarByID(id)
	if err == nil {
		if car.Deleted == nil {
			t := time.Now()
			car.Deleted = &t
		} else {
			car.Deleted = nil
		}
		err = db.SaveCar(&car, p.User.ID)
	}
	if err != nil {
		return errorApiResponse(c, err)
	} else {
		return okApiResponse(c, car)
	}
}
