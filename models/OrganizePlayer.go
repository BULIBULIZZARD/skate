package models

type OrganizePlayer struct {
	SOrganize `xorm:"extends"`
	SPlayer   `xorm:"extends"`
}

func NewOrganizePlayer() *OrganizePlayer  {
	return new(OrganizePlayer)
}

func MoreOrganizePlayer() []*OrganizePlayer  {
	return make([]*OrganizePlayer,0)
}