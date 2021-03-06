package model

import (
	"fmt"
	"strconv"
	"strings"
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

func (dt DateTime) After(d2 DateTime) bool {
	n1 := dt.Year*10000 + dt.Month*100 + dt.Day
	n2 := d2.Year*10000 + d2.Month*100 + d2.Day
	if n1 > n2 {
		return true
	} else {
		return false
	}
}

func (dt Date) AddDay() Date {
	y := dt.Year
	if y == 0 {
		y = time.Now().Year()
	}
	t := time.Date(y, time.Month(dt.Month), dt.Day, 0, 0, 0, 0, time.UTC)
	t = t.AddDate(0, 0, 1)
	return Date{t.Year(), int(t.Month()), t.Day()}
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

func Time2DateTime(t time.Time) DateTime {
	dt := DateTime{
		Date{
			Year:  t.Year(),
			Month: int(t.Month()),
			Day:   t.Day(),
		},
		t.Hour(),
		t.Minute(),
	}
	return dt
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

//2008-02-11 03:20
func String2DateTime(a string) (date *DateTime) {
	if len(a) < 16 {
		return nil
	}
	date = &DateTime{}
	date.Year, _ = strconv.Atoi(a[:4])
	date.Month, _ = strconv.Atoi(a[5:7])
	date.Day, _ = strconv.Atoi(a[8:10])
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
	if date != nil {
		a = fmt.Sprintf("%04d-%02d-%02d %02d:%02d", date.Year, date.Month, date.Day, date.Hour, date.Minute)
	}
	return
}

func DateTime2DTO(date *DateTime) (dto DateTimeDto) {
	if date != nil {
		dto.Date = fmt.Sprintf("%04d-%02d-%02d", date.Year, date.Month, date.Day)
		dto.Time = fmt.Sprintf("%02d:%02d", date.Hour, date.Minute)
		if strings.HasPrefix(dto.Date, "0000") {
			dto.Date = "2000-01-01"
		}
	}
	return
}

func DTO2DateTime(dto DateTimeDto) (d *DateTime) {
	d = String2DateTime(dto.Date + " " + dto.Time)
	fmt.Println("DTO", dto.Date, dto.Time, "->", d)
	return d
}
