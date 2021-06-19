package main

import (
	"fmt"
	"log"
	"net/http"

	"Junking/controllers"

	"github.com/gin-gonic/gin"
)

// func helloWorld(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()       //オプションを解析します。デフォルトでは解析しません。
// 	fmt.Println(r.Form) //このデータはサーバのプリント情報に出力されます。
// 	fmt.Println("path", r.URL.Path)
// 	fmt.Println("scheme", r.URL.Scheme)
// 	fmt.Println(r.Form["url_long"])
// 	fmt.Fprintf(w, "Hello, world!") //ここでwに入るものがクライアントに出力されます。
// }

func sampleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		name, name_err := c.Cookie("Name")
		id, id_err := c.Cookie("ID")
		email, email_err := c.Cookie(("Email"))
		fmt.Println("name =", name, "id =", id, ", email =", email, "name_err =", name_err, "id_err =", id_err, ", emai_err =", email_err)
		if name_err != nil || id_err != nil || email_err != nil {
			// c.Redirect()
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
	engine.Static("/TemplateData", "./views/TemplateData")

	engine.GET("/google/callback", controllers.GetToken)

	loginGroup := engine.Group("/")
	loginGroup.Use(sampleMiddleware())
	{
		loginGroup.GET("/", func(c *gin.Context) {
			// session := sessions.Default(c)
			// accessToken := session.Get("access_token")
			// accessSecret := session.Get("access_secret")
			// if accessToken == nil || accessSecret == nil {

			// requestToken, requestSecret, _ := config.RequestToken()
			// session.Set("request_secret", requestSecret)
			// session.Save()
			// c.Redirect(http.StatusFound, )

			// c.HTML(http.StatusOK, "index.html", gin.H{
			// 	"message": "Hi! here is root!",
			// })
			c.Redirect(http.StatusTemporaryRedirect, "/app")
		})
		// engine.GET("/", controllers.GetOauth2)
		loginGroup.GET("/app", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{})
		})

		loginGroup.GET("/ws:roomId", controllers.CheckIn)
	}

	// engine.Static("/view", "./views")

	// store := cookie.NewStore([]byte("secret"))
	// engine.Use(sessions.Sessions("mysession", store))

	// engine.StaticFS("/", http.Dir("./views/static"))

	// engine.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.html", gin.H{})
	// })

	err := engine.Run(":8080")
	if err != nil {
		log.Fatal("Run: ", err)
	}
}
