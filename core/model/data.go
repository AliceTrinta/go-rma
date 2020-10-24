package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Data struct {
	ID bson.ObjectId `json:"_id" bson:"_id"`
	Instant time.Time `json:"instant" bson:"instant"`
	DeviceName string `json:"deviceName" bson:"deviceName"`
	ResourceName string `json:"resourceName" bson:"resourceName"`
	Value string `json:"value" bson:"value"`
}