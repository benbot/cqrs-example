package events

import (
	"reflect"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// New Event Generator
func NewEvent(ev interface{}) EventRecord {
	return EventRecord{Type: reflect.TypeOf(ev).Name(), Timestamp: time.Now().Unix(), Event: ev}
}

// The EventRecord containing the event and metadata about it
type EventRecord struct {
	Id        bson.ObjectId `bson:"_id,omitempty"`
	Type      string        `bson:"type"`
	Timestamp int64         `bson:"created_at"`
	Event     interface{}   `bson:"event"`
}

//All of the event types
type PlayerAddedEvent struct {
	Id string `bson: id`
}
