package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"net/http"
)

type ClientManager struct {
	count    int
	register chan *Client
	message  chan Message
	clients  map[int]*Client
}

type Client struct {
	id      int
	pid     string
	read    chan string
	send    chan string
	socket  *websocket.Conn
	manager *ClientManager
}

type Message struct {
	To   string
	From string
	Msg  string
}


func (client *Client) readMessage() {
	defer func() {
		err := client.socket.Close()
		if err != nil {
			log.Print(err)
		}
	}()
	for {
		_, msg, err := client.socket.ReadMessage()
		if err != nil {
			log.Print(err)
			delete(client.manager.clients, client.id)
			break
		}
		fmt.Println(client.id, "read", string(msg))
		message := Message{}
		err = json.Unmarshal([]byte(string(msg)), &message)
		if err != nil {
			log.Print(err)
		}
		if message.To == "" {
			client.pid = message.From
			continue
		}
		client.manager.message <- message
	}
}

func (client *Client) sendMessage() {
	for {
		message := <-client.send
		if message == "" {
			continue
		}
		fmt.Println(client.id, "send", message)
		err := client.socket.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Print(err)
		}
	}
}

func (manager *ClientManager) RegisterClient() {
	for {
		client := <-manager.register
		fmt.Println(client.id, "RegisterClient")
		manager.clients[client.id] = client
		manager.count++
	}
}

func (manager *ClientManager) getClient(conn *websocket.Conn) {
	client := &Client{
		id:      manager.count,
		socket:  conn,
		read:    make(chan string),
		send:    make(chan string),
		manager: manager,
	}
	go client.readMessage()
	go client.sendMessage()
	manager.register <- client

}

func (manager *ClientManager) Push() {
	for {
		message := <-manager.message
		for _, v := range manager.clients {
			if v.pid == message.To {
				v.send <- message.Msg
			}
		}
	}
}

func (manager *ClientManager) Run() {
	go manager.RegisterClient()
	go manager.Push()
}

func GetClientManager() *ClientManager {
	manager := &ClientManager{
		clients:  make(map[int]*Client),
		register: make(chan *Client),
		message:  make(chan Message),
		count:    0,
	}
	manager.Run()
	return manager
}

func (manager *ClientManager) WebsocketServer(c echo.Context) error {
	up := websocket.Upgrader{
		// 解决跨域问题
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := up.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	manager.getClient(conn)
	return nil
}

func main() {
	var clientManager = GetClientManager()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/ws", clientManager.WebsocketServer)
	e.Logger.Fatal(e.Start(":1323"))
}
