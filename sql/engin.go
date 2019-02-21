package sql

import (
	"file/skate/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/labstack/gommon/log"
)

func GetSqlEngine() *xorm.Engine {

	engine, err := xorm.NewEngine(config.GetConfig().GetDatabase().Driver, config.GetConfig().GetDatabase().Address)
	if err != nil {
		log.Print(err.Error())
	}
	return engine
}
