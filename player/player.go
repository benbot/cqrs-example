package player

import (
	"cqrs-example/events"

	"github.com/graphql-go/graphql"
)

type Player struct {
	Id string `json:"id"`
}

var PlayerType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Player",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

func player_projection(ev events.DeconstructedEvent, p *Player) (*Player, error) {
	switch {
	case ev.Type == "PlayerAddedEvent":
		// This should be the first event for any player
		// so we define p here. If another event comes first
		// we will try to modify player and an error will be thrown
		// because p is an empty pointer
		addEv, err := ev.ConvertTo(PlayerAddedEvent{})
		if err != nil {
			panic(err)
		}

		p, err = addedEventHandler(addEv)
		if err != nil {
			return nil, err
		}
	}

	return p, nil
}
