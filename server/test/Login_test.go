package test

import (
	"strings"
	"testing"
	"time"

	"github.com/malaschitz/plamienok/server/model"
	"github.com/malaschitz/plamienok/server/utils"

	"github.com/malaschitz/plamienok/server"
	"github.com/malaschitz/plamienok/server/db"
)

//Created by Richard Malaschitz malaschitz@gmail.com
//30/11/2018 11:13
//Copyright (c) 2018. All Rights Reserved.
var email = "malaschitz@gmail.com"

func TestLogin(t *testing.T) {
	student := db.StudentByEmail(email)
	newStudent := student.ID == 0
	lang := "sk"
	if _, ok := model.LANGS_MAP[lang]; !ok {
		lang = "en"
	}

	if newStudent {
		name := email[:strings.Index(email, "@")]
		student = model.Student{
			Email:         email,
			Name:          name,
			FirstLanguage: lang,
		}
		db.InsertStudent(&student)
	}
	token := utils.UUID()
	tm := time.Now().Add(time.Hour * 2)
	db.InsertToken(student.ID, token, tm)
}

func init() {
	server.InitConst()
	db.InitDB()

}
