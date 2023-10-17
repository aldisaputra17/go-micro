package oauth20

import (
	"github.com/aldisaputra17/go-micro/helper"
	"golang.org/x/oauth2"
)

var (
	MicrosoftOAuthConfig = oauth2.Config{
		ClientID:     "67f5496e-0674-4178-ad20-d65ad137170c",
		ClientSecret: "fcS8Q~IzcXC2HQ9gaHQ-TXkYHnUq4lFTIZKF0bTt",
		RedirectURL:  "http://localhost:8000/callback", // Sesuaikan dengan URL callback Anda.
		Scopes:       []string{"user.read"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://login.microsoftonline.com/" + common + "/oauth2/v2.0/authorize",
			TokenURL: "https://login.microsoftonline.com/" + common + "/oauth2/v2.0/token",
		},
	}
	oauth2StateString, _ = helper.GenerateRandomString(32)
	common               = "common"
)

func Login() string {
	url := MicrosoftOAuthConfig.AuthCodeURL(oauth2StateString)
	return url
}
