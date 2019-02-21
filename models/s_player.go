package models

type SPlayer struct {
	Id         int    `json:"id" xorm:"not null pk autoincr INT(4)"`
	PlayerName string `json:"player_name" xorm:"not null default '' unique(player_name) VARCHAR(45)"`
	Organize   string `json:"organize" xorm:"not null default '' unique(player_name) VARCHAR(45)"`
}
