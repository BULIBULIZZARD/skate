package data

import (
	"file/skate/models"
	"file/skate/sql"
	"log"
)

type PlayerModel struct {
}

func NewPlayerModel() *PlayerModel {
	return new(PlayerModel)
}

func (p *PlayerModel) GetNameOrganizeById(id string) *models.SPlayer {
	engine := sql.GetSqlEngine()
	data := models.NewPlayer()
	_,err:=engine.Where("id=?",id).Get(data)
	if err != nil {
		log.Print(err.Error())
	}
	return data
}

func (p *PlayerModel) GetBestScoreById(id string,mName string) *models.MatchScore {
	engine := sql.GetSqlEngine()
	player := p.GetNameOrganizeById(id)
	data := models.NewMatchScore()
	_,err:=engine.Table("s_score").
		Join("INNER","s_match","s_score.match_id=s_match.id").
		Where("name=? and organize=? and match_name=?",player.PlayerName,player.Organize,mName).
		Asc("time_score").
		Get(data)
	if err != nil {
		log.Print(err.Error())
	}
	return data
}
