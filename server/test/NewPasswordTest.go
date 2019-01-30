package main

import (
	"fmt"
	"time"

	"github.com/malaschitz/plamienok/server/constants"
	"github.com/malaschitz/plamienok/server/db"
)

func main() {
	constants.InitConst()
	db.InitDB()

	//change
	admin, err := db.UserByEmail(constants.AdminEmail)
	if err != nil {
		panic(err)
	}
	pwd := time.Now().Format(time.RFC3339)
	db.SetNewPassword(admin, pwd)
	//check
	if !db.PasswordCheck(admin, pwd) {
		panic(err)
	}
	//login
	user, token, err := db.Login(constants.AdminEmail, pwd)
	if err != nil {
		panic(err)
	}
	if user.ID != admin.ID {
		panic(err)
	}
	fmt.Println("TOKEN", token)

	//change again
	db.SetNewPassword(admin, constants.AdminPassword)

}
