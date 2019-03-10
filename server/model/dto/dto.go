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
}

type TextValueDto struct {
	Text  string `json:"text"`
	Value string `json:"value"`
}

type VisitDto struct {
	model.Visit
	Typ string // H - Home P - Phone
}

//Created by Richard Malaschitz malaschitz@gmail.com
//2019-03-04 20:06
//Copyright (c) 2018. All Rights Reserved.
