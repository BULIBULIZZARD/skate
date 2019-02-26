package main

import (
	"file/skate/config"
	"file/skate/server/Organize"
	"file/skate/server/Player"
	"file/skate/server/index"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	//index page route
	e.GET(config.GetConfig().GetVersion()+"/index/getContest", index.NewIndexServer().GetIndexContest)
	e.GET(config.GetConfig().GetVersion()+"/index/getContestMatch/:cid", index.NewIndexServer().GetContestMatch)
	e.GET(config.GetConfig().GetVersion()+"/index/getMatchScore/:mid/:group", index.NewIndexServer().GetMatchScore)

	//player page route
	e.GET(config.GetConfig().GetVersion()+"/index/getPlayerScore/:pid", player.NewPlayerServer().GetPlayerScore)
	e.GET(config.GetConfig().GetVersion()+"/index/getPlayerBestScore/:pid", player.NewPlayerServer().GetPlayerBestScore)

	//organize page route
	e.GET(config.GetConfig().GetVersion()+"/index/getAllPlayer/:oid", Organize.NewOrganizeServer().GetAllPlayer)

	//server listening port 8000
	e.Logger.Fatal(e.Start(":8000"))
}
