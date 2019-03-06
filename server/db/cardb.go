/*
 * Developed by Richard Malaschitz on 3/4/19 1:21 PM
 * Last modified 3/4/19 9:49 AM
 * Copyright (c) 2019. All right reserved.
 *
 */

package db

import (
	"strings"

	"github.com/malaschitz/plamienok/server/model"
	"github.com/malaschitz/plamienok/server/utils"
)

func SaveCar(car *model.Car, authorID string) error {
	car.Basification(authorID)
	return _db.Save(car)
}

func CarByID(id string) (car model.Car, err error) {
	err = _db.One("ID", id, &car)
	return
}

func CarByEmail(email string) (car model.Car, err error) {
	email = strings.ToLower(email)
	err = _db.One("Email", email, &car)
	return
}

func Cars(all bool) (cars []model.Car, err error) {
	var z []model.Car
	err = _db.All(&z)
	if all {
		cars = z
	} else {
		for _, u := range z {
			if u.Deleted == nil {
				cars = append(cars, u)
			}
		}
	}
	return
}

func CarCount() int {
	count, err := _db.Count(&model.Car{})
	utils.LogErr(err)
	return count

}
