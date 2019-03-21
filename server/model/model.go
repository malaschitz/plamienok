package model

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
	IsPatient       bool // je patient inak je iba pribuzny alebo kontakt
	FirstName       string
	Surname         string
	BirthDate       *Date
	Sex             Sex
	BornPlace       string
	Death           *Date
	DeathPlace      string
	PlamPrijatie    *Date
	PlamPrepustenie *Date    // nie je uz v Plamienku
	RC              string   // Rodne Cislo
	DGN             []string //kody diagnoz
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

	//index
	FullText string
}

type Ciel struct {
	Base     `storm:"inline"`
	PersonID string `storm:"index"`
	Ciel     string
}

type Medikacia struct {
	Base      `storm:"inline"`
	PersonID  string `storm:"index"`
	Medikacia string
}

type PersonRelation struct {
	Base               `storm:"inline"`
	PersonID           string `storm:"index"`
	RelativeID         string `storm:"index"`
	RelationshipString string
}

type Car struct {
	Base  `storm:"inline"`
	Name  string
	Popis string
}

// stretnutie psych
type Session struct {
	Base        `storm:"inline"`
	UserID      string   `storm:"index"`
	Persons     []string //zoznam PersonID
	Datum       DateTime
	Duration    int //minutes
	Description string
}

type Visit struct {
	Base           `storm:"inline"`
	Datum          DateTime
	Duration       int //minutes
	Popis          string
	PopisDetail    string
	IsZdravotna    bool
	IsSprevadzanie bool
	IsSocial       bool
	IsPoUmrti      bool
	Users          []string //id uzivatel
	PersonID       string   //hlavna osoba
	Persons        []string //dalsi ucastnici
}

type VisitHome struct {
	Visit               `storm:"inline"`
	VyjazdFrom          DateTime
	VyjazdTo            DateTime
	IsPlanned           bool   // planovana navsteva
	CarID               string `storm:"index"`
	Vysetrenie          string
	MaterialnaPomoc     string
	SocialnePoradenstvo string
}

type VisitPhone struct {
	Visit `storm:"inline"`
	Smer  bool // true ak volal plamienok
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
	DatumOd  Date
}

type Work struct { // vykaz prace
	Base    `storm:"inline"`
	UserID  string
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
