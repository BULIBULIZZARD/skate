package player

import (
	"file/skate/config"
	"file/skate/data"
	"file/skate/models"
	"file/skate/redis"
	"file/skate/tools"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type Player struct {
	id string
}

func NewPlayerServer() *Player {
	return new(Player)
}

func (p *Player) GetPlayerScore(c echo.Context) error {
	if !p.checkToken(c) {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "fail",
		})
	}
	pid := p.id
	score := data.NewScoreModel().GetScoreByPlayerAndOrganize(pid)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": score,
	})
}
func (p *Player) GetPlayerBestScore(c echo.Context) error {
	if !p.checkToken(c) {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "fail",
		})
	}
	pid := p.id
	model := data.NewPlayerModel()
	m500 := model.GetBestScoreById(pid, "500米")
	m1000 := model.GetBestScoreById(pid, "1000米")
	m1500 := model.GetBestScoreById(pid, "1500米")
	m4 := model.GetBestScoreById(pid, "4圈")
	m7 := model.GetBestScoreById(pid, "7圈")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"0": p.checkBestScore(m500),
		"1": p.checkBestScore(m1000),
		"2": p.checkBestScore(m1500),
		"3": p.checkBestScore(m4),
		"4": p.checkBestScore(m7),
	})
}

func (p *Player) GetScoreByName(c echo.Context) error {
	if !p.checkToken(c) {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "fail",
		})
	}
	pid := p.id
	model := data.NewPlayerModel()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"4圈": p.buildEcharsData( model.GetAllScoreByMatchAndPlayer(pid, "4圈")),
		"7圈":    p.buildEcharsData(model.GetAllScoreByMatchAndPlayer(pid, "7圈")),
		"500米":  p.buildEcharsData(model.GetAllScoreByMatchAndPlayer(pid, "500米")),
		"1000米": p.buildEcharsData(model.GetAllScoreByMatchAndPlayer(pid, "1000米")),
		"1500米": p.buildEcharsData(model.GetAllScoreByMatchAndPlayer(pid, "1500米")),
	})
}

func (p *Player) PlayerLogin(c echo.Context) error {
	username := c.FormValue("username")
	psw := c.FormValue("password")
	model := data.NewPlayerModel()
	playerData, flag := model.PlayerLoginCheck(username, tools.NewTools().Sha1(psw))
	if flag {
		token := tools.NewTools().Sha1(config.GetConfig().GetSalt() +
			fmt.Sprintf("%d", playerData.Id) +
			fmt.Sprintf("%d", time.Now().Unix()))
		err := redis.NewRedis().SetValue(config.GetConfig().GetCachePre()+
			fmt.Sprintf("%d", playerData.Id),
			tools.NewTools().Sha1(token+config.GetConfig().GetSalt()), "259200")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "redis服务错误 ",
			})
		}
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
	cacheData, err := redis.NewRedis().GetValue(config.GetConfig().GetCachePre() + id)
	if err != nil {
		return false
	}
	if cacheData == "" {
		return false
	}
	p.id = id
	return cacheData == tools.NewTools().Sha1(token+config.GetConfig().GetSalt())
}

func (p *Player) checkBestScore(m *models.MatchScore) interface{} {
	if m.TimeScore == "00.00.00" || m.TimeScore == "" || m.TimeScore == "完成比赛" {
		return "filter"
	}
	return m
}

func (p *Player) buildEcharsData(data []*models.MatchScore) interface{} {
	count := 0
	var value []string
	var group []string
	var matchId []int
	var date []string
	var matchType []string
	for _, m := range data {
		count++
		value = append(value, "0000-00-00 00:"+m.TimeScore)
		group = append(group, m.SGroup)
		matchId = append(matchId, m.MatchId)
		date = append(date, m.Date)
		matchType = append(matchType, m.MatchType)
	}
	if count == 0 {
		return ""
	}
	result := map[string]interface{}{
		"count":      count,
		"value":      value,
		"group":      group,
		"match_id":   matchId,
		"date":       date,
		"match_type": matchType,
	}
	return result
}
