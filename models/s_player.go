package models

type SPlayer struct {
	Id              int    `json:"id" xorm:"not null pk autoincr unique INT(4)"`
	PlayerName      string `json:"player_name" xorm:"not null default '' unique(player_name) VARCHAR(45)"`
	Organize        string `json:"organize" xorm:"not null default '' unique(player_name) VARCHAR(45)"`
	OrganizeId      int    `json:"organize_id" xorm:"not null default 0 INT(3)"`
	PlayerGender    string `json:"player_gender" xorm:"not null default '' VARCHAR(10)"`
	PlayerGroupType string `json:"player_group_type" xorm:"not null default '' VARCHAR(10)"`
	Username        string `json:"username" xorm:"not null default '' VARCHAR(20)"`
	Password        string `json:"password" xorm:"not null default '' VARCHAR(40)"`
	CreateTime      string `json:"create_time" xorm:"not null default '' VARCHAR(45)"`
	UpdateTime      string `json:"update_time" xorm:"not null default '' VARCHAR(45)"`
	DeleteTime      string `json:"delete_time" xorm:"not null default '' VARCHAR(45)"`
}

func NewPlayer() *SPlayer {
	return new(SPlayer)
}

func MorePlayer() []*SPlayer {
	return make([]*SPlayer, 0)
}
