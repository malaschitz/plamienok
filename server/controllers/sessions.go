package controllers

import (
	"github.com/labstack/echo"
	"github.com/malaschitz/plamienok/server/db"
	"github.com/malaschitz/plamienok/server/model"
)

func Sessions(c echo.Context) error {
	sessions, err := db.Sessions()
	ret := make([]model.SessionDto, 0)
	for _, s := range sessions {
		s := sessionToDto(s)
		ret = append(ret, s)
	}
	if err != nil {
		return errorApiResponse(c, err)
	}
	return okApiResponse(c, ret)
}

func SessionPost(c echo.Context) error {
	p := c.(*PlContext)
	var data model.SessionDto
	err := c.Bind(&data)
	if err == nil {
		session := data.Session
		session.Datum = *model.String2DateTime(data.DtoDatum)
		err = db.SaveSession(&session, p.User.ID)
		if err == nil {
			return okApiResponse(c, sessionToDto(session))
		}
	}
	return errorApiResponse(c, err)
}

func sessionToDto(session model.Session) model.SessionDto {
	s := model.SessionDto{Session: session}
	s.DtoDatum = model.DateTime2String(&session.Datum)
	u, _ := db.UserByID(session.UserID)
	s.UserName = u.Name
	return s
}
