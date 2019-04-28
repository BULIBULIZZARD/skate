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
	chat.FromId, err = strconv.Atoi(from)
	chat.ToId, err = strconv.Atoi(to)
	chat.CreateTime = int(time.Now().Unix())
	chat.ReadStatus = status
	if err != nil {
		log.Print(err.Error())
	}
	flag, err := engine.InsertOne(chat)
	m.GetIsChatting(chat.FromId, chat.ToId)
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

func (m *MessageModel) FindChatting(with int, to int) int {
	engine := sql.GetSqlEngine()
	chatting := models.NewChatting()
	flag, err := engine.Where("with_id = ? and user_id =?", with, to).Get(chatting)
	if err != nil {
		log.Print(err)
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
	chatting.Status = 1
	chatting.NewTime = int(time.Now().Unix())
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
		Cols("s_chatting.with_id", "s_chatting.is_new", "s_player.player_name", "s_player.organize", "s_chatting.new_time").
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
	err := engine.Where("from_id = ? or to_id = ?", id, id).Find(&chat)
	if err != nil {
		log.Print(err.Error())
	}
	return chat
}

func (m *MessageModel) ChangeChattingIsNew(id string, with string) int {
	engine := sql.GetSqlEngine()
	chatting := models.NewChatting()
	var err error
	wthId, err := strconv.Atoi(with)
	userId, err := strconv.Atoi(id)
	chatting.IsNew = 0
	chatting.NewTime = int(time.Now().Unix())
	if err != nil {
		log.Print(err)
	}
	chattingID := m.FindChatting(wthId, userId)
	if chattingID == 0 {
		return 0
	}
	_, err = engine.Id(chattingID).Cols("is_new,new_time").Update(chatting)
	if err != nil {
		log.Print(err)
		return 0
	}
	return chattingID
}

func (m *MessageModel) CloseChatting(id string, with string) int {
	engine := sql.GetSqlEngine()
	chatting := models.NewChatting()
	var err error
	wthId, err := strconv.Atoi(with)
	userId, err := strconv.Atoi(id)
	chatting.Status = 0
	if err != nil {
		log.Print(err)
	}
	chattingID := m.FindChatting(wthId, userId)
	if chattingID == 0 {
		return 0
	}
	_, err = engine.Id(chattingID).Cols("status").Update(chatting)
	if err != nil {
		log.Print(err)
		return 0
	}
	return chattingID
}

func (m *MessageModel) ReadChatMessage(id string, with string) int {
	engine := sql.GetSqlEngine()
	chat := models.NewChat()
	chat.ReadStatus = 0
	flag, err := engine.Where("from_id = ? or to_id = ?", with, id).Cols("read_status").Update(chat)
	if err != nil {
		log.Print(err.Error())
	}
	return int(flag)
}
