package player

import (
	"cqrs-example/global"
	"cqrs-example/helpers"
	"errors"

	"github.com/kataras/iris"

	"gopkg.in/mgo.v2/bson"
)

type Player struct {
	Id string `json:"id"`
}

func player_projection(id string) (*Player, error) {
	c := global.Db.C(global.USER_COLLECTION)

	type evDeconstruct struct {
		Type  string                 `bson:"type"`
		Event map[string]interface{} `bson:"event"`
	}

	iter := c.Find(bson.M{"event.id": id}).Sort("_id").Iter()
	defer iter.Close()

	var (
		p   *Player
		ev  evDeconstruct
		err error
	)

	for iter.Next(&ev) {
		switch {
		case ev.Type == "PlayerAddedEvent":
			// This should be the first event for any player
			// so we define p here. If another event comes first
			// we will try to modify player and an error will be thrown
			// because p is an empty pointer
			addEv := &PlayerAddedEvent{}
			for s, v := range ev.Event {
				if err := helpers.SetField(addEv, s, v); err != nil {
					panic(err)
					iris.Logger.Println(err)
				}
			}

			p, err = added_event(addEv)
			if err != nil {
				return nil, err
			}
		}
	}

	if err = iter.Err(); err != nil {
		return nil, err
	} else if p == nil {
		return nil, errors.New("Error not found")
	}

	return p, nil
}
