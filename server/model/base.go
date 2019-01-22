package model

import (
	"fmt"
	"time"

	"github.com/malaschitz/plamienok/server/utils"
)

type Base struct {
	ID       string `storm:"id"` //ID zaznamu
	AuthorID string //ID autora
	Revision int    //cislo revizie
	Created  time.Time
	Updated  time.Time
	Deleted  *time.Time
}

type Token struct {
	ID       string `storm:"id"`
	UserID   string
	Created  time.Time
	Validity time.Time
}

type Date struct {
	Year  int
	Month time.Month
	Day   int
}

type Log struct {
	ID       string `storm:"id"`
	Created  time.Time
	AuthorID string
	ObjectID string
	Data     []byte
}

func (base *Base) Basification(authorID string) {
	now := time.Now()
	if base.ID == "" { //new
		base.ID = utils.UUID()
		base.Revision = 0
		base.Created = now
		base.AuthorID = authorID
		base.Updated = base.Created
	} else {
		base.Revision = base.Revision + 1
		base.AuthorID = authorID
		base.Updated = now
	}
}

func (u User) String() string {
	return fmt.Sprintf("[%s] <%s %s>", u.ID, u.Name, u.Email)
}

func Time2Date(t time.Time) Date {
	d := Date{
		Year:  t.Year(),
		Month: t.Month(),
		Day:   t.Day(),
	}
	return d
}
