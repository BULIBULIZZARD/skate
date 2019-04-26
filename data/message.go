package data

import (
	"file/skate/models"
	"file/skate/sql"
	"github.com/labstack/gommon/log"
	"strconv"
	"time"
)

type MessageModel struct {
}

func NewMessageModel() *MessageModel {
	return new(MessageModel)
}

func (m *MessageModel) SavePlayerChatLog(msg string, from string, to string, status int) int {
	engine := sql.GetSqlEngine()
	chat := models.NewChat()
	var err error
	chat.Message = msg
	chat.FormId, err = strconv.Atoi(from)
	chat.ToId, err = strconv.Atoi(to)
	chat.CreateTime = int(time.Now().Unix())
	chat.ReadStatus = status
	if err != nil {
		log.Print(err.Error())
	}
	flag, err := engine.InsertOne(chat)
	m.GetIsChatting(chat.FormId, chat.ToId)
	return int(flag)
}

func (m *MessageModel) GetIsChatting(from int, to int) {
	//todo 查询 返回id 没有->insert
	id := m.FindChatting(from, to)
	if id > 0 {
		m.ChangeChattingStatus(id)
	} else {
		m.InsChatting(from, to)
	}
}

func (m *MessageModel) InsChatting(from int, to int) {
	engine := sql.GetSqlEngine()
	chatting := models.NewChatting()
	chatting.UserId = to
	chatting.WithId = from
	chatting.NewTime = int(time.Now().Unix())
	chatting.IsNew = 1
	chatting.Status = 1
	_, err := engine.InsertOne(chatting)
	if err != nil {
		log.Print(err.Error())
	}
}

func (m *MessageModel) FindChatting(from int, to int) int {
	engine := sql.GetSqlEngine()
	chatting := models.NewChatting()
	flag, err := engine.Where("with_id = ? and user_id =?", from, to).Get(chatting)
	if err != nil {
		log.Print()
	}
	if flag {
		return chatting.Id
	}
	return 0
}

func (m *MessageModel) ChangeChattingStatus(id int) {
	engine := sql.GetSqlEngine()
	chatting := models.NewChatting()
	chatting.IsNew = 1
	_, err := engine.Id(id).Update(chatting)
	if err != nil {
		log.Print(err.Error())
	}
}

func (m *MessageModel) GetAllChatting(id string) []*models.ChattingPlayer {
	engine := sql.GetSqlEngine()
	chatting := models.MoreChattingPlayer()
	err := engine.
		Table("s_chatting").
		Join("INNER", "s_player", "s_chatting.with_id=s_player.id").
		Where("s_chatting.user_id = ? and s_chatting.status = ?", id, 1).
		Cols("s_chatting.with_id","s_chatting.is_new","s_player.player_name","s_player.organize","s_chatting.new_time").
		Desc("s_chatting.is_new").
		Desc("s_chatting.new_time").
		Find(&chatting)
	if err != nil {
		log.Print(err.Error())
	}
	return chatting
}

func (m *MessageModel) GetAllChatLog(id string) []*models.SChat {
	engine := sql.GetSqlEngine()
	chat := models.MoreChat()
	err := engine.Where("form_id = ? or to_id = ?", id, id).Find(&chat)
	if err != nil {
		log.Print(err.Error())
	}
	return chat
}
