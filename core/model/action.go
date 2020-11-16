package model

import (
	"encoding/json"
	"log"
	"time"
)

type Action struct {
	ID string `json:"_id" bson:"_id"`
	Date time.Time `json:"date" bson:"date"`
	DeviceName string `json:"deviceName" bson:"deviceName"`
	ResourceName string `json:"resourceName" bson:"resourceName"`
	Command string `json:"command" bson:"command"`
}

type IoTAction interface {
	DelegateAction(con MongoDB) (err error)
}

var ActionInstance IoTAction

func (a Action)DelegateAction (con MongoDB) (err error) {
	device, err := con.GetDeviceByID(a.DeviceName)
	if err != nil {
		log.Println("It was not possible find the device ")
		return
	}
	content, err := json.Marshal(a)
	if err != nil{
		log.Println("It was not possible to Marshal action ")
		return

	}
	err = con.CreateAction(a)
	if err != nil {
		log.Println("It was not possible to save this action ")
		return
	}
	log.Println("Action saved successfully!")
	sendAction(device.GatewayUUID, device.UUID, string(content))
	return
}

func sendAction(gatewayID string, receiverID string, content string){
	log.Println("Delivering Message...")
	log.Println(gatewayID)
	log.Println(receiverID)
	log.Println(content)
	return
}

