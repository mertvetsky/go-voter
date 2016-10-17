package main

import (
	"github.com/kataras/iris"
)

func actionIndex(ctx *iris.Context) {
	ctx.ServeFile("./static/index.html", false)
}

func actionUser(ctx *iris.Context) {
	ctx.JSON(iris.StatusOK, getUser(ctx))
}

func actionEntry(ctx *iris.Context) {
	var entries []Entry
	db.Find(&entries)
	ctx.JSON(iris.StatusOK, entries)
}

func actionEntryId(ctx *iris.Context) {
	var entry Entry
	db.Last(&entry, ctx.Param("id"))
	ctx.JSON(iris.StatusOK, entry)
}