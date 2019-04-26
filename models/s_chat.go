package models

type SChat struct {
	Id         int    `json:"id" xorm:"not null pk autoincr unique INT(11)"`
	Message    string `json:"message" xorm:"not null default '' VARCHAR(255)"`
	FromId     int    `json:"form_id" xorm:"not null default 0 INT(11)"`
	ToId       int    `json:"to_id" xorm:"not null default 0 INT(11)"`
	ReadStatus int    `json:"read_status" xorm:"not null default 0 TINYINT(1)"`
	CreateTime int    `json:"create_time" xorm:"not null default 0 INT(11)"`
}
func NewChat() *SChat {
	return new(SChat)
}

func MoreChat() []*SChat {
	return make([]*SChat, 0)
}