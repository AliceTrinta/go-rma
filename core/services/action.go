package services

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
	DelegateAction(e chan error)
}

func (a Action)DelegateAction (e chan error, con MongoDB) {
	err := con.CreateAction(a)
	if err != nil {
		log.Println("It was not possible to save this action ")
		e <- err
		return
	}
	log.Println("Action saved successfully!")
	device, err := con.GetDeviceByID(a.DeviceName)
	if err != nil {
		log.Println("It was not possible find the device's action ")
		e <- err
		return
	}
	content, err := json.Marshal(a)
	if err != nil{
		log.Println("It was not possible to Marshal action ")
		e <- err
		return

	}
	sendAction(device.GatewayUUID, device.UUID, content)
	return
}

func sendAction(gatewayID string, receiverID string, content []byte){
	log.Println("Delivering Message...")
	log.Println(gatewayID)
	log.Println(receiverID)
	log.Println(content)
	return
}

