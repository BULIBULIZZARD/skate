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
	id string
}

func NewOrganizeServer() *Organize {
	return new(Organize)
}

func (o *Organize) GetAllPlayer(c echo.Context) error {
	if !o.checkToken(c) {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "fail",
		})
	}
	model := data.NewOrganizeModel()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": model.GetAllPlayerById(o.id),
	})
}

func (o *Organize) GetAllPlayerScore(c echo.Context) error {
	if !o.checkToken(c) {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "fail",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": data.NewOrganizeModel().GetAllPlayerScore(o.id),
	})
}

func (o *Organize) GetOrganizeBestScore(c echo.Context) error {
	if !o.checkToken(c) {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "fail",
		})
	}
	model := data.NewOrganizeModel()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"4圈":  model.GetBestMatchScore(o.id, "4圈"),
		"7圈":     model.GetBestMatchScore(o.id, "7圈"),
		"500米":   model.GetBestMatchScore(o.id, "500米"),
		"1000米":  model.GetBestMatchScore(o.id, "1000米"),
		"1500米":  model.GetBestMatchScore(o.id, "1500米"),
	})
}




func (o *Organize) OrganizeLogin(c echo.Context) error {
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

func (o *Organize) checkToken(c echo.Context) bool {
	id := c.FormValue("id")
	token := c.FormValue("token")
	if id == "" || token == "" {
		return false
	}
	cacheData, err := redis.NewRedis().GetValue(config.GetConfig().GetOrganizePre() + id)
	if err != nil {
		return false
	}
	if cacheData == "" {
		return false
	}
	o.id = id
	return cacheData == tools.NewTools().Sha1(token+config.GetConfig().GetSalt())
}