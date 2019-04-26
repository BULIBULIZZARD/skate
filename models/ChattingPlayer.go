package models

type ChattingPlayer struct {
	SChatting `xorm:"extends"`
	SPlayer   `xorm:"extends"`
}

func NewChattingPlayer() *ChattingPlayer {
	return new(ChattingPlayer)
}

func MoreChattingPlayer() []*ChattingPlayer {
	return make([]*ChattingPlayer, 0)
}
