package main

import (
	"io/ioutil"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/dvsekhvalnov/jose2go/keys/rsa"
	"encoding/json"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"crypto/rsa"
)

var db *gorm.DB
var publicKey *rsa.PublicKey

func initDb() {
	var err error
	var config map[string]interface{}
	file, _ := ioutil.ReadFile("./config/db.json")

	if err := json.Unmarshal(file, &config); err != nil {
		panic(err)
	}

	sprintf := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		config["username"], config["password"], config["host"], config["dbname"])

	db, err = gorm.Open("mysql", sprintf)
	if err != nil {
		panic(err)
	}
	//db.Create(&Entry{Value:"sigurd", Status:"new"})
	db = db.Debug()
	db.AutoMigrate(&Entry{})
	db.AutoMigrate(&Vote{})
	db.AutoMigrate(&User{})

}

func initKey() {
	keyBytes, err := ioutil.ReadFile("./config/pubkey.pem")

	if (err != nil) {
		panic("invalid key file")
	}
	publicKey, err = Rsa.ReadPublic(keyBytes)

	if (err != nil) {
		panic("invalid key file")
	}
}