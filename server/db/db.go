package db

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/asdine/storm"
	"github.com/labstack/gommon/log"
	"github.com/malaschitz/plamienok/server"
	"github.com/malaschitz/plamienok/server/model"
	"github.com/malaschitz/plamienok/server/utils"
)

var _db *storm.DB

func InitDB() {
	var err error
	_db, err = storm.Open(server.DatabaseFile)
	if err != nil {
		panic(err)
	}
	log.Printf("Database initialized: %v", _db.Bolt.Path())
	// init
	_db.Init(&model.User{})

	// create admin ?
	user, err := UserByEmail(server.AdminEmail)
	if err != nil {
		if err == storm.ErrNotFound {
			hash, err := utils.EncodePassword(server.AdminPassword)
			if err == nil {
				user = model.User{Email: server.AdminEmail, Name: "admin", Password: hash}
				user.Basification("")
				err = SaveUser(&user, "")
			}
			if err != nil {
				log.Panic(err)
			}
		} else {
			log.Panic(err)
		}
	}

	// import lieky
	jsonFile, err := os.Open("data/lieky.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(bytes), &Lieky)
	if err != nil {
		panic(err)
	}

	// diagnozy
	jsonFile, err = os.Open("data/dgn.json")
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
	Diagnozy = make([]model.Diagnoza, 0)
	index := 0
	for _, value := range cd.PolozkyCiselnika {
		dg := model.Diagnoza{
			CiselnikPolozka: value,
			Podskupina:      value.Atributy[7],
			Skupina:         value.Atributy[8],
			Kapitola:        value.Atributy[9]}
		if _, ok := dg.DisplayNames[cd.Version]; ok {
			Diagnozy = append(Diagnozy, dg)
		}
		index++
	}
}
