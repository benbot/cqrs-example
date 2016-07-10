package main

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

func NewEvent(t string, data interface{}) Event {
	return Event{Type: t, Timestamp: time.Now().Unix(), Data: data}
}

type Event struct {
	Id        bson.ObjectId `bson:"_id,omitempty"`
	Type      string        `bson:"type"`
	Timestamp int64         `bson:"created_at"`
	Data      interface{}
}

type PlayerAddedEventData struct {
	Id string `bson: id`
}
