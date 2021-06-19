package controllers

import (
	"Junking/lib/google"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	v2 "google.golang.org/api/oauth2/v2"
)

func GetToken(c *gin.Context) {
	fmt.Println("Settin' Cookie!")

	if err := c.Request.ParseForm(); err != nil {
		panic(err)
	}

	config := google.GetConnection()
	tok, err := config.Exchange(c, c.Request.Form["code"][0])
	if err != nil {
		panic(err)
	}

	if tok.Valid() == false {
		panic(errors.New("vaild token"))
	}

	service, _ := v2.New(config.Client(c, tok))
	userinfo, _ := service.Userinfo.Get().Fields().Context(c).Do()
	tokenInfo, _ := service.Tokeninfo().AccessToken(tok.AccessToken).Context(c).Do()

	// c.JSON(http.StatusOK, gin.H{
	// 	"ID":    tokenInfo.UserId,
	// 	"Email": tokenInfo.Email,
	// })

	c.SetCookie("Name", userinfo.Name, 3600, "/", os.Getenv("HostName"), false, true)
	c.SetCookie("ID", tokenInfo.UserId, 3600, "/", os.Getenv("HostName"), false, true)
	c.SetCookie("Email", tokenInfo.Email, 3600, "/", os.Getenv("HostName"), false, true)
	fmt.Println("setted : ", tokenInfo.UserId, tokenInfo.Email)

	c.Redirect(http.StatusTemporaryRedirect, "http://"+os.Getenv("HostName"))
}
