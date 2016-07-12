package main

import (
	"cqrs-example/global"
	"cqrs-example/player"
	"flag"

	"gopkg.in/mgo.v2"

	"github.com/iris-contrib/middleware/logger"
	"github.com/kataras/iris"
)

func main() {
	dbAddr := flag.String("db", "0.0.0.0", "mongodb address")
	flag.Parse()

	iris.Logger.Println("Connecting to mogno...")
	session, err := mgo.Dial("mongodb://" + *dbAddr)
	if err != nil {
		panic(err)
	}

	iris.Logger.Println("Connected to mongo!")
	defer session.Close()

	global.G.Db = session.DB("game")

	iris.Use(logger.New(iris.Logger))

	iris.Post("/players", player.Player_add)

	iris.Listen(":8080")
}
