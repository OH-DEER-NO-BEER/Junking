package controllers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// var roomMessageChan = make(chan roomMessage)

// type Rate struct {
// 	Rock     float64 `json:rock`
// 	Scissors float64 `json:scissors`
// 	Paper    float64 `json:"paper"`
// }

type roomId struct {
	RoomId string `json:"roomId"`
}

type player struct {
	Name   string             `json:"name"`
	Rate   map[string]float64 `json:"rate"`
	Rank   int                `json:"rank"`
	RoomId string             `json:"roomid"`
}
type roomMessage struct {
	Message string `json:"message"`
	Player1 player `json:"player1"`
	Player2 player `json:"player2"`
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

// func (cl *client) write() {
// 	for {
// 		select {
// 		case msg := <-roomMessageChan:
// 			// t, _ := json.Marshal(msg)
// 			// fmt.Print(t)
// 			if err := cl.socket.WriteJSON(msg); err != nil {
// 				break
// 			}
// 		}
// 	}
// 	// c.socket.Close()
// }

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
	// go clt.write()

	// // con := &connection{send: make(chan []byte, 256), ws: ws}
	// // sub := subscription{conn: con, room: roomId}

	var msg roomId
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
			tmp_room_json := roomMessage{}
			if _, ok := rh.rooms[msg.RoomId]; !ok {
				fmt.Println("create room")
				var clients = map[*client]bool{clt: true}
				rh.rooms[msg.RoomId] = &room{clients: clients}
				fmt.Println(rh.rooms[msg.RoomId].clients[clt])

				// if len(rh.rooms[msg.roomId].clients) < 1 {
				// go rh.rooms["test"].run()
				// makeRoomMessageChan <- makeRoomMessage{"RoomMade", p1}

				var rates1 = map[string]float64{"Rock": 0.1, "Scissors": 0.2, "Paper": 0.7}
				p1_tmp_json := player{"P1", rates1, 1, msg.RoomId}
				p2_tmp_json := player{"", map[string]float64{"Rock": 0.0, "Scissors": 0.0, "Paper": 0.0}, 1, msg.RoomId}
				tmp_room_json = roomMessage{"created room", p1_tmp_json, p2_tmp_json}

			} else if len(rh.rooms[msg.RoomId].clients) < 2 {
				fmt.Println("room exists")
				rh.rooms[msg.RoomId].clients[clt] = true

				var rates1 = map[string]float64{"Rock": 0.1, "Scissors": 0.2, "Paper": 0.7}
				p1_tmp_json := player{"P1", rates1, 1, msg.RoomId}
				var rates2 = map[string]float64{"Rock": 0.2, "Scissors": 0.2, "Paper": 0.6}
				p2_tmp_json := player{"P2", rates2, 1, msg.RoomId}
				tmp_room_json = roomMessage{"checked in", p1_tmp_json, p2_tmp_json}
				// makeRoomMessageChan <- makeRoomMessage{"RoomEnter", p2}

			} else {
				fmt.Println("room's full!!!")

				p1_tmp_json := player{"", map[string]float64{"Rock": 0.0, "Scissors": 0.0, "Paper": 0.0}, 1, msg.RoomId}
				p2_tmp_json := player{"", map[string]float64{"Rock": 0.0, "Scissors": 0.0, "Paper": 0.0}, 1, msg.RoomId}
				tmp_room_json = roomMessage{"room's full!!!", p1_tmp_json, p2_tmp_json}

			}
			room_json, _ := json.Marshal(tmp_room_json)
			connectionErr := ws.WriteJSON(string(room_json))
			if connectionErr != nil {
				log.Println("write:", connectionErr)
			}
			break
		}
	}

	defer ws.Close()
}
