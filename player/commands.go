// Command dispatcher for player events
// each dispatcher validates the command then creates
// and saves an event to the event store
package player

import (
	"cqrs-example/events"
	"cqrs-example/global"
	"errors"

	"github.com/kataras/iris"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

func player_add() (string, error) {
	c := global.G.Db.C("users")
	event := PlayerAddedEvent{uuid.NewV4().String()}

	record := events.NewEvent(event)

	var records []events.EventRecord
	c.Find(bson.M{"event.id": event.Id}).All(&records)

	if len(records) > 0 {
		return "", errors.New("User exists")
	}

	err := c.Insert(&record)
	if err != nil {
		iris.Logger.Panicf("ERROR: " + err.Error())
		return "", err
	}

	return event.Id, nil
}
