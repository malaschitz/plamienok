package db

import (
	"sort"
	"strings"
	"time"

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
	persons = make([]model.Person, 0)
	var z []model.Person
	z, err = Persons()
	if err != nil {
		return
	}
	//
	ftFilter := utils.FullTextCreate(filter.FullText)
	for _, p := range z {
		test := true
		//fulltext test
		test = utils.FullTextTest(ftFilter, p.FullText)
		//clients
		if (filter.Clients == "d" && !p.IsPatient) || (filter.Clients == "p" && p.IsPatient) {
			test = false
		}
		//oddelenie
		if (filter.Oddelenie == "hc" && !p.IsHC) || (filter.Oddelenie == "cgt" && !p.IsCGT) {
			test = false
		}
		//stav
		if filter.Stav == "p" && !(p.PlamPrijatie != nil && p.PlamPrepustenie == nil) {
			test = false
		}
		if filter.Stav == "o" && p.PlamPrepustenie == nil {
			test = false
		}
		if filter.Stav == "z" && p.Death == nil {
			test = false
		}

		//append ?
		if test {
			persons = append(persons, p)
		}
	}

	//sort
	sort.Slice(persons, func(i, j int) bool {
		return persons[i].Updated.After(persons[j].Updated)
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

func Relatives(person model.Person) ([]model.PersonRelation, error) {
	var relations []model.PersonRelation
	err := _db.Find("PersonID", person.ID, &relations)
	return relations, err
}

func RelationByID(id string) (relation model.PersonRelation, err error) {
	err = _db.One("ID", id, &relation)
	return
}

func RelationDelete(id string, authorID string) error {
	relation, err := RelationByID(id)
	if err == nil {
		relation.Basification(authorID)
		t := time.Now()
		relation.Deleted = &t
		err = _db.Save(&relation)
	}
	return err
}

//Created by Richard Malaschitz
//2018-12-31 14:28
//Copyright (c) 2018. All Rights Reserved.
