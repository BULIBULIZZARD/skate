package main

import (
	"file/skate/data"
	"file/skate/websocket"
)

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
	//json字符中的"引号，需用\进行转义，否则编译出错
	message := websocket.Message{
		From:"1",
		Msg:"221",
		To:"2",
	}
	flag :=data.NewPlayerModel().SavePlayerChatLog(message.Msg,message.From,message.To)
	println(flag)
}
