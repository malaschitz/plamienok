package controllers

import (
	"github.com/labstack/echo"
	"github.com/malaschitz/plamienok/server/db"
	"github.com/malaschitz/plamienok/server/model"
)

// return all visists for user
func Visits(c echo.Context) error {
	id := c.Param("id")
	person, err := db.PersonByID(id)
	if err == nil {
		visits := db.DtoVisits(person)
		return okApiResponse(c, visits)
	}
	return errorApiResponse(c, err)
}

func VisitCall(c echo.Context) error {
	id := c.Param("id")
	call, err := db.VisitCallByID(id)
	if err != nil {
		return errorApiResponse(c, err)
	} else {
		return okApiResponse(c, callToDto(call))
	}
}

func VisitCallPost(c echo.Context) error {
	p := c.(*PlContext)
	var call model.VisitCallDto
	err := c.Bind(&call)

	if err != nil {
		return errorApiResponse(c, err)
	}
	call.VisitCall.Datum = *model.DTO2DateTime(call.DtoDatum)
	err = db.SaveVisitCall(&call.VisitCall, p.User.ID)
	if err != nil {
		return errorApiResponse(c, err)
	} else {
		return okApiResponse(c, callToDto(call.VisitCall))
	}
}

func VisitCalls(c echo.Context) error {
	visits, err := db.VisitCalls()
	if err != nil {
		return errorApiResponse(c, err)
	} else {
		return okApiResponse(c, visits)
	}
}

func VisitHomes(c echo.Context) error {
	visits, err := db.VisitHomes()
	if err != nil {
		return errorApiResponse(c, err)
	} else {
		return okApiResponse(c, visits)
	}
}

func VisitHome(c echo.Context) error {
	id := c.Param("id")
	call, err := db.VisitHomeByID(id)
	if err != nil {
		return errorApiResponse(c, err)
	} else {
		return okApiResponse(c, homeToDto(call))
	}
}

func VisitHomePost(c echo.Context) error {
	p := c.(*PlContext)
	var visit model.VisitHomeDto
	err := c.Bind(&visit)
	if err != nil {
		return errorApiResponse(c, err)
	}

	visit.VisitHome.Datum = *model.DTO2DateTime(visit.DtoDatum)
	visit.VisitHome.VyjazdFrom = *model.DTO2DateTime(visit.DtoVyjazdFrom)
	visit.VisitHome.VyjazdTo = *model.DTO2DateTime(visit.DtoVyjazdTo)

	err = db.SaveVisitHome(&visit.VisitHome, p.User.ID)
	if err != nil {
		return errorApiResponse(c, err)
	} else {
		return okApiResponse(c, homeToDto(visit.VisitHome))
	}
}

func callToDto(call model.VisitCall) model.VisitCallDto {
	d := model.VisitCallDto{VisitCall: call}
	d.DtoDatum = model.DateTime2DTO(&call.Datum)
	return d
}

func homeToDto(home model.VisitHome) model.VisitHomeDto {
	d := model.VisitHomeDto{VisitHome: home}
	d.DtoDatum = model.DateTime2DTO(&home.Datum)
	d.DtoVyjazdFrom = model.DateTime2DTO(&home.VyjazdFrom)
	d.DtoVyjazdTo = model.DateTime2DTO(&home.VyjazdTo)
	return d
}
