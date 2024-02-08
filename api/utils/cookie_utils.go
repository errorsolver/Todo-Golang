package utils

import (
	"os"

	"github.com/gin-gonic/gin"
)

func CreateJWTCookie(c *gin.Context, token string) string {
	c.SetCookie(
		os.Getenv("COOKIETITLE"), // name string
		token,                    // value string
		0,                        // maxAge int
		"/",                      // path string
		"localhost",              // domain string
		false,                    // secure bool
		true,                     // httpOnly bool
	)
	return token
}
func DeleteJWTCookie(c *gin.Context) {
	c.SetCookie(
		os.Getenv("COOKIETITLE"), // name string
		"",                       // value string
		-1,                       // maxAge int
		"/",                      // path string
		"localhost",              // domain string
		false,                    // secure bool
		true,                     // httpOnly bool
	)
}
