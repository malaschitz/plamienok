package main

import (
	"fmt"
	"log"

	"github.com/malaschitz/plamienok/server/constants"

	"github.com/malaschitz/plamienok/server/service"
	"github.com/malaschitz/plamienok/server/utils"
)

// this is not unit test
func main() {
	constants.InitConst()
	service.InitEmail()
	code6 := utils.RandomCode()
	err := service.SendCode("plamienok@mailinator.com", "Jozef Mrkvička", code6)
	if err != nil {
		fmt.Println("Sending email error")
		log.Fatal(err)
	}
}
