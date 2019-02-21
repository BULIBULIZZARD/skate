package data

import (
	"file/skate/models"
	"file/skate/sql"
	"github.com/labstack/gommon/log"
)

type ContestModel struct {
}

func NewContestModel() *ContestModel {
	return new(ContestModel)
}
func (m *ContestModel) GetAllContest() interface{} {
	engine := sql.GetSqlEngine()
	data := models.MoreContest()
	err := engine.OrderBy("id desc").Find(&data)
	if err != nil {
		log.Print(err.Error())
	}
	return data
}
