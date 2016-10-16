package main

import (
	"github.com/kataras/iris"
	"github.com/dvsekhvalnov/jose2go/keys/rsa"
	"github.com/dvsekhvalnov/jose2go"
	"io/ioutil"
	"fmt"
)

type User struct {
	Id int
	Name string
}

func main() {
	iris.Get("/", func(ctx *iris.Context) {
		ctx.ServeFile("./static/index.html", false)
		getUser(ctx)
	})

	iris.Get("/myjson", func(ctx *iris.Context){
		ctx.JSON(iris.StatusOK, iris.Map{
			"Name": "Iris",
			"Released": "13 March 2016",
			"Stars": 5525,
		})
	})
	iris.StaticWeb("/js","./static/js", 1)


	iris.Listen(":8080")
}

func getUser(ctx *iris.Context){

	token := ctx.GetCookie("hp_jwt")

	keyBytes, err := ioutil.ReadFile("./pubkey.pem")

	if(err!=nil) {
		panic("invalid key file")
	}

	publicKey, e:=Rsa.ReadPublic(keyBytes)

	if(e!=nil) {
		panic("invalid key format")
	}

	payload, headers, err := jose.Decode(token, publicKey)

	if(err==nil) {
		//go use token
		fmt.Printf("\npayload = %v\n",payload)

		//and/or use headers
		fmt.Printf("\nheaders = %v\n",headers)
	} else {
		fmt.Printf("%v", err)
	}
}