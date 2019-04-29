package data

import (
	"file/skate/config"
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
	for _, m := range data {
		if m.No == "9" {
			m.No = ""
		}
	}
	return data
}
func (m *ScoreModel) GetScoreByPlayerId(id string, page int) []*models.MatchScore {
	engine := sql.GetSqlEngine()
	data := models.MoreMatchScore()
	err := engine.Table("s_score").
		Join("INNER", "s_match", "s_score.match_id=s_match.id").
		Where("s_score.player_id=? ", id).
		Asc(`s_score.id`).
		Limit(config.GetConfig().GetPageSize(), page*config.GetConfig().GetPageSize()).
		Find(&data)
	if err != nil {
		log.Print(err.Error())
	}
	return data
}

func (m *ScoreModel) GetScoreCountByPlayerId(id string) int {
	engine := sql.GetSqlEngine()
	data := models.NewMatchScore()
	count, err := engine.Table("s_score").
		Join("INNER", "s_match", "s_score.match_id=s_match.id").
		Where("s_score.player_id=? ", id).
		Asc(`s_score.id`).
		Count(data)
	if err != nil {
		log.Print(err.Error())
	}
	if int(count)%config.GetConfig().GetPageSize() > 0 {
		return int(count)/config.GetConfig().GetPageSize() + 1
	}
	return int(count) / config.GetConfig().GetPageSize()
}
