package models

type OrganizePlayerScore struct {
	SOrganize `xorm:"extends"`
	SPlayer   `xorm:"extends"`
	SScore    `xorm:"extends"`
	SMatch    `xorm:"extends"`
}

func MoreOrganizePlayerScore() []*OrganizePlayerScore {
	return make([]*OrganizePlayerScore, 0)
}
