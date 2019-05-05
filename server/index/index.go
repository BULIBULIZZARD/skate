package index

import (
	"file/skate/data"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type Index struct {
}

func NewIndexServer() *Index {
	return new(Index)
}

func (x *Index) GetIndexContest(c echo.Context) error {
	model := data.NewContestModel()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": model.GetAllContest(),
	})
}
func (x *Index) GetContestMatch(c echo.Context) error {
	cid := c.FormValue("cid")
	model := data.NewMatchModel()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": model.GetMatchByContestId(cid),
	})
}
func (x *Index) GetMatchScore(c echo.Context) error {
	mid := c.FormValue("id")
	group := c.FormValue("group")
	model := data.NewScoreModel()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": model.GetScoreByMatchAndGroup(mid, group),
	})
}

func (x *Index) GetArticleList(c echo.Context) error {

	page, _ := strconv.Atoi(c.FormValue("page"))
	if page < 1 {
		page = 1
	}
	pageNum := data.NewContestModel().GetArticleCount()
	if pageNum < page {
		page = pageNum
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":     data.NewContestModel().GetArticleList(page - 1),
		"page":     page,
		"page_num": pageNum,
	})
}


func (x *Index) GetArticle(c echo.Context) error {
	id := c.FormValue("id")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":     data.NewContestModel().GetArticleContentById(id),
	})
}