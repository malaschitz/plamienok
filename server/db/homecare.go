package db

import (
	"github.com/malaschitz/plamienok/server/model"
	"github.com/malaschitz/plamienok/server/utils"
)

func VisitCount() int {
	var visits []model.VisitHome
	utils.LogErr(_db.All(&visits))
	return len(visits)
}
