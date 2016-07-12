package player

import (
	"cqrs-example/events"
	"cqrs-example/global"

	"gopkg.in/mgo.v2/bson"
)

type Player struct {
	Id string `json:"id"`
}

func player_projection(id string) (*Player, error) {
	c := global.G.Db.C("users")

	var records []events.EventRecord

	c.Find(bson.M{"event.id": id}).Sort("_id").All(&records)

	p := Player{}

	for _, e := range records {
		switch {
		case e.Type == "PlayerAddedEvent":
			ev := e.Event.(events.PlayerAddedEvent)
			p.Id = ev.Id
		}
	}

	return &p, nil
}
