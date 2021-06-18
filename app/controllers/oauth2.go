package controllers

import (
	"fmt"
	"net/http"

	"Junking/lib/google"

	"github.com/gin-gonic/gin"
)

func GetOauth2(c *gin.Context) {
	config := google.GetConnection()
	url := config.AuthCodeURL("help!")
	fmt.Println(string(url))
	c.Redirect(http.StatusTemporaryRedirect, url)
}
