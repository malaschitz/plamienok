package controllers

import (
	"time"

	"github.com/labstack/echo"
	"github.com/malaschitz/plamienok/server/db"
	"github.com/malaschitz/plamienok/server/model"
)

// return all tasks for user
func Ciele(c echo.Context) error {
	id := c.Param("id")
	person, err := db.PersonByID(id)
	if err == nil {
		tasks := db.DtoCiele(person)
		return okApiResponse(c, tasks)
	}
	return errorApiResponse(c, err)
}

func CielPost(c echo.Context) error {
	p := c.(*PlContext)
	var data model.CielDto
	err := c.Bind(&data)
	if err == nil {
		ciel := data.Ciel
		if data.IsDeleted && ciel.Deleted == nil {
			t := time.Now()
			ciel.Deleted = &t
		}
		err = db.SaveCiel(&ciel, p.User.ID)
		if err == nil {
			return okApiResponse(c, db.CielToDto(ciel))
		}
	}
	return errorApiResponse(c, err)
}
