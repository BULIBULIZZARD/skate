package Organize

import (
	"file/skate/config"
	"file/skate/data"
	"file/skate/redis"
	"file/skate/tools"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"time"
)

type Organize struct {
	id string
}

func NewOrganizeServer() *Organize {
	return new(Organize)
}

func (o *Organize) GetAllPlayer(c echo.Context) error {
	if !o.checkToken(c) {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "fail",
		})
	}
	model := data.NewOrganizeModel()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": model.GetAllPlayerById(o.id),
	})
}

func (o *Organize) GetAllPlayerScore(c echo.Context) error {
	if !o.checkToken(c) {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "fail",
		})
	}
	page, _ := strconv.Atoi(c.FormValue("page"))
	if page < 1 {
		page = 1
	}
	pageNum := data.NewOrganizeModel().GetAllPlayerScorePageNum(o.searchCond(c))
	if pageNum < page {
		page = pageNum
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":     data.NewOrganizeModel().GetAllPlayerScore(o.searchCond(c), page-1),
		"page":     page,
		"page_num": pageNum,
	})
}

func (o *Organize) GetOrganizeBestScore(c echo.Context) error {
	if !o.checkToken(c) {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "fail",
		})
	}
	model := data.NewOrganizeModel()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"4圈":    model.GetBestMatchScore(o.id, "4圈"),
		"7圈":    model.GetBestMatchScore(o.id, "7圈"),
		"500米":  model.GetBestMatchScore(o.id, "500米"),
		"1000米": model.GetBestMatchScore(o.id, "1000米"),
		"1500米": model.GetBestMatchScore(o.id, "1500米"),
	})
}

func (o *Organize) GetTreeData(c echo.Context) error {
	if !o.checkToken(c) {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "fail",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": o.buildTreeData(),
	})
}

func (o *Organize) GetPieData(c echo.Context) error {
	if !o.checkToken(c) {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "fail",
		})
	}
	name := []string{"4圈", "7圈", "500米", "1000米", "1500米"}
	count := []int{
		data.NewOrganizeModel().GetMatchCountById(o.id, "4圈"),
		data.NewOrganizeModel().GetMatchCountById(o.id, "7圈"),
		data.NewOrganizeModel().GetMatchCountById(o.id, "500米"),
		data.NewOrganizeModel().GetMatchCountById(o.id, "1000米"),
		data.NewOrganizeModel().GetMatchCountById(o.id, "1500米"),
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"name":  name,
		"value": count,
	})
}

func (o *Organize) OrganizeLogin(c echo.Context) error {
	username := c.FormValue("username")
	psw := c.FormValue("password")
	model := data.NewOrganizeModel()
	organize, flag := model.CheckOrganizeLogin(username, tools.NewTools().Sha1(psw))
	if flag {
		token := tools.NewTools().Sha1(config.GetConfig().GetSalt() +
			fmt.Sprintf("%d", organize.Id) +
			fmt.Sprintf("%d", time.Now().Unix()))
		err := redis.NewRedis().SetValue(config.GetConfig().GetOrganizePre()+
			fmt.Sprintf("%d", organize.Id),
			tools.NewTools().Sha1(token+config.GetConfig().GetSalt()), "259200")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "redis服务错误 ",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"flag":    flag,
			"id":      organize.Id,
			"name":    organize.OrganizeName,
			"token":   token,
			"message": "OK",
		})
	} else {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"flag":    flag,
			"message": "用户名或密码错误",
		})
	}
}

func (o *Organize) checkToken(c echo.Context) bool {
	id := c.FormValue("id")
	token := c.FormValue("token")
	if id == "" || token == "" {
		return false
	}
	cacheData, err := redis.NewRedis().GetValue(config.GetConfig().GetOrganizePre() + id)
	if err != nil {
		return false
	}
	if cacheData == "" {
		return false
	}
	o.id = id
	return cacheData == tools.NewTools().Sha1(token+config.GetConfig().GetSalt())
}

func (o *Organize) searchCond(c echo.Context) string {
	where := " s_organize.id = " + o.id
	playerId := c.FormValue("player_id")
	matchName := c.FormValue("match_name")

	if playerId != "" {
		where += "  and s_score.player_id=" + playerId
	}
	if matchName != "" {
		where += " and s_match.match_name like '" + matchName + "%'"
	}
	return where
}

func (o *Organize) buildTreeData() map[string]interface{} {
	model := data.NewOrganizeModel()
	var tree [] map[string]interface{}
	var boy [] map[string]interface{}
	var girl [] map[string]interface{}
	for _, v := range model.GetAllPlayerById(o.id) {
		if v.PlayerGender == "男子" {
			boy = append(boy, map[string]interface{}{
				"name":  v.PlayerName,
				"value": v.Id,
			})
		} else {
			girl = append(girl, map[string]interface{}{
				"name":  v.PlayerName,
				"value": v.Id,
			})
		}
	}
	tree = append(tree, map[string]interface{}{
		"name":     "男",
		"children": boy,
	})
	tree = append(tree, map[string]interface{}{
		"name":     "女",
		"children": girl,
	})
	return map[string]interface{}{
		"name":     model.GetOrganizeNameById(o.id).OrganizeName,
		"children": tree,
	}
}

func (o *Organize) ChangePassword(c echo.Context) error {
	if !o.checkToken(c) {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "fail",
		})
	}
	ordPass := c.FormValue("ord")
	newPass := c.FormValue("new")
	if newPass != c.FormValue("re") {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "两次密码不一致",
		})
	}
	flag := data.NewOrganizeModel().OrganizeChangePassword(o.id, ordPass, newPass)
	if !flag {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "密码错误",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "修改成功",
	})
}
