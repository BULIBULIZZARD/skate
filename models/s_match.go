package models

type SMatch struct {
	Id        int    `json:"id" xorm:"not null pk unique INT(6)"`
	Time      string `json:"time" xorm:"not null default '' VARCHAR(45)"`
	MatchName string `json:"match_name" xorm:"not null default '' VARCHAR(45)"`
	GroupType string `json:"group_type" xorm:"not null default '' VARCHAR(45)"`
	Gender    string `json:"gender" xorm:"not null default '' VARCHAR(45)"`
	MatchType string `json:"match_type" xorm:"not null default '' VARCHAR(45)"`
	PlayerNum string `json:"player_num" xorm:"not null default '' VARCHAR(45)"`
	GroupNum  string `json:"group_num" xorm:"not null default '' VARCHAR(45)"`
	Enter     string `json:"enter" xorm:"not null default '' VARCHAR(45)"`
	Remark    string `json:"remark" xorm:"not null default '' VARCHAR(45)"`
	ContestId int    `json:"contest_id" xorm:"not null INT(3)"`
}

func NewMatch() *SMatch {
	return new(SMatch)
}
