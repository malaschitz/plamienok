package db

import (
	"github.com/malaschitz/plamienok/server/model"
	"github.com/malaschitz/plamienok/server/model/dto"
)

func DtoCiele(person model.Person) (tasks []dto.CielDto) {
	tasks = make([]dto.CielDto, 0)
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

func CielToDto(ciel model.Ciel) dto.CielDto {
	s := dto.CielDto{Ciel: ciel}
	u, _ := UserByID(ciel.AuthorID)
	s.UserName = u.Name
	s.IsDeleted = (ciel.Deleted != nil)
	return s
}
