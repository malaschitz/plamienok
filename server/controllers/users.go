package controllers

import (
	"time"

	"github.com/labstack/echo"
	"github.com/malaschitz/plamienok/server/db"
	"github.com/malaschitz/plamienok/server/model"
	"github.com/malaschitz/plamienok/server/utils"
	"github.com/pkg/errors"
)

func Users(c echo.Context) error {
	users, err := db.Users(true)
	if err != nil {
		return errorApiResponse(c, err)
	} else {
		return okApiResponse(c, users)
	}
}

func UserPost(c echo.Context) error {
	p := c.(*PlContext)
	var user model.User
	err := c.Bind(&user)
	if err == nil {
		if user.Name == "" {
			err = errors.New("Nesprávne meno")
		} else if !utils.IsEmail(user.Email) {
			err = errors.New("Nesprávny email")
		}
		err = db.SaveUser(&user, p.User.ID)
	}
	if err != nil {
		return errorApiResponse(c, err)
	} else {
		return okApiResponse(c, user)
	}
}

func UserDelete(c echo.Context) error {
	p := c.(*PlContext)
	var id string
	id = c.Param("id")
	user, err := db.UserByID(id)
	if err == nil {
		if user.Deleted == nil {
			t := time.Now()
			user.Deleted = &t
		} else {
			user.Deleted = nil
		}
		err = db.SaveUser(&user, p.User.ID)
	}
	if err != nil {
		return errorApiResponse(c, err)
	} else {
		return okApiResponse(c, user)
	}
}
