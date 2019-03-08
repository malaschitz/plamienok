package controllers

import (
	"github.com/labstack/echo"
	"github.com/malaschitz/plamienok/server/db"
)

func Diagnozy(c echo.Context) error {
	filter := c.QueryParam("filter")
	max := 10
	dgns := db.GetDiagnozy(filter, max)
	return okApiResponse(c, dgns)

}

//Created by Richard Malaschitz malaschitz@gmail.com
//2019-03-08 22:39
//Copyright (c) 2018. All Rights Reserved.
