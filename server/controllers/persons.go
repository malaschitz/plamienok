package controllers

import (
	"fmt"

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

func PersonsAll(c echo.Context) error {
	persons, err := db.Persons()
	ret := make([]dto.TextValueDto, 0)
	for _, p := range persons {
		p := dto.TextValueDto{Text: p.FirstName + " " + p.Surname, Value: p.ID}
		ret = append(ret, p)
	}
	if err != nil {
		return errorApiResponse(c, err)
	}
	return okApiResponse(c, ret)
}

func PersonsTasks(c echo.Context) error {
	persons, err := db.Persons()
	ret := make([]dto.TextValueDto, 0)
	for _, p := range persons {
		if p.IsHC && p.PlamPrijatie != nil && p.PlamPrepustenie == nil {
			p := dto.TextValueDto{Text: p.FirstName + " " + p.Surname, Value: p.ID}
			ret = append(ret, p)
		}
	}
	if err != nil {
		return errorApiResponse(c, err)
	}
	return okApiResponse(c, ret)
}

func Person(c echo.Context) error {
	id := c.Param("id")
	person, err := db.PersonByID(id)
	if err != nil {
		return errorApiResponse(c, err)
	} else {
		return okApiResponse(c, personToDto(person))
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

func PersonPut(c echo.Context) error {
	p := c.(*PlContext)
	var data dto.PersonDto
	err := c.Bind(&data)
	if err == nil {
		person := data.Person
		person.BirthDate = model.String2Date(data.DtoBirthDate)
		person.PlamPrepustenie = model.String2Date(data.DtoPlamPrepustenie)
		person.PlamPrijatie = model.String2Date(data.DtoPlamPrijatie)
		person.Death = model.String2Date(data.DtoDeath)
		err = db.SavePerson(&person, p.User.ID)
		if err == nil {
			return okApiResponse(c, personToDto(person))
		}
	}
	return errorApiResponse(c, err)
}

func personToDto(person model.Person) dto.PersonDto {
	d := dto.PersonDto{Person: person}
	d.DtoBirthDate = model.Date2String(person.BirthDate)
	d.DtoMeniny = model.Date2String(person.Meniny)
	d.DtoDeath = model.Date2String(person.Death)
	d.DtoPlamPrijatie = model.Date2String(person.PlamPrijatie)
	d.DtoPlamPrepustenie = model.Date2String(person.PlamPrepustenie)
	//relatives
	relatives, err := db.Relatives(person)
	fmt.Println(relatives)
	fmt.Println(err)
	d.DtoRelatives = make([]dto.RelativeDto, 0)
	for _, r := range relatives {
		relative, err := db.PersonByID(r.RelativeID)
		if err != nil {
			continue
		}
		dr := dto.RelativeDto{
			ID:           r.ID,
			Relationship: r.Relationship,
			Person:       relative,
		}
		d.DtoRelatives = append(d.DtoRelatives, dr)
	}

	return d
}
