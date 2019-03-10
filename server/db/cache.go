package db

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sort"
	"time"

	"github.com/malaschitz/plamienok/server/model/dto"

	"github.com/labstack/gommon/log"
	"github.com/malaschitz/plamienok/server/model"
	"github.com/malaschitz/plamienok/server/utils"
)

var Diagnozy DiagnozyCache
var Lieky LiekyCache

type DiagnozyCache struct {
	Data []model.Diagnoza
}

type LiekyCache struct {
	Data []model.Liek
}

func GetDiagnozy(filter string, max int) []model.Diagnoza {
	dgns := make([]model.Diagnoza, 0)
	ftFilter := utils.FullTextCreate(filter)
	for _, d := range Diagnozy.Data {
		if utils.FullTextTest(ftFilter, d.Skratka+" "+d.Popis) {
			dgns = append(dgns, d)
		}
		if len(dgns) == max {
			break
		}
	}
	return dgns
}

func GetDiagnozyAll() []dto.TextValueDto {
	dgns := make([]dto.TextValueDto, 0)
	for _, d := range Diagnozy.Data {
		d := dto.TextValueDto{Text: d.Popis, Value: d.Skratka}
		dgns = append(dgns, d)
	}
	return dgns
}

func GetLieky(filterBy string, max int) (lieky []model.Liek) {
	lieky = make([]model.Liek, 0)
	ftFiler := utils.FullTextCreate(filterBy)
	for _, l := range Lieky.Data {
		if utils.FullTextTest(ftFiler, l.Nazov+" "+l.KodSukl) {
			lieky = append(lieky, l)
		}
		if len(lieky) == max {
			break
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
	sort.Slice(Diagnozy.Data, func(i, j int) bool {
		return Diagnozy.Data[i].Skratka < Diagnozy.Data[j].Skratka
	})
	log.Printf("Import Cache %v", time.Since(t))
}
