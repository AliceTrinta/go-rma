package model

import (
	"time"
)

type Action struct {
	ID string `json:"_id" bson:"_id"`
	Date time.Time `json:"date" bson:"date"`
	DeviceName string `json:"deviceName" bson:"deviceName"`
	ResourceName string `json:"resourceName" bson:"resourceName"`
	Command string `json:"command" bson:"command"`
}
