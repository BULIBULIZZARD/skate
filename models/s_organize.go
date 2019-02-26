package models

type SOrganize struct {
	Id           int    `json:"id" xorm:"not null pk autoincr INT(4)"`
	OrganizeName string `json:"organize_name" xorm:"not null default '' VARCHAR(45)"`
}
