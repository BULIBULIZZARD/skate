package player

import (
	"file/skate/config"
	"file/skate/data"
	"github.com/labstack/echo"
	"net/http"
)

type Player struct {
}

func NewPlayerServer() *Player {
	return new(Player)
}

func (p *Player) GetPlayerScore(c echo.Context) error {
	pid := c.Param("pid")
	player := data.NewPlayerModel().GetNameOrganizeById(pid)
	score := data.NewScoreModel().GetScoreByPlayerAndOrganize(player.PlayerName, player.Organize)
	config.GetConfig().SetAccessOriginUrl(c)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": score,
	})
}
func (p *Player) GetPlayerBestScore(c echo.Context) error {
	pid := c.Param("pid")
	config.GetConfig().SetAccessOriginUrl(c)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"500": data.NewPlayerModel().GetBestScoreById(pid,"500米"),
		"1000": data.NewPlayerModel().GetBestScoreById(pid,"1000米"),
		"1500": data.NewPlayerModel().GetBestScoreById(pid,"1500米"),
		"4": data.NewPlayerModel().GetBestScoreById(pid,"4圈"),
		"7": data.NewPlayerModel().GetBestScoreById(pid,"7圈"),
	})
}
