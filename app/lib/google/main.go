package google

import (
	"os"

	"golang.org/x/oauth2"
)

const (
	authorizeEndpoint = "https://accounts.google.com/o/oauth2/v2/auth"
	tokenEndpoint     = "https://www.googleapis.com/oauth2/v4/token"
)

func GetConnection() *oauth2.Config {
	// var protocol string
	// if os.Getenv("HostName") == "localhost" {
	// 	protocol = "http://"
	// } else {
	// 	protocol = "https://"
	// }

	config := &oauth2.Config{
		ClientID:     os.Getenv("GoogleClientID"),
		ClientSecret: os.Getenv("GoogleClientSecret"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  authorizeEndpoint,
			TokenURL: tokenEndpoint,
		},
		RedirectURL: "https://" + os.Getenv("HostName") + ":8080/google/callback",
		Scopes:      []string{"openid", "email", "profile"},
	}
	return config
}
