package entity

import (
	"encoding/json"
	"log"
	"time"
)

//An Action is an instance of a command that a resource can receive.
type Action struct {
	ID           int       `json:"_id" bson:"_id"`
	Date         time.Time `json:"date" bson:"date"`
	UUID         string    `json:"UUID" bson:"UUID"`
	ResourceName string    `json:"resourceName" bson:"resourceName"`
	Command      string    `json:"command" bson:"command"`
}

//IoTAction gather all the functions that an action can implement.
type IoTAction interface {
	DelegateAction(con MongoDB) (err error)
}

//ActionInstance is an instance of IoTAction.
var ActionInstance IoTAction

//DelegateAction function will delegate an action to it's respective Device to be executed.
func (a Action) DelegateAction(con MongoDB) (err error) {
	//Getting the UUID from the Device's action.
	device, err := con.GetDeviceByUUID(a.UUID)
	if err != nil {
		log.Println("It was not possible find the device ")
		return
	}
	//Transforming a model of an action into a JSON string.
	content, err := json.Marshal(a)
	if err != nil {
		log.Println("It was not possible to Marshal action ")
		return

	}
	//Registering the action in the database.
	err = con.CreateAction(a)
	if err != nil {
		log.Println("It was not possible to save this action ")
		return
	}
	//Sending the action to it's respective device.
	log.Println("Action saved successfully!")
	sendAction(device.GatewayUUID, device.UUID, string(content))
	return
}

//sendAction will send an action to a Device through and IoT network.
func sendAction(gatewayID string, receiverID string, content string) {
	log.Println("Delivering Message...")
	log.Println(gatewayID)
	log.Println(receiverID)
	log.Println(content)
	return
}
