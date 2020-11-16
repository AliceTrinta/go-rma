package services

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
	SaveData(e chan error)
}

func (d Data)SaveData(e chan error, con MongoDB){
	err := con.CreateData(d)
	if err != nil {
		log.Println("It was not possible to save data ")
		e <- err
		return
	}
	log.Println("Message registered Successfully!")
	e <- err
	return

}