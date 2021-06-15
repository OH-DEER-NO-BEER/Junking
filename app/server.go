package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //オプションを解析します。デフォルトでは解析しません。
	fmt.Println(r.Form) //このデータはサーバのプリント情報に出力されます。
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	fmt.Fprintf(w, "Hello, world!") //ここでwに入るものがクライアントに出力されます。
}

func main() {
	engine := gin.Default()
	// engine.Static("/", "./views")

	// engine.StaticFS("/", http.Dir("./views/static"))

	engine.LoadHTMLGlob("views/*")
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	err := engine.Run(":8080")
	if err != nil {
		log.Fatal("Run: ", err)
	}
}
