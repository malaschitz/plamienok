package db

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/malaschitz/plamienok/server/model"
)

var Diagnozy DiagnozyCache
var Lieky LiekyCache

type DiagnozyCache struct {
	Data []model.Diagnoza
}

type LiekyCache struct {
	Data []model.Liek
}

func ImportCache(dir string) {
	// import lieky
	jsonFile, err := os.Open(dir + string(os.PathSeparator) + "lieky.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(bytes), &Lieky.Data)
	if err != nil {
		panic(err)
	}

	// diagnozy
	jsonFile, err = os.Open(dir + string(os.PathSeparator) + "dgn.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	bytes, err = ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}
	var cd model.Codebook
	err = json.Unmarshal([]byte(bytes), &cd)
	if err != nil {
		panic(err)
	}
	Diagnozy.Data = make([]model.Diagnoza, 0)
	index := 0
	for _, value := range cd.PolozkyCiselnika {
		dg := model.Diagnoza{
			CiselnikPolozka: value,
			Podskupina:      value.Atributy[7],
			Skupina:         value.Atributy[8],
			Kapitola:        value.Atributy[9]}
		if _, ok := dg.DisplayNames[cd.Version]; ok {
			Diagnozy.Data = append(Diagnozy.Data, dg)
		}
		index++
	}
}

//Created by Richard Malaschitz
//23/10/2018 15:32
//Copyright (c) 2018. All Rights Reserved.
