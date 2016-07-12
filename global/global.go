package global

import "gopkg.in/mgo.v2"

type Global struct {
	Db *mgo.Database
}

var G Global
