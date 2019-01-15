package model

import "time"

type Codebook struct {
	ID                  string
	Name                string
	Version             int
	DatumPoslednejZmeny time.Time
	Oid                 string
	MD5                 string
	PlatnostOd          time.Time
	PolozkyCiselnika    map[string]CiselnikPolozka
	Platnost            map[int]CiselnikPlatnost
	Atributy            map[int]string
}

type CiselnikPolozka struct {
	ID           string
	DisplayName  string
	DisplayNames map[int]string
	Skratka      string
	Popis        string
	Atributy     map[int]string
}

type Diagnoza struct {
	CiselnikPolozka
	Podskupina string
	Skupina    string
	Kapitola   string
}

type Liek struct {
	Version                          int
	ID                               string
	Atc                              string
	AtcText                          string
	AntibiotikaChemoterapeutika      int
	CestaPodania                     string
	DlzkaPlatnostiReceptu            float64
	DoplnokNazvu                     string
	GenerickaPreskripcia             int
	ValidTime                        time.Time
	IndikacnaSkupina                 string
	Liecivo                          string
	LiecivoText                      string
	LiecivoMnozstvo                  string
	LiekovaForma                     string
	Nazov                            string
	OPL                              string
	OckovaciaLatka                   int
	OriginalitaLieku                 string
	PovolenieViacnasobnejPreskripcie string
	DatumPredlzenia                  time.Time
	DatumPrvejRegistracie            time.Time
	KodSukl                          string
	PlatnostDatumOd                  time.Time
	PlatnostDatumDo                  *time.Time
	RegistracneCislo                 string
	StavRegistracie                  string
	VyrobcaLieku                     string
	VelkostBalenia                   string
	ViazanostNaLP                    string
	KategorizacneUdajeLieku          []*LiekKategorizacia
	Links                            []*LiekLink
	IsSPC                            bool   //informacie ci existuje na serveri pribalovy letak SPC
	IsPIL                            bool   //informacie ci existuje na serveri pribalovy letak PIL
	LinkADC                          string //link na ADC
	Search                           string `storm:"index"`
}

type LiekKategorizacia struct {
	CenaLekaren              float64
	DoplatokPoistencaPerc    float64
	IndikacneObmedzeniaLieku string
	MaxCenaVyrobca           float64
	MaxDoplatokPoistenca     float64
	MaxUhradaZP              float64
	PlatnostDatumDo          *time.Time
	PlatnostDatumOd          time.Time
	PocetSDL                 float64
	SposobUhrady             string
	StavKategorizacie        string
	UhradaZPNaSuhlasRL       int
	PreskripcneObmedzenie    []string
}

type LiekLink struct {
	Typ   string
	Nazov string
	Link  string
	Cena  float64
}

type CiselnikPlatnost struct {
	PlatnostOd time.Time
	PlatnostDo time.Time
}

//Created by Richard Malaschitz
//2018-12-28 13:51
//Copyright (c) 2018. All Rights Reserved.
