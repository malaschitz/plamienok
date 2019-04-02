/*
 * Developed by Richard Malaschitz on 3/27/19 12:34 PM
 * Last modified 3/27/19 12:33 PM
 * Copyright (c) 2019. All right reserved.
 *
 */

package model

import (
	"github.com/malaschitz/plamienok/server/model/enum"
)

type PersonFilter struct {
	FullText  string
	Oddelenie string
	Stav      string
	Clients   string
}

type PersonDto struct {
	Person
	DtoBirthDate       string
	DtoDeath           string
	DtoPlamPrijatie    string
	DtoPlamPrepustenie string
	DtoRelatives       []RelativeDto
}

type RelativeDto struct {
	ID           string
	Relationship enum.Relationship
	Person       Person
}

type TextValueDto struct {
	Text  string `json:"text"`
	Value string `json:"value"`
}

type SessionDto struct {
	Session
	DtoDatum string
	UserName string
}

type CielDto struct {
	Ciel
	UserName  string
	IsDeleted bool
}

type MeninyDto struct {
	ID           string
	Person       Person
	Datum        Date
	Typ          string
	Relationship enum.Relationship
	Relative     Person
}

type VisitDto struct {
	Visit
	DtoTyp   string // H - Home P - Phone
	DtoDatum string
}

type VisitCallDto struct {
	VisitCall
	DtoDatum DateTimeDto
}

type VisitHomeDto struct {
	VisitHome
	DtoDatum, DtoVyjazdFrom, DtoVyjazdTo DateTimeDto
}

type DateTimeDto struct {
	Date string
	Time string
}

//Created by Richard Malaschitz malaschitz@gmail.com
//2019-03-04 20:06
//Copyright (c) 2018. All Rights Reserved.
