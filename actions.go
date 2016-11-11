package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ErrorAnswer struct {
	ErrorCode        int
	ErrorDescription string
}

func actionEntry(c *gin.Context) {
	var entries []Entry
	db.Preload("Votes.User").Find(&entries)
	c.JSON(200, entries)
}
func actionEntryAdd(c *gin.Context) {
	var body struct{ Value string `form:"value"` }
	c.BindJSON(&body)
	db.Create(&Entry{Value: body.Value, Status: "new"})
	var entry Entry
	db.Preload("Votes.User").Where("value = ?", body.Value).Last(&entry)
	c.JSON(200, entry)
}

func actionEntryId(c *gin.Context) {
	var entry Entry
	db.Preload("Votes.User").Last(&entry, c.Param("id"))
	c.JSON(200, entry)
}

func actionUser(c *gin.Context) {
	c.JSON(200, getUser(c))
}

func actionVote(weight int) func(*gin.Context) { // FP god
	return func(c *gin.Context) {
		var entry Entry
		var vote Vote
		parsedId, _ := strconv.Atoi(c.Param("id"))
		id := uint(parsedId)

		db.Find(&entry, Entry{Id: id})
		if entry.Id == 0 {
			c.JSON(404, ErrorAnswer{403, fmt.Sprintf("Entry %v not founded", c.Param("id"))})
			return
		}
		db.Where(Vote{UserID:getUser(c).Id, EntryID:entry.Id}).Assign(Vote{Weight: weight}).FirstOrCreate(&vote)
		db.Preload("Votes.User").Last(&entry, entry.Id)
		c.JSON(200, entry)
	}
}