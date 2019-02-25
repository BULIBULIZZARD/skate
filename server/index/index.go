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

func (x *Index) GetIndexContest(c echo.Context) error {
	model := data.NewContestModel()
	setAccessOriginUrl(c)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": model.GetAllContest(),
	})
}
func (x *Index) GetContestMatch(c echo.Context) error {
	cid := c.Param("cid")
	model := data.NewMatchModel()
	setAccessOriginUrl(c)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": model.GetMatchByContestId(cid),
	})
}
func (x *Index) GetMatchScore(c echo.Context) error {
	mid := c.Param("mid")
	group := c.Param("group")
	model := data.NewSocreModel()
	setAccessOriginUrl(c)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": model.GetScoreByMatchAndGroup(mid, group),
	})
}
func setAccessOriginUrl(contest echo.Context) {
	contest.Response().Header().Set("Access-Control-Allow-Origin", "*")
}