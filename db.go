package main

import "github.com/jinzhu/gorm"

type Entry struct {
	gorm.Model
	Status string
	Value  string
	Votes  []*Vote
}

type Vote struct {
	gorm.Model
	EntryID uint
	UserID  uint
	User    *User
	Weight  int
}

type User struct {
	gorm.Model
	Votes []*Vote
	Name  string
}