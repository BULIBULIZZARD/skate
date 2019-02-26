package models

type MatchScore struct {
	SScore `xorm:"extends"`
	SMatch `xorm:"extends"`
}

func NewMatchScore() *MatchScore {
	return new(MatchScore)
}

func MoreMatchScore() []*MatchScore {
	return make([]*MatchScore, 0)
}
