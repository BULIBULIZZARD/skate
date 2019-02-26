package index

import (
	"file/skate/config"
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

func (x *Index) GetIndexContest(c echo.Context) error {
	model := data.NewContestModel()
	config.GetConfig().SetAccessOriginUrl(c)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": model.GetAllContest(),
	})
}
func (x *Index) GetContestMatch(c echo.Context) error {
	cid := c.Param("cid")
	model := data.NewMatchModel()
	config.GetConfig().SetAccessOriginUrl(c)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": model.GetMatchByContestId(cid),
	})
}
func (x *Index) GetMatchScore(c echo.Context) error {
	mid := c.Param("mid")
	group := c.Param("group")
	model := data.NewScoreModel()
	config.GetConfig().SetAccessOriginUrl(c)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": model.GetScoreByMatchAndGroup(mid, group),
	})
}
