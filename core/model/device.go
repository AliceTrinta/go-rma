package model

import (
	"encoding/json"
	"log"
	"time"
)

type Device struct {
	ID string                      `json:"_id" bson:"_id"`
	UUID string                    `json:"UUID" bson:"UUID"`
	GatewayUUID string             `json:"gateway_UUID" bson:"gatewayUUID"`
	Name string                    `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	CycleDelayInMillis json.Number `json:"cycleDelayInMillis" bson:"cycleDelayInMillis"`
	ResourceList []Resource        `json:"resourceList" bson:"resourceList"`
	LastUpdate time.Time           `json:"lastUpdate" bson:"lastUpdate"`
}

type IoTObject interface {
	StartDevice(con MongoDB) (err error)
}

var DeviceInstance IoTObject

func (d Device)StartDevice(con MongoDB) (err error){
	_, err = con.GetDeviceByID(d.ID)
	if err != nil{
		log.Println("Registering new Device")
		err = con.CreateDevice(d)
		if err != nil {
			log.Println("Device was not registered ")
			return
		}
		return
	}
	log.Println("This device is already registered")
	return
}