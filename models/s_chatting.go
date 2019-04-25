package models

type SChatting struct {
	Id     int `json:"id" xorm:"not null pk autoincr unique INT(10)"`
	UserId int `json:"user_id" xorm:"not null default 0 INT(11)"`
	WithId int `json:"with_id" xorm:"not null default 0 INT(11)"`
	Status int `json:"status" xorm:"not null default 1 TINYINT(1)"`
	IsNew  int `json:"is_new" xorm:"not null default 1 TINYINT(1)"`
	NewTime int `json:"new_time" xorm:"not null default 0 INT(11)"`
}

func NewChatting() *SChatting {
	return new(SChatting)
}

func MoreChatting() []*SChatting {
	return make([]*SChatting, 0)
}

