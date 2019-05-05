package models

type SArticle struct {
	Id         int    `json:"id" xorm:"not null pk autoincr unique INT(11)"`
	Title      string `json:"title" xorm:"not null VARCHAR(45)"`
	Content    string `json:"content" xorm:"not null TEXT"`
	Status     int    `json:"status" xorm:"not null TINYINT(1)"`
	CreateTime int    `json:"create_time" xorm:"not null INT(11)"`
}

func NewArticle() *SArticle {
	return new(SArticle)
}

func MoreArticle() []*SArticle {
	return make([]*SArticle, 0)
}