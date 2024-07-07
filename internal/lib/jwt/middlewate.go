package jwt

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
)

func VerifyJwt() gin.HandlerFunc {
	return func(c *gin.Context) {

		authorization := c.Request.Header.Get("Authorization")
		parts := strings.Split(authorization, " ")

		if len(parts) == 2 {
			token, err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
				_, ok := token.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					c.Writer.WriteHeader(http.StatusUnauthorized)
					_, err := c.Writer.Write([]byte("Unauthorized"))
					if err != nil {
						return nil, err
					}
					return nil, fmt.Errorf("unexpected signing method")
				}

				return "", nil
			})
			if err != nil {
				c.Writer.WriteHeader(http.StatusUnauthorized)
				_, err := c.Writer.Write([]byte("You're Unauthorized due to error parsing token"))
				if err != nil {
					return
				}
			}

			if token.Valid {
				c.Next()

			} else {
				c.Writer.WriteHeader(http.StatusUnauthorized)
				_, err := c.Writer.Write([]byte("You're Unauthorized due to invalid  token"))
				if err != nil {
					return
				}
			}

		} else {
			c.Writer.WriteHeader(http.StatusUnauthorized)
			_, err := c.Writer.Write([]byte("You're Unauthorized due to No token in the header"))
			if err != nil {
				return
			}
		}

	}

}
