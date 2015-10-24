package model

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	UserId   bson.ObjectId `bson:"_id,omitempty"`
	UserName string
	Email    string
	IsAdmin  bool
	Lang     int //0 zh-CN 1 en
}

func (this *User) IsLogin() bool {
	return true
}

func LoadUser(userId string) *User {
	return nil
}

func LoadUserAll(userId string) *User {
	return nil
}
