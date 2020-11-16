package model

import (
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

type Data struct {
	ID bson.ObjectId `json:"_id" bson:"_id"`
	Instant time.Time `json:"instant" bson:"instant"`
	DeviceName string `json:"deviceName" bson:"deviceName"`
	ResourceName string `json:"resourceName" bson:"resourceName"`
	Value string `json:"value" bson:"value"`
}

type IoTData interface {
	SaveData (con MongoDB) (err error)
}

var DataInstance IoTData

func (d Data)SaveData (con MongoDB) (err error){
	err = con.CreateData(d)
	if err != nil {
		log.Println("It was not possible to save data ")
		return
	}
	log.Println("Message registered Successfully!")
	return

}