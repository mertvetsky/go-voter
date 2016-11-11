package main

type Entry struct {
	Id     uint
	Status string
	Value  string
	Votes  []*Vote
}

type Vote struct {
	EntryID uint
	UserID  uint
	User    *User
	Weight  int
}

type User struct {
	Id    uint
	Votes []*Vote
	Name  string
}