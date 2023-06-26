package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type header struct {
	Authorization string `header:"Authorization"`
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := header{}

		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "failed",
			})
			return
		}

		if h.Authorization == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Access is denied",
			})
			return
		}

		tokenString := strings.ReplaceAll(h.Authorization, "Bearer ", "")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			hmacSampleSecret := []byte("my_secret_key")
			return hmacSampleSecret, nil
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "InternalServerError",
			})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "InternalServerError",
			})
			return
		}

		// c.Set("userId", claims["foo"])
		id, _ := claims["id"].(float64)
		c.Set("userId", int(id))

		c.Next()

		// claims["foo"]
	}
}
