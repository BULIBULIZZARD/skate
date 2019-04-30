package main

import (
	"file/skate/config"
	"file/skate/server/Organize"
	"file/skate/server/Player"
	"file/skate/server/index"
	"file/skate/websocket"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	//index route
	e.GET(config.GetConfig().GetVersion()+"/index/getContest", index.NewIndexServer().GetIndexContest)
	e.GET(config.GetConfig().GetVersion()+"/index/getContestMatch", index.NewIndexServer().GetContestMatch)
	e.GET(config.GetConfig().GetVersion()+"/index/getMatchScore", index.NewIndexServer().GetMatchScore)

	//player route
	e.POST(config.GetConfig().GetVersion()+"/player/login", player.NewPlayerServer().PlayerLogin)
	e.GET(config.GetConfig().GetVersion()+"/player/getShowData", player.NewPlayerServer().GetScoreByName)
	e.GET(config.GetConfig().GetVersion()+"/player/getPlayerScore", player.NewPlayerServer().GetPlayerScore)
	e.POST(config.GetConfig().GetVersion()+"/player/changePassword", player.NewPlayerServer().ChangePassword)
	e.GET(config.GetConfig().GetVersion()+"/player/getPlayerBestScore", player.NewPlayerServer().GetPlayerBestScore)

	//player chat
	e.GET(config.GetConfig().GetVersion()+"/player/chatLog", player.NewPlayerServer().GetPlayerChat)
	e.GET(config.GetConfig().GetVersion()+"/player/chatNow", player.NewPlayerServer().ChattingStatus)
	e.GET(config.GetConfig().GetVersion()+"/player/status", player.NewPlayerServer().CheckPlayerStatus)
	e.GET(config.GetConfig().GetVersion()+"/player/chat", websocket.GetClientManager().WebsocketServer)
	e.GET(config.GetConfig().GetVersion()+"/player/notNew", player.NewPlayerServer().ChangeChattingIsNew)
	e.GET(config.GetConfig().GetVersion()+"/player/chatting", player.NewPlayerServer().GetPlayerChatting)
	e.GET(config.GetConfig().GetVersion()+"/player/closeChatting", player.NewPlayerServer().CloseChatting)
	e.GET(config.GetConfig().GetVersion()+"/player/readChatMessage", player.NewPlayerServer().ReadChatMessage)
	e.GET(config.GetConfig().GetVersion()+"/player/getChatName", player.NewPlayerServer().GetPlayerNameAndOrganizeById)

	//player follow
	e.GET(config.GetConfig().GetVersion()+"/player/nope", player.NewPlayerServer().PlayerNopeFan)
	e.GET(config.GetConfig().GetVersion()+"/player/follow", player.NewPlayerServer().PlayerFollow)
	e.GET(config.GetConfig().GetVersion()+"/player/fanList", player.NewPlayerServer().PlayerFanList)
	e.GET(config.GetConfig().GetVersion()+"/player/followList", player.NewPlayerServer().PlayerFollowList)

	//organize  route
	e.POST(config.GetConfig().GetVersion()+"/organize/login", Organize.NewOrganizeServer().OrganizeLogin)
	e.GET(config.GetConfig().GetVersion()+"/organize/getPieData", Organize.NewOrganizeServer().GetPieData)
	e.GET(config.GetConfig().GetVersion()+"/organize/getTreeData", Organize.NewOrganizeServer().GetTreeData)
	e.GET(config.GetConfig().GetVersion()+"/organize/getAllPlayer", Organize.NewOrganizeServer().GetAllPlayer)
	e.GET(config.GetConfig().GetVersion()+"/organize/getAllScore", Organize.NewOrganizeServer().GetAllPlayerScore)
	e.GET(config.GetConfig().GetVersion()+"/organize/getBestScore", Organize.NewOrganizeServer().GetOrganizeBestScore)

	//CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
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
