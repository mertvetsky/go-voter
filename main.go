package main

import (
	"github.com/kataras/iris"
)

func main() {
	initDb()
	initKey()
	iris.Get("/", actionIndex)
	iris.Get("/user/:id", actionUserId)
	iris.Get("/entry", actionEntry)
	iris.Post("/entry", actionEntryAdd)
	iris.Get("/entry/:id", actionEntryId)
	iris.Get("/entry/:id/pro", actionEntryPro)
	iris.Get("/entry/:id/con", actionEntryCon)
	iris.StaticWeb("/js", "./static/js", 1)
	iris.Listen(":8080")
}
