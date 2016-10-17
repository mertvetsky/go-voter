package main

import "github.com/jinzhu/gorm"

type Entry struct {
	gorm.Model
	Status string
	Value  string
}