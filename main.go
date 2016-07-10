package main

import (
	"cqrs-example/global"
	"flag"

	"github.com/iris-contrib/middleware/logger"
	"github.com/kataras/iris"
	"gopkg.in/mgo.v2"
)

var config global.Global

func main() {
	dbAddr := flag.String("db", "0.0.0.0", "mongodb address")
	flag.Parse()
	config = global.Global{}

	iris.Logger.Println("Connecting to mogno...")
	session, err := mgo.Dial("mongodb://" + *dbAddr)
	if err != nil {
		panic(err)
	}
	iris.Logger.Println("Connected to mongo!")
	defer session.Close()

	iris.Use(logger.New(iris.Logger))

	iris.Post("/players", player_add)

	iris.Listen(":8080")
}
