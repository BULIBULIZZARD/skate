package main

import (
	"file/skate/config"
	"file/skate/server/Organize"
	"file/skate/server/Player"
	"file/skate/server/index"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	//index page route
	e.GET(config.GetConfig().GetVersion()+"/index/getContest", index.NewIndexServer().GetIndexContest)
	e.GET(config.GetConfig().GetVersion()+"/index/getContestMatch/:cid", index.NewIndexServer().GetContestMatch)
	e.GET(config.GetConfig().GetVersion()+"/index/getMatchScore/:mid/:group", index.NewIndexServer().GetMatchScore)

	//player page route
	e.GET(config.GetConfig().GetVersion()+"/player/getPlayerScore/:pid", player.NewPlayerServer().GetPlayerScore)
	e.GET(config.GetConfig().GetVersion()+"/player/getPlayerBestScore/:pid", player.NewPlayerServer().GetPlayerBestScore)
	e.POST(config.GetConfig().GetVersion()+"/player/login",player.NewPlayerServer().PlayerLogin)

	//organize page route
	e.GET(config.GetConfig().GetVersion()+"/player/getAllPlayer/:oid", Organize.NewOrganizeServer().GetAllPlayer)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	//server listening port 8000
	e.Logger.Fatal(e.Start(":8000"))
}
