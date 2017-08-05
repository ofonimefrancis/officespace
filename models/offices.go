package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	//Offices represents an Office instance
	Offices struct {
		ID          bson.ObjectId `json:"id" bson:"_id"`
		Name        string        `json:"name" bson:"name"`
		Type        []string      `json:"type" bson:"type"`
		SeatCount   int32         `json:"seat_count" bson:"seat_count"`
		Slug        string        `json:"slug" bson:"slug"`
		Host        User          `json:"host" bson:"host"`
		Price       Price         `json:"price" bson:"price"`
		DateCreated time.Time     `json:"date_created" bson:"date_created"`
		DateUpdated time.Time     `json:"date_updated" bson:"date_updated"`
	}
)
