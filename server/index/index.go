package index

import (
	"github.com/labstack/echo"
	"net/http"
)

type Index struct {
}

func NewIndexServer() *Index {
	return new(Index)
}

func (x *Index) GetAllContest(c echo.Context) error {
	a := []int{1, 2, 3}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"name":     "Dolly!222",
		"document": "test______",
		"test":     a,
	})
}
