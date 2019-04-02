package db

import (
	"github.com/malaschitz/plamienok/server/model"
)

func DtoCiele(person model.Person) (tasks []model.CielDto) {
	tasks = make([]model.CielDto, 0)
	var ciele []model.Ciel
	err := _db.Find("PersonID", person.ID, &ciele)
	if err != nil {
		return
	}
	for _, c := range ciele {
		dto := CielToDto(c)
		tasks = append(tasks, dto)
	}
	return

}

func SaveCiel(ciel *model.Ciel, authorID string) error {
	ciel.Basification(authorID)
	err := _db.Save(ciel)
	return err
}

func CielToDto(ciel model.Ciel) model.CielDto {
	s := model.CielDto{Ciel: ciel}
	u, _ := UserByID(ciel.AuthorID)
	s.UserName = u.Name
	s.IsDeleted = (ciel.Deleted != nil)
	return s
}
