package main

import "file/skate/data"

type StuRead struct {
	Name  interface{} `json:"name"`
	Age   interface{}
	HIgh  interface{}
	sex   interface{}
	Class interface{} `json:"class"`
	Test  interface{}
}
type Message struct {
	To   string
	Msg  string
	From string
}
type Class struct {
	Name  string
	Grade int
}

func main() {
	p:=data.NewPlayerModel().GetPlayerNameAndOrganizeById("2")
	println(p.PlayerName,p.Organize)
}
