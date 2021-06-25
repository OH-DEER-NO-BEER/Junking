package controllers

import (
	"encoding/json"
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
	RoomId string `json:"roomId"`
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
	fmt.Println("checkin' in")
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	clt := &client{ws}

	// // go p.read()
	go clt.write()

	// // con := &connection{send: make(chan []byte, 256), ws: ws}
	// // sub := subscription{conn: con, room: roomId}

	var msg roomId
	var p_tmp_json = person{}
	for {
		err := ws.ReadJSON(&msg)
		if err != nil {
			fmt.Println(msg.RoomId)
			log.Printf("error: %v", err)
			// delete(clients, ws)
			break
		}
		if msg.RoomId != "" {
			fmt.Println(msg.RoomId)
			if _, ok := rh.rooms[msg.RoomId]; !ok {
				fmt.Println("create room")
				var clients = map[*client]bool{clt: true}
				rh.rooms[msg.RoomId] = &room{clients: clients}
				fmt.Println(rh.rooms[msg.RoomId].clients[clt])

				// if len(rh.rooms[msg.roomId].clients) < 1 {
				// go rh.rooms["test"].run()
				// makeRoomMessageChan <- makeRoomMessage{"RoomMade", p1}

				var rates = map[string]float64{"Rock": 0.1, "Scissors": 0.2, "Paper": 0.7}
				p_tmp_json = person{"P1", rates, 1, "test"}
			} else {
				fmt.Println("room exists")
				rh.rooms[msg.RoomId].clients[clt] = true
				var rates = map[string]float64{"Rock": 0.2, "Scissors": 0.2, "Paper": 0.6}
				p_tmp_json = person{"P2", rates, 1, "test"}
				// makeRoomMessageChan <- makeRoomMessage{"RoomEnter", p2}
			}
			break
		}
	}

	p_json, _ := json.Marshal(p_tmp_json)
	connectionErr := ws.WriteJSON(string(p_json))
	if connectionErr != nil {
		log.Println("write:", connectionErr)
	}

	defer ws.Close()
}
