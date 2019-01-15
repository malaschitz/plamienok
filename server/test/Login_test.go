package test

import (
	"fmt"
	"github.com/malaschitz/plamienok/server/model"
	"github.com/malaschitz/plamienok/server/utils"
	"testing"
	"time"
)

func TestEncoding(t *testing.T) {
	//password and hash
	encode, err := utils.EncodePassword("secret")
	if err != nil {
		panic(err)
	}
	fmt.Println("encode", encode)
	fmt.Println(utils.ComparePasswordAndHash("secret", encode))
}

func TestMeniny(t *testing.T) {
	tData := [][]interface{}{
		{"Tamara", 26, time.January},
		{"Richard", 3, time.April},
		{"Alenka", 27, time.March},
		{"Oskar", 8, time.August},
		{"Oskár", 8, time.August},
		{"Zdeno", 9, time.February},
		{"Zdenko", 9, time.February},
	}
	for _, m := range tData {
		meniny := model.Meniny(m[0].(string))
		if meniny.Day != m[1] || meniny.Month != m[2] {
			t.Fatalf("Nesprávne meniny pre %s. Má byť %d %s a je %d %s.", m[0], m[1], m[2], meniny.Day, meniny.Month)
		}
	}
}

//Created by Richard Malaschitz
//30/11/2018 11:13
//Copyright (c) 2018. All Rights Reserved.
