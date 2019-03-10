package model

import (
	"fmt"
	"strconv"
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
	Month int
	Day   int
}

type DateTime struct {
	Date
	Hour   int
	Minute int
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
		Month: int(t.Month()),
		Day:   t.Day(),
	}
	return d
}

// format 2002-02-20
func String2Date(a string) (date *Date) {
	if len(a) < 10 {
		return nil
	}
	date = &Date{}
	date.Year, _ = strconv.Atoi(a[:4])
	date.Month, _ = strconv.Atoi(a[5:7])
	date.Day, _ = strconv.Atoi(a[8:10])
	return
}

func String2DateTime(a string) (date *DateTime) {
	if len(a) < 16 {
		return nil
	}
	date = &DateTime{}
	date.Day, _ = strconv.Atoi(a[:2])
	date.Month, _ = strconv.Atoi(a[3:5])
	date.Year, _ = strconv.Atoi(a[6:10])
	date.Hour, _ = strconv.Atoi(a[11:13])
	date.Minute, _ = strconv.Atoi(a[14:16])
	return
}

func Date2String(date *Date) (a string) {
	if date == nil {
		a = ""
	} else {
		a = fmt.Sprintf("%04d-%02d-%02d", date.Year, date.Month, date.Day)
	}
	return
}

func DateTime2String(date *DateTime) (a string) {
	if date == nil {
		a = ""
	} else {
		a = fmt.Sprintf("%04d-%02d-%02d %02d:%02d", date.Year, date.Month, date.Day, date.Hour, date.Minute)
	}
	return
}
