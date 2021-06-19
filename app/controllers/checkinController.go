package controllers

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// connection is an middleman between the websocket connection and the hub.
type connection struct {
	// The websocket connection.
	ws *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

// readPump pumps messages from the websocket connection to the hub.
func (s subscription) readPump() {
	c := s.conn
	defer func() {
		h.unregister <- s
		c.ws.Close()
	}()
	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, msg, err := c.ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		m := message{msg, s.room}
		h.broadcast <- m
	}
}

// write writes a message with the given message type and payload.
func (c *connection) write(mt int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(mt, payload)
}

// writePump pumps messages from the hub to the websocket connection.
func (s *subscription) writePump() {
	c := s.conn
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

/////////////////////////////////////////////////////////////////////////////

type message struct {
	data []byte
	room string
}

type subscription struct {
	conn *connection
	room string
}

type Hub struct {
	// Registered connections.
	rooms map[string]map[*connection]bool

	// Inbound messages from the connections.
	broadcast chan message

	// Register requests from the connections.
	register chan subscription

	// Unregister requests from connections.
	unregister chan subscription
}

var clients = make(map[*websocket.Conn]bool) // 接続されるクライアント
var broadcast = make(chan Message)           // メッセージブロードキャストチャネル

// メッセージ用構造体
type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

func CheckIn(c *gin.Context) {
	roomId := c.Param("roomId")
	fmt.Print(roomId)
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	con := &connection{send: make(chan []byte, 256), ws: ws}
	s := subscrption{con, roomId}
	rooomsHub.register <- s

	go s.writePump()
	go s.readPump()

	// 関数が終わった際に必ずwebsocketnのコネクションを閉じる
	defer ws.Close()

	// クライアントを新しく登録
	clients[ws] = true

	for {
		var msg Message
		// 新しいメッセージをJSONとして読み込みMessageオブジェクトにマッピングする
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		// 新しく受信されたメッセージをブロードキャストチャネルに送る
		broadcast <- msg
	}
}

func HandleMessages() {
	for {
		// ブロードキャストチャネルから次のメッセージを受け取る
		msg := <-broadcast
		// 現在接続しているクライアント全てにメッセージを送信する
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
