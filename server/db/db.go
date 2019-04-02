package db

import (
	"time"

	"github.com/asdine/storm"
	"github.com/labstack/gommon/log"
	"github.com/malaschitz/plamienok/server/constants"
	"github.com/malaschitz/plamienok/server/model"
)

var _db *storm.DB

func InitDB() {
	var err error
	_db, err = storm.Open(constants.DatabaseFile)
	if err != nil {
		panic(err)
	}
	log.Printf("Database initialized: %v", _db.Bolt.Path())
	// init
	_db.Init(&model.User{})

	// create admin ?
	user, err := UserByEmail(constants.AdminEmail)
	if err != nil {
		if err == storm.ErrNotFound {
			user = model.User{Email: constants.AdminEmail, Name: "admin", RoleAdmin: true}
			user.Basification("")
			err = SaveUser(&user, "")
			if err == nil {
				SetNewPassword(user, constants.AdminPassword)
			} else {
				log.Panic(err)
			}
		} else {
			log.Panic(err)
		}
	}

	go func() {
		//reindex fulltext
		t := time.Now()
		var persons []*model.Person
		err = _db.All(&persons)
		if err != nil {
			panic(err)
		}
		for _, p := range persons {
			fullText(p)
			err = _db.Save(p)
			if err != nil {
				panic(err)
			}
		}
		log.Printf("Fulltext reindex: %v", time.Since(t))

		//cache
		ImportCache("data")
	}()

}
