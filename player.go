package main

import (
	"gopkg.in/mgo.v2/bson"

	"github.com/kataras/iris"
	"github.com/pborman/uuid"
)

type Player struct {
	Id string `json:"id"`
}

func player_add(ctx *iris.Context) {
	c := config.Db.C("users")
	data := PlayerAddedEventData{uuid.New()}

	ev := NewEvent("PLAYER_ADDED", data)

	var events []Event
	c.Find(bson.M{"data.id": data.Id}).All(&events)

	if len(events) > 0 {
		ctx.SetStatusCode(400)
		ctx.SetBodyString("User with id already exists")
		return
	}

	err := c.Insert(&ev)
	if err != nil {
		iris.Logger.Panicf("ERROR: " + err.Error())
		ctx.SetStatusCode(500)
	}

	ctx.SetStatusCode(201)
	ctx.SetContentType("game/user_id")
	ctx.SetBodyString(data.Id)
}

func player_projection(id string) (*Player, error) {
	c := config.Db.C("users")

	var events []struct {
		Event
		Data PlayerAddedEventData
	}
	c.Find(bson.M{"data.id": id}).Sort("_id").All(&events)

	p := Player{}

	for _, e := range events {
		data := e.Data
		switch {
		case e.Type == "PLAYER_ADDED":
			p.Id = data.Id
		}
	}

	return &p, nil
}
