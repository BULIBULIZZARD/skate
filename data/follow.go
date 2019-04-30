package data

import (
	"file/skate/models"
	"file/skate/sql"
	"log"
	"strconv"
	"time"
)

type FollowModel struct {
}

func NewFollowModel() *FollowModel {
	return new(FollowModel)
}

func (f *FollowModel) InsOrUpdate(uid string, fan string) {
	id := f.FindFan(uid, fan)
	if id != 0 {
		f.UpdateFan(id)
	}else {
		f.IncFan(uid, fan)
	}
}

func (f *FollowModel) FindFan(uid string, fan string) int {
	engine := sql.GetSqlEngine()
	follow := models.NewFollow()
	flag, err := engine.Where("user_id = ? and fan_id =?", uid, fan).Get(follow)
	if err != nil {
		log.Print(err.Error())
	}
	if flag {
		return follow.Id
	}
	return 0
}

func (f *FollowModel) IncFan(uid string, fan string) {
	engine := sql.GetSqlEngine()
	follow := models.NewFollow()
	follow.UserId, _ = strconv.Atoi(uid)
	follow.FanId, _ = strconv.Atoi(fan)
	follow.Status = 1
	follow.CreateTime = int(time.Now().Unix())
	_, err := engine.InsertOne(follow)
	if err != nil {
		log.Print(err.Error())
	}
}

func (f *FollowModel) UpdateFan(id int) {
	engine := sql.GetSqlEngine()
	follow := models.NewFollow()
	follow.Status = 1
	_, err := engine.Id(id).Update(follow)
	if err != nil {
		log.Print(err.Error())
	}
}

func (f *FollowModel) NopeFan(uid string, fan string) {
	engine := sql.GetSqlEngine()
	follow := models.NewFollow()
	follow.Status = 0
	_, err := engine.Where("user_id = ? and fan_id = ?", uid, fan).Cols("status").Update(follow)
	if err != nil {
		log.Print(err.Error())
	}
}

func (f *FollowModel) PlayerFunList(id string) []*models.PlayerFans {
	engine := sql.GetSqlEngine()
	follow := models.MorePlayerFans()
	err := engine.
		Table("s_follow").
		Join("INNER", "s_player", "s_follow.fan_id = s_player.id").
		Where("s_follow.user_id = ? and s_follow.status = ?", id, 1).
		Cols("s_follow.fan_id,s_follow.create_time,s_player.player_name,s_player.organize").
		Desc("s_follow.create_time").
		Find(&follow)
	if err != nil {
		log.Print(err.Error())
	}
	return follow
}
func (f *FollowModel) PlayerFollowList(id string) []*models.PlayerFans {
	engine := sql.GetSqlEngine()
	follow := models.MorePlayerFans()
	err := engine.
		Table("s_follow").
		Join("INNER", "s_player", "s_follow.user_id = s_player.id").
		Where("s_follow.fan_id = ? and s_follow.status = ?", id, 1).
		Cols("s_follow.user_id,s_follow.create_time,s_player.player_name,s_player.organize").
		Desc("s_follow.create_time").
		Find(&follow)
	if err != nil {
		log.Print(err.Error())
	}
	return follow
}
