package main

import (
	"github.com/kataras/iris"
)

func main() {
	initDb()
	initKey()
	iris.Get("/", actionIndex)
	iris.Get("/user", actionUser)
	iris.Get("/entry", actionEntry)
	iris.Get("/entry/:id", actionEntryId)
	iris.StaticWeb("/js", "./static/js", 1)
	iris.Listen(":8080")
}
