package db

import "github.com/malaschitz/plamienok/server/model"

var Diagnozy []model.Diagnoza
var Lieky []model.Liek

func CacheSet(key, value string) {
	_db.Set("cache", key, value)
}

func CacheGet(key string) (value string) {
	_db.Get("cache", key, &value)
	return
}

//Created by Richard Malaschitz malaschitz@gmail.com
//23/10/2018 15:32
//Copyright (c) 2018. All Rights Reserved.
