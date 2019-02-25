package data

import (
	"file/skate/models"
	"file/skate/sql"
	"log"
)

type MatchModel struct {
}

func NewMatchModel() *MatchModel {
	return new(MatchModel)
}

func (m *MatchModel) GetMatchByContestId(ContestId string) interface{} {
	engine := sql.GetSqlEngine()
	data := models.MoreMatch()
	err := engine.Where("contest_id=?", ContestId).Asc("match_name").Find(&data)
	if err != nil {
		log.Print(err.Error())
	}
	return data

}
