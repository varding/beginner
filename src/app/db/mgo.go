package db

import (
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

//mgo session
var DbSession *mgo.Session

// 帖子
var Topics *mgo.Database

// 用户
var User *mgo.Database

// Node
var Node *mgo.Database
