package player

import "fmt"

//All of the event types for Player
type PlayerAddedEvent struct {
	Id string `bson:"id"`
}

func addedEventHandler(ev *PlayerAddedEvent) (*Player, error) {
	fmt.Println(ev.Id)
	return &Player{
		Id: ev.Id,
	}, nil
}
