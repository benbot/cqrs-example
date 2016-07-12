package player

import (
	"cqrs-example/global"
	"errors"

	"github.com/kataras/iris"

	"gopkg.in/mgo.v2/bson"
)

type Player struct {
	Id string `json:"id"`
}

func player_projection(id string) (*Player, error) {
	defer func() {
		if e := recover(); e != nil {
			iris.Logger.Println(e)
		}
	}()

	c := global.G.Db.C("users")

	type evDeconstruct struct {
		Type  string                 `bson:"type"`
		Event map[string]interface{} `bson:"event"`
	}

	iter := c.Find(bson.M{"event.id": id}).Sort("_id").Iter()

	var ev evDeconstruct
	var p *Player

	for iter.Next(&ev) {
		switch {
		case ev.Type == "PlayerAddedEvent":
			// This should be the first event for any player
			// so we devine p here. If another event comes first
			// we will try to modify player and an error will be thrown
			p = &Player{}
			p.Id = ev.Event["id"].(string)
		}
	}

	if err := iter.Err(); err != nil {
		return nil, err
	} else if p == nil {
		return nil, errors.New("Error not found")
	}

	return p, nil
}
