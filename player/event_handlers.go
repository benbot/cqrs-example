package player

import "fmt"

func added_event(ev *PlayerAddedEvent) (*Player, error) {
	fmt.Println(ev.Id)
	return &Player{
		Id: ev.Id,
	}, nil
}
