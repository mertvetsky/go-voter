package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"github.com/gin-gonic/gin"
)

const (
	jwtName = "begetInnerJWT"
)

func getUser(c *gin.Context) (User) {
	user := User{}
	tokenString, err := c.Cookie(jwtName)

	if err == nil && tokenString != "" {
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return publicKey, nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			parsedId, _ := strconv.Atoi(claims["uid"].(string))
			db.FirstOrCreate(&user, User{Id: uint(parsedId), Name: claims["uname"].(string)})
		}
	}
	return user
}

func AuthMiddle(c *gin.Context) {
	if getUser(c).Id == 0 {
		c.JSON(403, ErrorAnswer{ErrorCode:403, ErrorDescription: "Уходи"})
		c.Abort()
	} else {
		c.Next()
	}
}