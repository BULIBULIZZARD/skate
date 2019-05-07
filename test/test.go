package main

import "strings"

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
	lll :=strings.Count("123123123","")
	println(lll-1)
}
