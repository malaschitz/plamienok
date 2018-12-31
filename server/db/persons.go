package db

import "github.com/malaschitz/plamienok/server/model"

func SavePerson(person *model.Person, authorID string) error {
	person.Basification(authorID)
	err := _db.Save(person)
	return err
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

//Created by Richard Malaschitz malaschitz@gmail.com
//2018-12-31 14:28
//Copyright (c) 2018. All Rights Reserved.
