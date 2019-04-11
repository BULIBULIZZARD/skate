package Organize

import (
	"file/skate/config"
	"file/skate/data"
	"file/skate/redis"
	"file/skate/tools"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type Organize struct {
}

func NewOrganizeServer() *Organize {
	return new(Organize)
}

func (o *Organize) GetAllPlayer(c echo.Context) error {
	oid := c.Param("oid")
	model := data.NewOrganizeModel()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": model.GetAllPlayerById(oid),
	})
}

func (o *Organize)OrganizeLogin(c echo.Context)  error{
	username := c.FormValue("username")
	psw := c.FormValue("password")
	model := data.NewOrganizeModel()
	organize, flag := model.CheckOrganizeLogin(username, tools.NewTools().Sha1(psw))
	if flag {
		token := tools.NewTools().Sha1(config.GetConfig().GetSalt() +
			fmt.Sprintf("%d", organize.Id) +
			fmt.Sprintf("%d", time.Now().Unix()))
		err := redis.NewRedis().SetValue(config.GetConfig().GetOrganizePre()+
			fmt.Sprintf("%d", organize.Id),
			tools.NewTools().Sha1(token+config.GetConfig().GetSalt()), "259200")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "redis服务错误 ",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"flag":    flag,
			"id":      organize.Id,
			"name":    organize.OrganizeName,
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
