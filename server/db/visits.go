package db

import (
	"sort"

	"github.com/malaschitz/plamienok/server/model"
	"github.com/malaschitz/plamienok/server/model/dto"
)

func DtoVisits(person model.Person) (visits []dto.VisitDto) {
	visits = make([]dto.VisitDto, 0)
	//h
	var vh []model.VisitHome
	_db.Find("PersonID", person.ID, &vh)
	for _, v := range vh {
		dto := dto.VisitDto{Visit: v.Visit, DtoTyp: "H", DtoDatum: model.Date2String(&v.Datum.Date)}
		visits = append(visits, dto)
	}

	//p
	var vp []model.VisitPhone
	_db.Find("PersonID", person.ID, &vp)
	for _, v := range vp {
		dto := dto.VisitDto{Visit: v.Visit, DtoTyp: "P", DtoDatum: model.Date2String(&v.Datum.Date)}
		visits = append(visits, dto)
	}

	//sort
	sort.Slice(visits, func(i, j int) bool {
		return visits[j].Datum.After(visits[i].Datum)
	})

	return
}

func SaveVisitHome(visit model.VisitHome, authorID string) (err error) {
	visit.Basification(authorID)
	err = _db.Save(&visit)
	return
}

func SaveVisitPhone(visit model.VisitPhone, authorID string) (err error) {
	visit.Basification(authorID)
	err = _db.Save(&visit)
	return
}
