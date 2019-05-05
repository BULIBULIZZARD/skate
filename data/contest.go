package data

import (
	"file/skate/config"
	"file/skate/models"
	"file/skate/sql"
	"github.com/labstack/gommon/log"
)

type ContestModel struct {
}

func NewContestModel() *ContestModel {
	return new(ContestModel)
}
func (m *ContestModel) GetAllContest() interface{} {
	engine := sql.GetSqlEngine()
	data := models.MoreContest()
	err := engine.Asc("id").Find(&data)
	if err != nil {
		log.Print(err.Error())
	}
	return data
}

func (m *ContestModel) GetArticleList(page int) []*models.SArticle {
	engine := sql.GetSqlEngine()
	data := models.MoreArticle()
	err := engine.Where("status = 1").
		Cols("id,title,create_time").
		Asc(`id`).
		Limit(config.GetConfig().GetPageSize(), page*config.GetConfig().GetPageSize()).
		Find(&data)
	if err != nil {
		log.Print(err.Error())
	}
	return data
}

func (m *ContestModel) GetArticleCount() int {
	engine := sql.GetSqlEngine()
	data := models.NewArticle()
	count, err := engine.Where("status = 1").
		Cols("id,title,create_time").
		Asc(`id`).
		Count(data)
	if err != nil {
		log.Print(err.Error())
	}
	if int(count)%config.GetConfig().GetPageSize() > 0 {
		return (int(count) / config.GetConfig().GetPageSize()) + 1
	}
	return int(count) / config.GetConfig().GetPageSize()
}

func (m *ContestModel) GetArticleContentById(id string) *models.SArticle {
	engine := sql.GetSqlEngine()
	data := models.NewArticle()
	_, err := engine.Where("status = 1 and id = ?", id).Get(data)
	if err != nil {
		log.Print(err.Error())
	}
	return data
}
