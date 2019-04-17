package data

import (
	"file/skate/models"
	"file/skate/sql"
	"file/skate/tools"
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
	_, err := engine.Where("id=?", id).Get(data)
	if err != nil {
		log.Print(err.Error())
	}
	return data
}

func (p *PlayerModel) GetBestScoreById(id string, mName string) *models.MatchScore {
	engine := sql.GetSqlEngine()
	data := models.NewMatchScore()
	_, err := engine.Table("s_score").
		Join("INNER", "s_match", "s_score.match_id=s_match.id").
		Where("player_id=? and match_name like ? and time_score <> ? and time_score <> ?", id, mName+"%", "00:00.000", "完成比赛").
		Asc("time_score").
		Get(data)
	if err != nil {
		log.Print(err.Error())
	}
	return data
}

func (p *PlayerModel) GetAllScoreByMatchAndPlayer(id string, mName string) []*models.MatchScore {
	engine := sql.GetSqlEngine()
	data := models.MoreMatchScore()
	err := engine.Table("s_score").
		Join("INNER", "s_match", "s_score.match_id=s_match.id").
		Where("s_score.player_id=? and s_match.match_name like ? and s_score.time_score <> ? and  s_score.time_score <> ?", id, mName+"%", "00:00.000", "完成比赛").
		Cols("s_group", "match_type", "time_score", "date", "match_id").
		Asc(`match_time`).
		Asc(`s_score.id`).
		Find(&data)
	if err != nil {
		log.Print(err.Error())
	}
	return data
}

func (p *PlayerModel) TreeData(id string) []*models.MatchScore {
	engine := sql.GetSqlEngine()
	data := models.MoreMatchScore()
	err := engine.Table("s_score").
		Join("INNER", "s_match", "s_score.match_id=s_match.id").
		Where("s_score.player_id=?  and s_score.time_score <> ? and  s_score.time_score <> ?", id, "00:00.000", "完成比赛").
		Cols("match_name", "time_score").
		Asc(`match_time`).
		Asc(`s_score.id`).
		Find(&data)
	if err != nil {
		log.Print(err.Error())
	}
	return data
}


func (p *PlayerModel) PlayerLoginCheck(username string, password string) (*models.SPlayer, bool) {
	engine := sql.GetSqlEngine()
	player := models.NewPlayer()
	flag, err := engine.Where("username=? and password=?", username, password).Get(player)
	if err != nil {
		log.Print(err.Error())
	}
	return player, flag
}

func (p *PlayerModel) PlayerChangePassword(id string, ordPass string, newPass string) bool {
	engine := sql.GetSqlEngine()
	player := models.NewPlayer()
	player.Password = tools.NewTools().Sha1(newPass)
	flag, _ := engine.Id(id).Where("password=?", tools.NewTools().Sha1(ordPass)).Update(player)
	return flag == 1
}
