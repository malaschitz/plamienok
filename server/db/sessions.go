package db

import (
	"github.com/malaschitz/plamienok/server/model"
)

func Sessions() (sessions []model.Session, err error) {
	err = _db.All(&sessions)
	return
}

func SaveSession(session *model.Session, authorID string) error {
	session.Basification(authorID)
	err := _db.Save(session)
	return err
}
