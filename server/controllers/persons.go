package controllers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/malaschitz/plamienok/server/utils"

	"github.com/labstack/echo"
	"github.com/malaschitz/plamienok/server/db"
	"github.com/malaschitz/plamienok/server/model"
	"github.com/malaschitz/plamienok/server/model/enum"
)

func Persons(c echo.Context) error {
	var filter model.PersonFilter
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
	ret := make([]model.TextValueDto, 0)
	for _, p := range persons {
		p := model.TextValueDto{Text: p.FirstName + " " + p.Surname, Value: p.ID}
		ret = append(ret, p)
	}
	if err != nil {
		return errorApiResponse(c, err)
	}
	return okApiResponse(c, ret)
}

func PersonsTasks(c echo.Context) error {
	persons, err := db.Persons()
	ret := make([]model.TextValueDto, 0)
	for _, p := range persons {
		if p.IsHC && p.PlamPrijatie != nil && p.PlamPrepustenie == nil {
			p := model.TextValueDto{Text: p.FirstName + " " + p.Surname, Value: p.ID}
			ret = append(ret, p)
		}
	}
	if err != nil {
		return errorApiResponse(c, err)
	}
	return okApiResponse(c, ret)
}

// get person
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
	var data model.PersonDto
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

func PersonRelationDelete(c echo.Context) error {
	id := c.Param("id")
	p := c.(*PlContext)
	err := db.RelationDelete(id, p.User.ID)
	if err != nil {
		return errorApiResponse(c, err)
	} else {
		return okApiResponse(c, "OK")
	}
}

func PersonRelationPost(c echo.Context) error {
	id := c.Param("id")
	p := c.(*PlContext)
	var data struct {
		Person, Relationship, Firstname, Surname, Phone string
	}
	err := c.Bind(&data)
	if err != nil {
		return errorApiResponse(c, err)
	}

	if data.Relationship == "" {
		return errorApiResponse(c, errors.New("Nie je zadaný vzťah"))
	}

	if data.Person == "" && (data.Firstname == "" || data.Surname == "") {
		return errorApiResponse(c, errors.New("Nie je zadaná osoba"))
	}

	var person model.Person
	person, err = db.PersonByID(id)
	if err != nil {
		return errorApiResponse(c, err)
	}

	var relationship enum.Relationship
	if r, ok := enum.RelationshipsMap[data.Relationship]; ok {
		relationship = r
	} else {
		return errorApiResponse(c, errors.New("Nie je zadaný správny vzťah"))
	}

	var relative model.Person
	if data.Person == "" {
		relative = model.Person{FirstName: data.Firstname, Surname: data.Surname, Phone: data.Phone, Sex: relationship.Sex}
		err = db.SavePerson(&relative, p.User.ID)
		if err != nil {
			return errorApiResponse(c, err)
		}
	} else {
		relative, err = db.PersonByID(data.Person)
		if relative.Sex != relationship.Sex {
			return errorApiResponse(c, errors.New("Asi nie je správne zadaný vzťah alebo je chybne zadané pohlavie osoby."))
		}
	}
	err = db.SaveRelation(person, relative, relationship, p.User.ID)
	if err != nil {
		return errorApiResponse(c, err)
	} else {
		return okApiResponse(c, "OK")
	}
}

func Meniny(c echo.Context) error {
	date := c.Param("date")
	z := strings.Split(date, ".")
	d, err := strconv.Atoi(z[0])
	if err != nil {
		return errorApiResponse(c, err)
	}
	m, err := strconv.Atoi(z[1])
	if err != nil {
		return errorApiResponse(c, err)
	}
	day := model.Date{Day: d, Month: m}
	// prejdi deti a najdi meniny aj u pribuznych...
	meniny := make([]model.MeninyDto, 0)
	persons, err := db.Persons()
	if err != nil {
		return errorApiResponse(c, err)
	}

	for i := 0; i < 7; i++ {

		for _, p := range persons {
			if p.Deleted == nil && p.IsPatient {
				//narodeniny
				if p.BirthDate != nil {
					if p.BirthDate.Day == day.Day && p.BirthDate.Month == day.Month {
						m := model.MeninyDto{
							ID:     utils.UUID(),
							Person: p,
							Datum:  *p.BirthDate,
							Typ:    "Narodeniny",
						}
						meniny = append(meniny, m)
					}
				}
				//meniny
				if p.FirstName != "" {
					mDate := model.Meniny(p.FirstName)
					if mDate.Day == day.Day && mDate.Month == day.Month {
						m := model.MeninyDto{
							ID:     utils.UUID(),
							Person: p,
							Datum:  mDate,
							Typ:    "Meniny",
						}
						meniny = append(meniny, m)
					}
				}
				//pribuzni
				rel1, rel2, _ := db.Relatives(p)
				ids := map[string]string{}
				for _, r := range rel1 {
					ids[r.RelativeID] = r.RelationshipString
				}
				for _, r := range rel2 {
					ids[r.PersonID] = enum.GetOpoRelation(r.RelationshipString, p.Sex).Relation
				}

				for id, relationship := range ids {
					//narodeniny
					pr, err := db.PersonByID(id)
					if pr.Deleted == nil {
						if err == nil {
							if pr.BirthDate != nil {
								if pr.BirthDate.Day == day.Day && pr.BirthDate.Month == day.Month {
									m := model.MeninyDto{
										ID:           utils.UUID(),
										Person:       p,
										Datum:        *pr.BirthDate,
										Typ:          "Narodeniny",
										Relative:     pr,
										Relationship: enum.RelationshipsMap[relationship],
									}
									meniny = append(meniny, m)
								}
							}
							//meniny
							if pr.FirstName != "" {
								mDate := model.Meniny(pr.FirstName)
								if mDate.Day == day.Day && mDate.Month == day.Month {
									m := model.MeninyDto{
										ID:           utils.UUID(),
										Person:       p,
										Datum:        mDate,
										Typ:          "Meniny",
										Relative:     pr,
										Relationship: enum.RelationshipsMap[relationship],
									}
									meniny = append(meniny, m)
								}
							}
						}
					}
				}
			}
		}
		day = day.AddDay()
	}
	return okApiResponse(c, meniny)
	//
}

func personToDto(person model.Person) model.PersonDto {
	d := model.PersonDto{Person: person}
	d.DtoBirthDate = model.Date2String(person.BirthDate)
	d.DtoDeath = model.Date2String(person.Death)
	d.DtoPlamPrijatie = model.Date2String(person.PlamPrijatie)
	d.DtoPlamPrepustenie = model.Date2String(person.PlamPrepustenie)
	//relatives
	rel1, rel2, err := db.Relatives(person)
	fmt.Println("REL1", rel1)
	fmt.Println("REL2", rel2)
	fmt.Println("ERRR", err)
	d.DtoRelatives = make([]model.RelativeDto, 0)
	for _, r := range rel1 {
		relative, err := db.PersonByID(r.RelativeID)
		if err != nil {
			continue
		}
		if relative.Deleted != nil {
			continue
		}
		dr := model.RelativeDto{
			ID:           r.ID,
			Relationship: enum.RelationshipsMap[r.RelationshipString],
			Person:       relative,
		}
		d.DtoRelatives = append(d.DtoRelatives, dr)
	}

	for _, r := range rel2 {
		relative, err := db.PersonByID(r.PersonID)
		if err != nil {
			continue
		}
		if relative.Deleted != nil {
			continue
		}
		fmt.Println("PROBLEM", r, relative)
		dr := model.RelativeDto{
			ID:           r.ID,
			Relationship: enum.GetOpoRelation(r.RelationshipString, relative.Sex),
			Person:       relative,
		}
		d.DtoRelatives = append(d.DtoRelatives, dr)
	}

	return d
}
