package model

import (
	"time"
)

type User struct {
	Base  `storm:"inline"`
	Name  string
	Email string `storm:"unique"`

	RoleAdmin  bool // administrator
	RoleDoctor bool // doctor
	RoleNurse  bool // nurse
	RoleSoc    bool // social worker
	RolePsych  bool // psychologist
}

type Person struct {
	Base            `storm:"inline"`
	IsHC            bool // could be patient or key person, could be relative
	IsCGT           bool // client of CGT
	FirstName       string
	Surname         string
	BirthDate       Date
	Sex             Sex
	BornPlace       string
	Meniny          *Date
	Death           *Date
	DeathPlace      string
	PlamPrijatie    *Date
	PlamPrepustenie *Date  // nie je uz v Plamienku
	RC              string // Rodne Cislo
	DGNkod          string
	DGNpopis        string
	AddrStreet      string
	AddrCity        string
	AddrPSC         string
	Phone           string
	Email           string
	Job             string // Povolanie kontaktu
	Odoslal         string
	LekarPKontaktu  string
	Laboratoria     string
	Alergia         string
	ZP              string // Zdravotna poistovna
	DennyRezim      string

	//index
	FullText string
}

type PersonRelation struct {
	Base         `storm:"inline"`
	PersonID     string `storm:"index"`
	RelativeID   string
	Relationship Relationship
}

type Car struct {
	Base  `storm:"inline"`
	Name  string
	Popis string
}

type Visit struct {
	Base           `storm:"inline"`
	VisitFrom      *time.Time
	VisitTo        *time.Time
	Popis          string
	IsZdravotna    bool
	IsSprevadzanie bool
	IsSocial       bool
	IsPoUmrti      bool
	Users          []string //id uzivatel
	Persons        []string
}

type VisitHome struct {
	Visit      `storm:"inline"`
	VyjazdFrom *time.Time
	VyjazdTo   *time.Time
	IsPlanned  bool // planovana navsteva
	CarID      string

	Vysetrenie          string
	Ciele               string
	MaterialnaPomoc     string
	SocialnePoradenstvo string
}

type VisitPhone struct {
	Visit  `storm:"inline"`
	UserID string
	Tema   string
	Smer   bool // true ak volal plamienok
}

type LabVysledky struct {
	Base     `storm:"inline"`
	PersonID string
	Popis    string
}

type Medikacie struct {
	Base     `storm:"inline"`
	PersonID string
	SUKL     string
	Popis    string
}

type Ciele struct {
	Base     `storm:"inline"`
	PersonID string
	Popis    string
}

type Work struct { // vykaz prace
	Base    `storm:"inline"`
	Date    Date
	Minutes int
	Popis   string
}

type VisitCGT struct {
	Base `storm:"inline"`
}

//Created by Richard Malaschitz
//2018-12-27 14:55
//Copyright (c) 2018. All Rights Reserved.
