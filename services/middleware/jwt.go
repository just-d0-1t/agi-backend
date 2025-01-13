package middleware

import (
	"agi-backend/utils"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	jwtSecret = "1111"
)

func GenerateToken(ID uint, name string) (string, error) {
	claims := jwt.MapClaims{
		"id":   ID,
		"name": name,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(jwtSecret))
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("X-token")

		if len(authHeader) == 0 {
			utils.ResponseError(c, fmt.Errorf("missing token.").Error())
			c.Abort()
			return
		}

		tokenStr := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			utils.ResponseError(c, err.Error())
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			utils.ResponseError(c, fmt.Errorf("invalid token claims").Error())
			c.Abort()
			return
		}

		c.Set("jwt", claims)
		c.Next()
	}
}
