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
func (m *ScoreModel) GetScoreByPlayerAndOrganize(id string) []*models.MatchScore {
	engine := sql.GetSqlEngine()
	data := models.MoreMatchScore()
	err := engine.Table("s_score").
		Join("INNER", "s_match", "s_score.match_id=s_match.id").
		Where("s_score.player_id=? ", id).
		Asc(`no`).
		Find(&data)
	if err != nil {
		log.Print(err.Error())
	}
	return data
}