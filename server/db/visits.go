package db

import (
	"sort"

	"github.com/malaschitz/plamienok/server/model"
)

func DtoVisits(person model.Person) (visits []model.VisitDto) {
	visits = make([]model.VisitDto, 0)
	//h
	var vh []model.VisitHome
	_db.Find("PersonID", person.ID, &vh)
	for _, v := range vh {
		dto := model.VisitDto{Visit: v.Visit, DtoTyp: "H", DtoDatum: model.Date2String(&v.Datum.Date)}
		visits = append(visits, dto)
	}

	//p
	var vp []model.VisitCall
	_db.Find("PersonID", person.ID, &vp)
	for _, v := range vp {
		dto := model.VisitDto{Visit: v.Visit, DtoTyp: "P", DtoDatum: model.Date2String(&v.Datum.Date)}
		visits = append(visits, dto)
	}

	//sort
	sort.Slice(visits, func(i, j int) bool {
		return visits[j].Datum.After(visits[i].Datum)
	})

	return
}

func VisitCalls() (visits []model.VisitCall, err error) {
	err = _db.All(&visits)
	sort.Slice(visits, func(i, j int) bool {
		return visits[i].Datum.After(visits[j].Datum)
	})
	return
}

func VisitHomes() (visits []model.VisitHome, err error) {
	err = _db.All(&visits)
	sort.Slice(visits, func(i, j int) bool {
		return visits[i].Datum.After(visits[j].Datum)
	})
	return
}

func VisitHomeByID(id string) (visit model.VisitHome, err error) {
	err = _db.One("ID", id, &visit)
	return
}

func SaveVisitHome(visit *model.VisitHome, authorID string) (err error) {
	visit.Basification(authorID)
	err = _db.Save(visit)
	return
}

func VisitCallByID(id string) (visit model.VisitCall, err error) {
	err = _db.One("ID", id, &visit)
	return
}

func SaveVisitCall(visit *model.VisitCall, authorID string) (err error) {
	visit.Basification(authorID)
	err = _db.Save(visit)
	return
}
