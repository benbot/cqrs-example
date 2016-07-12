// Command dispatcher for player events
// each dispatcher validates the command then creates
// and saves an event to the event store
package player

import (
	"cqrs-example/events"
	"cqrs-example/global"

	"github.com/kataras/iris"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

func Player_add(ctx *iris.Context) {
	c := global.G.Db.C("users")
	event := events.PlayerAddedEvent{uuid.NewV4().String()}

	record := events.NewEvent(event)

	var records []events.EventRecord
	c.Find(bson.M{"event.id": event.Id}).All(&records)

	if len(records) > 0 {
		ctx.SetStatusCode(400)
		ctx.SetBodyString("User with id already exists")
		return
	}

	err := c.Insert(&record)
	if err != nil {
		iris.Logger.Panicf("ERROR: " + err.Error())
		ctx.SetStatusCode(500)
	}

	ctx.SetStatusCode(201)
	ctx.SetContentType("game/user_id")
	ctx.SetBodyString(event.Id)
}
