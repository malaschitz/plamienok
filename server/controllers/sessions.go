package controllers

import (
	"github.com/labstack/echo"
	"github.com/malaschitz/plamienok/server/db"
	"github.com/malaschitz/plamienok/server/model"
	"github.com/malaschitz/plamienok/server/model/dto"
)

func Sessions(c echo.Context) error {
	sessions, err := db.Sessions()
	ret := make([]dto.SessionDto, 0)
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
	var data dto.SessionDto
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

func sessionToDto(session model.Session) dto.SessionDto {
	s := dto.SessionDto{Session: session}
	s.DtoDatum = model.DateTime2String(&session.Datum)
	u, _ := db.UserByID(session.UserID)
	s.UserName = u.Name
	return s
}
