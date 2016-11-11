package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	initDb()
	initKey()
	router := gin.Default()
	router.Use(AuthMiddle)
	router.StaticFile("/", "./static/index.html")
	router.GET("/user", actionUser)
	router.GET("/entry", actionEntry)
	router.POST("/entry", actionEntryAdd)
	router.GET("/entry/:id", actionEntryId)
	router.GET("/entry/:id/pro", actionVote(1))
	router.GET("/entry/:id/con", actionVote(-1))
	router.Static("/js", "./static/js")
	router.Run(":8085")
}
