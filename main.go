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
	//index route
	e.GET(config.GetConfig().GetVersion()+"/index/getContest", index.NewIndexServer().GetIndexContest)
	e.GET(config.GetConfig().GetVersion()+"/index/getContestMatch/:cid", index.NewIndexServer().GetContestMatch)
	e.GET(config.GetConfig().GetVersion()+"/index/getMatchScore/:mid/:group", index.NewIndexServer().GetMatchScore)

	//player route
	e.GET(config.GetConfig().GetVersion()+"/player/getShowData", player.NewPlayerServer().GetScoreByName)
	e.GET(config.GetConfig().GetVersion()+"/player/getPlayerScore", player.NewPlayerServer().GetPlayerScore)
	e.GET(config.GetConfig().GetVersion()+"/player/getPlayerBestScore", player.NewPlayerServer().GetPlayerBestScore)
	e.POST(config.GetConfig().GetVersion()+"/player/login",player.NewPlayerServer().PlayerLogin)
	e.POST(config.GetConfig().GetVersion()+"/player/changePassword",player.NewPlayerServer().ChangePassword)

	//organize  route
	e.GET(config.GetConfig().GetVersion()+"/organize/getPieData", Organize.NewOrganizeServer().GetPieData)
	e.GET(config.GetConfig().GetVersion()+"/organize/getAllPlayer", Organize.NewOrganizeServer().GetAllPlayer)
	e.GET(config.GetConfig().GetVersion()+"/organize/getAllScore", Organize.NewOrganizeServer().GetAllPlayerScore)
	e.GET(config.GetConfig().GetVersion()+"/organize/getBestScore", Organize.NewOrganizeServer().GetOrganizeBestScore)
	e.GET(config.GetConfig().GetVersion()+"/organize/getTreeData", Organize.NewOrganizeServer().GetTreeData)
	e.POST(config.GetConfig().GetVersion()+"/organize/login", Organize.NewOrganizeServer().OrganizeLogin)

	//CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	//
	//Error handler
	//e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
	//	return func(c echo.Context) error {
	//		// Extract the credentials from HTTP request header and perform a security
	//		// check
	//		// For invalid credentials
	//		return echo.NewHTTPError(http.StatusInternalServerError)
	//
	//		// For valid credentials call next
	//		// return next(c)
	//	}
	//})
	//server listening port 8000
	e.Logger.Fatal(e.Start(":8000"))
}
