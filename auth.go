package main

import (
	"github.com/kataras/iris"
	"fmt"
	"encoding/json"
	"github.com/dvsekhvalnov/jose2go"
)

type AuthUser struct {
	Id   int
	Name string
}

func getUser(ctx *iris.Context) (map[string]interface{}) {
	token := ctx.GetCookie("hp_jwt")

	payload, _, err := jose.Decode(token, publicKey)
	var dat map[string]interface{}

	if (err == nil) {
		json.Unmarshal([]byte(payload), &dat)
		//fmt.Printf("\npayload = %v\n", dat["uid"])
		return dat
	} else {
		dat = map[string]interface{}{"error": 123}
		fmt.Printf("%v", err)
	}

	return dat
}