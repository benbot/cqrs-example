package player

import (
	"cqrs-example/global"
	"errors"

	"gopkg.in/mgo.v2/bson"
)

type Player struct {
	Id string `json:"id"`
}

func player_projection(id string) (*Player, error) {
	c := global.G.Db.C("users")

	var records []struct {
		Type  string                 `bson:"type"`
		Event map[string]interface{} `bson:"event"`
	}

	c.Find(bson.M{"event.id": id}).Sort("_id").All(&records)

	if len(records) <= 0 {
		return nil, errors.New("player not found")
	}

	p := Player{}

	for _, e := range records {
		switch {
		case e.Type == "PlayerAddedEvent":
			p.Id = e.Event["id"].(string)
		}
	}

	return &p, nil
}
