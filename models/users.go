package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	//User represents a Single User in our Collection
	User struct {
		ID          bson.ObjectId `json:"id" bson: "_id"`
		FirtName    string        `json:"first_name" bson: "first_name"`
		MiddleName  string        `json:"middle_name" bson: "middle_name"`
		LastName    string        `json:"last_name" bson: "last_name"`
		Username    string        `json:"username" bson: "username"`
		Password    string        `json:"password" bson: "password"`
		Email       string        `json:"email" bson: "email"`
		IsAdmin     bool          `json:"is_admin" bson: "is_admin"`
		DateCreated time.Time     `json:"date_created" bson: "date_created"`
		DateUpdated time.Time     `json:"date_updated" bson: "date_updated"`
	}
)
