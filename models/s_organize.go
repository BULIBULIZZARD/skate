package models

type SOrganize struct {
	Id           int    `json:"id" xorm:"not null pk autoincr INT(4)"`
	OrganizeName string `json:"organize_name" xorm:"not null default '' VARCHAR(45)"`
	CreateTime   string `json:"create_time" xorm:"not null default '' VARCHAR(45)"`
	UpdateTime   string `json:"update_time" xorm:"not null default '' VARCHAR(45)"`
	DeleteTime   string `json:"delete_time" xorm:"not null default '' VARCHAR(45)"`
}