package db

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/labstack/gommon/log"
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

func GetLieky(filterBy string) (lieky []model.Liek) {
	lieky = make([]model.Liek, 0)
	filters := strings.Split(strings.ToLower(filterBy), " ")
	for _, l := range Lieky.Data {
		test := true
		for _, f := range filters {
			search := strings.ToLower(l.Nazov)
			test = test && strings.Contains(search, f)
		}
		if test {
			lieky = append(lieky, l)
			if len(lieky) >= 20 {
				break
			}
		}
	}
	return
}

func GetDiagnozy(filterBy string) (diagnozy []model.Diagnoza) {
	diagnozy = make([]model.Diagnoza, 0)
	filters := strings.Split(strings.ToLower(filterBy), " ")
	for _, d := range Diagnozy.Data {
		test := true
		for _, f := range filters {
			search := strings.ToLower(d.Skratka) + strings.ToLower(d.Popis)
			test = test && strings.Contains(search, f)
		}
		if test {
			diagnozy = append(diagnozy, d)
			if len(diagnozy) >= 20 {
				break
			}
		}
	}
	return
}

func ImportCache(dir string) {
	t := time.Now()
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
	log.Printf("Import Cache %v", time.Since(t))
}
