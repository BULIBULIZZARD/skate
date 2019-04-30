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
	"strconv"
	"strings"
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
	page, _ := strconv.Atoi(c.FormValue("page"))
	if page < 1 {
		page = 1
	}
	pageNum := data.NewScoreModel().GetScoreCountByPlayerId(p.id)
	if pageNum < page {
		page = pageNum
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":     data.NewScoreModel().GetScoreByPlayerId(p.id, page-1),
		"page":     page,
		"page_num": pageNum,
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
		"4圈":    p.buildEcharsData(model.GetAllScoreByMatchAndPlayer(pid, "4圈")),
		"7圈":    p.buildEcharsData(model.GetAllScoreByMatchAndPlayer(pid, "7圈")),
		"500米":  p.buildEcharsData(model.GetAllScoreByMatchAndPlayer(pid, "500米")),
		"1000米": p.buildEcharsData(model.GetAllScoreByMatchAndPlayer(pid, "1000米")),
		"1500米": p.buildEcharsData(model.GetAllScoreByMatchAndPlayer(pid, "1500米")),
	})
}

func (p *Player) CheckPlayerStatus(c echo.Context) error {
	if !p.checkToken(c) {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "fail",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "OK",
	})
}

func (p *Player) GetPlayerChatting(c echo.Context) error {
	if !p.checkToken(c) {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "fail",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": data.NewMessageModel().GetAllChatting(p.id),
	})
}

func (p *Player) GetPlayerChat(c echo.Context) error {
	if !p.checkToken(c) {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "fail",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": data.NewMessageModel().GetAllChatLog(p.id),
	})
}

func (p *Player) GetPlayerNameAndOrganizeById(c echo.Context) error {
	if !p.checkToken(c) {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "fail",
		})
	}
	playerId := c.FormValue("with_id")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": data.NewPlayerModel().GetPlayerNameAndOrganizeById(playerId),
	})
}

func (p *Player) ChangeChattingIsNew(c echo.Context) error {
	if !p.checkToken(c) {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "fail",
		})
	}
	playerId := c.FormValue("with_id")
	data.NewMessageModel().ChangeChattingIsNew(p.id, playerId)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "OK",
	})
}

func (p *Player) CloseChatting(c echo.Context) error {
	if !p.checkToken(c) {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "fail",
		})
	}
	playerId := c.FormValue("with_id")
	data.NewMessageModel().CloseChatting(p.id, playerId)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "OK",
	})
}

func (p *Player) ReadChatMessage(c echo.Context) error {
	if !p.checkToken(c) {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "fail",
		})
	}
	playerId := c.FormValue("with_id")
	data.NewMessageModel().ReadChatMessage(p.id, playerId)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "OK",
	})
}

func (p *Player) ChattingStatus(c echo.Context) error {
	if !p.checkToken(c) {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "fail",
		})
	}

	playerId := c.FormValue("with_id")
	id, _ := strconv.Atoi(p.id)
	with, _ := strconv.Atoi(playerId)
	data.NewMessageModel().GetIsChatting(with, id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "OK",
	})
}

func (p *Player) PlayerFollow(c echo.Context) error {
	if !p.checkToken(c) {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "fail",
		})
	}
	playerId := c.FormValue("player_id")
	data.NewFollowModel().InsOrUpdate(playerId, p.id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "OK",
	})
}
func (p *Player) PlayerFanList(c echo.Context) error {
	if !p.checkToken(c) {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "fail",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": data.NewFollowModel().PlayerFunList(p.id),
	})
}

func (p *Player) PlayerFollowList(c echo.Context) error {
	if !p.checkToken(c) {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "fail",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": data.NewFollowModel().PlayerFollowList(p.id),
	})
}
func (p *Player) PlayerNopeFan(c echo.Context) error {
	if !p.checkToken(c) {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "fail",
		})
	}

	playerId := c.FormValue("player_id")
	data.NewFollowModel().NopeFan(playerId, p.id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "OK",
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
		err := redis.NewRedis().SetValue(config.GetConfig().GetPlayerPre()+
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

func (p *Player) ChangePassword(c echo.Context) error {
	if !p.checkToken(c) {
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
	if ordPass != newPass {
		flag := data.NewPlayerModel().PlayerChangePassword(p.id, ordPass, newPass)
		if !flag {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "密码错误",
			})
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "修改成功",
	})
}

func (p *Player) checkToken(c echo.Context) bool {
	id := c.FormValue("id")
	token := c.FormValue("token")
	if id == "" || token == "" {
		return false
	}
	cacheData, err := redis.NewRedis().GetValue(config.GetConfig().GetPlayerPre() + id)
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
		group = append(group, strings.Replace(strings.Replace(m.SGroup, "第", "", 1), "组", "", 1))
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
