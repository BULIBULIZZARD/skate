package tools

import (
	"crypto/sha1"
	"fmt"
	"github.com/labstack/gommon/log"
	"io"
)

type Tools struct {
}

func NewTools() *Tools {
	return new(Tools)
}

func (o *Tools) Sha1(data string) string {
	t := sha1.New()
	_, err := io.WriteString(t, data)
	if err != nil {
		log.Print(err.Error())
	}
	return fmt.Sprintf("%x", t.Sum(nil))
}


