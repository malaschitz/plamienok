package main

import (
	"fmt"
	"time"

	"github.com/malaschitz/plamienok/server/constants"

	"github.com/malaschitz/plamienok/server/db"
	"github.com/malaschitz/plamienok/server/model"

	"github.com/malaschitz/plamienok/server/utils"
)

// this is not unit test
func main() {
	constants.InitConst()
	db.InitDB()
	data := model.Code6{Code6: utils.RandomCode(), Validity: time.Now().Add(time.Hour * 2)}
	db.SetCode6("test", data)
	//
	var data2 model.Code6
	data2 = db.GetCode6("test")
	utils.Jprint(data)
	utils.Jprint(data2)

	//again
	data = model.Code6{Code6: utils.RandomCode(), Validity: time.Now().Add(time.Hour * 2)}
	db.SetCode6("test", data)
	data2 = db.GetCode6("test")
	utils.Jprint(data)
	utils.Jprint(data2)

	//wrong
	data2 = db.GetCode6("testX")
	if data2.Code6 == "" {
		fmt.Println("It is OK")
	}
	utils.Jprint(data2)

	//password
	ID := utils.UUID()
	password := "secret"
	hash, err := utils.EncodePassword(password)
	if err != nil {
		panic(err)
	}
	db.SetPassword(ID, hash)
	hash2 := db.GetPassword(ID)
	if ok, err := utils.ComparePasswordAndHash(password, hash2); ok {
		fmt.Println("Password and hash are OK")
	} else {
		fmt.Println("problem with password and hash", err, password, hash, hash2)
	}
	utils.Jprint(hash)

}
