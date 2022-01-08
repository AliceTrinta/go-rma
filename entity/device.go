package entity

import (
	"log"
)

//Device corresponds to an IoTObject at the Resource Management Layer (RMA).
type Device struct {
	UUID               string     `json:"UUID" bson:"UUID"`
	GatewayUUID        string     `json:"gateway_UUID" bson:"gatewayUUID"`
	Name               string     `json:"name" bson:"name"`
	Description        string     `json:"description" bson:"description"`
	CycleDelayInMillis string     `json:"cycleDelayInMillis" bson:"cycleDelayInMillis"`
	ResourceList       []Resource `json:"resourceList" bson:"resourceList"`
	Status             bool       `json:"status" bson:"status"`
}

//IoTObject Gather all the functions that a Device can implement.
type IoTObject interface {
	StartDevice(con MongoDB) (err error)
}

//DeviceInstance is and instance of an IoTObject
var DeviceInstance IoTObject

//StartDevice Function will register a Device in the database.
func (d Device) StartDevice(con MongoDB) (err error) {
	//Verifying if the Device already exist.
	_, err = con.GetDeviceByUUID(d.UUID)
	if err != nil {
		log.Println("Registering new Device")
		//If does not exist, register a new device in the database.
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
