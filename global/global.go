package global

import "gopkg.in/mgo.v2"

var (
	Db              *mgo.Database
	USER_COLLECTION = "users"
)
