package controllers

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var makeRoomMessageChan = make(chan makeRoomMessage)

// type Rate struct {
// 	Rock     float64 `json:rock`
// 	Scissors float64 `json:scissors`
// 	Paper    float64 `json:"paper"`
// }

type roomId struct {
	roomId string `json:"roomId"`
}

type person struct {
	Name   string             `json:"name"`
	Rate   map[string]float64 `json:"rate"`
	Rank   int                `json:"rank"`
	RoomId string             `json:"roomid"`
}
type makeRoomMessage struct {
	Message string `json:"message"`
	P1      person `json:"P1"`
}

type message struct {
	data []byte
	room string
}

type subscription struct {
	conn *connection
	room string
}

type connection struct {
	ws   *websocket.Conn
	send chan []byte
}

type client struct {
	socket *websocket.Conn
	// send   chan []byte
	// room   *room
}

type room struct {
	join    chan *client
	clients map[*client]bool
	// conn    map[socket]bool
}

var RoomsHub = roomsHub{
	broadcast:  make(chan message),
	register:   make(chan subscription),
	unregister: make(chan subscription),
	rooms:      make(map[string]*room),
}

type roomsHub struct {
	rooms      map[string]*room
	broadcast  chan message
	register   chan subscription
	unregister chan subscription
}

func (cl *client) write() {
	for {
		select {
		case msg := <-makeRoomMessageChan:
			// t, _ := json.Marshal(msg)
			// fmt.Print(t)
			if err := cl.socket.WriteJSON(msg); err != nil {
				break
			}
		}
	}
	// c.socket.Close()
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
		}
	}
}

func (rh *roomsHub) CheckIn(c *gin.Context) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	// con := &connection{send: make(chan []byte, 256), ws: ws}
	// sub := subscription{conn: con, room: roomId}

	msg := roomId{}
	for {
		err := ws.ReadJSON(&msg)
		if msg.roomId != "" {
			fmt.Println(msg.roomId)
			log.Printf("error: %v", err)
			// delete(clients, ws)
			break
		}
	}

	p := &client{ws}
	// go p.read()
	go p.write()

	if len(rh.rooms[msg.roomId].clients) < 1 {
		go rh.rooms[msg.roomId].run()
		var rates = map[string]float64{"Rock": 0.1, "Scissors": 0.2, "Paper": 0.7}
		p1 := person{"P1", rates, 1, msg.roomId}
		makeRoomMessageChan <- makeRoomMessage{"RoomMade", p1}
	} else {
		var rates = map[string]float64{"Rock": 0.2, "Scissors": 0.2, "Paper": 0.6}
		p2 := person{"P2", rates, 1, msg.roomId}
		makeRoomMessageChan <- makeRoomMessage{"RoomEnter", p2}
	}

	defer ws.Close()
}
