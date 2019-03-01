package models

type SScore struct {
	Id        int    `json:"id" xorm:"not null pk autoincr INT(5)"`
	SGroup    string `json:"s_group" xorm:"not null default '' VARCHAR(45)"`
	No        string `json:"no" xorm:"not null default '' VARCHAR(45)"`
	RowNum    string `json:"row_num" xorm:"not null default '' VARCHAR(45)"`
	HeadNum   string `json:"head_num" xorm:"not null default '' VARCHAR(45)"`
	Name      string `json:"name" xorm:"not null default '' unique(name) VARCHAR(45)"`
	Organize  string `json:"organize" xorm:"not null default '' VARCHAR(45)"`
	TimeScore string `json:"time_score" xorm:"not null default '' VARCHAR(45)"`
	Remark    string `json:"remark" xorm:"not null default '' VARCHAR(45)"`
	MatchId   int    `json:"match_id" xorm:"not null unique(name) INT(6)"`
	CreateTime string `json:"create_time" xorm:"not null default '' VARCHAR(45)"`
	UpdateTime string `json:"update_time" xorm:"not null default '' VARCHAR(45)"`
	DeleteTime string `json:"delete_time" xorm:"not null default '' VARCHAR(45)"`
}

func NewScore() *SScore {
	return new(SScore)
}

func MoreScore() []*SScore {
	return make([]*SScore, 0)
}
