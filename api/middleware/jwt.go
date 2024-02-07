package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type JWTDataClaims struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func CreateJwt(name string, id uint) (string, error) {
	// idString := strconv.FormatUint(uint64(id), 10)
	claim := jwt.MapClaims{
		"exp":    time.Now().AddDate(0, 0, 5).Unix(),
		"issuer": "Todo-Golang",
		"data": JWTDataClaims{
			ID:   id,
			Name: name,
		},
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256, // Method
		claim,                  // Claim
		// Option
	)

	tokenString, err := token.SignedString([]byte(os.Getenv("SIGNINGKEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := GetJWTData(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"code":    http.StatusBadRequest,
				"error":   err.Error(),
			})
			return
		}
		c.Next()
	}
}

func GetJWTData(c *gin.Context) (jwt.MapClaims, error) {
	tokenString, err := GetCookie(c)
	if err != nil {
		return nil, err
	}

	result, err := ParseJWT(tokenString)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetCookie(c *gin.Context) (string, error) {
	tokenString, err := c.Cookie(os.Getenv("COOKIETITLE"))
	if err != nil {
		return "", fmt.Errorf("You are not logged in yet")
	}
	return tokenString, nil
}

func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("SIGNINGKEY")), nil
	})

	if err != nil {
		return jwt.MapClaims{}, err
	}

	claims := token.Claims.(jwt.MapClaims)
	return claims, nil
}
