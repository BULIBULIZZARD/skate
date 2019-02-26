package data

import (
	"file/skate/models"
	"file/skate/sql"
	"log"
)

type ScoreModel struct {
}

func NewScoreModel() *ScoreModel {
	return new(ScoreModel)
}

func (m *ScoreModel) GetScoreByMatchAndGroup(mid string, group string) []*models.SScore {
	engine := sql.GetSqlEngine()
	data := models.MoreScore()
	err := engine.Where("match_id=? and s_group=? ", mid, "第"+group+"组").Asc(`no`).Find(&data)
	if err != nil {
		log.Print(err.Error())
	}
	return data
}
func (m *ScoreModel) GetScoreByPlayerAndOrganize(PlayerName string,Organize string) []*models.SScore  {
	engine := sql.GetSqlEngine()
	data := models.MoreScore()
	err := engine.Where("name=? and organize=? ", PlayerName, Organize).Asc(`no`).Find(&data)
	if err != nil {
		log.Print(err.Error())
	}
	return data
}