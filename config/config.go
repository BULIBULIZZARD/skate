package config

import "github.com/labstack/echo"

type database struct {
	Driver  string
	Address string
}

type config struct {
}

func GetConfig() *config {

	return new(config)
}

func (c *config) GetDatabase() database {
	return database{
		Driver:  `mysql`,
		Address: `root:root@/skate?charset=utf8`,
	}
}

func (c *config) GetVersion() string {
	return "/v1"
}

func (c *config) SetAccessOriginUrl(contest echo.Context) {
	contest.Response().Header().Set("Access-Control-Allow-Origin", "*")
	contest.Response().Header().Set("Access-Control-Allow-Credentials","true")
}

func (c *config) GetCacheTableName() string {
	return "myCache"
}

func (c *config) GetSalt() string {
	return "_skate_"
}
func (c *config) GetCachePre() string {
	return "player"
}
