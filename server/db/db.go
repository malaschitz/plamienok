package db

import (
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
	ImportCache("data")
}
