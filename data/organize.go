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
	data := models.MoreOrganizePlayer()
	err := engine.Table("s_organize").
		Join("INNER", "s_player", "s_organize.organize_name=s_player.organize").
		Where("s_organize.id=?", oid).
		Find(&data)
	if err != nil {
		log.Print(err.Error())
	}
	player := models.MorePlayer()

	for _, v := range data {
		p := models.SPlayer{
			Id:         v.SPlayer.Id,
			PlayerName: v.SPlayer.PlayerName,
			Organize:   v.SPlayer.Organize,
		}
		player = append(player, &p)
	}
	return player
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
