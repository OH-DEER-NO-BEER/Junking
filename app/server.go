package main

import (
	"fmt"
	"log"
	"net/http"

	"Junking/controllers"

	"github.com/gin-gonic/gin"
)

func sampleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		name, name_err := c.Cookie("Name")
		id, id_err := c.Cookie("ID")
		email, email_err := c.Cookie(("Email"))
		fmt.Println("name =", name, "id =", id, ", email =", email, "name_err =", name_err, "id_err =", id_err, ", emai_err =", email_err)
		if name_err != nil || id_err != nil || email_err != nil {
			fmt.Println("No Cookie!")
			controllers.GetOauth2(c)
			c.Abort()
		}
		c.Next()
	}
}

func main() {
	engine := gin.Default()

	engine.LoadHTMLGlob("views/*.html")
	engine.Static("/js", "./views/dist")
	engine.Static("/unity", "./views/unity")
	engine.Static("/Build", "./views/Build")
	engine.Static("/css", "./views/css")
	engine.Static("/TemplateData", "./views/TemplateData")

	engine.GET("/google/callback", controllers.GetToken)

	loginGroup := engine.Group("/")
	loginGroup.Use(sampleMiddleware())
	{
		loginGroup.GET("/", func(c *gin.Context) {
			c.Redirect(http.StatusTemporaryRedirect, "/app")
		})

		loginGroup.GET("/app", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{})
		})

		loginGroup.GET("/ws:roomId", controllers.CheckIn)
	}

	err := engine.Run(":8080")
	if err != nil {
		log.Fatal("Run: ", err)
	}
}
