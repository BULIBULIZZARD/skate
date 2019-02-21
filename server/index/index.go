package index

import (
	"file/skate/data"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"net/http"
)

type Index struct {
}

func NewIndexServer() *Index {
	return new(Index)
}



func (x *Index) GetAllContest(c echo.Context) error {
	model := data.NewContestModel()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": model.GetAllContest(),
	})
}
