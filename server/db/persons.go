package db

import (
	"sort"
	"strings"

	"github.com/malaschitz/plamienok/server/model"
	"github.com/malaschitz/plamienok/server/model/dto"
	"github.com/malaschitz/plamienok/server/utils"
)

func Persons() (persons []model.Person, err error) {
	err = _db.All(&persons)
	return
}

func PersonByID(id string) (person model.Person, err error) {
	err = _db.One("ID", id, &person)
	return
}

func PersonsFiltered(filter dto.PersonFilter) (persons []model.Person, err error) {
	utils.JsonPrint(filter)

	var z []model.Person
	z, err = Persons()
	if err != nil {
		return
	}
	for _, p := range z {
		test := true
		//fulltext test
		if filter.FullText != "" {
			test = false
			fullTexts := utils.RemoveDiacritics(filter.FullText)
			z := strings.Split(fullTexts, " ")
			for _, f := range z {
				if strings.Contains(p.FullText, f) {
					test = true
					break
				}
			}
		}
		//clients
		if (filter.Clients == "d" && !p.IsPatient) || (filter.Clients == "p" && p.IsPatient) {
			test = false
		}

		//append ?
		if test {
			persons = append(persons, p)
		}
	}

	//sort
	sort.Slice(persons, func(i, j int) bool {
		return persons[j].Updated.After(persons[i].Updated)
	})

	return
}

func SavePerson(person *model.Person, authorID string) error {
	person.Basification(authorID)
	fullText(person)
	err := _db.Save(person)
	return err
}

func fullText(person *model.Person) {
	data := []string{person.FirstName, person.Surname, person.Phone, person.Email, person.RC}
	fulltext := strings.Join(data, " ")
	fulltext = utils.RemoveDiacritics(fulltext)
	person.FullText = fulltext
}

func SaveRelation(person, relative model.Person, relationship model.Relationship, authorID string) error {
	pr := model.PersonRelation{
		PersonID:     person.ID,
		RelativeID:   relative.ID,
		Relationship: relationship,
	}
	pr.Basification(authorID)
	return _db.Save(&pr)
}

func PersonsCount() int {
	count, err := _db.Count(&model.Person{})
	utils.LogErr(err)
	return count
}

//Created by Richard Malaschitz
//2018-12-31 14:28
//Copyright (c) 2018. All Rights Reserved.
