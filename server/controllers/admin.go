package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/malaschitz/plamienok/server/constants"
	"github.com/malaschitz/plamienok/server/db"
	"github.com/malaschitz/plamienok/server/model"
	"github.com/malaschitz/plamienok/server/service"
	"github.com/malaschitz/plamienok/server/utils"
)

type PlContext struct {
	echo.Context
	User *model.User
}

func Status(c echo.Context) error {
	fmt.Println("Status is ok")
	data := map[string]interface{}{"status": "OK"}
	return okApiResponse(c, data)
}

func Info(c echo.Context) error {
	v := c.(*PlContext)
	info := struct {
		App     string
		Version string
		User    *model.User
	}{
		App:     constants.AppName,
		Version: constants.Version,
		User:    v.User,
	}

	return okApiResponse(c, info)
}

func Login(c echo.Context) error {
	fmt.Println("LOGIN")
	var data struct {
		Email    string
		Password string
	}
	err := c.Bind(&data)
	if err != nil {
		return errorApiResponse(c, err)
	}

	user, token, err := db.Login(data.Email, data.Password)
	if err != nil {
		return errorApiResponse(c, err)
	} else {
		info := struct {
			User  model.User
			Token string
		}{
			User:  user,
			Token: token,
		}
		return okApiResponse(c, info)
	}
}

func ForgottenPassword(c echo.Context) error {
	var data struct {
		Email string
	}
	err := c.Bind(&data)
	if err != nil {
		return errorApiResponse(c, err)
	}
	user, err := db.UserByEmail(data.Email)
	if err != nil {
		return errorApiResponse(c, constants.ERR_NO_USER)
	} else {
		code := model.Code6{Code6: utils.RandomCode(), Validity: time.Now().Add(constants.TokenDuration)}
		db.SetCode6(user.ID, code)
		go func() {
			if constants.ServerDebug {
				log.Printf("%s %s %s", user.Email, user.Name, code.Code6)
			}
			err := service.SendCode(user.Email, user.Name, code.Code6)
			if err != nil {
				log.Printf("Chyba pri posielaní emailu %s", err.Error())
			}
		}()
		return okApiResponse(c, "OK")
	}
}

func CheckCode6(c echo.Context) error {
	var data struct {
		Email    string
		Code6    string
		Password string
	}
	err := c.Bind(&data)
	if err != nil {
		return errorApiResponse(c, err)
	}
	user, err := db.UserByEmail(data.Email)
	if err != nil {
		return errorApiResponse(c, constants.ERR_NO_USER)
	} else {
		code := db.GetCode6(user.ID)
		if time.Now().After(code.Validity) || code.Code6 != data.Code6 {
			return errorApiResponse(c, constants.ERR_WRONG_CODE)
		}
		if data.Password != "" {
			db.SetNewPassword(user, data.Password)
		}
		info := struct {
			User  model.User
			Token string
		}{
			User:  user,
			Token: db.CreateToken(user.ID),
		}
		return okApiResponse(c, info)
	}
}

func errorApiResponse(c echo.Context, err error) error {
	return c.JSON(http.StatusAccepted, struct{ Error string }{Error: err.Error()})
}

func okApiResponse(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, data)
}
