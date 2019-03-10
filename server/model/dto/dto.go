package dto

import "github.com/malaschitz/plamienok/server/model"

type PersonFilter struct {
	FullText  string
	Oddelenie string
	Stav      string
	Clients   string
}

type PersonDto struct {
	model.Person
	DtoBirthDate       string
	DtoMeniny          string
	DtoDeath           string
	DtoPlamPrijatie    string
	DtoPlamPrepustenie string
	DtoRelatives       []RelativeDto
}

type RelativeDto struct {
	ID           string
	Relationship model.Relationship
	Person       model.Person
}

type TextValueDto struct {
	Text  string `json:"text"`
	Value string `json:"value"`
}

type SessionDto struct {
	model.Session
	DtoDatum string
	UserName string
}

type VisitDto struct {
	model.Visit
	DtoTyp   string // H - Home P - Phone
	DtoDatum string
}

type CielDto struct {
	model.Ciel
	UserName  string
	IsDeleted bool
}

//Created by Richard Malaschitz malaschitz@gmail.com
//2019-03-04 20:06
//Copyright (c) 2018. All Rights Reserved.
