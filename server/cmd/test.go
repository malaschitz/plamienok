package main

import (
	"fmt"

	"github.com/malaschitz/plamienok/server/utils"

	"github.com/malaschitz/plamienok/server"

	"github.com/malaschitz/plamienok/server/db"

	"github.com/malaschitz/plamienok/server/model"
)

func main() {
	fmt.Println(model.Meniny("Tamara"))
	fmt.Println(model.Meniny("Richard"))
	fmt.Println(model.Meniny("Alenka"))
	fmt.Println(model.Meniny("Oskar"))
	fmt.Println(model.Meniny("Zdeno"))

	server.InitConst()
	db.InitDB()

	fmt.Println(db.Lieky[100])
	fmt.Println(db.Diagnozy[100])

	//password and hash
	encode, err := utils.EncodePassword("secret")
	if err != nil {
		panic(err)
	}
	fmt.Println("encode", encode)
	fmt.Println(utils.ComparePasswordAndHash("secret", encode))

}

//Created by Richard Malaschitz malaschitz@gmail.com
//2018-12-27 17:17
//Copyright (c) 2018. All Rights Reserved.
