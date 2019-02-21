package models

type SContest struct {
	Id          int    `json:"id" xorm:"not null pk INT(3)"`
	Name        string `json:"name" xorm:"not null default '' VARCHAR(45)"`
	Station     string `json:"station" xorm:"not null default '' VARCHAR(45)"`
	Type        string `json:"type" xorm:"not null default '' VARCHAR(45)"`
	ContestTime string `json:"contest_time" xorm:"not null default '' VARCHAR(45)"`
}
