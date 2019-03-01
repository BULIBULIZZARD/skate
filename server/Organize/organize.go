package Organize

import (
	"file/skate/data"
	"github.com/labstack/echo"
	"net/http"
)

type Organize struct {
}

func NewOrganizeServer() *Organize {
	return new(Organize)
}

func (o *Organize) GetAllPlayer(c echo.Context) error {
	oid := c.Param("oid")
	model := data.NewOrganizeModel()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": model.GetAllPlayerById(oid),
	})

}
