package models

type SFollow struct {
	Id         int `json:"id" xorm:"not null pk autoincr unique INT(11)"`
	FanId     int `json:"fan_id" xorm:"not null default 0 INT(11)"`
	UserId     int `json:"user_id" xorm:"not null default 0 INT(11)"`
	Status     int `json:"status" xorm:"not null default 1 TINYINT(1)"`
	CreateTime int `json:"create_time" xorm:"not null default 0 INT(11)"`
}

func NewFollow() *SFollow {
	return new(SFollow)
}
func MoreFollow() []*SFollow {
	return make([]*SFollow, 0)
}
