package index

import (
	"github.com/labstack/echo"
	"net/http"
)

type Index struct {

}

func NewIndexServer() *Index  {
	return new(Index)
}


func (x *Index)GetAllContest(c echo.Context) error {
	return c.Render(http.StatusOK, "template.html", map[string]interface{}{
		"name":     "Dolly!222",
		"document": "test______",
	})
}
