package main

import (
	"github.com/kataras/iris"
)

var myid uint = 35

func actionIndex(ctx *iris.Context) {
	ctx.ServeFile("./static/index.html", false)
}

func actionEntry(ctx *iris.Context) {
	var entries []Entry
	db.Preload("Votes.User").Find(&entries)
	ctx.JSON(iris.StatusOK, entries)
}
func actionEntryAdd(ctx *iris.Context) {
	var body struct{ Value string `form:"value"` }
	ctx.ReadJSON(&body)
	db.Create(&Entry{Value: body.Value, Status: "new"})
	var entry Entry
	db.Preload("Votes.User").Where("value = ?", body.Value).Last(&entry)
	ctx.JSON(iris.StatusOK, entry)
}

func actionEntryId(ctx *iris.Context) {
	var entry Entry
	db.Preload("Votes.User").Last(&entry, ctx.Param("id"))
	ctx.JSON(iris.StatusOK, entry)
}
func actionEntryPro(ctx *iris.Context) {
	var entry Entry
	entryId := ctx.Param("id")
	db.Where("user_id = ? and entry_id = ?", myid, entryId).Delete(&Vote{})
	db.Last(&entry, entryId).Association("Votes").Append(Vote{UserID:myid, Weight: 1})
	db.Preload("Votes.User").Last(&entry, entryId)
	ctx.JSON(iris.StatusOK, entry)
}
func actionEntryCon(ctx *iris.Context) {
	var entry Entry
	entryId := ctx.Param("id")
	db.Where("user_id = ? and entry_id = ?", myid, entryId).Delete(&Vote{})
	db.Last(&entry, entryId).Association("Votes").Append(Vote{UserID:myid, Weight: -1})
	db.Preload("Votes.User").Last(&entry, entryId)
	ctx.JSON(iris.StatusOK, entry)
}

func actionUserId(ctx *iris.Context) {
	var user User
	db.Last(&user, ctx.Param("id"))
	ctx.JSON(iris.StatusOK, user)
}
