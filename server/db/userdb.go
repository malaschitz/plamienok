package db

import (
	"errors"
	"log"
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
	if !PasswordCheck(user, password) {
		err = errors.New("Nesprávne heslo")
		return
	}
	token = CreateToken(user.ID)
	return
}

func CreateToken(userid string) string {
	t := model.Token{ID: utils.UUID(), UserID: userid, Created: time.Now(), Validity: time.Now().Add(constants.TokenDuration)}
	err := _db.Save(&t)
	if err != nil {
		utils.LogErr(err)
	}
	return t.ID
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

func SetNewPassword(user model.User, password string) {
	hash, err := utils.EncodePassword(password)
	if err != nil {
		log.Panic(err)
	}
	setPassword(user.ID, hash)
}

func PasswordCheck(user model.User, password string) bool {
	hash := getPassword(user.ID)
	test, err := utils.ComparePasswordAndHash(password, hash)
	if err != nil {
		log.Panic(err)
	}
	return test
}

func SetCode6(key string, value model.Code6) {
	_db.Set(string(model.PROPERTY_CODE6), key, value)
}

func GetCode6(key string) (value model.Code6) {
	_db.Get(string(model.PROPERTY_CODE6), key, &value)
	return
}

func setPassword(key string, value string) {
	_db.Set(string(model.PROPERTY_PASSWORD_HASH), key, value)
}

func getPassword(key string) (value string) {
	_db.Get(string(model.PROPERTY_PASSWORD_HASH), key, &value)
	return
}
