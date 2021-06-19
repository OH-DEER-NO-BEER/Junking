package controllers

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/gorilla/websocket"
)

func CheckIn(c *gin.Context) {
	var upgrader = websocket.Upgrader{}
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatal(err)
	}
	// 関数が終わった際に必ずwebsocketnのコネクションを閉じる
	defer ws.Close()
}
