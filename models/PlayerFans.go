package models

type PlayerFans struct {
	SPlayer `xorm:"extends"`
	SFollow `xorm:"extends"`
}

func NewPlayerFans() *PlayerFans {
	return new(PlayerFans)
}

func MorePlayerFans() []*PlayerFans {
	return make([]*PlayerFans, 0)
}
