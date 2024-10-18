package utils

import (
	"log"

	"github.com/gin-gonic/gin"
)

type CookieConfig struct {
	Name     string
	Value    string
	MaxAge   int // MaxAge in seconds
	HttpOnly bool
	Secure   bool
}

func SetCookie(c *gin.Context, config CookieConfig) {
	log.Printf("Setting cookie: Name=%s, Value=%s", config.Name, config.Value)
	if config.MaxAge > 0 {
		c.SetCookie(
			config.Name,
			config.Value,
			config.MaxAge,
			"/",
			"",
			config.Secure,
			config.HttpOnly,
		)
	} else {
		c.SetCookie(config.Name, "", -1, "/", "", config.Secure, config.HttpOnly)
	}
}

func GetContentFromCookie(c *gin.Context, key string) (string, error) {
	cookie, err := c.Cookie("token")
	if err != nil {
		return "", err
	}

	userID, err := VerifyTokenUsingClaims(cookie)
	return userID, err
}
