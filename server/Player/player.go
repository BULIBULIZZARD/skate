package player

import (
	"file/skate/cache"
	"file/skate/config"
	"file/skate/data"
	"file/skate/models"
	"file/skate/tools"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type Player struct {
}

func NewPlayerServer() *Player {
	return new(Player)
}

func (p *Player) GetPlayerScore(c echo.Context) error {
	config.GetConfig().SetAccessOriginUrl(c)
	if !p.checkToken(c) {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "fail",
		})
	}
	pid := c.Param("pid")
	player := data.NewPlayerModel().GetNameOrganizeById(pid)
	score := data.NewScoreModel().GetScoreByPlayerAndOrganize(player.PlayerName, player.Organize)
	config.GetConfig().SetAccessOriginUrl(c)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": score,
	})
}
func (p *Player) GetPlayerBestScore(c echo.Context) error {
	config.GetConfig().SetAccessOriginUrl(c)

	if !p.checkToken(c) {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "fail",
		})
	}
	pid := c.Param("pid")
	m500 := data.NewPlayerModel().GetBestScoreById(pid, "500米")
	m1000 := data.NewPlayerModel().GetBestScoreById(pid, "1000米")
	m1500 := data.NewPlayerModel().GetBestScoreById(pid, "1500米")
	m4 := data.NewPlayerModel().GetBestScoreById(pid, "4圈")
	m7 := data.NewPlayerModel().GetBestScoreById(pid, "7圈")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"0": p.checkBestScore(m500),
		"1": p.checkBestScore(m1000),
		"2": p.checkBestScore(m1500),
		"3": p.checkBestScore(m4),
		"4": p.checkBestScore(m7),
	})
}

func (p *Player) PlayerLogin(c echo.Context) error {
	username := c.FormValue("username")
	psw := c.FormValue("password")
	model := data.NewPlayerModel()
	playerData, flag := model.PlayerLoginCheck(username, tools.NewTools().Sha1(psw))
	config.GetConfig().SetAccessOriginUrl(c)
	if flag {
		token := tools.NewTools().Sha1(config.GetConfig().GetSalt() +
			fmt.Sprintf("%d", playerData.Id) +
			fmt.Sprintf("%d", time.Now().Unix()))
		cache.NewCache().SetCache(config.GetConfig().GetCachePre()+
			fmt.Sprintf("%d", playerData.Id),
			tools.NewTools().Sha1(token+config.GetConfig().GetSalt()))
		return c.JSON(http.StatusOK, map[string]interface{}{
			"flag":    flag,
			"id":      playerData.Id,
			"name":    playerData.PlayerName,
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

func (p *Player) checkToken(c echo.Context) bool {
	id := c.FormValue("id")
	token := c.FormValue("token")
	if id == "" || token == "" {
		return false
	}
	cacheData := cache.NewCache().GetCache(config.GetConfig().GetCachePre() + id)
	if cacheData=="" {
		return false
	}
	return cacheData == tools.NewTools().Sha1(token+config.GetConfig().GetSalt())
}

func (p *Player) checkBestScore(m *models.MatchScore) interface{} {
	if m.TimeScore == "" || m.TimeScore == "完成比赛" || m.TimeScore == "弃权" || m.TimeScore == "犯规" || m.TimeScore == "退赛" ||
		m.TimeScore == "判进" || m.TimeScore == "伤病" || m.TimeScore == "黄牌" {
		return "filter"
	}
	return m
}
