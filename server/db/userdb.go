package db

import (
	"errors"
	"strings"
	"time"

	"github.com/malaschitz/plamienok/server/constants"

	"github.com/malaschitz/plamienok/server/model"
	"github.com/malaschitz/plamienok/server/utils"
)

func SaveUser(user *model.User, authorID string) error {
	user.Email = strings.ToLower(user.Email)
	user.Basification(authorID)
	return _db.Save(user)
}

func UserByID(id string) (user model.User, err error) {
	err = _db.One("ID", id, &user)
	return
}

func UserByEmail(email string) (user model.User, err error) {
	email = strings.ToLower(email)
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

func Login(email string, password string) (user model.User, token string, err error) {
	email = strings.ToLower(email)
	user, err = UserByEmail(email)
	if err != nil {
		return
	}
	var test bool
	passwordHash := GetPassword(user.ID)
	test, err = utils.ComparePasswordAndHash(password, passwordHash)
	if err != nil {
		return
	}
	if !test {
		err = errors.New("Nesprávne heslo")
		return
	}

	t := model.Token{ID: utils.UUID(), UserID: user.ID, Created: time.Now(), Validity: time.Now().Add(constants.TokenDuration)}
	err = _db.Save(&t)
	if err != nil {
		return
	}
	token = t.ID
	return
}

func ValidToken(tokenKey string) (user model.User, err error) {
	var token model.Token
	err = _db.One("ID", tokenKey, &token)
	if err != nil {
		return
	}
	if time.Now().After(token.Validity) {
		err = errors.New("token expired")
		return
	}
	token.Validity = time.Now().Add(constants.TokenDuration)
	_db.Save(token) //error is not important
	user, err = UserByID(token.UserID)
	return
}
