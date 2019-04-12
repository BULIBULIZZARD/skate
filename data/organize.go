package data

import (
	"file/skate/models"
	"file/skate/sql"
	"log"
)

type OrganizeModel struct {
}

func NewOrganizeModel() *OrganizeModel {
	return new(OrganizeModel)
}

func (o *OrganizeModel) GetAllPlayerById(oid string) []*models.SPlayer {
	engine := sql.GetSqlEngine()
	data := models.MorePlayer()
	err := engine.Table("s_player").
		Where("organize_id=?", oid).
		Cols("id","player_name","gender","group_type").
		Find(&data)
	if err != nil {
		log.Print(err.Error())
	}
	return data
}
func (o *OrganizeModel) CheckOrganizeLogin(username string, password string) (*models.SOrganize, bool) {
	engine := sql.GetSqlEngine()
	organize := models.NewOrganize()
	flag, err := engine.Where("username=? and password=?", username, password).Get(organize)
	if err != nil {
		log.Print(err.Error())
	}
	return organize, flag
}

func (o *OrganizeModel)  GetAllPlayerScore(oid string) []*models.OrganizePlayerScore{
	engine := sql.GetSqlEngine()
	data:=models.MoreOrganizePlayerScore()
	err := engine.Table("s_organize").
		Join("INNER", "s_player", "s_organize.id=s_player.organize_id").
		Join("INNER", "s_score","s_score.player_id=s_player.id").
		Join("INNER", "s_match","s_score.match_id=s_match.id").
		Cols("player_name","s_group","match_id","match_type","date","time_score","match_name","group_type","no").
		Where("s_organize.id=?", oid).
		Asc("s_score.id").
		Find(&data)
	if err != nil {
		log.Print(err.Error())
	}
	return data
}