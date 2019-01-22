package db

import (
	"github.com/asdine/storm"
	"github.com/labstack/gommon/log"
	"github.com/malaschitz/plamienok/server/constants"
	"github.com/malaschitz/plamienok/server/model"
	"github.com/malaschitz/plamienok/server/utils"
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
			hash, err := utils.EncodePassword(constants.AdminPassword)
			if err == nil {
				user = model.User{Email: constants.AdminEmail, Name: "admin"}
				user.Basification("")
				err = SaveUser(&user, "")
				if err == nil {
					SetPassword(user.ID, hash)
				}
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

func SetCode6(key string, value model.Code6) {
	_db.Set(string(model.PROPERTY_CODE6), key, value)
}

func GetCode6(key string) (value model.Code6) {
	_db.Get(string(model.PROPERTY_CODE6), key, &value)
	return
}

func SetPassword(key string, value string) {
	_db.Set(string(model.PROPERTY_PASSWORD_HASH), key, value)
}

func GetPassword(key string) (value string) {
	_db.Get(string(model.PROPERTY_PASSWORD_HASH), key, &value)
	return
}
