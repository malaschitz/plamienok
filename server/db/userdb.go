package db

import (
	"errors"
	"github.com/malaschitz/plamienok/server/utils"
	"time"

	"github.com/malaschitz/plamienok/server/model"
)

func SaveUser(user *model.User, authorID string) error {
	user.Basification(authorID)
	return _db.Save(user)
}

func UserByID(id string) (user model.User, err error) {
	err = _db.One("ID", id, &user)
	return
}

func UserByEmail(email string) (user model.User, err error) {
	err = _db.One("Email", email, &user)
	return
}

func UserCount() int {
	var users []model.User
	utils.LogErr(_db.All(&users))
	count := 0
	for _, u := range users {
		if u.Deleted == nil {
			count++
		}
	}
	return count
}

func ValidToken(tokenKey string) (user model.User, err error) {
	var token model.Token
	err = _db.One("ID", tokenKey, &token)
	if err != nil {
		return
	}
	if time.Now().After(token.Platnost) {
		err = errors.New("token expired")
		return
	}
	user, err = UserByID(token.UserID)
	return
}
